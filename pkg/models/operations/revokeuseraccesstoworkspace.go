// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RevokeUserAccessToWorkspaceRequest struct {
	// Unique identifier of the user.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
	// Unique identifier of the workspace.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *RevokeUserAccessToWorkspaceRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *RevokeUserAccessToWorkspaceRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type RevokeUserAccessToWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RevokeUserAccessToWorkspaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RevokeUserAccessToWorkspaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RevokeUserAccessToWorkspaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
