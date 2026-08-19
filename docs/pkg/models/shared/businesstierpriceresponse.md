# BusinessTierPriceResponse


## Fields

| Field                              | Type                               | Required                           | Description                        |
| ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- |
| `Currency`                         | `string`                           | :heavy_check_mark:                 | The currency code (e.g., usd)      |
| `Interval`                         | `string`                           | :heavy_check_mark:                 | The billing interval (e.g., month) |
| `PriceID`                          | `string`                           | :heavy_check_mark:                 | The Stripe price ID                |
| `ProductName`                      | `string`                           | :heavy_check_mark:                 | The product name from Stripe       |
| `UnitAmount`                       | `int64`                            | :heavy_check_mark:                 | The price amount in cents          |