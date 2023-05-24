package main

import (
	"log"
	"net/http"

	"example.com/controller"
	"github.com/gorilla/mux"
)

func main() {
    

   

    // Define routes
	router := mux.NewRouter()
	router.HandleFunc("/getproject", controller.AllProject).Methods("GET")
	router.HandleFunc("/insertproject", controller.InsertProject).Methods("POST")
	http.Handle("/", router)
   

    // Start server
    log.Println("Server started on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
