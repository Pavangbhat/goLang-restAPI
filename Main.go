package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavangbhat/gotest/controllers"
)



func main(){
	var router=mux.NewRouter()

	router.HandleFunc("/",controllers.GenerateGoals).Methods("GET")
	router.HandleFunc("/goals/create",controllers.CreateGoal).Methods("POST")
	router.HandleFunc("/goals",controllers.GetGoals).Methods("GET")
	router.HandleFunc("/goal/{id}",controllers.GetGoal).Methods("GET")
	router.HandleFunc("/goal/update/{id}",controllers.UpdateGoal).Methods("PATCH")
	router.HandleFunc("/goal/delete/{id}",controllers.DeleteGoal).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
	
}