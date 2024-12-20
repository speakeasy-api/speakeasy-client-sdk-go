// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type RevokeUserAccessToWorkspaceGlobals struct {
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *RevokeUserAccessToWorkspaceGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type RevokeUserAccessToWorkspaceRequest struct {
	// Unique identifier of the user.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
	// Unique identifier of the workspace.
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *RevokeUserAccessToWorkspaceRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *RevokeUserAccessToWorkspaceRequest) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type RevokeUserAccessToWorkspaceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *RevokeUserAccessToWorkspaceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
