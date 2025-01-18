// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type MethodPaths struct {
	Method components.HTTPMethod `queryParam:"name=method"`
	Path   string                `queryParam:"name=path"`
}

func (o *MethodPaths) GetMethod() components.HTTPMethod {
	if o == nil {
		return components.HTTPMethod("")
	}
	return o.Method
}

func (o *MethodPaths) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type GetCodeSamplesRequest struct {
	// The languages to retrieve snippets for.
	Languages []string `queryParam:"style=form,explode=true,name=languages"`
	// The method paths to retrieve snippets for.
	MethodPaths []MethodPaths `queryParam:"style=form,explode=true,name=method_paths"`
	// The operation IDs to retrieve snippets for.
	OperationIds []string `queryParam:"style=form,explode=true,name=operation_ids"`
	// The registry URL from which to retrieve the snippets.
	RegistryURL string `queryParam:"style=form,explode=true,name=registry_url"`
}

func (o *GetCodeSamplesRequest) GetLanguages() []string {
	if o == nil {
		return nil
	}
	return o.Languages
}

func (o *GetCodeSamplesRequest) GetMethodPaths() []MethodPaths {
	if o == nil {
		return nil
	}
	return o.MethodPaths
}

func (o *GetCodeSamplesRequest) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *GetCodeSamplesRequest) GetRegistryURL() string {
	if o == nil {
		return ""
	}
	return o.RegistryURL
}

type GetCodeSamplesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	UsageSnippets *components.UsageSnippets
}

func (o *GetCodeSamplesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCodeSamplesResponse) GetUsageSnippets() *components.UsageSnippets {
	if o == nil {
		return nil
	}
	return o.UsageSnippets
}
