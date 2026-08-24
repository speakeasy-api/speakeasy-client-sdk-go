<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
	ctx := context.Background()

	s := speakeasyclientsdkgo.New()

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