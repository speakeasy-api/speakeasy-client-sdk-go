# speakeasy-client-sdk-go

This is the Speakeasy API Client SDK for Go. It is generated from our OpenAPI spec found at https://docs.speakeasyapi.dev/openapi.yaml and used for interacting with the [Speakeasy API](https://docs.speakeasyapi.dev/docs/speakeasy-api/speakeasy-api).

This SDK was generated using Speakeasy's SDK Generator. For more information on how to use the generator to generate your own SDKs, please see the [Speakeasy Client SDK Generator Docs](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks).

## Installation

```bash
go get github.com/speakeasy-api/speakeasy-client-sdk-go
```

## Example usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
	ctx := context.Background()

	opts := []sdk.SDKOption{
		sdk.WithSecurity(shared.Security{
			APIKey: shared.SchemeAPIKey{
				APIKey: "YOUR_API_KEY", // Replace with your API key from your Speakeasy Workspace
			},
		}),
	}

	s := sdk.New(opts...)

	res, err := s.GetApis(ctx, operations.GetApisRequest{
		QueryParams: operations.GetApisQueryParams{
			Metadata: map[string][]string{
				"label": {"1"},
			},
			Op: &operations.GetApisOp{
				And: true,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Apis)
	}
}
```
