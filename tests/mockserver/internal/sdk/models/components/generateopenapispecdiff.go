// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GenerateOpenAPISpecDiff struct {
	CurrentSchema string `json:"current_schema"`
	NewSchema     string `json:"new_schema"`
}

func (o *GenerateOpenAPISpecDiff) GetCurrentSchema() string {
	if o == nil {
		return ""
	}
	return o.CurrentSchema
}

func (o *GenerateOpenAPISpecDiff) GetNewSchema() string {
	if o == nil {
		return ""
	}
	return o.NewSchema
}
