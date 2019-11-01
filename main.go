package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

const winCount = 0

func winHandler() (string, error) {
	startDate := Date(2018, 7, 1)
	currentDate := time.Now()
	days := fmt.Sprintf("Days since last win: %f number of wins: %d.", currentDate.Sub(startDate).Hours()/24, winCount)
	return days, nil
}

func main() {
	lambda.Start(winHandler)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
