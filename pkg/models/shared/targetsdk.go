// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

type TargetSDK struct {
	// Remote commit ID.
	CommitHead *string `json:"commit_head,omitempty"`
	// Name of the CI environment.
	ContinuousIntegrationEnvironment *string `json:"continuous_integration_environment,omitempty"`
	// Error message if the last event was not successful.
	Error *string `json:"error,omitempty"`
	// Version of the generated target (post generation)
	GenerateConfigPostVersion *string `json:"generate_config_post_version,omitempty"`
	// gen.lock ID (expected to be a uuid). The same as `id`. A unique identifier for the target.
	GenerateGenLockID string `json:"generate_gen_lock_id"`
	// Features prior to generation
	GenerateGenLockPreFeatures *string `json:"generate_gen_lock_pre_features,omitempty"`
	// Artifact version for the Previous Generation
	GenerateGenLockPreVersion *string `json:"generate_gen_lock_pre_version,omitempty"`
	// Indicates whether the target was considered published.
	GeneratePublished *bool `json:"generate_published,omitempty"`
	// eg `typescript`, `terraform`, `python`
	GenerateTarget string `json:"generate_target"`
	// The name of the target as defined by the user.
	GenerateTargetName *string `json:"generate_target_name,omitempty"`
	// The version of the Speakeasy generator for this target eg v2 of the typescript generator.
	GenerateTargetVersion *string `json:"generate_target_version,omitempty"`
	// GitHub organization of the action.
	GhActionOrganization *string `json:"gh_action_organization,omitempty"`
	// GitHub repository of the action.
	GhActionRepository *string `json:"gh_action_repository,omitempty"`
	// Link to the GitHub action run.
	GhActionRunLink *string `json:"gh_action_run_link,omitempty"`
	// Version of the GitHub action.
	GhActionVersion *string `json:"gh_action_version,omitempty"`
	// Current working directory relative to the git root.
	GitRelativeCwd *string `json:"git_relative_cwd,omitempty"`
	// Default owner for git remote.
	GitRemoteDefaultOwner *string `json:"git_remote_default_owner,omitempty"`
	// Default repository name for git remote.
	GitRemoteDefaultRepo *string `json:"git_remote_default_repo,omitempty"`
	// User email from git configuration.
	GitUserEmail *string `json:"git_user_email,omitempty"`
	// User's name from git configuration. (not GitHub username)
	GitUserName *string `json:"git_user_name,omitempty"`
	// Remote hostname.
	Hostname *string `json:"hostname,omitempty"`
	// Unique identifier of the target the same as `generate_gen_lock_id`
	ID string `json:"id"`
	// Timestamp when the event was created in the database.
	LastEventCreatedAt time.Time `json:"last_event_created_at"`
	// Unique identifier of the last event for the target
	LastEventID string `json:"last_event_id"`
	// Type of interaction.
	LastEventInteractionType InteractionType `json:"last_event_interaction_type"`
	// Label of the git repository.
	RepoLabel *string `json:"repo_label,omitempty"`
	// The blob digest of the source.
	SourceBlobDigest *string `json:"source_blob_digest,omitempty"`
	// The namespace name of the source.
	SourceNamespaceName *string `json:"source_namespace_name,omitempty"`
	// The revision digest of the source.
	SourceRevisionDigest *string `json:"source_revision_digest,omitempty"`
	// Indicates whether the event was successful.
	Success *bool `json:"success,omitempty"`
}

func (t TargetSDK) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TargetSDK) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TargetSDK) GetCommitHead() *string {
	if o == nil {
		return nil
	}
	return o.CommitHead
}

func (o *TargetSDK) GetContinuousIntegrationEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.ContinuousIntegrationEnvironment
}

func (o *TargetSDK) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *TargetSDK) GetGenerateConfigPostVersion() *string {
	if o == nil {
		return nil
	}
	return o.GenerateConfigPostVersion
}

func (o *TargetSDK) GetGenerateGenLockID() string {
	if o == nil {
		return ""
	}
	return o.GenerateGenLockID
}

func (o *TargetSDK) GetGenerateGenLockPreFeatures() *string {
	if o == nil {
		return nil
	}
	return o.GenerateGenLockPreFeatures
}

func (o *TargetSDK) GetGenerateGenLockPreVersion() *string {
	if o == nil {
		return nil
	}
	return o.GenerateGenLockPreVersion
}

func (o *TargetSDK) GetGeneratePublished() *bool {
	if o == nil {
		return nil
	}
	return o.GeneratePublished
}

func (o *TargetSDK) GetGenerateTarget() string {
	if o == nil {
		return ""
	}
	return o.GenerateTarget
}

func (o *TargetSDK) GetGenerateTargetName() *string {
	if o == nil {
		return nil
	}
	return o.GenerateTargetName
}

func (o *TargetSDK) GetGenerateTargetVersion() *string {
	if o == nil {
		return nil
	}
	return o.GenerateTargetVersion
}

func (o *TargetSDK) GetGhActionOrganization() *string {
	if o == nil {
		return nil
	}
	return o.GhActionOrganization
}

func (o *TargetSDK) GetGhActionRepository() *string {
	if o == nil {
		return nil
	}
	return o.GhActionRepository
}

func (o *TargetSDK) GetGhActionRunLink() *string {
	if o == nil {
		return nil
	}
	return o.GhActionRunLink
}

func (o *TargetSDK) GetGhActionVersion() *string {
	if o == nil {
		return nil
	}
	return o.GhActionVersion
}

func (o *TargetSDK) GetGitRelativeCwd() *string {
	if o == nil {
		return nil
	}
	return o.GitRelativeCwd
}

func (o *TargetSDK) GetGitRemoteDefaultOwner() *string {
	if o == nil {
		return nil
	}
	return o.GitRemoteDefaultOwner
}

func (o *TargetSDK) GetGitRemoteDefaultRepo() *string {
	if o == nil {
		return nil
	}
	return o.GitRemoteDefaultRepo
}

func (o *TargetSDK) GetGitUserEmail() *string {
	if o == nil {
		return nil
	}
	return o.GitUserEmail
}

func (o *TargetSDK) GetGitUserName() *string {
	if o == nil {
		return nil
	}
	return o.GitUserName
}

func (o *TargetSDK) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *TargetSDK) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TargetSDK) GetLastEventCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastEventCreatedAt
}

func (o *TargetSDK) GetLastEventID() string {
	if o == nil {
		return ""
	}
	return o.LastEventID
}

func (o *TargetSDK) GetLastEventInteractionType() InteractionType {
	if o == nil {
		return InteractionType("")
	}
	return o.LastEventInteractionType
}

func (o *TargetSDK) GetRepoLabel() *string {
	if o == nil {
		return nil
	}
	return o.RepoLabel
}

func (o *TargetSDK) GetSourceBlobDigest() *string {
	if o == nil {
		return nil
	}
	return o.SourceBlobDigest
}

func (o *TargetSDK) GetSourceNamespaceName() *string {
	if o == nil {
		return nil
	}
	return o.SourceNamespaceName
}

func (o *TargetSDK) GetSourceRevisionDigest() *string {
	if o == nil {
		return nil
	}
	return o.SourceRevisionDigest
}

func (o *TargetSDK) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}
