// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"net/http"
)

type CheckAccessRequest struct {
	Org  string `queryParam:"style=form,explode=true,name=org"`
	Repo string `queryParam:"style=form,explode=true,name=repo"`
}

func (o *CheckAccessRequest) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *CheckAccessRequest) GetRepo() string {
	if o == nil {
		return ""
	}
	return o.Repo
}

type CheckAccessResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *sdkerrors.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CheckAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckAccessResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *CheckAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
