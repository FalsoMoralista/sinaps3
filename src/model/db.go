package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const(
	db = "https://hackathonsanar5.herokuapp.com"
	pacientes = "/pacientes"
	agendar_procedimentos = "/agendar-procedimentos"
	procedimentos_agendados = "/agendamentos/procedimentos"
)

func GetPacientes()([]Paciente, error){
	response,err := http.Get(db+pacientes)
	pacientes := []Paciente{}
	body,err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body,&pacientes)
	return pacientes, err
}

func GetProcedimentos()([]Procedimento,error){
	response,err := http.Get(db+procedimentos_agendados)
	procedimentos := []Procedimento{}
	body,err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body,&procedimentos)
	return procedimentos, err
}

func GetConsultas()([]Consulta,error){
	response,err := http.Get(db+agendar_procedimentos)
	procedimentos := []Consulta{}
	body,err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body,&procedimentos)
	return procedimentos, err
}
