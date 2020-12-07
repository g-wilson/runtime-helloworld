package app

import (
	"context"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Message string `json:"message"`
}

var GreetRequestSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,
	"required": [ "name" ],
	"properties": {
		"name": {
			"type": "string",
			"minLength": 1
		}
	}
}`)

func (a *App) Greet(ctx context.Context, req *GreetRequest) (res *GreetResponse, err error) {
	res = &GreetResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}

	return
}
