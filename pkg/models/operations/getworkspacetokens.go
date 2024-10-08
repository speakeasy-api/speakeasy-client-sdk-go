// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetWorkspaceTokensRequest struct {
	// Unique identifier of the workspace.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *GetWorkspaceTokensRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type GetWorkspaceTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.WorkspaceToken
}

func (o *GetWorkspaceTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWorkspaceTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWorkspaceTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWorkspaceTokensResponse) GetClasses() []shared.WorkspaceToken {
	if o == nil {
		return nil
	}
	return o.Classes
}