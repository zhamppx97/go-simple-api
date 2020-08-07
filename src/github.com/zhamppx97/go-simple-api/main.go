package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	UserID    int
	Firstname string
	Lastname  string
}

func getUserAll(w http.ResponseWriter, r *http.Request) {
	addUser := user{
		UserID:    1,
		Firstname: "Steve",
		Lastname:  "Rogers",
	}
	json.NewEncoder(w).Encode(addUser)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GO Simple API")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getuser", getUserAll)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
