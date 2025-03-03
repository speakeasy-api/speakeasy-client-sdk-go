// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"io"
)

type SchemaFile struct {
	Content  io.Reader `multipartForm:"content"`
	FileName string    `multipartForm:"name=fileName"`
}

func (o *SchemaFile) GetContent() io.Reader {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *SchemaFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

type CodeSampleSchemaInput struct {
	// The language to generate code samples for
	Language string `multipartForm:"name=language"`
	// A list of operations IDs to generate code samples for
	OperationIds []string `multipartForm:"name=operation_ids"`
	// The name of the package
	PackageName *string `multipartForm:"name=package_name"`
	// The OpenAPI file to be uploaded
	SchemaFile SchemaFile `multipartForm:"file"`
	// The SDK class name
	SDKClassName *string `multipartForm:"name=sdk_class_name"`
}

func (o *CodeSampleSchemaInput) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

func (o *CodeSampleSchemaInput) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *CodeSampleSchemaInput) GetPackageName() *string {
	if o == nil {
		return nil
	}
	return o.PackageName
}

func (o *CodeSampleSchemaInput) GetSchemaFile() SchemaFile {
	if o == nil {
		return SchemaFile{}
	}
	return o.SchemaFile
}

func (o *CodeSampleSchemaInput) GetSDKClassName() *string {
	if o == nil {
		return nil
	}
	return o.SDKClassName
}
