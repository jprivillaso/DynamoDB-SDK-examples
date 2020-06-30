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

func updateTable(tableName string) error {
    dynamoDBClient := dynamodb.New(getSession())

    response, err := dynamoDBClient.UpdateTable(&dynamodb.UpdateTableInput{
        TableName:             aws.String(tableName),
        BillingMode:           aws.String("PROVISIONED"),
        ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
            ReadCapacityUnits:  aws.Int64(5),
            WriteCapacityUnits: aws.Int64(10),
        },
    })

    if (err != nil) {
        return err
    }

    err = dynamoDBClient.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
    });

    if err != nil {
        fmt.Println("An error occurred updating the table.", err)
		return err
	}

    fmt.Println(response)
    return nil
}

func main() {
    fmt.Println("Listing Tables ...")

    tableName := "Music"
    updateTable(tableName)

    fmt.Println("Finished ...")
}
