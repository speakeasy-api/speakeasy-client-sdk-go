# CodeSampleSchemaInput


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Language`                                                    | *string*                                                      | :heavy_check_mark:                                            | The language to generate code samples for                     |
| `OperationIds`                                                | []*string*                                                    | :heavy_minus_sign:                                            | A list of operations IDs to generate code samples for         |
| `PackageName`                                                 | **string*                                                     | :heavy_minus_sign:                                            | The name of the package                                       |
| `SchemaFile`                                                  | [shared.SchemaFile](../../../pkg/models/shared/schemafile.md) | :heavy_check_mark:                                            | The OpenAPI file to be uploaded                               |
| `SDKClassName`                                                | **string*                                                     | :heavy_minus_sign:                                            | The SDK class name                                            |