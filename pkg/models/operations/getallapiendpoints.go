// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetAllAPIEndpointsRequest struct {
	// The ID of the Api to retrieve ApiEndpoints for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

func (o *GetAllAPIEndpointsRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

type GetAllAPIEndpointsResponse struct {
	// OK
	APIEndpoints []shared.APIEndpoint
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAllAPIEndpointsResponse) GetAPIEndpoints() []shared.APIEndpoint {
	if o == nil {
		return nil
	}
	return o.APIEndpoints
}

func (o *GetAllAPIEndpointsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllAPIEndpointsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllAPIEndpointsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
