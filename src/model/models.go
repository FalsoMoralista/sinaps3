package model

type Paciente struct{
	Id string `json:"id"`
	Nome string `json:"nome"`
	Endereco string `json:"endereco"`
	Cidade string	`json:"cidade"`
	Estado string	`json:"estado"`
	Telefone string `json:"telefone"`
	Cartao_sus string `json:"cartao_sus"`
	Email string	`json:"email"`
	Senha string	`json:"senha"`
}

type Procedimento struct {
	Procedimento_id string`json:"procedimento_id"`
	Uds_id string`json:"uds_id"`
	Paciente_id string`json:"paciente_id"`
	Datahora string	`json:"datahora"`
}

type Consulta struct {
	Especialidade string `json:"especialidade"`
	Uds_id string		 `json:"uds_id"`
	Paciente_id string `json:"paciente_id"`
	Datahora string	`json:"datahora"`
}

type Uds struct {
	Nome string `json:"nome"`
	Endereco string`json:"endereco"`
	Cidade string`json:"cidade"`
	Estado string`json:"estado"`
	Telefone string`json:"telefone"`
}

type Join_procedimento_unidade struct {
	Unidade Uds `json:"unidade"`
	Procedimento Procedimento `json:"procedimento"`
}