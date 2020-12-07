package lambda

import (
	app "github.com/g-wilson/runtime-helloworld/app"
)

func SQSEntrypoint() (interface{}, error) {
	a, err := app.New()
	if err != nil {
		return nil, err
	}

	return a.Service().WrapSQSHandler("greet"), nil
}
