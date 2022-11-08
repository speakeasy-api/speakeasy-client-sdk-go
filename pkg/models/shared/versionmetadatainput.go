package shared

// VersionMetadataInput
// A set of keys and associated values, attached to a particular version of an Api.
type VersionMetadataInput struct {
	MetaKey   string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
}
