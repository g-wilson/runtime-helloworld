package service

import (
	"github.com/g-wilson/runtime/rpcservice"
	"github.com/sirupsen/logrus"
)

// App keeps our dependencies and is the base for methods
type App struct {
	Logger *logrus.Entry
}

// New creates a new app
func NewApp(log *logrus.Entry) (*App, error) {
	return &App{
		Logger: log,
	}, nil
}

// NewRPC returns a runtime RPCService for interacting with the service
func NewRPC(a *App) *rpcservice.Service {
	return rpcservice.NewService(a.Logger).
		WithIdentityProvider(SetIdentity).
		WithContextProvider(IdentityLogger).
		AddMethod("greet", a.Greet, GreetRequestSchema)
}
