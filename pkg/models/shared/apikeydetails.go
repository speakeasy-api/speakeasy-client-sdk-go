// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountType string

const (
	AccountTypeFree       AccountType = "free"
	AccountTypeScaleUp    AccountType = "scale-up"
	AccountTypeEnterprise AccountType = "enterprise"
)

func (e AccountType) ToPointer() *AccountType {
	return &e
}

func (e *AccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "free":
		fallthrough
	case "scale-up":
		fallthrough
	case "enterprise":
		*e = AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountType: %v", v)
	}
}

type APIKeyDetails struct {
	AccountType               AccountType `json:"account_type"`
	GenerationAccessUnlimited *bool       `json:"generation_access_unlimited,omitempty"`
	WorkspaceID               string      `json:"workspace_id"`
}

func (o *APIKeyDetails) GetAccountType() AccountType {
	if o == nil {
		return AccountType("")
	}
	return o.AccountType
}

func (o *APIKeyDetails) GetGenerationAccessUnlimited() *bool {
	if o == nil {
		return nil
	}
	return o.GenerationAccessUnlimited
}

func (o *APIKeyDetails) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
