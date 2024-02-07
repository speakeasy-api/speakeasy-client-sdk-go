<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
		speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("string")),
	)

	ctx := context.Background()
	res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
		Metadata: map[string][]string{
			"key": []string{
				"string",
			},
		},
		Op: &operations.QueryParamOp{
			And: false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Apis != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->