package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *App) NewHandler() http.Handler {
	r := chi.NewRouter()

	//tasks endpoints
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", app.CreateTaskHandler)
		r.Get("/{id}", app.GetTaskHandler)
		r.Put("/{id}", app.UpdateTaskHandler)
		r.Delete("{id}", app.DeleteTaskHandler)

	})

	//reminder endpoints
	r.Route("/reminders", func(r chi.Router) {
		r.Post("/reminders", app.CreateReminderHandler)
		r.Get("/reminders/{id}", app.GetReminderHandler)
		r.Put("/reminders/{id}", app.UpdateReminderHandler)
		r.Delete("/reminders/{id}", app.DeleteReminderHandler)
	})

	return r
}
