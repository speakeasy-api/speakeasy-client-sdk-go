// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type FindAPIEndpointRequest struct {
	// The ID of the Api the ApiEndpoint belongs to.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The displayName of the ApiEndpoint to find (set by operationId from OpenAPI schema).
	DisplayName string `pathParam:"style=simple,explode=false,name=displayName"`
	// The version ID of the Api the ApiEndpoint belongs to.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *FindAPIEndpointRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *FindAPIEndpointRequest) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *FindAPIEndpointRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type FindAPIEndpointResponse struct {
	// OK
	APIEndpoint *shared.APIEndpoint
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *FindAPIEndpointResponse) GetAPIEndpoint() *shared.APIEndpoint {
	if o == nil {
		return nil
	}
	return o.APIEndpoint
}

func (o *FindAPIEndpointResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FindAPIEndpointResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FindAPIEndpointResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
