package main

import (
	"log"
	"net/http"
	"suncity/reg"
	"suncity/reps"

	"github.com/gorilla/mux"
)

var regController *reg.Controller

var cntx *reps.DBContext

func main() {
	cnt, err := reps.NewDB()

	if err != nil {
		panic(err)
	}

	cntx = cnt
	router := mux.NewRouter()
	regController = reg.InitRegController(reps.InitRegRep(cnt), router)
	log.Fatal(http.ListenAndServe(":8844", router))
}
