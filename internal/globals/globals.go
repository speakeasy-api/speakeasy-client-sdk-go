// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package globals

type Globals struct {
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (o *Globals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}
