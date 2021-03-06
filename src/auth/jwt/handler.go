package jwt

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/auth/sign_in", SignIn).
		Methods("GET").
		Schemes("http")
	router.HandleFunc("/auth/refresh", Refresh).
		Methods("GET").
		Schemes("http")
	router.HandleFunc("/auth/sign_out", SignOut).
		Methods("GET").
		Schemes("http")

	return router
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	AuthenticateSignIn(w, r)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	RefreshToken(w, r)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	InvalidateToken(w)
}