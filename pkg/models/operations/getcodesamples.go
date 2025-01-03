// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetCodeSamplesRequest struct {
	Languages    []string `queryParam:"style=form,explode=true,name=languages"`
	OperationIds []string `queryParam:"style=form,explode=true,name=operation_ids"`
	// The registry URL from which to retrieve the snippets. E.g. https://spec.speakeasy.com/org/ws/my-source
	RegistryURL string `queryParam:"style=form,explode=true,name=registry_url"`
}

func (o *GetCodeSamplesRequest) GetLanguages() []string {
	if o == nil {
		return nil
	}
	return o.Languages
}

func (o *GetCodeSamplesRequest) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *GetCodeSamplesRequest) GetRegistryURL() string {
	if o == nil {
		return ""
	}
	return o.RegistryURL
}

type GetCodeSamplesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UsageSnippets *shared.UsageSnippets
}

func (o *GetCodeSamplesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCodeSamplesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCodeSamplesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCodeSamplesResponse) GetUsageSnippets() *shared.UsageSnippets {
	if o == nil {
		return nil
	}
	return o.UsageSnippets
}
