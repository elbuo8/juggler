package app

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/elbuo8/juggler/app/models"
	"os"
	"os/signal"
)

var REGIONS = []string{"us-east-1"}

type App struct {
	AWSSessions map[string]*session.Session
	DBSession   *models.DBSession
	Logger      *log.Logger
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
	app.AWSSessions = make(map[string]*session.Session)
	for _, region := range REGIONS {
		// https://github.com/aws/aws-sdk-go/blob/master/aws/session/session.go#L328
		val, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
		app.AWSSessions[region] = val
	}
	app.Logger = log.New()
	db, err := models.NewDB()
	if db != nil {
		listenAndCloseGracefully(db, app.Logger)
	}
	app.DBSession = db
	return &app, err
}
