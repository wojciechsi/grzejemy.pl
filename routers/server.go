package routers

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/wojciechsi/grzejemy.pl/models"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dorzuć do pieca!")
}

func offerHandler(w http.ResponseWriter, r *http.Request) {
	
//	jacuś := models.NewBuyer("Jacek")
	jacek := models.NewVendor("dr Szedel")
	jacex := models.NewSalesPoint("Gliwice", jacek)

	taniosprzedam := models.NewOffer(jacex, models.FuelType{Name: "groszek"}, 3560.75)

	t, _ := template.ParseFiles("offers.html")
	t.Execute(w, taniosprzedam)
}

func RunServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homePageHandler))
	http.ListenAndServe(":8080", mux)
	mux.Handle("/offer", http.HandlerFunc(offerHandler))
}
