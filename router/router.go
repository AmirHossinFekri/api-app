package router

import "net/http"


func New()http.Handler{
	mux:= http.NewServeMux()

	// mux.HandleFunc("GET /users/{id}", userH.Get)
	// mux.HandleFunc("POST /users", userH.Create)

	// mux.HandleFunc("GET /courses/{id}", courseH.Get)
	// mux.HandleFunc("POST /courses/{id}/enroll", courseH.Enroll)

	return mux

}