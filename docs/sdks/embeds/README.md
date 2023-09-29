# Embeds
(*Embeds*)

## Overview

REST APIs for managing embeds

### Available Operations

* [GetEmbedAccessToken](#getembedaccesstoken) - Get an embed access token for the current workspace.
* [GetValidEmbedAccessTokens](#getvalidembedaccesstokens) - Get all valid embed access tokens for the current workspace.
* [RevokeEmbedAccessToken](#revokeembedaccesstoken) - Revoke an embed access EmbedToken.

## GetEmbedAccessToken

Returns an embed access token for the current workspace. This can be used to authenticate access to externally embedded content.
Filters can be applied allowing views to be filtered to things like particular customerIds.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Embeds.GetEmbedAccessToken(ctx, operations.GetEmbedAccessTokenRequest{
        Description: speakeasyclientsdkgo.String("Versatile asynchronous leverage"),
        Duration: speakeasyclientsdkgo.Int64(554373),
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "<key>",
                    Operator: "Account",
                    Value: "unless to frame",
                },
            },
            Limit: 792058,
            Offset: 425010,
            Operator: "kilogram innovative Virginia",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EmbedAccessTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetEmbedAccessTokenRequest](../../models/operations/getembedaccesstokenrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetEmbedAccessTokenResponse](../../models/operations/getembedaccesstokenresponse.md), error**


## GetValidEmbedAccessTokens

Get all valid embed access tokens for the current workspace.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Embeds.GetValidEmbedAccessTokens(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EmbedTokens != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetValidEmbedAccessTokensResponse](../../models/operations/getvalidembedaccesstokensresponse.md), error**


## RevokeEmbedAccessToken

Revoke an embed access EmbedToken.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Embeds.RevokeEmbedAccessToken(ctx, operations.RevokeEmbedAccessTokenRequest{
        TokenID: "Small",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RevokeEmbedAccessTokenRequest](../../models/operations/revokeembedaccesstokenrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RevokeEmbedAccessTokenResponse](../../models/operations/revokeembedaccesstokenresponse.md), error**

