<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"log"
)

func main() {
	ctx := context.Background()

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.CreatePublishingToken(ctx, shared.PublishingToken{
		CreatedAt:      types.MustTimeFromString("2025-10-25T02:17:15.413Z"),
		CreatedBy:      "<value>",
		ID:             "<id>",
		OrganizationID: "<id>",
		TargetID:       "<id>",
		TargetResource: shared.TargetResourceDocument,
		Token:          "<value>",
		WorkspaceID:    "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PublishingToken != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->