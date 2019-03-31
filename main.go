package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

//Request http request
type Request struct {
	Param string `json:"param"`
}

//Handle lambda hander
func Handle(ctx context.Context, param Request) (string, error) {
	return "OK-go", nil
}

func main() {
	lambda.Start(Handle)
}
