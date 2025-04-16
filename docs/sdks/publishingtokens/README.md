# PublishingTokens
(*PublishingTokens*)

## Overview

### Available Operations

* [Create](#create) - Create a publishing token for a workspace
* [Delete](#delete) - Delete a specific publishing token
* [Get](#get) - Get a specific publishing token
* [List](#list) - Get publishing tokens for a workspace
* [ResolveMetadata](#resolvemetadata) - Get metadata about the token
* [ResolveTarget](#resolvetarget) - Get a specific publishing token target
* [Update](#update) - Updates the validitity period of a publishing token

## Create

Creates a publishing token for the current workspace

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.PublishingTokens.Create(ctx, &operations.CreatePublishingTokenRequestBody{
        TargetID: "<id>",
        TargetResource: "document",
        TokenName: "<value>",
        ValidUntil: types.MustTimeFromString("2024-06-17T07:14:55.338Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PublishingToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreatePublishingTokenRequestBody](../../pkg/models/operations/createpublishingtokenrequestbody.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreatePublishingTokenResponse](../../pkg/models/operations/createpublishingtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## Delete

Delete a particular publishing token.

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.PublishingTokens.Delete(ctx, operations.DeletePublishingTokenRequest{
        TokenID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeletePublishingTokenRequest](../../pkg/models/operations/deletepublishingtokenrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.DeletePublishingTokenResponse](../../pkg/models/operations/deletepublishingtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## Get

Get information about a particular publishing token.

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.PublishingTokens.Get(ctx, operations.GetPublishingTokenByIDRequest{
        TokenID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PublishingToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetPublishingTokenByIDRequest](../../pkg/models/operations/getpublishingtokenbyidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetPublishingTokenByIDResponse](../../pkg/models/operations/getpublishingtokenbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## List

Returns a publishing token for the current workspace

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.PublishingTokens.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.GetPublishingTokenResponse](../../pkg/models/operations/getpublishingtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## ResolveMetadata

Get information about a particular publishing token.

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.PublishingTokens.ResolveMetadata(ctx, operations.GetPublishingTokenPublicMetadataRequest{
        TokenID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetPublishingTokenPublicMetadataRequest](../../pkg/models/operations/getpublishingtokenpublicmetadatarequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetPublishingTokenPublicMetadataResponse](../../pkg/models/operations/getpublishingtokenpublicmetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## ResolveTarget

Get information about a particular publishing token target.

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.PublishingTokens.ResolveTarget(ctx, operations.GetPublishingTokenTargetByIDRequest{
        TokenID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetPublishingTokenTargetByIDRequest](../../pkg/models/operations/getpublishingtokentargetbyidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.GetPublishingTokenTargetByIDResponse](../../pkg/models/operations/getpublishingtokentargetbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## Update

Updates the validity period of a particular publishing token.

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.PublishingTokens.Update(ctx, operations.UpdatePublishingTokenExpirationRequest{
        TokenID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.UpdatePublishingTokenExpirationRequest](../../pkg/models/operations/updatepublishingtokenexpirationrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpdatePublishingTokenExpirationResponse](../../pkg/models/operations/updatepublishingtokenexpirationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |