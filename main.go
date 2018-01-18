package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is a Lambda function handler that logs the request context from the
// triggering event.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("START: golang handler")

	// Log the request context to AWS CloudWatch Logs for this function
	log.Printf("Request Body: %s\n", request.RequestContext)

	log.Println("END: golang handler")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
