// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GithubConfigureMintlifyRepoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GithubConfigureMintlifyRepoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
