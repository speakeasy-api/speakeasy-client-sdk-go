// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// OrganizationBillingAddOnResponse - Billing add on response
type OrganizationBillingAddOnResponse struct {
	AddOns []BillingAddOn `json:"add_ons"`
}

func (o *OrganizationBillingAddOnResponse) GetAddOns() []BillingAddOn {
	if o == nil {
		return []BillingAddOn{}
	}
	return o.AddOns
}
