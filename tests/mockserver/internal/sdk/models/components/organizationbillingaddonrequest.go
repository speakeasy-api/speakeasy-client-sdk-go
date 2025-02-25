// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// OrganizationBillingAddOnRequest - A request to add billing add ons
type OrganizationBillingAddOnRequest struct {
	AddOns []BillingAddOn `json:"add_ons"`
}

func (o *OrganizationBillingAddOnRequest) GetAddOns() []BillingAddOn {
	if o == nil {
		return []BillingAddOn{}
	}
	return o.AddOns
}
