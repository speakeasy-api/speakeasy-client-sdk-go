// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UsageSnippet struct {
	// The code snippet
	Code string `json:"code"`
	// The language of the snippet
	Language string `json:"language"`
	// The HTTP method of the operation
	Method any `json:"method"`
	// The operation ID for the snippet
	OperationID string `json:"operationId"`
	// The path of the operation
	Path string `json:"path"`
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

func (o *UsageSnippet) GetMethod() any {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *UsageSnippet) GetOperationID() string {
	if o == nil {
		return ""
	}
	return o.OperationID
}

func (o *UsageSnippet) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}
