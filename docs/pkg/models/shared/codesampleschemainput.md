# CodeSampleSchemaInput


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Languages`                                                   | []*string*                                                    | :heavy_check_mark:                                            | A list of languages to generate code samples for              |
| `PackageName`                                                 | **string*                                                     | :heavy_minus_sign:                                            | The name of the package                                       |
| `SchemaFile`                                                  | [shared.SchemaFile](../../../pkg/models/shared/schemafile.md) | :heavy_check_mark:                                            | The OpenAPI file to be uploaded                               |
| `SDKClassName`                                                | **string*                                                     | :heavy_minus_sign:                                            | The SDK class name                                            |