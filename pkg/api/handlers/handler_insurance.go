package handlers

import (
	"encoding/json"
	"github.com/JhonasMutton/izzi/internal/infra"
	"github.com/JhonasMutton/izzi/pkg/api/render"
	"github.com/JhonasMutton/izzi/pkg/models"
	"github.com/JhonasMutton/izzi/pkg/usecase"
	"net/http"
)

type InsuranceHandler struct {
	insuranceUseCase usecase.InsuranceUseCase
}

func NewInsuranceHandler(insuranceUseCase usecase.InsuranceUseCase) *InsuranceHandler {
	return &InsuranceHandler{insuranceUseCase: insuranceUseCase}
}

func (i *InsuranceHandler) Simulation(w http.ResponseWriter, r *http.Request) {
	var simulation models.SimulacoesDTO

	if err := json.NewDecoder(r.Body).Decode(&simulation); err != nil {
		infra.Logger.Errorw("Parsing simulacoes json error", "error", err.Error())
		render.ResponseError(w, err, http.StatusUnprocessableEntity)
		return
	}

	simulado, err := i.insuranceUseCase.Simulation(models.SimulacoesModel{}.ToModel(simulation))
	if err != nil{
		infra.Logger.Errorw("Error to simulate a insurance", "error", err.Error())
		render.ResponseError(w, err, http.StatusInternalServerError)
	}

	render.Response(w, simulado, 200)
}