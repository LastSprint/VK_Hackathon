package main

import (
	"log"
	"net/http"
	"os"
	"suncity/auth"
	"suncity/feedback"
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
var feedbackController *feedback.FeedbackController

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
	feedbackController = feedback.InitFeedbackController(reps.InitFeedbackRep(cntx), router)

	auth.Init(reps.InitAuthRep(cnt))

	// chat.StartChat(router)

	log.Fatal(http.ListenAndServe(":8844", handlers.LoggingHandler(os.Stdout, router)))
}
