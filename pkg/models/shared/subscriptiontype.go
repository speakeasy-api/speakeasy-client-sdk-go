// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SubscriptionType string

const (
	SubscriptionTypeCli SubscriptionType = "cli"
)

func (e SubscriptionType) ToPointer() *SubscriptionType {
	return &e
}
func (e *SubscriptionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cli":
		*e = SubscriptionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SubscriptionType: %v", v)
	}
}
