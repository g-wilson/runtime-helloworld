package app

import (
	"context"
	"fmt"
)

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Message string `json:"message"`
}

func (a *App) Greet(ctx context.Context, req *GreetRequest) (res *GreetResponse, err error) {
	res = &GreetResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}

	return
}
