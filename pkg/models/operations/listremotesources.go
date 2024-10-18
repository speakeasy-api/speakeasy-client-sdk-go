// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type ListRemoteSourcesRequest struct {
	NamespaceName string `queryParam:"style=form,explode=true,name=namespace_name"`
}

func (o *ListRemoteSourcesRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

type ListRemoteSourcesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	RemoteSource *shared.RemoteSource
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListRemoteSourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRemoteSourcesResponse) GetRemoteSource() *shared.RemoteSource {
	if o == nil {
		return nil
	}
	return o.RemoteSource
}

func (o *ListRemoteSourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRemoteSourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
