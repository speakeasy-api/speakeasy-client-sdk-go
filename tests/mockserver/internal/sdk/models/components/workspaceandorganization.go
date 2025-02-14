// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// WorkspaceAndOrganization - A workspace and organization
type WorkspaceAndOrganization struct {
	// A speakeasy organization
	Organization Organization `json:"organization"`
	// A speakeasy workspace
	Workspace Workspace `json:"workspace"`
}

func (o *WorkspaceAndOrganization) GetOrganization() Organization {
	if o == nil {
		return Organization{}
	}
	return o.Organization
}

func (o *WorkspaceAndOrganization) GetWorkspace() Workspace {
	if o == nil {
		return Workspace{}
	}
	return o.Workspace
}
