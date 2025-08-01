<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := v3.New(
		v3.WithSecurity(shared.Security{
			APIKey: v3.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.CreateRemoteSource(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->