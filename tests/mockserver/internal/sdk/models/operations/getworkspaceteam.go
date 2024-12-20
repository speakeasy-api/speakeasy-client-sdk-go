// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetWorkspaceTeamGlobals struct {
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *GetWorkspaceTeamGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type GetWorkspaceTeamRequest struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *GetWorkspaceTeamRequest) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type GetWorkspaceTeamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	WorkspaceTeamResponse *components.WorkspaceTeamResponse
}

func (o *GetWorkspaceTeamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWorkspaceTeamResponse) GetWorkspaceTeamResponse() *components.WorkspaceTeamResponse {
	if o == nil {
		return nil
	}
	return o.WorkspaceTeamResponse
}
