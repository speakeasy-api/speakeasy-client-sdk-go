# Organizations

## Overview

REST APIs for managing Organizations (speakeasy L1 Tenancy construct)

### Available Operations

* [ActivateLanguage](#activatelanguage) - Activate language
* [CancelSubscription](#cancelsubscription) - Cancel subscription
* [Create](#create) - Create an organization
* [CreateBillingAddOns](#createbillingaddons) - Create billing add ons
* [CreateFreeTrial](#createfreetrial) - Create a free trial for an organization
* [CreateLanguageCheckoutSession](#createlanguagecheckoutsession) - Create language checkout session
* [DeactivateLanguage](#deactivatelanguage) - Deactivate language
* [DeleteBillingAddOn](#deletebillingaddon) - Delete billing add ons
* [Get](#get) - Get organization
* [GetAll](#getall) - Get organizations for a user
* [GetBillingAddOns](#getbillingaddons) - Get billing add ons
* [GetBillingEmail](#getbillingemail) - Get billing email for an organization
* [GetBillingOperations](#getbillingoperations) - Get billing operations breakdown for an organization
* [GetBusinessTierPrices](#getbusinesstierprices) - Get business tier prices
* [GetLanguages](#getlanguages) - Get language billing configurations
* [GetSubscription](#getsubscription) - Get organization subscription
* [GetTrialTargets](#gettrialtargets) - Get trial targets
* [GetUsage](#getusage) - Get billing usage summary for a particular organization
* [HandleCheckoutCallback](#handlecheckoutcallback) - Checkout callback
* [RevertSubscriptionCancellation](#revertsubscriptioncancellation) - Revert subscription cancellation
* [UpsertBillingEmail](#upsertbillingemail) - Create or update billing email

## ActivateLanguage

Activates a language for billing. If the language was previously deactivated,
this will reactivate it. If the language is new, it may require checkout.

### Example Usage

<!-- UsageSnippet language="go" operationID="activateLanguage" method="post" path="/v1/organization/languages/{language}/activate" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.ActivateLanguage(ctx, operations.ActivateLanguageRequest{
        Language: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LanguageActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ActivateLanguageRequest](../../pkg/models/operations/activatelanguagerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ActivateLanguageResponse](../../pkg/models/operations/activatelanguageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## CancelSubscription

Requests cancellation of the organization's self-service business subscription.
The subscription will be cancelled at the end of the current billing period.
Specify a target to keep after downgrade to free tier when active targets remain.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelSubscription" method="post" path="/v1/organization/billing/subscription/cancel" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.CancelSubscription(ctx, shared.CancelSubscriptionRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.CancelSubscriptionRequest](../../pkg/models/shared/cancelsubscriptionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CancelSubscriptionResponse](../../pkg/models/operations/cancelsubscriptionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## Create

Creates an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="createOrganization" method="post" path="/v1/organization" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.Create(ctx, shared.Organization{
        AccountType: shared.AccountTypeBusiness,
        CreatedAt: types.MustTimeFromString("2026-10-26T09:05:00.560Z"),
        ID: "<id>",
        Name: "<value>",
        Slug: "<value>",
        SsoActivated: false,
        TelemetryDisabled: false,
        UpdatedAt: types.MustTimeFromString("2024-12-22T08:00:51.380Z"),
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
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
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
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
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

## CreateLanguageCheckoutSession

Creates a Stripe checkout session for per-language billing.
Used when upgrading to business tier with per-language billing model.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLanguageCheckoutSession" method="post" path="/v1/organization/billing/create_language_checkout_session" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.CreateLanguageCheckoutSession(ctx, shared.CreateLanguageCheckoutSessionRequest{
        CancelURL: "https://bony-toothbrush.com/",
        Languages: []string{},
        SuccessURL: "https://writhing-complication.net",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCheckoutSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [shared.CreateLanguageCheckoutSessionRequest](../../pkg/models/shared/createlanguagecheckoutsessionrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateLanguageCheckoutSessionResponse](../../pkg/models/operations/createlanguagecheckoutsessionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## DeactivateLanguage

Deactivates a language. Subject to 2-week cooldown between deactivations.

### Example Usage

<!-- UsageSnippet language="go" operationID="deactivateLanguage" method="post" path="/v1/organization/languages/{language}/deactivate" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.DeactivateLanguage(ctx, operations.DeactivateLanguageRequest{
        Language: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LanguageActivationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeactivateLanguageRequest](../../pkg/models/operations/deactivatelanguagerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.DeactivateLanguageResponse](../../pkg/models/operations/deactivatelanguageresponse.md), error**

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
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
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
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
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
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
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
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
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

## GetBillingEmail

Returns the billing email and Stripe customer status for the current organization

### Example Usage

<!-- UsageSnippet language="go" operationID="getBillingEmail" method="get" path="/v1/organization/billing_email" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.GetBillingEmail(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEmailResponse != nil {
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

**[*operations.GetBillingEmailResponse](../../pkg/models/operations/getbillingemailresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetBillingOperations

Returns a breakdown of billing operations by language and generated SDK target
for an organization. Each language row is sourced from generation events,
and target rows optionally include the source spec namespace when available.

### Example Usage

<!-- UsageSnippet language="go" operationID="getBillingOperations" method="get" path="/v1/organization/billing_operations" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.GetBillingOperations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingOperationsResponse != nil {
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

**[*operations.GetBillingOperationsResponse](../../pkg/models/operations/getbillingoperationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetBusinessTierPrices

Returns both monthly and annual business tier prices from Stripe

### Example Usage

<!-- UsageSnippet language="go" operationID="getBusinessTierPrices" method="get" path="/v1/organization/billing/business_tier_prices" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.GetBusinessTierPrices(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.BusinessTierPricesResponse != nil {
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

**[*operations.GetBusinessTierPricesResponse](../../pkg/models/operations/getbusinesstierpricesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetLanguages

Returns all language billing configurations for the organization.
Only returns languages with self-serve billing configured.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLanguages" method="get" path="/v1/organization/languages" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.GetLanguages(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLanguagesResponse != nil {
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

**[*operations.GetLanguagesResponse](../../pkg/models/operations/getlanguagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetSubscription

Returns the subscription status for the current organization

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrganizationSubscription" method="get" path="/v1/organization/billing/subscription" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.GetSubscription(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationSubscriptionResponse != nil {
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

**[*operations.GetOrganizationSubscriptionResponse](../../pkg/models/operations/getorganizationsubscriptionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetTrialTargets

Returns the list of target languages with available and used trials for the organization.
Available trials are languages that have not yet been trialed.
Used trials are languages that have already been trialed.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTrialTargets" method="get" path="/v1/organization/billing/trial_targets" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.GetTrialTargets(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.TrialTargetsResponse != nil {
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

**[*operations.GetTrialTargetsResponse](../../pkg/models/operations/gettrialtargetsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetUsage

Returns a billing usage summary by target languages for a particular organization

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrganizationUsage" method="get" path="/v1/organization/usage" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
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

## HandleCheckoutCallback

Handles the Stripe checkout success callback. This endpoint is called by Stripe
after a successful checkout, runs reconciliation (idempotent), and redirects
to the original client success URL stored in session metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleCheckoutCallback" method="get" path="/v1/billing/checkout/callback" -->
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

    res, err := s.Organizations.HandleCheckoutCallback(ctx, operations.HandleCheckoutCallbackRequest{
        SessionID: "<id>",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.HandleCheckoutCallbackRequest](../../pkg/models/operations/handlecheckoutcallbackrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.HandleCheckoutCallbackResponse](../../pkg/models/operations/handlecheckoutcallbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## RevertSubscriptionCancellation

Reverts a pending subscription cancellation.
The subscription will continue to renew automatically.

### Example Usage

<!-- UsageSnippet language="go" operationID="revertSubscriptionCancellation" method="post" path="/v1/organization/billing/subscription/cancel/revert" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.RevertSubscriptionCancellation(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelSubscriptionResponse != nil {
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

**[*operations.RevertSubscriptionCancellationResponse](../../pkg/models/operations/revertsubscriptioncancellationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## UpsertBillingEmail

Creates a Stripe customer if one does not exist, or updates the billing email for an existing Stripe customer

### Example Usage

<!-- UsageSnippet language="go" operationID="upsertBillingEmail" method="post" path="/v1/organization/billing_email" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Organizations.UpsertBillingEmail(ctx, shared.BillingEmailRequest{
        BillingEmail: "Graciela.Sauer@hotmail.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEmailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.BillingEmailRequest](../../pkg/models/shared/billingemailrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.UpsertBillingEmailResponse](../../pkg/models/operations/upsertbillingemailresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |