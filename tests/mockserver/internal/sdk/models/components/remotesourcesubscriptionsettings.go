// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RemoteSourceSubscriptionSettings struct {
	BaseSpecNamespaces []string `json:"base_spec_namespaces"`
	IgnoredNamespaces  []string `json:"ignored_namespaces,omitempty"`
	OutputNamespace    string   `json:"output_namespace"`
	OverlayNamespaces  []string `json:"overlay_namespaces"`
}

func (o *RemoteSourceSubscriptionSettings) GetBaseSpecNamespaces() []string {
	if o == nil {
		return []string{}
	}
	return o.BaseSpecNamespaces
}

func (o *RemoteSourceSubscriptionSettings) GetIgnoredNamespaces() []string {
	if o == nil {
		return nil
	}
	return o.IgnoredNamespaces
}

func (o *RemoteSourceSubscriptionSettings) GetOutputNamespace() string {
	if o == nil {
		return ""
	}
	return o.OutputNamespace
}

func (o *RemoteSourceSubscriptionSettings) GetOverlayNamespaces() []string {
	if o == nil {
		return []string{}
	}
	return o.OverlayNamespaces
}
