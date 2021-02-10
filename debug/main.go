package main

import (
	"log"
	"os"

	service "github.com/g-wilson/runtime-helloworld/service"

	"github.com/g-wilson/runtime/auth"
	"github.com/g-wilson/runtime/devserver"
	"github.com/g-wilson/runtime/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	listenAddr := "127.0.0.1:" + os.Getenv("HTTP_PORT")
	log := logger.Create("debug", os.Getenv("LOG_FORMAT"), os.Getenv("LOG_LEVEL"))

	authn, err := auth.New(os.Getenv("OPENID_CONFIG_URL"))
	if err != nil {
		panic(err)
	}

	app, err := service.NewApp(log)
	if err != nil {
		panic(err)
	}

	svc := service.NewRPC(app)

	devserver.New(listenAddr, authn).
		AddService("helloworld", svc).
		Listen()
}
