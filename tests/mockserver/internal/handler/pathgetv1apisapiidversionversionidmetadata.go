// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"net/http"
)

func pathGetV1ApisAPIIDVersionVersionIDMetadata(dir *logging.HTTPFileDirectory) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")

		switch test {
		case "getVersionMetadata":
			dir.HandlerFunc("getVersionMetadata", testGetVersionMetadataGetVersionMetadata)(w, req)
		default:
			http.Error(w, "Unknown test: "+test, http.StatusBadRequest)
		}
	}
}

func testGetVersionMetadataGetVersionMetadata(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, "Bearer"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.SecurityHeader(req, "x-api-key", true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.SecurityHeader(req, "x-workspace-identifier", true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respBody := []components.VersionMetadata{
		components.VersionMetadata{
			APIID:       "<id>",
			CreatedAt:   types.MustTimeFromString("2022-06-27T04:40:43.913Z"),
			MetaKey:     "<value>",
			MetaValue:   "<value>",
			VersionID:   "<id>",
			WorkspaceID: "<id>",
		},
	}
	respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

	if err != nil {
		http.Error(
			w,
			"Unable to encode response body as JSON: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBodyBytes)
}
