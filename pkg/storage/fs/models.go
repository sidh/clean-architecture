package fs

import (
	"encoding/json"

	"github.com/sidhman/clean-architecture/pkg/models"
)

type intPtr struct {
	Int   int64 `json:"int,omitempty"`
	Valid bool  `json:"valid"`
}

type storedValue struct {
	Data  string            `json:"my_best_data_name"`
	Meta  map[string]string `json:"THE_stuff"`
	Count intPtr            `json:"count"`
}

func (sv *storedValue) Marshal() ([]byte, error) {
	return json.Marshal(sv)
}

func (sv *storedValue) Unmarshal(data []byte) error {
	return json.Unmarshal(data, sv)
}

func fromValueModel(v models.Value) storedValue {
	var count intPtr
	if v.Count != nil {
		count.Int = int64(*v.Count)
		count.Valid = true
	}

	return storedValue{
		Data:  v.Data,
		Meta:  v.Meta,
		Count: count,
	}
}

func toValueModel(sv storedValue) models.Value {
	var count *int
	if sv.Count.Valid {
		c := int(sv.Count.Int)
		count = &c
	}

	return models.Value{
		Data:  sv.Data,
		Meta:  sv.Meta,
		Count: count,
	}
}
