package main

import (
	"log"
	"net/http"
	"os"
	"suncity/auth"
	"suncity/reg"
	"suncity/reps"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type AuthModel struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

var regController *reg.Controller
var authController *auth.AuthController

var cntx *reps.DBContext

func main() {
	cnt, err := reps.NewDB()

	if err != nil {
		panic(err)
	}

	cntx = cnt
	router := mux.NewRouter()
	regController = reg.InitRegController(reps.InitRegRep(cnt), router)
	authController = auth.InitAuthService(reps.InitAuthRep(cnt), router)

	log.Fatal(http.ListenAndServe(":8844", handlers.LoggingHandler(os.Stdout, router)))
}
