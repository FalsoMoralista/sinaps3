package main

import (
	"fmt"
	"model"
)

const url = "https://hackathonsanar5.herokuapp.com/pacientes"

func main(){
	pacientes,_ := model.GetPacientes()
	procedimentos, _:= model.GetProcedimentos()
	unidades,_ := model.GetUds()
	fmt.Println(pacientes[0])
	fmt.Println(procedimentos[0])
	fmt.Println(unidades[0])
}

