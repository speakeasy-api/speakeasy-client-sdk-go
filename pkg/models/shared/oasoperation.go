// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type OASOperation struct {
	Description string   `json:"description"`
	Method      string   `json:"method"`
	OperationID string   `json:"operation_id"`
	Path        string   `json:"path"`
	Tags        []string `json:"tags"`
}

func (o *OASOperation) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *OASOperation) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *OASOperation) GetOperationID() string {
	if o == nil {
		return ""
	}
	return o.OperationID
}

func (o *OASOperation) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *OASOperation) GetTags() []string {
	if o == nil {
		return []string{}
	}
	return o.Tags
}
