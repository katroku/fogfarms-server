package main

import (
	"github.com/KitaPDev/fogfarms-server/src/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/modulegroup_management"
	"github.com/KitaPDev/fogfarms-server/src/plant_management"
	"github.com/KitaPDev/fogfarms-server/src/user_management"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	mux := http.NewServeMux()

	jwtAuthHandler := jwt.MakeHTTPHandler()
	mux.Handle("/auth", jwtAuthHandler)

	moduleGroupManagementHandler := modulegroup_management.MakeHTTPHandler()
	mux.Handle("/modulegroup_management", moduleGroupManagementHandler)
	mux.Handle("/modulegroup_management/", moduleGroupManagementHandler)

	userManagementHandler := user_management.MakeHTTPHandler()
	mux.Handle("/user_management", userManagementHandler)
	mux.Handle("/user_management/", userManagementHandler)

	plantManagementHandler := plant_management.MakeHTTPHandler()
	mux.Handle("/plant_management", plantManagementHandler)
	mux.Handle("/plant_management/", plantManagementHandler)

	return http.ListenAndServe(":9090", mux)
}

