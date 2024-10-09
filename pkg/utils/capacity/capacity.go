package capacity

import (
	"encoding/json"

	"github.com/MicroOps-cn/fuck/capacity"
	"github.com/gogo/protobuf/jsonpb"
)

func (m *Capacity) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return json.Unmarshal(b, (*capacity.Capacities)(&m.Capacity))
}

func NewCapacity(c int64) *Capacity {
	return &Capacity{Capacity: c}
}
