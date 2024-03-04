package server

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.con/reward-rabieth/Task-Api/core/components/Reminder/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateReminderHandler(t *testing.T) {
	app := &App{}
	r := chi.NewRouter()
	r.Post("/reminders", app.CreateReminderHandler)

	reminder := models.Reminder{
		ReminderTime: time.Now(),
		ContactName:  "John Doe",
	}
	reminderJSON, _ := json.Marshal(reminder)
	req, err := http.NewRequest("POST", "/reminders", bytes.NewBuffer(reminderJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

}
