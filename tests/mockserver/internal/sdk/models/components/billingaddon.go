// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type BillingAddOn string

const (
	BillingAddOnWebhooks          BillingAddOn = "webhooks"
	BillingAddOnSDKTesting        BillingAddOn = "sdk_testing"
	BillingAddOnCustomCodeRegions BillingAddOn = "custom_code_regions"
	BillingAddOnSnippetAi         BillingAddOn = "snippet_ai"
)

func (e BillingAddOn) ToPointer() *BillingAddOn {
	return &e
}
