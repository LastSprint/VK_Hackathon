package reg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"suncity/commod"
	rm "suncity/reg/models"
	"suncity/reps"

	"github.com/gorilla/mux"
)

// Controller контроллер для регистрации
type Controller struct {
	// Контекст базы данных
	rep    *reps.RegRep
	router *mux.Router
}

// InitRegController Инициаллизирует контроллер
func InitRegController(rep *reps.RegRep, router *mux.Router) *Controller {
	res := &Controller{rep: rep, router: router}
	router.HandleFunc("/reg/form", res.postRegForm).Methods("POST")
	router.HandleFunc("/reg/status", res.checkStatus).Methods("POST")
	return res
}

func (reg *Controller) postRegForm(w http.ResponseWriter, r *http.Request) {

	var form rm.FormModel

	fmt.Println(r.Body)

	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&err)
		return
	}

	if !form.IsValid() {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(commod.ErrorModel{Msg: "Неправильно заполнены поля формы"})
		return
	}

	fmt.Println(form)

	err = reg.rep.PostForm(form)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&err)
		return
	}

	w.WriteHeader(204)
	return
}

func (reg *Controller) checkStatus(w http.ResponseWriter, r *http.Request) {
	var model rm.CheckStatusRequestModel

	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil || !model.IsValid() {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&err)
		return
	}

	res, err := reg.rep.CheckFormStatus(*model.Value)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&err)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}
