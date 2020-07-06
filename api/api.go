package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serkandemirel0420/fintech/users"
)

//Login login
type Login struct {
	Username string
	Password string
}

//ErrResponse err
type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	login := users.Login(r.FormValue("Username"), r.FormValue("Password"))

	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

//StartAPI start server
func StartAPI() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}
