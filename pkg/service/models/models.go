package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
)

func (x *UserMeta_UserStatus) UnmarshalJSON(bytes []byte) error {
	if strings.HasPrefix(string(bytes), `"`) {
		var name string
		err := json.Unmarshal(bytes, &name)
		if err != nil {
			return err
		}
		s, ok := UserMeta_UserStatus_value[name]
		if !ok {
			return errors.ParameterError(fmt.Sprintf("unknown status: %s", name))
		}
		*x = UserMeta_UserStatus(s)
	} else {
		var val int32
		err := json.Unmarshal(bytes, &val)
		if err != nil {
			return err
		}

		if _, ok := UserMeta_UserStatus_name[val]; !ok {
			return errors.ParameterError(fmt.Sprintf("unknown status: %d", val))
		}
		*x = UserMeta_UserStatus(val)
	}

	return nil
}

func (x UserMeta_UserStatus) Is(s ...UserMeta_UserStatus) bool {
	for _, status := range s {
		if x&status != status {
			return false
		}
	}
	return true
}

func (x UserMeta_UserStatus) IsAnyOne(s ...UserMeta_UserStatus) bool {
	for _, status := range s {
		if x&status == status {
			return true
		}
	}
	return false
}

const (
	UserMetaStatusAll UserMeta_UserStatus = -1
)
