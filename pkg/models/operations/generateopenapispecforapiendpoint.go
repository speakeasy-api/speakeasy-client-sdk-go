// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GenerateOpenAPISpecForAPIEndpointRequest struct {
	// The ID of the ApiEndpoint to generate an OpenAPI specification for.
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	// The ID of the Api to generate an OpenAPI specification for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to generate an OpenAPI specification for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GenerateOpenAPISpecForAPIEndpointRequest) GetAPIEndpointID() string {
	if o == nil {
		return ""
	}
	return o.APIEndpointID
}

func (o *GenerateOpenAPISpecForAPIEndpointRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GenerateOpenAPISpecForAPIEndpointRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GenerateOpenAPISpecForAPIEndpointResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GenerateOpenAPISpecForAPIEndpointResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateOpenAPISpecForAPIEndpointResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GenerateOpenAPISpecForAPIEndpointResponse) GetGenerateOpenAPISpecDiff() *shared.GenerateOpenAPISpecDiff {
	if o == nil {
		return nil
	}
	return o.GenerateOpenAPISpecDiff
}

func (o *GenerateOpenAPISpecForAPIEndpointResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateOpenAPISpecForAPIEndpointResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
