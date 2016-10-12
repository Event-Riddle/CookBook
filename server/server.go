package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PathPrefix = "/api/v1/"

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix+"config/{type}", errorHandler(GetConfig)).Methods("GET")
	r.HandleFunc(PathPrefix+"config/{type}", errorHandler(SaveConfig)).Methods("POST")

	middleware := NewMiddleware(r)
	middleware.PreProcessing()

	http.Handle(PathPrefix, middleware)
}

type badRequest struct{ error }

type notFound struct{ error }

type readError struct{ error }

func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func GetConfig(w http.ResponseWriter, r *http.Request) error {
	configType, _ := mux.Vars(r)["type"]
	file, err := ioutil.ReadFile("server/configs/" + configType + ".json")
	if err != nil {
		return readError{err}
	}
	var c = GetConfigStruct(configType)
	json.Unmarshal(file, &c)
	return json.NewEncoder(w).Encode(c)
}

func SaveConfig(w http.ResponseWriter, r *http.Request) error {
	configType := mux.Vars(r)["type"]
	configStruct := GetConfigStruct(configType)
	if err := json.NewDecoder(r.Body).Decode(&configStruct); err != nil {
		return badRequest{err}
	}
	return writeJsonToFile(configStruct, configType)
}

func GetConfigStruct(ct string) interface{} {
	switch ct {
	case "filter":
		return []Filter{}
	case "user":
		return User{}
	}
	return nil
}

func writeJsonToFile(configStruct interface{}, configType string) error {
	jsonFile, _ := json.Marshal(configStruct)
	return ioutil.WriteFile("server/configs/"+configType+".json", jsonFile, 0777)
}
