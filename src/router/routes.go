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
	root := http.FileServer(http.Dir("www/startbootstrap-landing-page-gh-pages/"))
	css := http.FileServer(http.Dir("www/startbootstrap-landing-page-gh-pages/css"))
 	router.Handle("/",root)
	router.Handle("/index.html",root)
	router.Handle("/css",css)

	router.HandleFunc("/paciente.html", Paciente)
	router.HandleFunc("/medico.html", Medico)
	router.HandleFunc("/paciente", Paciente)
	router.HandleFunc("/medico", Medico)
	router.HandleFunc("/procedimentos/disponiveis", ProcedimentosDisponiveis)
	router.HandleFunc("/img/sinaps_logo1.png",logoSinaps)
	router.HandleFunc("/vendor/jquery/jquery.min.js", jquery)
	router.HandleFunc("/vendor/bootstrap/js/bootstrap.bundle.min.js", bootstrap_js)
	router.HandleFunc("/vendor/bootstrap/css/bootstrap.min.css", bootstrap)
	router.HandleFunc("/vendor/fontawesome-free/css/all.min.css", all)
	router.HandleFunc("/vendor/simple-line-icons/css/simple-line-icons.css", line_icons)
	router.HandleFunc("/vendor/simple-line-icons/fonts/Simple-Line-Icons.woff2?v=2.4.0",simple_icons)
	router.HandleFunc("/css/landing-page.min.css", landing_page)
	router.HandleFunc("/js/handy-collapse.js", handy_collapse)
	router.HandleFunc("/js/index.js", index_js)
	router.HandleFunc("/css/index.css", index_css)
	router.HandleFunc("", line_icons)
	fmt.Println("listening ...")
	log.Fatal(http.ListenAndServe(GetPort(), router))
}


func simple_icons(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/font")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/simple-line-icons/fonts/Simple-Line-Icons.woff2?v=2.4.0")
}

func jquery(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/vendor/jquery/jquery.min.js")
}

func bootstrap_js(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/vendor/bootstrap/js/bootstrap.bundle.min.js")
}
func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/html")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/index.html")
}


func logoSinaps(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","image/png")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/sinaps_logo1.png")
}

func Medico(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","text/html")
	http.ServeFile(w,req,"www/startbootstrap-landing-page-gh-pages/medico.html")
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

func Paciente(w http.ResponseWriter, req *http.Request) {
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
