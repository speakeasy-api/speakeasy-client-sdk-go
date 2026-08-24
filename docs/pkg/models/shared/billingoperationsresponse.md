# BillingOperationsResponse

Contains the billing operations breakdown for an organization


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Languages`                                                                                 | [][shared.LanguageBillingBreakdown](../../../pkg/models/shared/languagebillingbreakdown.md) | :heavy_check_mark:                                                                          | Billing breakdown for each language                                                         |
| `TotalBillableUnits`                                                                        | `int64`                                                                                     | :heavy_check_mark:                                                                          | Total billable units across all generated targets                                           |
| `TotalUniqueOperations`                                                                     | `int64`                                                                                     | :heavy_check_mark:                                                                          | Total count of unique operations across all languages                                       |