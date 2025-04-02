// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatYaml Format = "yaml"
)

func (e Format) ToPointer() *Format {
	return &e
}
func (e *Format) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "yaml":
		*e = Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Format: %v", v)
	}
}

type SchemaStoreItem struct {
	CreatedAt    time.Time `json:"created_at"`
	Format       Format    `json:"format"`
	ID           string    `json:"id"`
	PackageName  string    `json:"packageName"`
	SDKClassname string    `json:"sdkClassname"`
	Spec         string    `json:"spec"`
}

func (s SchemaStoreItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemaStoreItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemaStoreItem) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *SchemaStoreItem) GetFormat() Format {
	if o == nil {
		return Format("")
	}
	return o.Format
}

func (o *SchemaStoreItem) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SchemaStoreItem) GetPackageName() string {
	if o == nil {
		return ""
	}
	return o.PackageName
}

func (o *SchemaStoreItem) GetSDKClassname() string {
	if o == nil {
		return ""
	}
	return o.SDKClassname
}

func (o *SchemaStoreItem) GetSpec() string {
	if o == nil {
		return ""
	}
	return o.Spec
}
