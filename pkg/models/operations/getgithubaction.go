// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
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
	// HTTP response content type for this operation
	ContentType string
	// OK
	GithubGetActionResponse *shared.GithubGetActionResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetGitHubActionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGitHubActionResponse) GetGithubGetActionResponse() *shared.GithubGetActionResponse {
	if o == nil {
		return nil
	}
	return o.GithubGetActionResponse
}

func (o *GetGitHubActionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGitHubActionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}