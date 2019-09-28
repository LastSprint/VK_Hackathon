package main

import (
	"log"
	"net/http"
	"os"
	"suncity/auth"
	"suncity/commod"
	"suncity/feedback"
	"suncity/notifications"
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

	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("/static/"))),
	)

	router.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		notifications.SendNotification(nil, &commod.ServiceUser{Apns: "f217d876f98f78330ff3da4ac72adaf542defd5f4ab01ace1eeded41cb1a5a6b"})
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8844", handlers.LoggingHandler(os.Stdout, router)))
}
