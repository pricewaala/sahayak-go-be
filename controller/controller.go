package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anujsharma13/model"
	"github.com/anujsharma13/repository"
	"github.com/gorilla/mux"
)

func GetMyAllWorkers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allWorkers := repository.GetAllWorkers()
	json.NewEncoder(w).Encode(allWorkers)
}

func CreateWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var worker model.Workers
	_ = json.NewDecoder(r.Body).Decode(&worker)
	fmt.Println(worker)
	repository.InsertOneWorker(worker)
	json.NewEncoder(w).Encode(worker)
}

func UpdateWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	repository.UpdateOneWorker(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	repository.DeleteOneWorker(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
