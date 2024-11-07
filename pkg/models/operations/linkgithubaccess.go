// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type LinkGithubAccessRequest struct {
	GithubOidcToken *string `queryParam:"style=form,explode=true,name=github_oidc_token"`
	GithubOrg       *string `queryParam:"style=form,explode=true,name=github_org"`
	InstallationID  *string `queryParam:"style=form,explode=true,name=installation_id"`
}

func (o *LinkGithubAccessRequest) GetGithubOidcToken() *string {
	if o == nil {
		return nil
	}
	return o.GithubOidcToken
}

func (o *LinkGithubAccessRequest) GetGithubOrg() *string {
	if o == nil {
		return nil
	}
	return o.GithubOrg
}

func (o *LinkGithubAccessRequest) GetInstallationID() *string {
	if o == nil {
		return nil
	}
	return o.InstallationID
}

type LinkGithubAccessResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *LinkGithubAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LinkGithubAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LinkGithubAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
