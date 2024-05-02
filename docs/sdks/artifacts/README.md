# Artifacts
(*Artifacts*)

### Available Operations

* [GetBlob](#getblob) - Get blob for a particular digest
* [GetManifest](#getmanifest) - Get manifest for a particular reference
* [GetNamespaces](#getnamespaces) - Each namespace contains many revisions.
* [GetRevisions](#getrevisions)
* [GetTags](#gettags)
* [PostTags](#posttags) - Add tags to an existing revision
* [Preflight](#preflight) - Get access token for communicating with OCI distribution endpoints

## GetBlob

Get blob for a particular digest

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Artifacts.GetBlob(ctx, operations.GetBlobRequest{
        Digest: "<value>",
        NamespaceName: "<value>",
        OrganizationSlug: "<value>",
        WorkspaceSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Blob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetBlobRequest](../../pkg/models/operations/getblobrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetBlobResponse](../../pkg/models/operations/getblobresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetManifest

Get manifest for a particular reference

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Artifacts.GetManifest(ctx, operations.GetManifestRequest{
        NamespaceName: "<value>",
        OrganizationSlug: "<value>",
        RevisionReference: "<value>",
        WorkspaceSlug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Manifest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetManifestRequest](../../pkg/models/operations/getmanifestrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetManifestResponse](../../pkg/models/operations/getmanifestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetNamespaces

Each namespace contains many revisions.

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
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Artifacts.GetNamespaces(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetNamespacesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetNamespacesResponse](../../pkg/models/operations/getnamespacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRevisions

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Artifacts.GetRevisions(ctx, operations.GetRevisionsRequest{
        NamespaceName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRevisionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetRevisionsRequest](../../pkg/models/operations/getrevisionsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetRevisionsResponse](../../pkg/models/operations/getrevisionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTags

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Artifacts.GetTags(ctx, operations.GetTagsRequest{
        NamespaceName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTagsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetTagsRequest](../../pkg/models/operations/gettagsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetTagsResponse](../../pkg/models/operations/gettagsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostTags

Add tags to an existing revision

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Artifacts.PostTags(ctx, operations.PostTagsRequest{
        NamespaceName: "<value>",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.PostTagsRequest](../../pkg/models/operations/posttagsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.PostTagsResponse](../../pkg/models/operations/posttagsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Preflight

Get access token for communicating with OCI distribution endpoints

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
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Artifacts.Preflight(ctx, &shared.PreflightRequest{
        NamespaceName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PreflightToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.PreflightRequest](../../pkg/models/shared/preflightrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.PreflightResponse](../../pkg/models/operations/preflightresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
