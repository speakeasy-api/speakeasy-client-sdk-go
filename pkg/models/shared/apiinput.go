package shared

// APIInput
// An Api is representation of a API (a collection of API Endpoints) within the Speakeasy Platform.
type APIInput struct {
	APIID       string              `json:"api_id"`
	Description string              `json:"description"`
	MetaData    map[string][]string `json:"meta_data,omitempty"`
	VersionID   string              `json:"version_id"`
}
