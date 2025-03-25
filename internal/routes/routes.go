package routes

import (
	"beginnerGo/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workout/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workout", app.WorkoutHandler.HandleCreateWorkout)

	return r
}
