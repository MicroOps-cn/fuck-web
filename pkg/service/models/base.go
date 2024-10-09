package models

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"

	g "github.com/MicroOps-cn/fuck/generator"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (m *Model) BeforeCreate(db *gorm.DB) error {
	if m.Id == "" {
		id := g.NewId(db.Statement.Table)
		if len(id) != 36 {
			return errors.New("Failed to generate ID: " + id)
		}
		db.Statement.SetColumn("Id", id)
	}
	if m.UpdateTime.IsZero() {
		db.Statement.SetColumn("UpdateTime", time.Now().UTC())
	}
	if m.CreateTime.IsZero() {
		db.Statement.SetColumn("CreateTime", time.Now().UTC())
	}
	if m.IsDelete && !m.DeleteTime.Valid {
		m.DeleteTime = gorm.DeletedAt{Time: time.Now().UTC(), Valid: true}
		db.Statement.SetColumn("DeleteTime", m.DeleteTime)
	}
	return nil
}

func (m *Model) BeforeSave(db *gorm.DB) error {
	db.Statement.SetColumn("UpdateTime", time.Now().UTC())
	if m.IsDelete && !m.DeleteTime.Valid {
		m.DeleteTime = gorm.DeletedAt{Time: time.Now().UTC(), Valid: true}
		db.Statement.SetColumn("DeleteTime", m.DeleteTime)
	}
	return nil
}

func (m *Model) AfterFind(_ *gorm.DB) error {
	if m.DeleteTime.Valid {
		m.IsDelete = true
	}
	return nil
}

type CompressField []byte

func (c *CompressField) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*c))
}

func (c *CompressField) GormDataType() string {
	return "blob"
}

// Scan implements the Scanner interface.
func (c *CompressField) Scan(value any) error {
	var val []byte
	switch vt := value.(type) {
	case []uint8:
		val = vt
	case string:
		val = []byte(vt)
	default:
		return fmt.Errorf("failed to resolve field, type exception: %T", value)
	}
	if len(val) > 0 {
		if val[0] == 0x78 {
			r, err := zlib.NewReader(bytes.NewBuffer(val))
			if err != nil {
				return err
			}
			*c, err = io.ReadAll(r)
			if err != nil && err != io.ErrUnexpectedEOF {
				return err
			}
		} else {
			*c = val
		}
	}
	return nil
}

func (c CompressField) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	buf := bytes.NewBuffer(nil)
	w := zlib.NewWriter(buf)
	_, err := w.Write(c)
	if err != nil {
		_ = db.AddError(err)
	} else if err = w.Flush(); err != nil {
		_ = db.AddError(err)
	} else {
		return clause.Expr{
			SQL:  "from_base64(?)",
			Vars: []interface{}{base64.StdEncoding.EncodeToString(buf.Bytes())},
		}
	}
	return clause.Expr{}
}

// Value implements the driver Valuer interface.
//func (c CompressField) Value() (driver.Value, error) {
//	buf := bytes.NewBuffer(nil)
//	w := zlib.NewWriter(buf)
//	_, err := w.Write(c)
//	if err != nil {
//		return nil, err
//	}
//	if err = w.Flush(); err != nil {
//		return nil, err
//	}
//	return buf.String(), err
//}
