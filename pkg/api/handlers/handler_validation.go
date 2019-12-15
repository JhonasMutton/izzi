package handlers

import (
	"encoding/json"
	"github.com/JhonasMutton/izzi/internal/infra"
	"github.com/JhonasMutton/izzi/pkg/api/render"
	"github.com/JhonasMutton/izzi/pkg/models"
	"github.com/JhonasMutton/izzi/pkg/usecase"
	"net/http"
)

type ValidationHandler struct {
	validationsUseCase usecase.ValidationsUseCase
}

func NewValidationHandler(validationsUseCase usecase.ValidationsUseCase) *ValidationHandler {
	return &ValidationHandler{validationsUseCase: validationsUseCase}

}



func (v *ValidationHandler) VerifyRgData(w http.ResponseWriter, r *http.Request) {
	var validationRGDTO models.ValidationRGDTO

	if err := json.NewDecoder(r.Body).Decode(&validationRGDTO); err != nil {
		infra.Logger.Errorw("Parsing validation RG json error", "error", err.Error())
		render.ResponseError(w, err, http.StatusUnprocessableEntity)
		return
	}

	validationRGModel := models.ValidationRGModel{}.ToModel(validationRGDTO)
	result, err := v.validationsUseCase.ValidateRG(validationRGModel)
	if err != nil{
		infra.Logger.Errorw("Error to validate rg", "error", err.Error())
		render.ResponseError(w, err, http.StatusInternalServerError)
	}else {
		render.Response(w, models.RGVerifiedResult{IsValid:result}, 200)
	}

}