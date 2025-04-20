package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/invite", service.InviteAPIHandler)
	http.HandleFunc("/api/activity", service.CreateActivityHandler)
	http.HandleFunc("/api/user", service.CreateUserHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
