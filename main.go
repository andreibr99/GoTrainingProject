package main

import (
	"GoTrainingProject/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/user/", handlers.GetUserPermissions)
	http.HandleFunc("/v1/user/service/", handlers.GetServicePermissions)
	http.HandleFunc("/v1/service/", handlers.GetPermissionUsers)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
