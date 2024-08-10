package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/seunghoon34/goapi/internal/handlers"
	log "github.com/siruspen/logrus"
)

func main() {
	log.SetReportCaller((true))
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("starting api service")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
