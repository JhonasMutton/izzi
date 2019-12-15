package pkg

import (
	"github.com/JhonasMutton/izzi/internal/infra"
	"github.com/JhonasMutton/izzi/pkg/api/config"
	"github.com/JhonasMutton/izzi/pkg/api/router"
	"github.com/JhonasMutton/izzi/pkg/client"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

type Application struct {
	Routes        *router.Routes
	HealthCheck   config.HealthCheck
	bigIdClient   client.BigIDClient
	mongeralAegon client.MongeralAegonClient
}

func NewApplication(routes *router.Routes, healthCheck config.HealthCheck, bigIdClient client.BigIDClient, mongeralAegon client.MongeralAegonClient) Application {
	return Application{Routes: routes, HealthCheck: healthCheck, bigIdClient: bigIdClient, mongeralAegon: mongeralAegon}
}

func (app Application) SetupHealthCheck() {
	if err := app.HealthCheck.SetupHealthCheck(); err != nil {
		panic("Error to setup health check!")
	}
}

func (app Application) SetupRoutes() *mux.Router {
	infra.Logger.Debug("Registering handlers")
	routers := app.Routes.MakeHandlers()
	return routers
}

func (app Application) SetupServerApplication(router *mux.Router) {
	infra.Logger.Debug("Setting up server application")
	serverPort := os.Getenv("SERVER_PORT")
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + serverPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	infra.Logger.Info("Server starting on port: ", serverPort)
	if e := srv.ListenAndServe(); e != nil {
		infra.Logger.Fatal("Failed to startup the HTTP server", "error", e.Error())
	}
}


