package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.SQSEvent) error {
	if err := json.NewEncoder(os.Stdout).Encode(event); err != nil {
		return err
	}
	fmt.Println("record num: %d", len(event.Records))
	return nil
}

func main() {
	lambda.Start(handler)
}
