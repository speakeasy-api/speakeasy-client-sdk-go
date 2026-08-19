# GithubStorePublishingSecretsRequest

A request to store publishing secrets for a github target


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `GenerateGenLockID`                            | `string`                                       | :heavy_check_mark:                             | The generation lock ID                         |
| `Secrets`                                      | map[string]`string`                            | :heavy_minus_sign:                             | A map of secrets to store in the GitHub target |