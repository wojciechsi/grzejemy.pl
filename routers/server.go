package routers

import (
	"fmt"
	"html/template"
	"net/http"
	//"github.com/wojciechsi/grzejemy.pl/models"
)

type cat struct {
	Name     string
	Question string
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dorzuć do pieca!")
}

var templates = template.Must(template.ParseFiles("routers/offers.html"))

func offerHandler(w http.ResponseWriter, r *http.Request) {
	p := cat{Name: "Mr.Meow", Question: "I can has cheezburger?"}
	templates.ExecuteTemplate(w, "offers.html", p)
}

func RunServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homePageHandler))
	mux.Handle("/offer/", http.HandlerFunc(offerHandler))
	http.ListenAndServe(":8080", mux)
}
