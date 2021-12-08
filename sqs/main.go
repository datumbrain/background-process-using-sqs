package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	lambda.Start(handleSqsRequest)
}

func handleSqsRequest(sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		var request SqsTriggerMessage
		err := json.Unmarshal([]byte(message.Body),&request)
		if err!=nil {
			log.Println(err)
			continue
		}
		fmt.Println(request.Message)
	}

	return nil
}

type SqsTriggerMessage struct {
	Message string
}
