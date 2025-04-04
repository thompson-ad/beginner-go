package api

import (
	"beginnerGo/internal/store"
	"beginnerGo/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore // api only knows about the interface
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutId := chi.URLParam(r, "id")

	if paramsWorkoutId == "" {
		wh.logger.Printf("ERROR: No workout ID provided")
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "No workout ID provided"})
		return
	}

	workoutId, err := strconv.ParseInt(paramsWorkoutId, 10, 64)

	if err != nil {
		wh.logger.Printf("ERROR: Invalid workout ID: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid workout ID"})
		return
	}

	workout, err := wh.workoutStore.GetWorkoutByID(workoutId)

	if err != nil {
		wh.logger.Printf("ERROR: Failed to get the workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to get the workout"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": workout})
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: Failed to decode workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to decode workout"})
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)

	if err != nil {
		wh.logger.Printf("ERROR: Failed to create workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to create workout"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"workout": createdWorkout})
}

func (wh *WorkoutHandler) HandleUpdateWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutId := chi.URLParam(r, "id")

	if paramsWorkoutId == "" {
		wh.logger.Printf("ERROR: No workout ID provided")
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "No workout ID provided"})
		return
	}

	workoutId, err := strconv.ParseInt(paramsWorkoutId, 10, 64)

	if err != nil {
		wh.logger.Printf("ERROR: Invalid workout ID: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid workout ID"})
		return
	}

	existingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutId)

	if err != nil {
		wh.logger.Printf("ERROR: Failed to get the workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to get the workout"})
		return
	}

	if existingWorkout == nil {
		wh.logger.Printf("ERROR: Workout not found")
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "Workout not found"})
		return
	}

	var updateWorkoutRequest struct {
		Title           *string              `json:"title"`
		Description     *string              `json:"description"`
		DurationMinutes *int                 `json:"duration_minutes"`
		CaloriesBurned  *int                 `json:"calories_burned"`
		Entries         []store.WorkoutEntry `json:"entries"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)

	if err != nil {
		wh.logger.Printf("ERROR: Failed to decode workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to decode workout"})
		return
	}

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}

	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}

	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}

	if updateWorkoutRequest.CaloriesBurned != nil {
		existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}

	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}

	err = wh.workoutStore.UpdateWorkout(existingWorkout)

	if err != nil {
		wh.logger.Printf("ERROR: Failed to update workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to update workout"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": existingWorkout})
}
