// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

// QueryParamOp - Configuration for filter operations
type QueryParamOp struct {
	// Whether to AND or OR the filters
	And bool `queryParam:"name=and"`
}

func (o *QueryParamOp) GetAnd() bool {
	if o == nil {
		return false
	}
	return o.And
}

type GetApisRequest struct {
	// Metadata to filter Apis on
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	// Configuration for filter operations
	Op *QueryParamOp `queryParam:"style=deepObject,explode=true,name=op"`
}

func (o *GetApisRequest) GetMetadata() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *GetApisRequest) GetOp() *QueryParamOp {
	if o == nil {
		return nil
	}
	return o.Op
}

type GetApisResponse struct {
	// OK
	Apis []shared.API
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *sdkerrors.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetApisResponse) GetApis() []shared.API {
	if o == nil {
		return nil
	}
	return o.Apis
}

func (o *GetApisResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetApisResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetApisResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetApisResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
