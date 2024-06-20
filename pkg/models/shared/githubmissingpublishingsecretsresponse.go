// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GithubMissingPublishingSecretsResponse - A valid response containing MISSING publishing secret keys for a github target
type GithubMissingPublishingSecretsResponse struct {
	MissingSecrets []string `json:"missing_secrets,omitempty"`
}

func (o *GithubMissingPublishingSecretsResponse) GetMissingSecrets() []string {
	if o == nil {
		return nil
	}
	return o.MissingSecrets
}
