package server

import (
	"github.con/reward-rabieth/Task-Api/config"
	"github.con/reward-rabieth/Task-Api/core/components/Reminder"
	"github.con/reward-rabieth/Task-Api/core/components/Task"
	"github.con/reward-rabieth/Task-Api/db"
	"gorm.io/gorm"
	"log/slog"
)

type App struct {
	DB                *gorm.DB
	shutdownCallbacks []func()

	// logger
	logger *slog.Logger

	TaskRepo     Task.Repo
	ReminderRepo Reminder.Repo
}

func NewApp() (app *App, err error) {
	app = &App{}
	app.logger = slog.Default()
	if app.DB, err = db.Connect(app.logger, config.GetDatabaseConfig()); err != nil {
		return
	}

	app.shutdownCallbacks = []func(){}

	//repositories initialization
	if app.TaskRepo, err = Task.NewRepo(app.DB); err != nil {
		return app, err
	}
	if app.ReminderRepo, err = Reminder.NewRepo(app.DB); err != nil {
		return app, err
	}
	return app, nil
}

func (app *App) Shutdown() {
	for _, callback := range app.shutdownCallbacks {
		callback()
	}
}
