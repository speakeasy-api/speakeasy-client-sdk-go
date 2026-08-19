# CancelSubscriptionResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `EffectiveAt`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | When the cancellation takes effect (from Stripe current_period_end) |
| `Message`                                                           | `*string`                                                           | :heavy_minus_sign:                                                  | Additional message about the operation                              |
| `Success`                                                           | `bool`                                                              | :heavy_check_mark:                                                  | Whether the operation was successful                                |
| `TargetToKeep`                                                      | `*string`                                                           | :heavy_minus_sign:                                                  | The gen_lock_id of the target to keep                               |