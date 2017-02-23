package app

import (
	log "github.com/Sirupsen/logrus"
	"github.com/elbuo8/juggler/app/models"
	"os"
	"os/signal"
)

type App struct {
	DBSession *models.DBSession
	Logger    *log.Logger
}

func listenAndCloseGracefully(s *models.DBSession, logger *log.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			logger.Info(sig, " captured - Closing database connection")
			s.Close()
			os.Exit(1)
		}
	}()
}

func NewApp() (*App, error) {
	app := App{}
	app.Logger = log.New()
	db, err := models.NewDB()
	if db != nil {
		listenAndCloseGracefully(db, app.Logger)
	}
	app.DBSession = db
	return &app, err
}
