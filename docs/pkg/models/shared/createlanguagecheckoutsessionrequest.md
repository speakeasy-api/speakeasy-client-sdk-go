# CreateLanguageCheckoutSessionRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `BillingInterval`                                                        | [*shared.BillingInterval](../../../pkg/models/shared/billinginterval.md) | :heavy_minus_sign:                                                       | Billing interval (defaults to month)                                     |
| `CancelURL`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | URL to redirect to if checkout is canceled                               |
| `Languages`                                                              | []`string`                                                               | :heavy_check_mark:                                                       | List of languages to activate                                            |
| `SuccessURL`                                                             | `string`                                                                 | :heavy_check_mark:                                                       | URL to redirect to on successful checkout                                |