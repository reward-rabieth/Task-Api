package server

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.con/reward-rabieth/Task-Api/core/components/Reminder/models"
	"net/http"
	"strconv"
	"time"
)

type CreateReminderRequest struct {
	ContactName  string    `json:"contactName"`
	ReminderTime time.Time `json:"reminderTime"`
}

func (app *App) CreateReminderHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and validate request body
	var reqBody CreateReminderRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdReminder, err := app.ReminderRepo.CreateReminder(r.Context(), &models.Reminder{
		ReminderTime: reqBody.ReminderTime,
		ContactName:  reqBody.ContactName,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdReminder)
}

func (app *App) GetReminderHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reminder, err := app.ReminderRepo.GetReminder(r.Context(), uint(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reminder)
}

func (app *App) UpdateReminderHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reminder models.Reminder
	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := app.ReminderRepo.UpdateReminder(r.Context(), uint(idInt), &reminder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) DeleteReminderHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := app.ReminderRepo.DeleteReminder(r.Context(), uint(idInt)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
