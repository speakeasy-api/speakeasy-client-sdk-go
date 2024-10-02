// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetSchemaDiffRequest struct {
	// The ID of the Api to retrieve schemas for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The base revision ID of the schema to retrieve.
	BaseRevisionID string `pathParam:"style=simple,explode=false,name=baseRevisionID"`
	// The target revision ID of the schema to retrieve.
	TargetRevisionID string `pathParam:"style=simple,explode=false,name=targetRevisionID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GetSchemaDiffRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetSchemaDiffRequest) GetBaseRevisionID() string {
	if o == nil {
		return ""
	}
	return o.BaseRevisionID
}

func (o *GetSchemaDiffRequest) GetTargetRevisionID() string {
	if o == nil {
		return ""
	}
	return o.TargetRevisionID
}

func (o *GetSchemaDiffRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GetSchemaDiffResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	SchemaDiff *shared.SchemaDiff
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSchemaDiffResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSchemaDiffResponse) GetSchemaDiff() *shared.SchemaDiff {
	if o == nil {
		return nil
	}
	return o.SchemaDiff
}

func (o *GetSchemaDiffResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSchemaDiffResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
