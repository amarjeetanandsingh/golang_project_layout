package main

import (
	"git.comp.com/foo/foo/internal/auth"
	"git.comp.com/foo/foo/internal/cfg"
	"git.comp.com/foo/foo/internal/db/arango"
	"git.comp.com/foo/foo/internal/routeplan"
	"git.comp.com/foo/foo/internal/trip"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	cfg.Load()

	eServer := echo.New()
	eServer.Use(auth.Handler)
	routeplan.RegisterAPI(eServer)
	trip.RegisterAPI(eServer)

	arango.InitializeArangoDatabase()

	httperr := http.ListenAndServe(":8080", eServer)
	if httperr != nil {
		log.Fatal("error starting server", httperr.Error())
	}
}
