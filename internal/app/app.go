package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jul-cesar/learning-golang/internal/api"
	"github.com/jul-cesar/learning-golang/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	Db 				*sql.DB
}

func NewApplication() (*Application, error) {
	pgDb, err :=  store.Open()
	
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	} 
	
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// Our stores will be here

	// our handlers will go here

	workoutHandler := api.NewWorkoutHandler()
	
	app := &Application{
		Logger:          logger,
		WorkoutHandler: workoutHandler,
		Db: pgDb,
	}
	return app, nil
}
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available and running")
}
