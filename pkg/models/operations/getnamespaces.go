// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetNamespacesRequest struct {
	IsComposite *bool `queryParam:"style=form,explode=true,name=is_composite"`
}

func (o *GetNamespacesRequest) GetIsComposite() *bool {
	if o == nil {
		return nil
	}
	return o.IsComposite
}

type GetNamespacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	GetNamespacesResponse *shared.GetNamespacesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetNamespacesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNamespacesResponse) GetGetNamespacesResponse() *shared.GetNamespacesResponse {
	if o == nil {
		return nil
	}
	return o.GetNamespacesResponse
}

func (o *GetNamespacesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNamespacesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
