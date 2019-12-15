package models
type SimulacoesDTO struct {
	Simulacoes []Simulacao `json:"simulacoes"`
}

type Proponente struct {
Nome           string `json:"nome"`
Cpf            string `json:"cpf"`
DataNascimento string `json:"dataNascimento"`
ProfissaoCbo   string `json:"profissaoCbo"`
Renda          int    `json:"renda"`
SexoID         int    `json:"sexoId"`
Uf             string `json:"uf"`
}

type Simulacao struct {
		Proponente `json:"proponente"`
		PeriodicidadeCobrancaID int `json:"periodicidadeCobrancaId"`

}

type SimulacoesModel struct {
	Simulacoes []Simulacao `json:"simulacoes"`
}

func (s SimulacoesModel) ToModel (dto SimulacoesDTO) SimulacoesModel{
	for k, v := range dto.Simulacoes {
		s.Simulacoes[k]= v
		s.Simulacoes[k].PeriodicidadeCobrancaID = 30
		s.Simulacoes[k].ProfissaoCbo = "3171-10" //Como é um pubico alvo a profissão é a mesma
	}
	return s
}

