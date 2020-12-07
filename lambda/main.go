package lambda

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type EntrypointFn func() (interface{}, error)

var Functions = map[string]EntrypointFn{
	"api":    APIEntrypoint,
	"worker": SQSEntrypoint,
}

func main() {
	err := (func() error {

		entrypoint := os.Getenv("LAMBDA_ENTRYPOINT")
		if entrypoint == "" {
			return fmt.Errorf("no entrypoint defined, LAMBDA_ENTRYPOINT=%s", entrypoint)
		}

		entrypointFn, ok := Functions[entrypoint]
		if !ok {
			return fmt.Errorf("entrypoint %s not found", entrypoint)
		}

		handler, err := entrypointFn()
		if err != nil {
			return err
		}

		lambda.Start(handler)

		return nil
	})()
	if err != nil {
		log.Fatal(fmt.Errorf("error: %w", err))
	}
}
