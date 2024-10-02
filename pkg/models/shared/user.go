// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

type User struct {
	// Indicates whether the user is an admin.
	Admin bool `json:"admin"`
	// Indicates whether the user has been confirmed.
	Confirmed bool `json:"confirmed"`
	// Timestamp of the user's creation.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the default workspace.
	DefaultWorkspaceID *string `json:"default_workspace_id,omitempty"`
	// Display name of the user.
	DisplayName string `json:"display_name"`
	// Email address of the user.
	Email string `json:"email"`
	// Indicates whether the email address has been verified.
	EmailVerified bool `json:"email_verified"`
	// GitHub handle of the user.
	GithubHandle *string `json:"github_handle,omitempty"`
	// Unique identifier for the user.
	ID string `json:"id"`
	// Timestamp of the last login.
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	// URL of the user's photo.
	PhotoURL *string `json:"photo_url,omitempty"`
	// Timestamp of the user's last update.
	UpdatedAt time.Time `json:"updated_at"`
	// Indicates whether the user has been whitelisted.
	Whitelisted bool `json:"whitelisted"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *User) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *User) GetAdmin() bool {
	if o == nil {
		return false
	}
	return o.Admin
}

func (o *User) GetConfirmed() bool {
	if o == nil {
		return false
	}
	return o.Confirmed
}

func (o *User) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *User) GetDefaultWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.DefaultWorkspaceID
}

func (o *User) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *User) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *User) GetEmailVerified() bool {
	if o == nil {
		return false
	}
	return o.EmailVerified
}

func (o *User) GetGithubHandle() *string {
	if o == nil {
		return nil
	}
	return o.GithubHandle
}

func (o *User) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *User) GetLastLoginAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastLoginAt
}

func (o *User) GetPhotoURL() *string {
	if o == nil {
		return nil
	}
	return o.PhotoURL
}

func (o *User) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *User) GetWhitelisted() bool {
	if o == nil {
		return false
	}
	return o.Whitelisted
}
