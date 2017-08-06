package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"html/template"
	"os"
	"path/filepath"
)

type Request struct {
	Name    string
	Body    string
	Created time.Time
}

var (
	appPath, _ = os.Getwd()
	templatePath = filepath.Join(appPath, "templates")
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(templatePath, "home.html"))
	if err != nil {
		w.WriteHeader(500)
		println(err)
		return
	}
	t.Execute(w, nil)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/create", CreateHandler)
	r.HandleFunc("view", ViewHandler)
    http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
