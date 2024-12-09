// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"mockserver/internal/sdk/models/components"
)

type RegisterSchemaFile struct {
	Content  io.Reader `multipartForm:"content"`
	FileName string    `multipartForm:"name=file"`
}

func (o *RegisterSchemaFile) GetContent() io.Reader {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *RegisterSchemaFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// RegisterSchemaRequestBody - The schema file to upload provided as a multipart/form-data file segment.
type RegisterSchemaRequestBody struct {
	File RegisterSchemaFile `multipartForm:"file"`
}

func (o *RegisterSchemaRequestBody) GetFile() RegisterSchemaFile {
	if o == nil {
		return RegisterSchemaFile{}
	}
	return o.File
}

type RegisterSchemaRequest struct {
	// The schema file to upload provided as a multipart/form-data file segment.
	RequestBody RegisterSchemaRequestBody `request:"mediaType=multipart/form-data"`
	// The ID of the Api to get the schema for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *RegisterSchemaRequest) GetRequestBody() RegisterSchemaRequestBody {
	if o == nil {
		return RegisterSchemaRequestBody{}
	}
	return o.RequestBody
}

func (o *RegisterSchemaRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *RegisterSchemaRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type RegisterSchemaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *RegisterSchemaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
