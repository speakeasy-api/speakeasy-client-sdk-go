// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

type Claims struct {
}

type AccessTokenUser struct {
	Admin         *bool      `json:"admin,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	DisplayName   *string    `json:"display_name,omitempty"`
	Email         *string    `json:"email,omitempty"`
	EmailVerified *bool      `json:"email_verified,omitempty"`
	ID            *string    `json:"id,omitempty"`
}

func (a AccessTokenUser) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccessTokenUser) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccessTokenUser) GetAdmin() *bool {
	if o == nil {
		return nil
	}
	return o.Admin
}

func (o *AccessTokenUser) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccessTokenUser) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AccessTokenUser) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AccessTokenUser) GetEmailVerified() *bool {
	if o == nil {
		return nil
	}
	return o.EmailVerified
}

func (o *AccessTokenUser) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Workspaces struct {
	AccountType *AccountType `json:"account_type,omitempty"`
	ID          *string      `json:"id,omitempty"`
	Name        *string      `json:"name,omitempty"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
}

func (w Workspaces) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *Workspaces) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Workspaces) GetAccountType() *AccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *Workspaces) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Workspaces) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Workspaces) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// An AccessToken is a token that can be used to authenticate with the Speakeasy API.
type AccessToken struct {
	AccessToken  string          `json:"access_token"`
	Claims       Claims          `json:"claims"`
	FeatureFlags []FeatureFlag   `json:"feature_flags,omitempty"`
	User         AccessTokenUser `json:"user"`
	Workspaces   []Workspaces    `json:"workspaces,omitempty"`
}

func (o *AccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *AccessToken) GetClaims() Claims {
	if o == nil {
		return Claims{}
	}
	return o.Claims
}

func (o *AccessToken) GetFeatureFlags() []FeatureFlag {
	if o == nil {
		return nil
	}
	return o.FeatureFlags
}

func (o *AccessToken) GetUser() AccessTokenUser {
	if o == nil {
		return AccessTokenUser{}
	}
	return o.User
}

func (o *AccessToken) GetWorkspaces() []Workspaces {
	if o == nil {
		return nil
	}
	return o.Workspaces
}
