// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

// FeatureFlag - A feature flag is a key-value pair that can be used to enable or disable features.
type FeatureFlag struct {
	FeatureFlag string     `json:"feature_flag"`
	TrialEndsAt *time.Time `json:"trial_ends_at,omitempty"`
}

func (f FeatureFlag) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FeatureFlag) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FeatureFlag) GetFeatureFlag() string {
	if o == nil {
		return ""
	}
	return o.FeatureFlag
}

func (o *FeatureFlag) GetTrialEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TrialEndsAt
}
