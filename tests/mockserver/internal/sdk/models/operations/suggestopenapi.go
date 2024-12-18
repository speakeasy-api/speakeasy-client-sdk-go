// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"mockserver/internal/sdk/models/components"
)

type Schema struct {
	Content  io.Reader `multipartForm:"content"`
	FileName string    `multipartForm:"name=fileName"`
}

func (o *Schema) GetContent() io.Reader {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *Schema) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// SuggestOpenAPIRequestBody - The schema file to upload provided as a multipart/form-data file segment.
type SuggestOpenAPIRequestBody struct {
	Opts   *components.SuggestOptsOld `multipartForm:"name=opts,json"`
	Schema Schema                     `multipartForm:"file"`
}

func (o *SuggestOpenAPIRequestBody) GetOpts() *components.SuggestOptsOld {
	if o == nil {
		return nil
	}
	return o.Opts
}

func (o *SuggestOpenAPIRequestBody) GetSchema() Schema {
	if o == nil {
		return Schema{}
	}
	return o.Schema
}

type SuggestOpenAPIRequest struct {
	// The schema file to upload provided as a multipart/form-data file segment.
	RequestBody SuggestOpenAPIRequestBody `request:"mediaType=multipart/form-data"`
	XSessionID  string                    `header:"style=simple,explode=false,name=x-session-id"`
}

func (o *SuggestOpenAPIRequest) GetRequestBody() SuggestOpenAPIRequestBody {
	if o == nil {
		return SuggestOpenAPIRequestBody{}
	}
	return o.RequestBody
}

func (o *SuggestOpenAPIRequest) GetXSessionID() string {
	if o == nil {
		return ""
	}
	return o.XSessionID
}

type SuggestOpenAPIResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// An overlay containing the suggested spec modifications.
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Schema io.ReadCloser
}

func (o *SuggestOpenAPIResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SuggestOpenAPIResponse) GetSchema() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Schema
}
