// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GrantUserAccessToWorkspaceRequest struct {
	// Email of the user to grant access to.
	Email string `pathParam:"style=simple,explode=false,name=email"`
	// Unique identifier of the workspace.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *GrantUserAccessToWorkspaceRequest) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *GrantUserAccessToWorkspaceRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type GrantUserAccessToWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	WorkspaceInviteResponse *shared.WorkspaceInviteResponse
}

func (o *GrantUserAccessToWorkspaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GrantUserAccessToWorkspaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GrantUserAccessToWorkspaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GrantUserAccessToWorkspaceResponse) GetWorkspaceInviteResponse() *shared.WorkspaceInviteResponse {
	if o == nil {
		return nil
	}
	return o.WorkspaceInviteResponse
}