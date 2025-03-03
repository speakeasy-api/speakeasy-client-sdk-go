// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GithubCheckPublishingPRsRequest struct {
	GenerateGenLockID string `queryParam:"style=form,explode=true,name=generate_gen_lock_id"`
	Org               string `queryParam:"style=form,explode=true,name=org"`
	Repo              string `queryParam:"style=form,explode=true,name=repo"`
}

func (o *GithubCheckPublishingPRsRequest) GetGenerateGenLockID() string {
	if o == nil {
		return ""
	}
	return o.GenerateGenLockID
}

func (o *GithubCheckPublishingPRsRequest) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *GithubCheckPublishingPRsRequest) GetRepo() string {
	if o == nil {
		return ""
	}
	return o.Repo
}

type GithubCheckPublishingPRsResponse struct {
	// OK
	GithubPublishingPRResponse *components.GithubPublishingPRResponse
	HTTPMeta                   components.HTTPMetadata `json:"-"`
}

func (o *GithubCheckPublishingPRsResponse) GetGithubPublishingPRResponse() *components.GithubPublishingPRResponse {
	if o == nil {
		return nil
	}
	return o.GithubPublishingPRResponse
}

func (o *GithubCheckPublishingPRsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
