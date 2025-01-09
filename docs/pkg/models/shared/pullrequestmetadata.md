# PullRequestMetadata

This can only be populated when the github app is installed for a repo


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `BaseBranch`                               | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `CanMerge`                                 | **bool*                                    | :heavy_minus_sign:                         | N/A                                        |
| `CreatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Description`                              | **string*                                  | :heavy_minus_sign:                         | truncated to first 1000 characters         |
| `HeadBranch`                               | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `Labels`                                   | []*string*                                 | :heavy_minus_sign:                         | List of github labels                      |
| `RequestedReviewers`                       | []*string*                                 | :heavy_minus_sign:                         | List of github handles                     |
| `Status`                                   | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `Title`                                    | **string*                                  | :heavy_minus_sign:                         | N/A                                        |