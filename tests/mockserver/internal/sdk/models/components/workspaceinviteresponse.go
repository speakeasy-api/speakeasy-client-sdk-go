// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Relationship struct {
	UserID      string `json:"user_id"`
	WorkspaceID string `json:"workspace_id"`
}

func (o *Relationship) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *Relationship) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

// WorkspaceInviteResponse - A response for workspace user invite
type WorkspaceInviteResponse struct {
	InviteLink   *string      `json:"invite_link,omitempty"`
	Relationship Relationship `json:"relationship"`
}

func (o *WorkspaceInviteResponse) GetInviteLink() *string {
	if o == nil {
		return nil
	}
	return o.InviteLink
}

func (o *WorkspaceInviteResponse) GetRelationship() Relationship {
	if o == nil {
		return Relationship{}
	}
	return o.Relationship
}
