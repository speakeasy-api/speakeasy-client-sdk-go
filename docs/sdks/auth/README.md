# Auth
(*Auth*)

## Overview

REST APIs for managing Authentication

### Available Operations

* [ValidateAPIKey](#validateapikey) - Validate the current api key.

## ValidateAPIKey

Validate the current api key.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("string")),
    )

    ctx := context.Background()
    res, err := s.Auth.ValidateAPIKey(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIKeyDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ValidateAPIKeyResponse](../../pkg/models/operations/validateapikeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
