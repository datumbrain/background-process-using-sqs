package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

type ProvisioningTriggerMethod string

const (
	TriggerMethodCreate ProvisioningTriggerMethod = "CREATE"
	TriggerMethodUpdate ProvisioningTriggerMethod = "UPDATE"
	TriggerMethodDelete ProvisioningTriggerMethod = "DELETE"
)

type SqsTriggerMessage struct {
	RegistrationId string
	Name           string
	Method         ProvisioningTriggerMethod
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sqsRequest := SqsTriggerMessage{
		RegistrationId: "ABC",
		Name:           "Hamza",
		Method:         TriggerMethodCreate,
	}
	b, _ := json.Marshal(sqsRequest)

	fmt.Println(string(b))

	svc := sqs.New(
		session.Must(
			session.NewSession(
				&aws.Config{
					Credentials: credentials.NewStaticCredentials("aws-access-key", "aws-secret-key", ""),
					Region:      aws.String("us-east-1"),
				},
			),
		),
	)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("queue-name"),
	})
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	_, err = svc.SendMessage(
		&sqs.SendMessageInput{
			MessageBody:            aws.String(string(b)),
			QueueUrl:               result.QueueUrl,
			MessageGroupId:         aws.String("1"),
			MessageDeduplicationId: aws.String("1"),
		})
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "Message Sent", StatusCode: 200}, nil
}

func main() {
	// Register handler function in the main function using the AWS Lambda for Go library.
	lambda.Start(handleRequest)
}
