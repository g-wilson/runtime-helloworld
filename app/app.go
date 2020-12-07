package app

import (
	"os"

	"github.com/g-wilson/runtime/logger"
	"github.com/g-wilson/runtime/rpcservice"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger *logrus.Entry
}

func New() (*App, error) {
	appLogger := logger.Create("helloworld", os.Getenv("LOG_FORMAT"), os.Getenv("LOG_LEVEL"))

	// create any dependencies etc here and add them to the app instance

	// for example...
	// awsConfig := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION"))
	// awsSession := session.Must(session.NewSession())

	return &App{
		Logger: appLogger,
	}, nil
}

func (a *App) Service() *rpcservice.Service {
	return rpcservice.NewService(a.Logger).
		AddMethod("greet", a.Greet, GreetRequestSchema)
}
