package models

type ValidationRGDTO struct {
	RgImage string `json:"rgImage"`
	RGData  RGdata `json:"rgData"`
}

type RGdata struct {
	Name       string `json:"name"`
	FatherName string `json:"fatherName"`
	MotherName string `json:"motherName"`
	BirthDate  string `json:"birthDate"`
}

type ValidationRGModel struct {
	Parameters  []string `json:"Parameters"`
	Validations RGdata
}

func (v ValidationRGModel) ToModel(validationRGDTO ValidationRGDTO) ValidationRGModel {
	v.Parameters = []string{
		"DOC_IMG=" + validationRGDTO.RgImage,
		"DOC_TYPE=RG", "SIDE=C",
	}
	v.Validations = validationRGDTO.RGData

	return v
}

type VerifiedRG struct {
	DocInfo struct {
		BIRTHDATE  string `json:"BIRTHDATE"`
		DOCTYPE    string `json:"DOCTYPE"`
		FATHERNAME string `json:"FATHERNAME"`
		MOTHERNAME string `json:"MOTHERNAME"`
		NAME       string `json:"NAME"`
	} `json:"DocInfo"`
	TicketID      string        `json:"TicketId"`
	ResultCode    int           `json:"ResultCode"`
	ResultMessage string        `json:"ResultMessage"`
	Questions     []interface{} `json:"Questions"`
}

type RGVerifiedResult struct {
	IsValid bool `json:"isValid"`
}