package rest

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/sidhman/clean-architecture/pkg/models"
)

// Valid URL query params
const (
	QueryParamUser = "user"
	QueryParamKey  = "key"
)

// Possible HTTP errors
var (
	ErrKeyMissing    = errors.New("key missing")
	ErrUserMissing   = errors.New("user missing")
	ErrValueMissing  = errors.New("value missing")
	ErrValueInvalid  = errors.New("invalid value")
	ErrForbidden     = errors.New("forbidden")
	ErrInternalError = errors.New("internal server error")
)

type valueAttrs struct {
	Meta  map[string]string `json:"meta"`
	Count *int              `json:"count"`
}

type userValue struct {
	Data  string `json:"data"`
	Attrs valueAttrs
}

func (uv *userValue) Marshal() ([]byte, error) {
	return json.Marshal(uv)
}

func (uv *userValue) Unmarshal(data []byte) error {
	return json.Unmarshal(data, uv)
}

func fromValueModel(v models.Value) userValue {
	return userValue{
		Data: v.Data,
		Attrs: valueAttrs{
			Meta:  v.Meta,
			Count: v.Count,
		},
	}
}

func toValueModel(uv userValue) models.Value {
	return models.Value{
		Data:  uv.Data,
		Meta:  uv.Attrs.Meta,
		Count: uv.Attrs.Count,
	}
}
