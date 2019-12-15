package router

import (
	"github.com/JhonasMutton/izzi/pkg/api/handlers"
	"github.com/gorilla/mux"
	"github.com/hellofresh/health-go"
	"net/http"
)

type Routes struct {
	insuranceHandler *handlers.InsuranceHandler
	validationHandler *handlers.ValidationHandler
}

func NewRoutes(insuranceHandler *handlers.InsuranceHandler, validationHandler *handlers.ValidationHandler) *Routes {
	return &Routes{insuranceHandler: insuranceHandler, validationHandler: validationHandler}
}



func (routes *Routes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/simulacao", routes.insuranceHandler.Simulation).Methods(http.MethodPost)
	r.HandleFunc("/verificarRG", routes.validationHandler.VerifyRgData).Methods(http.MethodPost)
	r.HandleFunc("/health", health.HandlerFunc).Methods(http.MethodGet)

	return r
}
