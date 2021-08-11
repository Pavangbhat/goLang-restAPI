package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pavangbhat/gotest/models"
)


var goals []models.Goals

// controllers
func GenerateGoals(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	goals = append(goals, models.Goals{Id:1,Title: "Read on event loop",Status: "completed"})
	goals = append(goals, models.Goals{Id:2,Title: "read just javaScript articles",Status: "active"})
	goals = append(goals, models.Goals{Id:3,Title: "learn go",Status: "closed"})

	var status = map[string]string{"message":"Succesfully generated a goals"}
	json.NewEncoder(w).Encode(status)
}

func CreateGoal(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var goal models.Goals
	_ =json.NewDecoder(r.Body).Decode(&goal)
	goals=append(goals,goal)
	var status = map[string]string{"message":"Succesfully created a goal"}
	json.NewEncoder(w).Encode(status)
}

func GetGoals(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	if len(goals) ==0 {
		var status = map[string]string{"message":"Generate goals by hiting / route"}
		json.NewEncoder(w).Encode(status)
	}else {
		json.NewEncoder(w).Encode(goals)
	}
}

func GetGoal(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,goal:= range goals{
		if strconv.Itoa(goal.Id) == params["id"] {
			json.NewEncoder(w).Encode(goal)
			return
		}
	}
	http.Error(w,"No goal found",http. StatusBadRequest)
}

func UpdateGoal(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,goal:= range goals{
		if strconv.Itoa(goal.Id) == params["id"] {
			goals = append(goals[:index], goals[index+1:]...) //remove a goal from map
			var goal models.Goals
			_ =json.NewDecoder(r.Body).Decode(&goal)
			goals=append(goals,goal)
			var status = map[string]string{"message":"Succesfully updated a goal"}
			json.NewEncoder(w).Encode(status)
			return
		}
	}
	http.Error(w,"No goal found",http. StatusBadRequest)
}

func DeleteGoal(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,goal:= range goals{
		if strconv.Itoa(goal.Id) == params["id"] {
			goals = append(goals[:index], goals[index+1:]...)
			var status = map[string]string{"message":"Succesfully deleted a goal"}
			json.NewEncoder(w).Encode(status)
			return
		}
	}
	http.Error(w,"No goal found",http. StatusBadRequest)
}
