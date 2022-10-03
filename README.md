# speakeasy-client-sdk-go

## Example usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdk "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

func main() {
    ctx := context.Background()

	sdk := sdk.New(
		sdk.WithSecurity(shared.Security{
			APIKey: shared.APIKey{
				APIKey: "YOUR_API_KEY", // Replace with your API key from your Speakeasy Workspace
			},
		}),
	)

	getApisRes, err := sdk.GetApisV1(ctx, operations.GetApisV1Request{
		QueryParams: operations.GetApisV1QueryParams{
			Metadata: map[string][]string{
				"label": {"1"},
			},
			Op: &operations.GetApisV1Op{
				And: true,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if getApisRes.StatusCode != 200 {
		fmt.Println(getApisRes.Responses[getApisRes.StatusCode][getApisRes.ContentType].Error)
	} else {
		fmt.Println(getApisRes.Responses[getApisRes.StatusCode][getApisRes.ContentType].Apis)
	}
}
```
