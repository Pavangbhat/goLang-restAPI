package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pavangbhat/gotest/controllers"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string {

  // SetConfigFile explicitly defines the path, name and extension of the config file.
  // Viper will use this and not check any of the config paths.
  // .env - It will search for the .env file in the current directory
  viper.SetConfigFile(".env")

  // Find and read the config file
  err := viper.ReadInConfig()

  if err != nil {
    log.Fatalf("Error while reading config file %s", err)
  }

  // viper.Get() returns an empty interface{}
  // to get the underlying type of the key,
  // we have to do the type assertion, we know the underlying value is string
  // if we type assert to other type it will throw an error
  value, ok := viper.Get(key).(string)

  // If the type is a string then ok will be true
  // ok will make sure the program not break
  if !ok {
    log.Fatalf("Invalid type assertion")
  }

  return value
}

func main(){
	// viper package read .env
  viperenv := viperEnvVariable("PRIVATE_KEY")

  fmt.Printf("viper : %s = %s \n", "PRIVATE_KEY", viperenv)

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	println(os.Getenv("PRIVATE_KEY"))

	var router=mux.NewRouter()

	router.HandleFunc("/",controllers.GenerateGoals).Methods("GET")
	router.HandleFunc("/goals/create",controllers.CreateGoal).Methods("POST")
	router.HandleFunc("/goals",controllers.GetGoals).Methods("GET")
	router.HandleFunc("/goal/{id}",controllers.GetGoal).Methods("GET")
	router.HandleFunc("/goal/update/{id}",controllers.UpdateGoal).Methods("PATCH")
	router.HandleFunc("/goal/delete/{id}",controllers.DeleteGoal).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:3000", router))
	
}