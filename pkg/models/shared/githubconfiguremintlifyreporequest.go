// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GithubConfigureMintlifyRepoRequest - A request to configure a GitHub repository for mintlify
type GithubConfigureMintlifyRepoRequest struct {
	// The input OpenAPI document
	Input string `json:"input"`
	// The GitHub organization name
	Org string `json:"org"`
	// The overlays to apply
	Overlays []string `json:"overlays"`
	// The GitHub repository name
	Repo string `json:"repo"`
}

func (o *GithubConfigureMintlifyRepoRequest) GetInput() string {
	if o == nil {
		return ""
	}
	return o.Input
}

func (o *GithubConfigureMintlifyRepoRequest) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *GithubConfigureMintlifyRepoRequest) GetOverlays() []string {
	if o == nil {
		return []string{}
	}
	return o.Overlays
}

func (o *GithubConfigureMintlifyRepoRequest) GetRepo() string {
	if o == nil {
		return ""
	}
	return o.Repo
}
