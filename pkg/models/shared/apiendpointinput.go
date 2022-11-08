package shared

// APIEndpointInput
// An ApiEndpoint is a description of an Endpoint for an API.
type APIEndpointInput struct {
	APIEndpointID string `json:"api_endpoint_id"`
	Description   string `json:"description"`
	DisplayName   string `json:"display_name"`
	Method        string `json:"method"`
	Path          string `json:"path"`
	VersionID     string `json:"version_id"`
}
