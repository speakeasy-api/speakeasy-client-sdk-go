// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetGitHubActionRequest struct {
	Org  string `queryParam:"style=form,explode=true,name=org"`
	Repo string `queryParam:"style=form,explode=true,name=repo"`
	// The targetName of the workflow target.
	TargetName *string `queryParam:"style=form,explode=true,name=targetName"`
}

func (o *GetGitHubActionRequest) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *GetGitHubActionRequest) GetRepo() string {
	if o == nil {
		return ""
	}
	return o.Repo
}

func (o *GetGitHubActionRequest) GetTargetName() *string {
	if o == nil {
		return nil
	}
	return o.TargetName
}

type GetGitHubActionResponse struct {
	// OK
	GithubGetActionResponse *components.GithubGetActionResponse
	HTTPMeta                components.HTTPMetadata `json:"-"`
}

func (o *GetGitHubActionResponse) GetGithubGetActionResponse() *components.GithubGetActionResponse {
	if o == nil {
		return nil
	}
	return o.GithubGetActionResponse
}

func (o *GetGitHubActionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
