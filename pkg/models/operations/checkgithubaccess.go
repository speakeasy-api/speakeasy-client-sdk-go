// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CheckGithubAccessRequest struct {
	Org  string `queryParam:"style=form,explode=true,name=org"`
	Repo string `queryParam:"style=form,explode=true,name=repo"`
}

func (o *CheckGithubAccessRequest) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *CheckGithubAccessRequest) GetRepo() string {
	if o == nil {
		return ""
	}
	return o.Repo
}

type CheckGithubAccessResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CheckGithubAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckGithubAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckGithubAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}