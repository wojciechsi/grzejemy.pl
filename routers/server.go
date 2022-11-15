package routers

import (
	"fmt"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dorzuć do pieca!")
}

func RunServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homePageHandler))
	http.ListenAndServe(":8080", mux)
}
