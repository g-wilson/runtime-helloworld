package main

import (
	app "github.com/g-wilson/runtime-helloworld/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	appInstance, err := app.New()
	if err != nil {
		panic(err)
	}

	lambda.Start(appInstance.RPCService().WrapAPIGatewayHTTP())
}
