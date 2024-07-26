// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetOrganizationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *sdkerrors.Error
	// OK
	Organizations []shared.Organization
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetOrganizationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOrganizationsResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetOrganizationsResponse) GetOrganizations() []shared.Organization {
	if o == nil {
		return nil
	}
	return o.Organizations
}

func (o *GetOrganizationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOrganizationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
