//go:build !make_test

package config

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/MicroOps-cn/fuck/clients/gorm"
	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/MicroOps-cn/fuck-web/pkg/testutils"
)

func TestConfig(t *testing.T) {
	logger := logs.New(logs.WithConfig(logs.MustNewConfig("debug", "logfmt")))
	logs.SetDefaultLogger(logger)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()
	ctx, logger = logs.NewContextLogger(ctx)
	tablePrefix := "gm_" + rand.String(10)
	schema := "gm_" + rand.String(10)
	var rawCfg string
	testutils.RunWithMySQLContainer(ctx, t, func(host, rootPassword string) {
		var testCfg Config
		t.Run("Test Marshal Config", func(t *testing.T) {
			mysqlOptions := gorm.NewMySQLOptions()
			mysqlOptions.Host = host
			mysqlOptions.Username = "root"
			mysqlOptions.Password.SetValue(rootPassword)
			mysqlOptions.TablePrefix = tablePrefix
			mysqlOptions.Schema = schema
			marshaler := jsonpb.Marshaler{
				Indent:   "    ",
				OrigName: true,
			}
			client := &gorm.MySQLClient{}
			client.SetOptions(mysqlOptions)
			testCfg.Storage = &Storages{
				Default: &Storage{
					Source: &Storage_Mysql{
						Mysql: client,
					},
				},
			}
			buf := bytes.NewBuffer(nil)
			err := marshaler.Marshal(buf, &testCfg)
			require.NoError(t, err, "Failed to Marshal config")
			fmt.Println(buf.String())
			rawCfg = buf.String()
		})
		t.Run("Test Unmarshal Config", func(t *testing.T) {
			err := safeCfg.ReloadConfigFromYamlReader(logger, NewConverter("./", bytes.NewReader([]byte(rawCfg))))
			require.NoError(t, err, "Failed to Unmarshal config")
			dftSource, ok := safeCfg.C.Storage.Default.Source.(*Storage_Mysql)
			require.True(t, ok)
			require.Equal(t, dftSource.Mysql.Options().Host, host)
			require.Equal(t, dftSource.Mysql.Options().Username, "root")
			require.Equal(t, dftSource.Mysql.Options().Schema, schema)
			require.Equal(t, dftSource.Mysql.Options().Password, rootPassword)
			require.Equal(t, dftSource.Mysql.Options().TablePrefix, tablePrefix)
		})
	})
}
