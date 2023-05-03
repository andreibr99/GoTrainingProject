package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetPermissionUsers(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/v1/service/"), "/")
	service := parts[0]
	feature := parts[1]
	permission := parts[2]

	var users []string
	for user, perms := range Permissions {
		if svcPerms, ok := perms[service]; ok {
			for _, perm := range svcPerms {
				if perm.Key == fmt.Sprintf("%s.%s.%s", service, feature, permission) &&
					strings.HasPrefix(perm.Key, fmt.Sprintf("%s.%s", service, feature)) &&
					perm.Value == true {
					users = append(users, user)
				}
			}
		}
	}

	json.NewEncoder(w).Encode(users)
}
