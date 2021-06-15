package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tmp string

type Data struct {
	Link string `json:"link" bson:"link"`
}

func ControllerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	data.Link = ""
	json.NewEncoder(w).Encode(dataControl(1, data))
}

func ControllerSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	json.NewEncoder(w).Encode(dataControl(2, data))

}

func dataControl(check int, data Data) map[string]interface{} {
	response := make(map[string]interface{})
	if check == 1 {
		response["link"] = tmp
	}
	if check == 2 {
		tmp = data.Link
		response["updated"] = "Yes!"
		response["link"] = tmp
	}
	return response
}

func main() {
	tmp = "http://amssamples.streaming.mediaservices.windows.net/91492735-c523-432b-ba01-faba6c2206a2/AzureMediaServicesPromo.ism/manifest"
	port := os.Getenv("PORT")
	//port := "8080"
	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{
		"X-Requested-With", "Accept", "Authorization", "Content-Type", "X-CSRF-Token",
	})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	api := r.PathPrefix("/api/video").Subrouter()
	api.HandleFunc("/get", ControllerGet).Methods("GET")
	api.HandleFunc("/set", ControllerSet).Methods("POST")

	_ = http.ListenAndServe(":"+port, handlers.CORS(headers, origins, methods)(r))
}
