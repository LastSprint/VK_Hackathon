package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"suncity/reps"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type AuthModel struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type AuthController struct {
	rep *reps.AuthRep
}

func InitAuthService(rep *reps.AuthRep, router *mux.Router) *AuthController {
	res := &AuthController{rep: rep}
	router.HandleFunc("/auth", res.Auth).Methods("POST")
	return res
}

func (contr *AuthController) Auth(w http.ResponseWriter, r *http.Request) {

	var user AuthModel

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println(user)

	res, err := contr.rep.GetUser(user.Email, user.Pass)

	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(err)
		return
	}

	claims := jwt.MapClaims{
		"id":   res.ID.Hex(),
		"type": res.UserType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println(token)

	tokenString, err := token.SignedString([]byte("mySigningKey"))

	err = contr.rep.RegToken(tokenString, res.ID)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
