# Webhooks

## Overview

Webhook endpoints for external service integrations

### Available Operations

* [HandleStripeWebhook](#handlestripewebhook) - Handle Stripe webhook

## HandleStripeWebhook

Receives and processes Stripe webhook events for subscription management.
This endpoint is called by Stripe and uses webhook signature verification instead of API authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleStripeWebhook" method="post" path="/v1/webhooks/stripe" -->
```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New()

    res, err := s.Webhooks.HandleStripeWebhook(ctx, operations.HandleStripeWebhookRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.HandleStripeWebhookRequestBody](../../pkg/models/operations/handlestripewebhookrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.HandleStripeWebhookResponse](../../pkg/models/operations/handlestripewebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |