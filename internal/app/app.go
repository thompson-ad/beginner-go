package app

import (
	"beginnerGo/internal/api"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Stores here
	// Handlers here
	workoutHandler := api.NewWorkoutHandler()

	app := &App{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

// r is a pointer because we want to persist the data that they send rather than what we respond with
func (app *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
