// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PreflightRequest struct {
	NamespaceName string `json:"namespace_name"`
}

func (o *PreflightRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}
