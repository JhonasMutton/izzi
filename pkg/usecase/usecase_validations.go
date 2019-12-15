package usecase

import (
	"bytes"
	"encoding/json"
	"github.com/JhonasMutton/izzi/internal/errors"
	"github.com/JhonasMutton/izzi/internal/infra"
	"github.com/JhonasMutton/izzi/pkg/client"
	"github.com/JhonasMutton/izzi/pkg/models"
	"io/ioutil"
	"strings"
)

type IValidationsUseCase interface {
	ValidateRG(model models.ValidationRGModel) (bool, error)
}

type ValidationsUseCase struct {
	bigIdClient client.BigIDClient
}

func (v ValidationsUseCase) ValidateRG(model models.ValidationRGModel) (bool, error) {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(model)

	if err != nil{
		infra.Logger.Errorw("Error to encode ValidationRG", "error", err.Error())
		return false, errors.WrapWithMessage(errors.ErrInternalServer, err.Error())
	}

	response, err := v.bigIdClient.VerifyRG(reqBodyBytes.Bytes())
	if err != nil{
		infra.Logger.Errorw("Error to call BigId", "error", err.Error())
		return false, errors.WrapWithMessage(errors.ErrInternalServer, err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		infra.Logger.Errorw("Error to read body returned of Big id", "error", err.Error())
		return false, errors.WrapWithMessage(errors.ErrInternalServer, err.Error())
	}

	var result models.VerifiedRG
	err = json.Unmarshal(body,&result)
	return validate(result, model.Validations), nil
}

func validate(rgVerified models.VerifiedRG, rgData models.RGdata) (result bool) {
	if strings.ToLower(rgVerified.DocInfo.NAME) != strings.ToLower(rgData.Name) {
		return false
	}
	if strings.ToLower(rgVerified.DocInfo.BIRTHDATE) != strings.ToLower(rgData.BirthDate) {
		return false
	}
	if strings.ToLower(rgVerified.DocInfo.FATHERNAME) != strings.ToLower(rgData.FatherName) {
		return false
	}
	if strings.ToLower(rgVerified.DocInfo.MOTHERNAME) != strings.ToLower(rgData.MotherName) {
		return false
	}

	return true
}

func NewValidationsUseCase(bigIdClient client.BigIDClient) ValidationsUseCase {
	return ValidationsUseCase{bigIdClient: bigIdClient}
}