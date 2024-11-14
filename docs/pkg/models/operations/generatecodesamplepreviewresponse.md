# GenerateCodeSamplePreviewResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `TwoHundredApplicationJSONResponseStream`               | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Successfully returned codeSample overlay file           |
| `TwoHundredApplicationXYamlResponseStream`              | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Successfully returned codeSample overlay file           |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_check_mark:                                      | Raw HTTP response; suitable for custom response parsing |