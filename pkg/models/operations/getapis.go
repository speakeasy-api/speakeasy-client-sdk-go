package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetApisOp struct {
	And bool `queryParam:"name=and"`
}

type GetApisQueryParams struct {
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	Op       *GetApisOp          `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetApisRequest struct {
	Retries     *utils.RetryConfig
	QueryParams GetApisQueryParams
}

type GetApisResponse struct {
	Apis        []shared.API
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
