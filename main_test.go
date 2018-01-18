package main_test

import (
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
			// Test that the handler responds with a 200 OK response
			request:    events.APIGatewayProxyRequest{Body: ""},
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
