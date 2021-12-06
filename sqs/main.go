package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleSqsRequest)
}

func handleSqsRequest(sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Println(message.Body)
	}

	return nil
}
