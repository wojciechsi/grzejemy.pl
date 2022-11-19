package routers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/wojciechsi/grzejemy.pl/db"
	"github.com/wojciechsi/grzejemy.pl/models"
)

type offersList struct{
	OList [] models.Offer
}
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dorzuć do pieca!")
}

var templates = template.Must(template.ParseFiles("routers/offers.html"))

func offerHandler(w http.ResponseWriter, r *http.Request) {
	o:= offersList{ db.GetTestOffer()}
	err := templates.ExecuteTemplate(w, "offers.html", o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RunServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homePageHandler))
	mux.Handle("/offer/", http.HandlerFunc(offerHandler))
	http.ListenAndServe(":8080", mux)
}
