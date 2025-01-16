# GetCodeSamplesRequest


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Languages`                                              | []*string*                                               | :heavy_minus_sign:                                       | The languages to retrieve snippets for.                  |                                                          |
| `OperationIds`                                           | []*string*                                               | :heavy_minus_sign:                                       | The operation IDs to retrieve snippets for.              | getPetById                                               |
| `RegistryURL`                                            | *string*                                                 | :heavy_check_mark:                                       | The registry URL from which to retrieve the snippets.    | https://spec.speakeasy.com/my-org/my-workspace/my-source |