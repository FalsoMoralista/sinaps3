package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"model"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/procedimentos-disponiveis", ProcedimentosDisponiveis)
	fmt.Println("listening...")
	log.Fatal(http.ListenAndServe(GetPort(), router))
}

func ProcedimentosDisponiveis(w http.ResponseWriter, req *http.Request) {
	procedimentos, err := model.GetProcedimentos()
	if err == nil {
		var nao_agendados []model.Procedimento
		for _, procedimento := range procedimentos {
			if procedimento.Paciente_id == "" {
				nao_agendados = append(nao_agendados, procedimento)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		bArray, _ := json.Marshal(nao_agendados)
		w.Write(bArray)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
