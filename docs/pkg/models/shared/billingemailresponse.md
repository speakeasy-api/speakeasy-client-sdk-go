# BillingEmailResponse


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `BillingEmail`                                                  | `*string`                                                       | :heavy_minus_sign:                                              | The current billing email address (empty if no Stripe customer) |
| `HasStripeCustomer`                                             | `bool`                                                          | :heavy_check_mark:                                              | Whether the organization has a linked Stripe customer           |