package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jul-cesar/learning-golang/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/workout/{id}", app.WorkoutHandler.GetWorkoutById)
	r.Post("/workout", app.WorkoutHandler.HandleCreateWorkout)
	return r
}
