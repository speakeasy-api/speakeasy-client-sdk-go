// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetNamespacesResponse struct {
	Items []Namespace `json:"items"`
}

func (o *GetNamespacesResponse) GetItems() []Namespace {
	if o == nil {
		return []Namespace{}
	}
	return o.Items
}
