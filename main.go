package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func printJson(out io.Writer, data interface{}) {
	if err := json.NewEncoder(out).Encode(data); err != nil {
		log.Println(err)
	}
}

func handler(ctx context.Context, event events.SQSEvent) error {
	s3Records := make([]events.S3EventRecord, 0, 10)
	for _, r := range event.Records {
		body := r.Body
		r.Body = ""
		printJson(os.Stdout, r)

		var record events.SNSEntity
		if err := json.Unmarshal([]byte(body), &record); err != nil {
			return fmt.Errorf("unmarshal sns entity error: %w", err)
		}
		printJson(os.Stdout, record)

		var records events.S3Event
		if err := json.Unmarshal([]byte(record.Message), &records); err != nil {
			return fmt.Errorf("unmarshal s3 event error: %w", err)
		}
		printJson(os.Stdout, records)

		s3Records = append(s3Records, records.Records...)
	}
	fmt.Printf("record num: %d", len(event.Records))
	fmt.Printf("s3 record num: %d", len(s3Records))
	return nil
}

func main() {
	lambda.Start(handler)
}
