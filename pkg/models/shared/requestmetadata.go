// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// RequestMetadata - Key-Value pairs associated with a request
type RequestMetadata struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

func (o *RequestMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *RequestMetadata) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
