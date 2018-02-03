package main_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	main "github.com/bvarberg/comment-limit-webhook"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request    events.APIGatewayProxyRequest
		statusCode int
		body       string
		err        error
	}{
		{
			// Responds "Bad Request" when given invalid JSON
			request:    events.APIGatewayProxyRequest{Body: `{`},
			statusCode: 400,
			body:       "Bad Request",
			err:        &json.SyntaxError{},
		},
		{
			// Responds "Too Many Comments" when event issue's comments exceed the threshold
			request: events.APIGatewayProxyRequest{Body: `
				{
					"issue": {
						"comments": 11
					}
				}
			`},
			statusCode: 200,
			body:       "Too Many Comments",
			err:        nil,
		},
		{
			// Responds "Too Many Comments" when event issue's comments match the threshold
			request: events.APIGatewayProxyRequest{Body: `
				{
					"issue": {
						"comments": 10
					}
				}
			`},
			statusCode: 200,
			body:       "Too Many Comments",
			err:        nil,
		},
		{
			// Responds "OK" when event issue's comments are below the threshold
			request: events.APIGatewayProxyRequest{Body: `
				{
					"issue": {
						"comments": 9
					}
				}
			`},
			statusCode: 200,
			body:       "OK",
			err:        nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.body, response.Body)
		assert.Equal(t, test.statusCode, response.StatusCode)
	}
}
