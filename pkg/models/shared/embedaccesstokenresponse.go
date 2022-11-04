package shared

// EmbedAccessTokenResponse
// An EmbedAccessTokenResponse contains a token that can be used to embed a Speakeasy dashboard.
type EmbedAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}
