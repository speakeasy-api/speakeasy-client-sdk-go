// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"io"
	"net/http"
)

type GetBlobRequest struct {
	Digest           string `pathParam:"style=simple,explode=false,name=digest"`
	NamespaceName    string `pathParam:"style=simple,explode=false,name=namespace_name"`
	OrganizationSlug string `pathParam:"style=simple,explode=false,name=organization_slug"`
	WorkspaceSlug    string `pathParam:"style=simple,explode=false,name=workspace_slug"`
}

func (o *GetBlobRequest) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

func (o *GetBlobRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *GetBlobRequest) GetOrganizationSlug() string {
	if o == nil {
		return ""
	}
	return o.OrganizationSlug
}

func (o *GetBlobRequest) GetWorkspaceSlug() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceSlug
}

type GetBlobResponse struct {
	// OK
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Blob io.ReadCloser
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *sdkerrors.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBlobResponse) GetBlob() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Blob
}

func (o *GetBlobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBlobResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetBlobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBlobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
