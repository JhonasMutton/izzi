package usecase

import (
	"bytes"
	"encoding/json"
	"github.com/JhonasMutton/izzi/internal/errors"
	"github.com/JhonasMutton/izzi/internal/infra"
	"github.com/JhonasMutton/izzi/pkg/client"
	"github.com/JhonasMutton/izzi/pkg/models"
	"io/ioutil"
)

type IInsuranceUseCase interface {
	Simulation(simulacoes models.SimulacoesModel) (*map[string]interface{}, error)
}

type InsuranceUseCase struct {
	mongeralAegonClient client.MongeralAegonClient
}

func (i InsuranceUseCase) Simulation(simulacoes models.SimulacoesModel) (*map[string]interface{}, error) {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(simulacoes)

	if err != nil{
		infra.Logger.Errorw("Error to encode Simulacoes", "error", err.Error())
		return nil, errors.WrapWithMessage(errors.ErrInternalServer, err.Error())
	}


	response, err := i.mongeralAegonClient.Simulation(reqBodyBytes.Bytes())
	if err != nil{
		infra.Logger.Errorw("Error to call Mongeral Aegon", "error", err.Error())
		return nil, errors.WrapWithMessage(errors.ErrInternalServer, err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		infra.Logger.Errorw("Error to read body returned of Mongeral Aegon", "error", err.Error())
		return nil, errors.WrapWithMessage(errors.ErrInternalServer, err.Error())
	}

	var result map[string]interface{}
	err = json.Unmarshal(body,&result)
	return &result, nil
}

func NewInsuranceUseCase(mongeralAegonClient client.MongeralAegonClient) InsuranceUseCase {
	return InsuranceUseCase{mongeralAegonClient: mongeralAegonClient}
}



