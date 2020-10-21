package main

import (
	"github.com/hokurikucar/PeriodTweetTrigger/src/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
