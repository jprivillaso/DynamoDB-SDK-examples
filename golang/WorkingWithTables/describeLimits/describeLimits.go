package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "fmt"
)

func getSession() (*session.Session) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        // Provide SDK Config options, such as Region and Endpoint
        Config: aws.Config{
            Region: aws.String("us-west-2"),
	    },
    }))

    return sess
}

func describeLimits() error {
    dynamoDBClient := dynamodb.New(getSession())
    response, err := dynamoDBClient.DescribeLimits(&dynamodb.DescribeLimitsInput{})

    if (err != nil) {
        return err
    }

    fmt.Println("Table Limits ...", response)
    return nil
}

func main() {
    fmt.Println("Describing DynamoDB Limits ...")
    describeLimits()
    fmt.Println("Finished ...")
}
