package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetUserPermissions(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/v1/user/")
	if perms, ok := Permissions[user]; ok {
		json.NewEncoder(w).Encode(perms)
	} else {
		http.NotFound(w, r)
	}
}
