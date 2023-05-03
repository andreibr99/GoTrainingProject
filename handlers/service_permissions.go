package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetServicePermissions(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/v1/user/"), "/")
	user := parts[0]
	service := parts[1]

	if perms, ok := Permissions[user]; ok {
		if svcPerms, ok := perms[service]; ok {
			json.NewEncoder(w).Encode(svcPerms)
		} else {
			http.NotFound(w, r)
		}
	} else {
		http.NotFound(w, r)
	}
}
