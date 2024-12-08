// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"context"
	"net/http"

	"mockserver/internal/logging"
)

// GeneratedHandlers returns all generated handlers.
func GeneratedHandlers(ctx context.Context, dir *logging.HTTPFileDirectory) []*GeneratedHandler {
	return []*GeneratedHandler{
		NewGeneratedHandler(ctx, http.MethodDelete, "/v1/apis/{apiID}/version/{versionID}", pathDeleteV1ApisAPIIDVersionVersionID(dir)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", pathDeleteV1ApisAPIIDVersionVersionIDAPIEndpointsAPIEndpointID(dir)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/v1/apis/{apiID}/version/{versionID}/metadata/{metaKey}/{metaValue}", pathDeleteV1ApisAPIIDVersionVersionIDMetadataMetaKeyMetaValue(dir)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}", pathDeleteV1ApisAPIIDVersionVersionIDSchemaRevisionID(dir)),
		NewGeneratedHandler(ctx, http.MethodDelete, "/v1/workspace/embed-access-tokens/{tokenID}", pathDeleteV1WorkspaceEmbedAccessTokensTokenID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis", pathGetV1Apis(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}", pathGetV1ApisAPIID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/api_endpoints", pathGetV1ApisAPIIDAPIEndpoints(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/api_endpoints", pathGetV1ApisAPIIDVersionVersionIDAPIEndpoints(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/find/{displayName}", pathGetV1ApisAPIIDVersionVersionIDAPIEndpointsFindDisplayName(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", pathGetV1ApisAPIIDVersionVersionIDAPIEndpointsAPIEndpointID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}/generate/openapi", pathGetV1ApisAPIIDVersionVersionIDAPIEndpointsAPIEndpointIDGenerateOpenapi(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}/generate/postman", pathGetV1ApisAPIIDVersionVersionIDAPIEndpointsAPIEndpointIDGeneratePostman(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/generate/openapi", pathGetV1ApisAPIIDVersionVersionIDGenerateOpenapi(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/generate/postman", pathGetV1ApisAPIIDVersionVersionIDGeneratePostman(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/metadata", pathGetV1ApisAPIIDVersionVersionIDMetadata(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/schema", pathGetV1ApisAPIIDVersionVersionIDSchema(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/schema/{baseRevisionID}/diff/{targetRevisionID}", pathGetV1ApisAPIIDVersionVersionIDSchemaBaseRevisionIDDiffTargetRevisionID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}", pathGetV1ApisAPIIDVersionVersionIDSchemaRevisionID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/apis/{apiID}/version/{versionID}/schemas", pathGetV1ApisAPIIDVersionVersionIDSchemas(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/namespaces", pathGetV1ArtifactsNamespaces(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/namespaces/{namespace_name}/revisions", pathGetV1ArtifactsNamespacesNamespaceNameRevisions(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/namespaces/{namespace_name}/tags", pathGetV1ArtifactsNamespacesNamespaceNameTags(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/artifacts/remote_sources", pathGetV1ArtifactsRemoteSources(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/auth/access_token", pathGetV1AuthAccessToken(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/auth/validate", pathGetV1AuthValidate(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/code_sample/preview/async/{jobID}", pathGetV1CodeSamplePreviewAsyncJobID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/eventlog/{requestID}", pathGetV1EventlogRequestID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/eventlog/{requestID}/generate/postman", pathGetV1EventlogRequestIDGeneratePostman(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/github/check_access", pathGetV1GithubCheckAccess(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/github/setup", pathGetV1GithubSetup(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/oci/v2/{organization_slug}/{workspace_slug}/{namespace_name}/blobs/{digest}", pathGetV1OciV2OrganizationSlugWorkspaceSlugNamespaceNameBlobsDigest(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/oci/v2/{organization_slug}/{workspace_slug}/{namespace_name}/manifests/{revision_reference}", pathGetV1OciV2OrganizationSlugWorkspaceSlugNamespaceNameManifestsRevisionReference(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/organization/usage", pathGetV1OrganizationUsage(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/organization/{organizationID}", pathGetV1OrganizationOrganizationID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/organizations", pathGetV1Organizations(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/reports/changes/{documentChecksum}", pathGetV1ReportsChangesDocumentChecksum(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/reports/linting/{documentChecksum}", pathGetV1ReportsLintingDocumentChecksum(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/user", pathGetV1User(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace", pathGetV1Workspace(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/access", pathGetV1WorkspaceAccess(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/embed-access-token", pathGetV1WorkspaceEmbedAccessToken(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/{workspace_id}", pathGetV1WorkspaceWorkspaceID(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/{workspace_id}/feature_flags", pathGetV1WorkspaceWorkspaceIDFeatureFlags(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/{workspace_id}/registry_subscriptions", pathGetV1WorkspaceWorkspaceIDRegistrySubscriptions(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspace/{workspace_id}/settings", pathGetV1WorkspaceWorkspaceIDSettings(dir)),
		NewGeneratedHandler(ctx, http.MethodGet, "/v1/workspaces", pathGetV1Workspaces(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/apis/{apiID}/version/{versionID}/metadata", pathPostV1ApisAPIIDVersionVersionIDMetadata(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/apis/{apiID}/version/{versionID}/schema", pathPostV1ApisAPIIDVersionVersionIDSchema(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/artifacts/namespaces/{namespace_name}/tags", pathPostV1ArtifactsNamespacesNamespaceNameTags(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/artifacts/namespaces/{namespace_name}/visibility", pathPostV1ArtifactsNamespacesNamespaceNameVisibility(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/artifacts/remote_sources", pathPostV1ArtifactsRemoteSources(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/code_sample/preview", pathPostV1CodeSamplePreview(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/code_sample/preview/async", pathPostV1CodeSamplePreviewAsync(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/configure_code_samples", pathPostV1GithubConfigureCodeSamples(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/configure_mintlify_repo", pathPostV1GithubConfigureMintlifyRepo(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/configure_target", pathPostV1GithubConfigureTarget(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/link", pathPostV1GithubLink(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/publishing_secrets", pathPostV1GithubPublishingSecrets(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/github/trigger_action", pathPostV1GithubTriggerAction(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/organization", pathPostV1Organization(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/organization/free_trial", pathPostV1OrganizationFreeTrial(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/reports", pathPostV1Reports(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/short_urls", pathPostV1ShortUrls(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/suggest/items", pathPostV1SuggestItems(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/workspace", pathPostV1Workspace(dir)),
		NewGeneratedHandler(ctx, http.MethodPost, "/v1/workspace/{workspace_id}/registry_subscriptions", pathPostV1WorkspaceWorkspaceIDRegistrySubscriptions(dir)),
		NewGeneratedHandler(ctx, http.MethodPut, "/v1/apis/{apiID}", pathPutV1ApisAPIID(dir)),
		NewGeneratedHandler(ctx, http.MethodPut, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", pathPutV1ApisAPIIDVersionVersionIDAPIEndpointsAPIEndpointID(dir)),
		NewGeneratedHandler(ctx, http.MethodPut, "/v1/workspace/{workspace_id}/settings", pathPutV1WorkspaceWorkspaceIDSettings(dir)),
	}
}
