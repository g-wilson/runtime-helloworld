package lambda

import (
	"os"

	service "github.com/g-wilson/runtime-helloworld/service"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-wilson/runtime/logger"
)

func main() {
	log := logger.Create("my-service", os.Getenv("LOG_FORMAT"), os.Getenv("LOG_LEVEL"))

	// create any dependencies etc here and add them to the app instance
	// for example...
	// awsConfig := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION"))
	// awsSession := session.Must(session.NewSession())

	app, err := service.NewApp(log)
	if err != nil {
		panic(err)
	}

	svc := service.NewRPC(app)

	lambda.Start(svc.WrapAPIGatewayHTTP())
}
