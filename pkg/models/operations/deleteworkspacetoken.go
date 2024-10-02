// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteWorkspaceTokenRequest struct {
	// Unique identifier of the token.
	TokenID string `pathParam:"style=simple,explode=false,name=tokenID"`
	// Unique identifier of the workspace.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *DeleteWorkspaceTokenRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

func (o *DeleteWorkspaceTokenRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type DeleteWorkspaceTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteWorkspaceTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteWorkspaceTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteWorkspaceTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
