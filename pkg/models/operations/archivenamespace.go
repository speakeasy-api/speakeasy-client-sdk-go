// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"net/http"
)

// ArchiveNamespaceRequestBody - Archived status
type ArchiveNamespaceRequestBody struct {
	Archived *bool `default:"true" json:"archived"`
}

func (a ArchiveNamespaceRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ArchiveNamespaceRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ArchiveNamespaceRequestBody) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

type ArchiveNamespaceRequest struct {
	// Archived status
	RequestBody   *ArchiveNamespaceRequestBody `request:"mediaType=application/json"`
	NamespaceName string                       `pathParam:"style=simple,explode=false,name=namespace_name"`
}

func (o *ArchiveNamespaceRequest) GetRequestBody() *ArchiveNamespaceRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *ArchiveNamespaceRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

type ArchiveNamespaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ArchiveNamespaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ArchiveNamespaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ArchiveNamespaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
