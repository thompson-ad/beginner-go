package app

import (
	"beginnerGo/internal/api"
	"beginnerGo/internal/store"
	"beginnerGo/migrations"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	Logger         *log.Logger
	DB             *sql.DB
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := store.Open()
	if err != nil {
		return nil, err
	}
	migrateErr := store.MigrateFS(db, migrations.FS, ".")
	if migrateErr != nil {
		panic(migrateErr)
	}

	workoutStore := store.NewPostgresWorkoutStore(db)

	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)

	app := &App{
		Logger:         logger,
		DB:             db,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

// r is a pointer because we want to persist the data that they send rather than what we respond with
func (app *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
