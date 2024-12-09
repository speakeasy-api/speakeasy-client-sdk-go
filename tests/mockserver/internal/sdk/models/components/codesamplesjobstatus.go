// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// CodeSamplesJobStatus - The current status of the job. Possible values are `pending` or `running`.
type CodeSamplesJobStatus string

const (
	CodeSamplesJobStatusPending CodeSamplesJobStatus = "pending"
	CodeSamplesJobStatusRunning CodeSamplesJobStatus = "running"
)

func (e CodeSamplesJobStatus) ToPointer() *CodeSamplesJobStatus {
	return &e
}
func (e *CodeSamplesJobStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "running":
		*e = CodeSamplesJobStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CodeSamplesJobStatus: %v", v)
	}
}
