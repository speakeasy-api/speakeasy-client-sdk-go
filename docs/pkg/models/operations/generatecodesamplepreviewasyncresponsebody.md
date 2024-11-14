# GenerateCodeSamplePreviewAsyncResponseBody

Job accepted, returns a job ID to poll for status and result


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `JobID`                                                                           | *string*                                                                          | :heavy_check_mark:                                                                | The job ID for polling                                                            |
| `Status`                                                                          | [shared.CodeSamplesJobStatus](../../../pkg/models/shared/codesamplesjobstatus.md) | :heavy_check_mark:                                                                | The current status of the job. Possible values are `pending` or `running`.        |