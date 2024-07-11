// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GithubTriggerActionRequest - A request to trigger an action on a GitHub target
type GithubTriggerActionRequest struct {
	// The generation lock ID
	GenLockID string `json:"gen_lock_id"`
	// The GitHub organization name
	Org string `json:"org"`
	// The GitHub repository name
	RepoName string `json:"repo_name"`
	// A version to override the SDK too in workflow dispatch
	SetVersion *string `json:"set_version,omitempty"`
	// The target name for the action
	TargetName *string `json:"target_name,omitempty"`
}

func (o *GithubTriggerActionRequest) GetGenLockID() string {
	if o == nil {
		return ""
	}
	return o.GenLockID
}

func (o *GithubTriggerActionRequest) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *GithubTriggerActionRequest) GetRepoName() string {
	if o == nil {
		return ""
	}
	return o.RepoName
}

func (o *GithubTriggerActionRequest) GetSetVersion() *string {
	if o == nil {
		return nil
	}
	return o.SetVersion
}

func (o *GithubTriggerActionRequest) GetTargetName() *string {
	if o == nil {
		return nil
	}
	return o.TargetName
}
