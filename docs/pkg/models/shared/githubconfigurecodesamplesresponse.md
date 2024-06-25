# GithubConfigureCodeSamplesResponse

A response to configure GitHub code samples


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `CodeSampleOverlayRegistryURL`                                            | *string*                                                                  | :heavy_check_mark:                                                        | The URL of the code sample overlay registry                               |
| `GhActionID`                                                              | **string*                                                                 | :heavy_minus_sign:                                                        | The ID of the GitHub action that was dispatched                           |
| `Source`                                                                  | [shared.WorkflowDocument](../../../pkg/models/shared/workflowdocument.md) | :heavy_check_mark:                                                        | A document referenced by a workflow                                       |