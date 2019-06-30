package main

import (
	"fmt"
	"model"
)

const url = "https://hackathonsanar5.herokuapp.com/pacientes"

func main(){
	pacientes,_ := model.GetPacientes()
	procedimentos, _:= model.GetProcedimentos()
	fmt.Println(pacientes[0])
	fmt.Println(procedimentos[0])
	fmt.Println(procedimentos[1])
	fmt.Println(procedimentos[2])
	fmt.Println(procedimentos[3])

}

