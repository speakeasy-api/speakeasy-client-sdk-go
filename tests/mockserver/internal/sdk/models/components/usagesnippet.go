// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UsageSnippet struct {
	// The code snippet
	Code string `json:"code"`
	// The language of the snippet
	Language string `json:"language"`
	// The operation ID for the snippet
	OperationID string `json:"operationId"`
}

func (o *UsageSnippet) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *UsageSnippet) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

func (o *UsageSnippet) GetOperationID() string {
	if o == nil {
		return ""
	}
	return o.OperationID
}
