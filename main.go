package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type event struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, e event) error {
	fmt.Printf("Event: %s", e.Name)
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
