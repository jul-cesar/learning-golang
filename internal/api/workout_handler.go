package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) GetWorkoutById(w http.ResponseWriter, r *http.Request) {

	paramsWorkoudId := chi.URLParam(r, "id")
	if paramsWorkoudId == "" {
		http.NotFound(w, r)
		return
	}

	workoudId, err := strconv.ParseInt(paramsWorkoudId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Workout id: %d\n", workoudId)

}


func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	// Handle the creation of a new workout
	fmt.Fprintf(w, "Create a new workout\n")
}
