// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"mockserver/internal/sdk/models/components"
)

type GetCodeSamplePreviewAsyncRequest struct {
	// The ID of the job to check the status and retrieve results
	JobID string `pathParam:"style=simple,explode=false,name=jobID"`
}

func (o *GetCodeSamplePreviewAsyncRequest) GetJobID() string {
	if o == nil {
		return ""
	}
	return o.JobID
}

// GetCodeSamplePreviewAsyncResponseBody - Job is still in progress
type GetCodeSamplePreviewAsyncResponseBody struct {
	// The current status of the job. Possible values are `pending` or `running`.
	Status components.CodeSamplesJobStatus `json:"status"`
}

func (o *GetCodeSamplePreviewAsyncResponseBody) GetStatus() components.CodeSamplesJobStatus {
	if o == nil {
		return components.CodeSamplesJobStatus("")
	}
	return o.Status
}

type GetCodeSamplePreviewAsyncResponse struct {
	// Successfully returned codeSample overlay file
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	TwoHundredApplicationJSONResponseStream io.ReadCloser
	// Successfully returned codeSample overlay file
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	TwoHundredApplicationXYamlResponseStream io.ReadCloser
	// Job is still in progress
	TwoHundredAndTwoApplicationJSONObject *GetCodeSamplePreviewAsyncResponseBody
	HTTPMeta                              components.HTTPMetadata `json:"-"`
}

func (o *GetCodeSamplePreviewAsyncResponse) GetTwoHundredApplicationJSONResponseStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONResponseStream
}

func (o *GetCodeSamplePreviewAsyncResponse) GetTwoHundredApplicationXYamlResponseStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationXYamlResponseStream
}

func (o *GetCodeSamplePreviewAsyncResponse) GetTwoHundredAndTwoApplicationJSONObject() *GetCodeSamplePreviewAsyncResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredAndTwoApplicationJSONObject
}

func (o *GetCodeSamplePreviewAsyncResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
