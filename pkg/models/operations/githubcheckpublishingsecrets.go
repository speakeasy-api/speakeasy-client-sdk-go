// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GithubCheckPublishingSecretsRequest struct {
	GenerateGenLockID string `queryParam:"style=form,explode=true,name=generate_gen_lock_id"`
}

func (o *GithubCheckPublishingSecretsRequest) GetGenerateGenLockID() string {
	if o == nil {
		return ""
	}
	return o.GenerateGenLockID
}

type GithubCheckPublishingSecretsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	GithubMissingPublishingSecretsResponse *shared.GithubMissingPublishingSecretsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GithubCheckPublishingSecretsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GithubCheckPublishingSecretsResponse) GetGithubMissingPublishingSecretsResponse() *shared.GithubMissingPublishingSecretsResponse {
	if o == nil {
		return nil
	}
	return o.GithubMissingPublishingSecretsResponse
}

func (o *GithubCheckPublishingSecretsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GithubCheckPublishingSecretsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
