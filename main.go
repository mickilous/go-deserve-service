package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	Deserve string `json:"deserve"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/deserve/{userId}", func(writer http.ResponseWriter, request *http.Request) {
		userId := mux.Vars(request)["userId"]
		log.Printf("Does user %v deserve greetings ?", userId)
		ret := false
		if userId != "666" {
			ret = true
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(Message{Deserve: strconv.FormatBool(ret)})
	})
	http.ListenAndServe(":8090", router)
}
