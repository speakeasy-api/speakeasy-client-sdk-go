// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetSchemasRequest struct {
	// The ID of the Api to retrieve schemas for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GetSchemasRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetSchemasRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GetSchemasResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *sdkerrors.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.Schema
}

func (o *GetSchemasResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSchemasResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetSchemasResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSchemasResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSchemasResponse) GetClasses() []shared.Schema {
	if o == nil {
		return nil
	}
	return o.Classes
}
