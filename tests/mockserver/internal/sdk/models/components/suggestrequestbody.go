// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type SuggestRequestBodySuggestionType string

const (
	SuggestRequestBodySuggestionTypeMethodNames     SuggestRequestBodySuggestionType = "method-names"
	SuggestRequestBodySuggestionTypeDiagnosticsOnly SuggestRequestBodySuggestionType = "diagnostics-only"
)

func (e SuggestRequestBodySuggestionType) ToPointer() *SuggestRequestBodySuggestionType {
	return &e
}
func (e *SuggestRequestBodySuggestionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "method-names":
		fallthrough
	case "diagnostics-only":
		*e = SuggestRequestBodySuggestionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SuggestRequestBodySuggestionType: %v", v)
	}
}

type SuggestRequestBody struct {
	Diagnostics    []Diagnostic                     `json:"diagnostics"`
	OasSummary     OASSummary                       `json:"oas_summary"`
	SuggestionType SuggestRequestBodySuggestionType `json:"suggestion_type"`
}

func (o *SuggestRequestBody) GetDiagnostics() []Diagnostic {
	if o == nil {
		return []Diagnostic{}
	}
	return o.Diagnostics
}

func (o *SuggestRequestBody) GetOasSummary() OASSummary {
	if o == nil {
		return OASSummary{}
	}
	return o.OasSummary
}

func (o *SuggestRequestBody) GetSuggestionType() SuggestRequestBodySuggestionType {
	if o == nil {
		return SuggestRequestBodySuggestionType("")
	}
	return o.SuggestionType
}
