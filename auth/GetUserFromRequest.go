package auth

import (
	"net/http"
	"suncity/commod"
	"suncity/reps"

	"github.com/sirupsen/logrus"
)

var authRepo *reps.AuthRep

func Init(rep *reps.AuthRep) {
	authRepo = rep
}

func GetUserFromRequest(r *http.Request) (*commod.ServiceUser, error) {

	if len(r.Header["Authorization"]) < 1 {
		return nil, nil
	}

	token := r.Header["Authorization"][0]

	return authRepo.GetUserByToken(token)
}

func AuthHandler(next func(w http.ResponseWriter, r *http.Request, user *commod.ServiceUser)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		user, err := GetUserFromRequest(r)

		if err != nil {
			logrus.Error(err)
			w.WriteHeader(401)
			return
		}

		if user == nil {
			logrus.Error("User not found")
			w.WriteHeader(401)
			return
		}

		next(w, r, user)
	}
}
