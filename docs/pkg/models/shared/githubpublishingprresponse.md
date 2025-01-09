# GithubPublishingPRResponse

Open generation PRs pending publishing


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `PendingVersion`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `PullRequest`                                                                    | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `PullRequestMetadata`                                                            | [*shared.PullRequestMetadata](../../../pkg/models/shared/pullrequestmetadata.md) | :heavy_minus_sign:                                                               | This can only be populated when the github app is installed for a repo           |