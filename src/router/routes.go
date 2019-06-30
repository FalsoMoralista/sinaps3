package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"model"
	"net/http"
)


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/",home)
	router.HandleFunc("/teste",GetProcedimentos)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/html")
	http.ServeFile(w,req,"www/index.html")
}

func GetProcedimentos(w http.ResponseWriter, req *http.Request) {
	procedimentos, err := model.GetProcedimentos()
	if err == nil {
		var nao_agendados []model.Procedimento
		for _,procedimento := range procedimentos{
			if procedimento.Paciente_id == "" {
				nao_agendados = append(nao_agendados, procedimento)
			}
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		bArray,_ := json.Marshal(nao_agendados)
		w.Write(bArray)
	}else{
		w.WriteHeader(http.StatusNoContent)
	}
}


