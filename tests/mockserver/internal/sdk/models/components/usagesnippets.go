// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UsageSnippets struct {
	Snippets []UsageSnippet `json:"snippets"`
}

func (o *UsageSnippets) GetSnippets() []UsageSnippet {
	if o == nil {
		return []UsageSnippet{}
	}
	return o.Snippets
}
