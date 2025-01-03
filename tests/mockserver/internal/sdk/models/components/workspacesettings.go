// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

type WorkspaceSettings struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	WebhookURL  string    `json:"webhook_url"`
	WorkspaceID string    `json:"workspace_id"`
}

func (w WorkspaceSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspaceSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WorkspaceSettings) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *WorkspaceSettings) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *WorkspaceSettings) GetWebhookURL() string {
	if o == nil {
		return ""
	}
	return o.WebhookURL
}

func (o *WorkspaceSettings) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}