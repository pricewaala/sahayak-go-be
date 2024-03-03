package router

import (
	"github.com/anujsharma13/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/workers", controller.GetMyAllWorkers).Methods("GET")
	router.HandleFunc("/api/worker", controller.CreateWorker).Methods("POST")
	router.HandleFunc("/api/worker/{id}", controller.UpdateWorker).Methods("PUT")
	router.HandleFunc("/api/worker/{id}", controller.DeleteOneWorker).Methods("DELETE")
	return router
}
