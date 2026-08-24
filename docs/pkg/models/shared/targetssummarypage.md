# TargetsSummaryPage

Paginated response containing a list of target summaries


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `HasMore`                                                                   | `bool`                                                                      | :heavy_check_mark:                                                          | Whether there are more results available                                    |
| `NextCursor`                                                                | `*string`                                                                   | :heavy_minus_sign:                                                          | Opaque cursor for the next page. Null if no more pages.                     |
| `Targets`                                                                   | [][shared.TargetSDKSummary](../../../pkg/models/shared/targetsdksummary.md) | :heavy_check_mark:                                                          | List of target summaries for the current page                               |