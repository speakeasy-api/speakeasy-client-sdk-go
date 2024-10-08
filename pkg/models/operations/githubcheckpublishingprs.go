// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
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
	// HTTP response content type for this operation
	ContentType string
	// OK
	GithubPublishingPRResponse *shared.GithubPublishingPRResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GithubCheckPublishingPRsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GithubCheckPublishingPRsResponse) GetGithubPublishingPRResponse() *shared.GithubPublishingPRResponse {
	if o == nil {
		return nil
	}
	return o.GithubPublishingPRResponse
}

func (o *GithubCheckPublishingPRsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GithubCheckPublishingPRsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
