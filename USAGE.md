<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.GenerateCodeSamplePreview(ctx, shared.CodeSampleSchemaInput{
		Languages: []string{
			"<value>",
			"<value>",
		},
		SchemaFile: shared.SchemaFile{
			Content:  content,
			FileName: "example.file",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.TwoHundredApplicationJSONResponseStream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->