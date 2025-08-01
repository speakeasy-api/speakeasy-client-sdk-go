# Organizations
(*Organizations*)

## Overview

REST APIs for managing Organizations (speakeasy L1 Tenancy construct)

### Available Operations

* [Create](#create) - Create an organization
* [CreateBillingAddOns](#createbillingaddons) - Create billing add ons
* [CreateFreeTrial](#createfreetrial) - Create a free trial for an organization
* [DeleteBillingAddOn](#deletebillingaddon) - Delete billing add ons
* [Get](#get) - Get organization
* [GetAll](#getall) - Get organizations for a user
* [GetBillingAddOns](#getbillingaddons) - Get billing add ons
* [GetUsage](#getusage) - Get billing usage summary for a particular organization

## Create

Creates an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="createOrganization" method="post" path="/v1/organization" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.Create(ctx, shared.Organization{
        AccountType: shared.AccountTypeScaleUp,
        CreatedAt: types.MustTimeFromString("2025-10-26T09:05:00.560Z"),
        ID: "<id>",
        Name: "<value>",
        Slug: "<value>",
        SsoActivated: false,
        TelemetryDisabled: false,
        UpdatedAt: types.MustTimeFromString("2023-12-23T08:00:51.380Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Organization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.Organization](../../pkg/models/shared/organization.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../pkg/models/operations/option.md)   | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.CreateOrganizationResponse](../../pkg/models/operations/createorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## CreateBillingAddOns

Create billing add ons

### Example Usage

<!-- UsageSnippet language="go" operationID="createBillingAddOns" method="post" path="/v1/organization/add_ons" -->
```go
package main

import(
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

    res, err := s.Organizations.CreateBillingAddOns(ctx, shared.OrganizationBillingAddOnRequest{
        AddOns: []shared.BillingAddOn{
            shared.BillingAddOnCustomCodeRegions,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationBillingAddOnResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.OrganizationBillingAddOnRequest](../../pkg/models/shared/organizationbillingaddonrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateBillingAddOnsResponse](../../pkg/models/operations/createbillingaddonsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## CreateFreeTrial

Creates a free trial for an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="createFreeTrial" method="post" path="/v1/organization/free_trial" -->
```go
package main

import(
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

    res, err := s.Organizations.CreateFreeTrial(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.CreateFreeTrialResponse](../../pkg/models/operations/createfreetrialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## DeleteBillingAddOn

Delete billing add ons

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteBillingAddOn" method="delete" path="/v1/organization/add_ons/{add_on}" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.DeleteBillingAddOn(ctx, operations.DeleteBillingAddOnRequest{
        AddOn: shared.BillingAddOnSDKTesting,
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteBillingAddOnRequest](../../pkg/models/operations/deletebillingaddonrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.DeleteBillingAddOnResponse](../../pkg/models/operations/deletebillingaddonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## Get

Get information about a particular organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrganization" method="get" path="/v1/organization/{organizationID}" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.Get(ctx, operations.GetOrganizationRequest{
        OrganizationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Organization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetOrganizationRequest](../../pkg/models/operations/getorganizationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetOrganizationResponse](../../pkg/models/operations/getorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetAll

Returns a list of organizations a user has access too

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrganizations" method="get" path="/v1/organizations" -->
```go
package main

import(
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

    res, err := s.Organizations.GetAll(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Organizations != nil {
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

**[*operations.GetOrganizationsResponse](../../pkg/models/operations/getorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetBillingAddOns

Get billing add ons

### Example Usage

<!-- UsageSnippet language="go" operationID="getBillingAddOns" method="get" path="/v1/organization/add_ons" -->
```go
package main

import(
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

    res, err := s.Organizations.GetBillingAddOns(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationBillingAddOnResponse != nil {
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

**[*operations.GetBillingAddOnsResponse](../../pkg/models/operations/getbillingaddonsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## GetUsage

Returns a billing usage summary by target languages for a particular organization

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrganizationUsage" method="get" path="/v1/organization/usage" -->
```go
package main

import(
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

    res, err := s.Organizations.GetUsage(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationUsageResponse != nil {
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

**[*operations.GetOrganizationUsageResponse](../../pkg/models/operations/getorganizationusageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |