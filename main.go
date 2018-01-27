package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type issueCommentEvent struct {
	Action string `json:"action"`
	Issue  issue  `json:"issue"`
}

type issue struct {
	Comments    int    `json:"comments"`
	CommentsURL string `json:"comments_url"`
	State       string `json:"state"`
}

// Handler is a Lambda function handler that logs the request context from the
// triggering event.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Mark the start of the handler execution in CloudWatch logs
	log.Println("START: golang handler")

	// Log the request context to AWS CloudWatch Logs for this function
	log.Printf("Request Body: %s\n", request.RequestContext)

	// Parse JSON into useful structure
	var e issueCommentEvent
	if err := json.Unmarshal([]byte(request.Body), &e); err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Bad Request",
		}, err // NOTE: this is probably not what I want to pass here...
	}

	// Check if the number of comments exceeds the threshold
	if e.Issue.Comments >= 10 {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Too Many Comments",
		}, nil
	}

	// Mark the end of the handler execution in CloudWatch logs
	log.Println("END: golang handler")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
