package feedback

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"suncity/auth"
	"suncity/commod"
	"suncity/reps"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type FeedbackController struct {
	rep *reps.FeedbackRepo
}

func InitFeedbackController(rep *reps.FeedbackRepo, router *mux.Router) *FeedbackController {
	controller := &FeedbackController{rep: rep}

	router.HandleFunc("/feedback/comment/{id}", auth.AuthHandler(controller.commentFeedback)).Methods("POST")
	router.HandleFunc("/feedback/comment", auth.AuthHandler(controller.createNewPos)).Methods("POST")
	router.HandleFunc("/feedback", auth.AuthHandler(controller.getAllPosts)).Methods("GET")

	return controller
}

func (contrl *FeedbackController) commentFeedback(w http.ResponseWriter, r *http.Request, user *commod.ServiceUser) {

	var comment reps.MessageModel

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		log.Println(400)
		log.Println(err)
		return
	}

	id := mux.Vars(r)["id"]

	comment.Name = user.Name
	comment.Date = time.Time{}
	comment.Image = user.Image
	comment.UserID = user.ID

	err = contrl.rep.AddComment(&comment, id)

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(204)

	return
}

func (contrl *FeedbackController) createNewPos(w http.ResponseWriter, r *http.Request, user *commod.ServiceUser) {
	r.ParseMultipartForm(100 << 29)

	ad, err := ksuid.NewRandom()

	audioPath := "/static/audio/" + ad.String() + ".m4a"

	err = saveAudio(r, audioPath)

	if err != nil {
		audioPath = ""
		log.Println(err)
	}

	images, err := savePhotos(r, "/static/photos/")

	fmt.Println()

	post := reps.CreateFeedbackModel{
		UserId: user.ID,
		Audio:  audioPath,
		Images: images,
		Text:   r.FormValue("text"),
	}

	fmt.Println(contrl)
	err = contrl.rep.CreatePost(&post)

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.WriteHeader(204)
}

func (contrl *FeedbackController) getAllPosts(w http.ResponseWriter, r *http.Request, user *commod.ServiceUser) {
	data, err := contrl.rep.GetAllPosts()

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	js := map[string]interface{}{
		"user":    user,
		"payload": data,
	}

	err = json.NewEncoder(w).Encode(js)

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	return
}

func saveAudio(r *http.Request, filepath string) error {

	file, _, err := r.FormFile("audio")

	if err != nil {
		return err
	}

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	io.Copy(f, file)
	f.Close()
	file.Close()
	return nil
}

func savePhotos(r *http.Request, dirpath string) ([]string, error) {

	prefix := "photo"

	result := []string{}

	for i := 0; ; i++ {

		name := prefix + strconv.Itoa(i)

		file, _, err := r.FormFile(name)

		if err != nil {
			fmt.Println(err)
			break
		}

		resultName := dirpath + name + ".jpg"

		f, err := os.OpenFile(resultName, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			break
		}

		result = append(result, resultName)

		io.Copy(f, file)
		f.Close()
		file.Close()
	}

	return result, nil
}
