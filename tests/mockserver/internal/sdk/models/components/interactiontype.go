// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// InteractionType - Type of interaction.
type InteractionType string

const (
	InteractionTypeCiExec         InteractionType = "CI_EXEC"
	InteractionTypeCliExec        InteractionType = "CLI_EXEC"
	InteractionTypeLint           InteractionType = "LINT"
	InteractionTypeOpenapiDiff    InteractionType = "OPENAPI_DIFF"
	InteractionTypeTargetGenerate InteractionType = "TARGET_GENERATE"
	InteractionTypeTombstone      InteractionType = "TOMBSTONE"
	InteractionTypeAuthenticate   InteractionType = "AUTHENTICATE"
	InteractionTypeQuickstart     InteractionType = "QUICKSTART"
	InteractionTypeRun            InteractionType = "RUN"
	InteractionTypeConfigure      InteractionType = "CONFIGURE"
	InteractionTypePublish        InteractionType = "PUBLISH"
	InteractionTypeTest           InteractionType = "TEST"
)

func (e InteractionType) ToPointer() *InteractionType {
	return &e
}
func (e *InteractionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CI_EXEC":
		fallthrough
	case "CLI_EXEC":
		fallthrough
	case "LINT":
		fallthrough
	case "OPENAPI_DIFF":
		fallthrough
	case "TARGET_GENERATE":
		fallthrough
	case "TOMBSTONE":
		fallthrough
	case "AUTHENTICATE":
		fallthrough
	case "QUICKSTART":
		fallthrough
	case "RUN":
		fallthrough
	case "CONFIGURE":
		fallthrough
	case "PUBLISH":
		fallthrough
	case "TEST":
		*e = InteractionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InteractionType: %v", v)
	}
}
