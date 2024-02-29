package responselib

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Result struct {
	Message any `json:"message"`
}
type Response events.APIGatewayProxyResponse

func Generate(event events.APIGatewayProxyRequest, statusCode int, value any) (Response, error) {
	JSONResponse, err := json.Marshal(Result{Message: value})
	if err != nil {
		return Response{Body: "err in marshal", StatusCode: 500}, err
	}
	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(JSONResponse),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}, nil
}
