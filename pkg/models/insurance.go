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

type PropostaModel struct {
	PROPOSTA struct {
		NUMERO          string `json:"NUMERO"`
		DTPROTOCOLO     string `json:"DT_PROTOCOLO"`
		DTASSINATURA    string `json:"DT_ASSINATURA"`
		DTINDEXACAO     string `json:"DT_INDEXACAO"`
		DADOSPROPONENTE struct {
			MATRICULA    string `json:"MATRICULA"`
			NOME         string `json:"NOME"`
			DTNASCIMENTO string `json:"DT_NASCIMENTO"`
			IDADE        string `json:"IDADE"`
			SEXO         string `json:"SEXO"`
			ESTADOCIVIL  string `json:"ESTADO_CIVIL"`
			CPF          string `json:"CPF"`
			TITULARCPF   string `json:"TITULAR_CPF"`
			EMAIL        string `json:"EMAIL"`
			RESIDEBRASIL string `json:"RESIDE_BRASIL"`
			RENDAMENSAL  string `json:"RENDA_MENSAL"`
			NUMFILHOS    string `json:"NUM_FILHOS"`
			PPE          string `json:"PPE"`
			DOCUMENTOS   struct {
				DOCUMENTO struct {
					NATUREZADOC    string `json:"NATUREZA_DOC"`
					DOCUMENTO      string `json:"DOCUMENTO"`
					ORGAOEXPEDIDOR string `json:"ORGAO_EXPEDIDOR"`
					DATAEXPEDICAO  string `json:"DATA_EXPEDICAO"`
				} `json:"DOCUMENTO"`
			} `json:"DOCUMENTOS"`
			ENDERECOS struct {
				TPCORRESPONDENCIA string `json:"TP_CORRESPONDENCIA"`
				ENDERECO          []struct {
					TIPO        string `json:"TIPO"`
					LOGRADOURO  string `json:"LOGRADOURO"`
					NUMERO      string `json:"NUMERO"`
					COMPLEMENTO string `json:"COMPLEMENTO"`
					BAIRRO      string `json:"BAIRRO"`
					CIDADE      string `json:"CIDADE"`
					ESTADO      string `json:"ESTADO"`
					CEP         string `json:"CEP"`
				} `json:"ENDERECO"`
			} `json:"ENDERECOS"`
			TELEFONES struct {
				TELEFONE []struct {
					TIPO   string `json:"TIPO"`
					DDI    string `json:"DDI"`
					DDD    string `json:"DDD"`
					NUMERO string `json:"NUMERO"`
				} `json:"TELEFONE"`
			} `json:"TELEFONES"`
			PROFISSAO struct {
				CODIGO    string `json:"CODIGO"`
				DESCRICAO string `json:"DESCRICAO"`
				CATEGORIA string `json:"CATEGORIA"`
				EMPRESA   struct {
					NOME string `json:"NOME"`
				} `json:"EMPRESA"`
			} `json:"PROFISSAO"`
		} `json:"DADOS_PROPONENTE"`
		PLANOS struct {
			VLTOTAL string `json:"VL_TOTAL"`
			PLANO   []struct {
				CODIGO          string `json:"CODIGO"`
				NOME            string `json:"NOME"`
				VLAPINICIAL     string `json:"VL_AP_INICIAL"`
				VLPORTAB        string `json:"VL_PORTAB"`
				TPTRIBUTACAO    string `json:"TP_TRIBUTACAO"`
				DTCONCESSAO     string `json:"DT_CONCESSAO"`
				PRAZOCERTO      string `json:"PRAZO_CERTO"`
				PRAZODECRESCIMO string `json:"PRAZO_DECRESCIMO"`
				COBERTURAS      struct {
					COBERTURA struct {
						CODIGO      string `json:"CODIGO"`
						VLCONTRIB   string `json:"VL_CONTRIB"`
						VLCOBERTURA string `json:"VL_COBERTURA"`
					} `json:"COBERTURA"`
				} `json:"COBERTURAS"`
			} `json:"PLANO"`
		} `json:"PLANOS"`
		BENEFICIARIOS struct {
			BENEFICIARIO []struct {
				NOME         string `json:"NOME"`
				NASCIMENTO   string `json:"NASCIMENTO"`
				PARENTESCO   string `json:"PARENTESCO"`
				PARTICIPACAO string `json:"PARTICIPACAO"`
				CDPLANO      string `json:"CD_PLANO"`
			} `json:"BENEFICIARIO"`
		} `json:"BENEFICIARIOS"`
		DECLARACOES struct {
			DPS struct {
				TIPODPS   string `json:"TIPO_DPS"`
				PESO      string `json:"PESO"`
				ALTURA    string `json:"ALTURA"`
				PERGUNTAS struct {
					PERGUNTA []struct {
						NUMERO      string `json:"NUMERO"`
						QUESTAO     string `json:"QUESTAO"`
						RESPOSTA    string `json:"RESPOSTA"`
						OBSRESPOSTA string `json:"OBS_RESPOSTA"`
					} `json:"PERGUNTA"`
				} `json:"PERGUNTAS"`
			} `json:"DPS"`
		} `json:"DECLARACOES"`
		DADOSCOBRANCA struct {
			PERIODICIDADE string `json:"PERIODICIDADE"`
			TIPOCOBRANCA  string `json:"TIPO_COBRANCA"`
			DIAVENCIMENTO string `json:"DIA_VENCIMENTO"`
			COMPDEBITO    string `json:"COMP_DEBITO"`
			NUMCONVENIO   string `json:"NUM_CONVENIO"`
		} `json:"DADOS_COBRANCA"`
		USOMONGERAL struct {
			CONVADESAO        string `json:"CONV_ADESAO"`
			ACAOMARKETING     string `json:"ACAO_MARKETING"`
			ALTERNATIVA       string `json:"ALTERNATIVA"`
			SUCURSAL          string `json:"SUCURSAL"`
			DIRREGIONAL       string `json:"DIR_REGIONAL"`
			GERSUCURSAL       string `json:"GER_SUCURSAL"`
			GERCOMERCIAL      string `json:"GER_COMERCIAL"`
			AGENTE            string `json:"AGENTE"`
			CORRETOR1         string `json:"CORRETOR1"`
			CORRETOR2         string `json:"CORRETOR2"`
			AGENTEFIDELIZACAO string `json:"AGENTE_FIDELIZACAO"`
			MODELOPROPOSTA    string `json:"MODELO_PROPOSTA"`
			MODELOPROPOSTAGED string `json:"MODELO_PROPOSTA_GED"`
			TIPOCOMISSAO      string `json:"TIPO_COMISSAO"`
		} `json:"USO_MONGERAL"`
	} `json:"PROPOSTA"`
}