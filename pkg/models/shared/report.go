// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypeLinting Type = "linting"
	TypeChanges Type = "changes"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linting":
		fallthrough
	case "changes":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Report struct {
	Type *Type `json:"type,omitempty"`
}

func (o *Report) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
