// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DownloadSchemaRequest struct {
	// The ID of the Api to download the schema for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *DownloadSchemaRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *DownloadSchemaRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type DownloadSchemaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	Schema []byte
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DownloadSchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadSchemaResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DownloadSchemaResponse) GetSchema() []byte {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DownloadSchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadSchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
