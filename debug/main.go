package main

import (
	"log"
	"os"

	app "github.com/g-wilson/runtime-helloworld/app"

	"github.com/g-wilson/runtime/devserver"
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

	server := devserver.New(listenAddr)

	helloworldApp, err := app.New()
	if err != nil {
		panic(err)
	}

	server.AddService("helloworld", helloworldApp.RPCService(), nil)

	server.Listen()
}
