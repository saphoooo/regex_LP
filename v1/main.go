package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/wordfind", wordFinding).Methods("POST")

	fmt.Print("Start listening on :8000...")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func wordFinding(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("an error occurred: " + err.Error())
	}
	msg := message{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Println("an error occurred: " + err.Error())
	}
	result, err := regexp.MatchString(msg.Word, msg.Sentence)
	if err != nil {
		fmt.Println("an error occurred: " + err.Error())
	}

	if result {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "word found!")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "word not found...")
}

type message struct {
	Sentence string `json:"sentence"`
	Word     string `json:"word"`
}
