package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"

	"context"
	"encoding/json"
	"log"
)

func main() {
	// Register handler function in the main function using the AWS Lambda for Go library.
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := SqsTriggerMessage{
		Message: "You can add different fields here according to your data requirements in the background task.",
	}

	err := sendMessage(message)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "Message Sent", StatusCode: 200}, nil
}

type SqsTriggerMessage struct {
	Message string
}

func sendMessage(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	svc := sqs.New(
		session.Must(
			session.NewSession(
				&aws.Config{
					Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
					Region:      aws.String(os.Getenv("AWS_REGION")),
				},
			),
		),
	)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("queue-name"),
	})
	if err != nil {
		return err
	}

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:            aws.String(string(b)),
		QueueUrl:               result.QueueUrl,
		MessageGroupId:         aws.String("group-id"),
		MessageDeduplicationId: aws.String("deduplication-id"),
	})

	return err
}
