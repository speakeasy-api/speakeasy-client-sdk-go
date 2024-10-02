// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type CreateWorkspaceTokenRequest struct {
	WorkspaceToken shared.WorkspaceToken `request:"mediaType=application/json"`
	// Unique identifier of the workspace.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *CreateWorkspaceTokenRequest) GetWorkspaceToken() shared.WorkspaceToken {
	if o == nil {
		return shared.WorkspaceToken{}
	}
	return o.WorkspaceToken
}

func (o *CreateWorkspaceTokenRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type CreateWorkspaceTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateWorkspaceTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWorkspaceTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWorkspaceTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
