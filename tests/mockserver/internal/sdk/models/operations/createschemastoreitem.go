// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
)

// Format - The format of the OpenAPI specification.
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

type CreateSchemaStoreItemRequestBody struct {
	// The format of the OpenAPI specification.
	Format Format `json:"format"`
	// The package name to use in code snippets / quickstart.
	PackageName string `json:"packageName"`
	// The classname of the SDK to use in code snippets / quickstart.
	SDKClassname string `json:"sdkClassname"`
	// The OpenAPI specification to store.
	Spec string `json:"spec"`
}

func (o *CreateSchemaStoreItemRequestBody) GetFormat() Format {
	if o == nil {
		return Format("")
	}
	return o.Format
}

func (o *CreateSchemaStoreItemRequestBody) GetPackageName() string {
	if o == nil {
		return ""
	}
	return o.PackageName
}

func (o *CreateSchemaStoreItemRequestBody) GetSDKClassname() string {
	if o == nil {
		return ""
	}
	return o.SDKClassname
}

func (o *CreateSchemaStoreItemRequestBody) GetSpec() string {
	if o == nil {
		return ""
	}
	return o.Spec
}

type CreateSchemaStoreItemResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	SchemaStoreItem *components.SchemaStoreItem
}

func (o *CreateSchemaStoreItemResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateSchemaStoreItemResponse) GetSchemaStoreItem() *components.SchemaStoreItem {
	if o == nil {
		return nil
	}
	return o.SchemaStoreItem
}
