package routers

import (
	"fmt"
	"net/http"
	"html/template"
	//"github.com/wojciechsi/grzejemy.pl/models"
)

type cat struct {
	Name string
	Question string
}
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dorzuć do pieca!")
}

func offerHandler(w http.ResponseWriter, r *http.Request) {
	p := cat{Name: "Mr.Meow", Question: "I can has cheezburger?"}
	t, _ := template.ParseFiles("offers.html")
	t.Execute(w, p)
}

func RunServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homePageHandler))
	mux.Handle("/offer/", http.HandlerFunc(offerHandler))
	http.ListenAndServe(":8080", mux)
}
