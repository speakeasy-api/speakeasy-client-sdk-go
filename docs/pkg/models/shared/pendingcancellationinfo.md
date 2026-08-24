# PendingCancellationInfo

Information about a pending subscription cancellation


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `EffectiveAt`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | When the cancellation takes effect (from Stripe current_period_end) |
| `TargetToKeep`                                                      | `*string`                                                           | :heavy_minus_sign:                                                  | gen_lock_id of the target that will be kept after downgrade         |
| `TargetToKeepName`                                                  | `*string`                                                           | :heavy_minus_sign:                                                  | Display name of the target to keep                                  |