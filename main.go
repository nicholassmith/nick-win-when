package main

import(
	"time"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func winHandler()(string, error) {
	startDate := Date(2018, 7, 1)
	currentDate := time.Now()
	days := fmt.Sprintf("%f", currentDate.Sub(startDate).Hours() / 24)
	fmt.Println("Days since win: ", days)
	return days, nil
}

func main() {
	lambda.Start(winHandler)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}