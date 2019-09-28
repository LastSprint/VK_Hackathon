package feedback

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"suncity/reps"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type MessageModel struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type CreateFeedbackModel struct {
	Filepath *string `json:"аilepath"`
	Text     string  `json:"text"`
}

type FeedbackDockModel struct {
	ID string `json:"id"`

	Filepath *string `json:"аilepath"`
	Text     string  `json:"text"`

	Messages *[]MessageModel `json:"text"`
}

type FeedbackController struct {
	rep *reps.FeedbackRepo
}

func InitFeedbackController(router *mux.Router, rep *reps.FeedbackRepo) *FeedbackController {
	controller := &FeedbackController{}

	router.HandleFunc("/feedback/comment/{id}", controller.commentFeedback).Methods("POST")
	router.HandleFunc("/feedback/comment", controller.createNewPos).Methods("POST")
	router.HandleFunc("/feedback/comment", controller.getAllPosts).Methods("GET")

	return controller
}

func (contrl *FeedbackController) commentFeedback(w http.ResponseWriter, r *http.Request) {

	var comment MessageModel

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		log.Println(400)
		log.Println(err)
		return
	}

	id := mux.Vars(r)["id"]

	err = contrl.rep.AddComment(&comment, id)

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(204)

	return
}

func (contrl *FeedbackController) createNewPos(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("audio")
	if err != nil {
		log.Println(err)
		return
	}

	ad, err := ksuid.NewRandom()

	if err != nil {
		w.WriteHeader(500)
		return
	}

	filepath := "/static/audio/" + ad.String()

	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		filepath = ""
	} else {
		io.Copy(f, file)
		f.Close()
	}

	post := CreateFeedbackModel{Filepath: &filepath, Text: r.FormValue("text")}

	err = contrl.rep.CreatePost(&post)

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(204)
}

func (contrl *FeedbackController) getAllPosts(w http.ResponseWriter, r *http.Request) {
	data, err := contrl.rep.GetAllPosts()

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	return
}
