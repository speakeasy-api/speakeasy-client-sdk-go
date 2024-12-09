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

func pathGetV1ApisAPIID(dir *logging.HTTPFileDirectory) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")

		switch test {
		case "getAllApiVersions":
			dir.HandlerFunc("getAllApiVersions", testGetAllAPIVersionsGetAllAPIVersions)(w, req)
		default:
			http.Error(w, "Unknown test: "+test, http.StatusBadRequest)
		}
	}
}

func testGetAllAPIVersionsGetAllAPIVersions(w http.ResponseWriter, req *http.Request) {
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
	respBody := []components.API{
		components.API{
			APIID:       "<id>",
			CreatedAt:   types.MustTimeFromString("2022-12-01T14:20:45.655Z"),
			Description: "cavernous so continually qua",
			UpdatedAt:   types.MustTimeFromString("2023-05-30T05:25:29.047Z"),
			VersionID:   "<id>",
			WorkspaceID: "<id>",
		},
		components.API{
			APIID:       "<id>",
			CreatedAt:   types.MustTimeFromString("2023-04-23T11:41:22.396Z"),
			Description: "readmit chubby oof gym sedately micromanage trench",
			UpdatedAt:   types.MustTimeFromString("2022-10-16T03:11:00.742Z"),
			VersionID:   "<id>",
			WorkspaceID: "<id>",
		},
		components.API{
			APIID:       "<id>",
			CreatedAt:   types.MustTimeFromString("2023-03-06T08:55:20.384Z"),
			Description: "aw even but peter abaft",
			UpdatedAt:   types.MustTimeFromString("2024-07-27T04:38:48.315Z"),
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
