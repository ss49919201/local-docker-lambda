package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type event struct {
	Name string `json:"name"`
}

func handleRequest(ctx context.Context, e event) (string, error) {
	_, err := fmt.Println("Event: ", e.Name)
	if err != nil {
		return "Failed", err
	}
	return "Success", nil
}

func main() {
	lambda.Start(handleRequest)
}
