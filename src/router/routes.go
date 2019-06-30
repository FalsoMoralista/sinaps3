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

	router.HandleFunc("/procedimentos/disponiveis", ProcedimentosDisponiveis)
	router.HandleFunc("/paciente", Pacientes)

	router.HandleFunc("/vendor/bootstrap/css/bootstrap.min.css", bootstrap)
	router.HandleFunc("/vendor/fontawesome-free/css/all.min.css", all)
	router.HandleFunc("/vendor/simple-line-icons/css/simple-line-icons.css", line_icons)
	router.HandleFunc("/css/landing-page.min.css", landing_page)
	router.HandleFunc("/js/handy-collapse.js", handy_collapse)
	router.HandleFunc("/js/index.js", index_js)
	router.HandleFunc("/css/index.css", index_css)
	router.HandleFunc("", line_icons)
	fmt.Println("listening ...")
	log.Fatal(http.ListenAndServe(GetPort(), router))
}

func index_js(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","application/javascript")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/js/index.js")
}

func handy_collapse(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","application/javascript")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/js/handy-collapse.js")
}


func landing_page(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/css")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/css/landing-page.min.css")
}


func line_icons(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/css")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/vendor/simple-line-icons/css/simple-line-icons.css")
}

func index_css(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/css")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/css/index.css")
}

func all(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/css")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/vendor/fontawesome-free/css/all.min.css")
}

func bootstrap(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/css")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/vendor/bootstrap/css/bootstrap.min.css")
}

func Pacientes(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/html")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/paciente.html")
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
