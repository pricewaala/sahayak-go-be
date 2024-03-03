package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anujsharma13/router"
)

func main() {
	fmt.Println("Mongo Api")
	fmt.Println("Server getting started on port 4000")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4002", r))

}
