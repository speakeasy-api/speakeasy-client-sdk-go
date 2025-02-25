// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"context"
	"mockserver/internal/logging"
	"mockserver/internal/tracking"
	"net/http"
)

// GeneratedHandlers returns all generated handlers.
func GeneratedHandlers(ctx context.Context, dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) []*GeneratedHandler {
	return []*GeneratedHandler{
		NewGeneratedHandler(ctx, http.MethodDelete, "/v1/organization/add_ons/{add_on}", pathDeleteV1OrganizationAddOnsAddOn(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/namespaces", pathGetV1ArtifactsNamespaces(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/namespaces/{namespace_name}/revisions", pathGetV1ArtifactsNamespacesNamespaceNameRevisions(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/namespaces/{namespace_name}/tags", pathGetV1ArtifactsNamespacesNamespaceNameTags(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/remote_sources", pathGetV1ArtifactsRemoteSources(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/auth/access_token", pathGetV1AuthAccessToken(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/auth/validate", pathGetV1AuthValidate(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/code_sample", pathGetV1CodeSample(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/code_sample/preview/async/{jobID}", pathGetV1CodeSamplePreviewAsyncJobID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/github/check_access", pathGetV1GithubCheckAccess(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/github/setup", pathGetV1GithubSetup(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/oci/v2/{organization_slug}/{workspace_slug}/{namespace_name}/blobs/{digest}", pathGetV1OciV2OrganizationSlugWorkspaceSlugNamespaceNameBlobsDigest(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/oci/v2/{organization_slug}/{workspace_slug}/{namespace_name}/manifests/{revision_reference}", pathGetV1OciV2OrganizationSlugWorkspaceSlugNamespaceNameManifestsRevisionReference(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/organization/add_ons", pathGetV1OrganizationAddOns(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/organization/usage", pathGetV1OrganizationUsage(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/organization/{organizationID}", pathGetV1OrganizationOrganizationID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/organizations", pathGetV1Organizations(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/reports/changes/{documentChecksum}", pathGetV1ReportsChangesDocumentChecksum(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/reports/linting/{documentChecksum}", pathGetV1ReportsLintingDocumentChecksum(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/user", pathGetV1User(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace", pathGetV1Workspace(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/access", pathGetV1WorkspaceAccess(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/{workspace_id}", pathGetV1WorkspaceWorkspaceID(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/{workspace_id}/feature_flags", pathGetV1WorkspaceWorkspaceIDFeatureFlags(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/{workspace_id}/settings", pathGetV1WorkspaceWorkspaceIDSettings(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspaces", pathGetV1Workspaces(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/artifacts/namespaces/{namespace_name}/archive", pathPostV1ArtifactsNamespacesNamespaceNameArchive(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/artifacts/namespaces/{namespace_name}/tags", pathPostV1ArtifactsNamespacesNamespaceNameTags(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/artifacts/namespaces/{namespace_name}/visibility", pathPostV1ArtifactsNamespacesNamespaceNameVisibility(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/artifacts/remote_sources", pathPostV1ArtifactsRemoteSources(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/code_sample/preview", pathPostV1CodeSamplePreview(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/code_sample/preview/async", pathPostV1CodeSamplePreviewAsync(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/configure_code_samples", pathPostV1GithubConfigureCodeSamples(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/configure_mintlify_repo", pathPostV1GithubConfigureMintlifyRepo(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/configure_target", pathPostV1GithubConfigureTarget(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/link", pathPostV1GithubLink(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/publishing_secrets", pathPostV1GithubPublishingSecrets(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/trigger_action", pathPostV1GithubTriggerAction(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/organization", pathPostV1Organization(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/organization/add_ons", pathPostV1OrganizationAddOns(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/organization/free_trial", pathPostV1OrganizationFreeTrial(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/reports", pathPostV1Reports(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/short_urls", pathPostV1ShortUrls(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/subscriptions/{subscriptionID}/{namespaceName}/activate", pathPostV1SubscriptionsSubscriptionIDNamespaceNameActivate(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/subscriptions/{subscriptionID}/{namespaceName}/ignore", pathPostV1SubscriptionsSubscriptionIDNamespaceNameIgnore(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/suggest/items", pathPostV1SuggestItems(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/workspace", pathPostV1Workspace(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/workspace/feature_flags", pathPostV1WorkspaceFeatureFlags(dir, rt)),
		NewGeneratedHandler(ctx, http.MethodPut, "/v1/workspace/{workspace_id}/settings", pathPutV1WorkspaceWorkspaceIDSettings(dir, rt)),
	}
}
