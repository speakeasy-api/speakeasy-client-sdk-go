// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RemoteDocument - A document hosted in the registry
type RemoteDocument struct {
	RegistryURL string `json:"registry_url"`
}

func (o *RemoteDocument) GetRegistryURL() string {
	if o == nil {
		return ""
	}
	return o.RegistryURL
}
