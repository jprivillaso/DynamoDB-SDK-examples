package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "fmt"
)

func configureSession() (*dynamodb.DynamoDB) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        // Provide SDK Config options, such as Region and Endpoint
        Config: aws.Config{
            Region: aws.String("us-west-2"),
            Endpoint: aws.String("http://localhost:8000"),
	    },
    }))

    client := dynamodb.New(sess)

    return client
}

func updateTable() error {
    client := configureSession()

    table := "Music"

    provisionedThroughput := &dynamodb.ProvisionedThroughput{
        ReadCapacityUnits:  aws.Int64(20),
        WriteCapacityUnits: aws.Int64(5),
    }

    _, err := client.UpdateTable(&dynamodb.UpdateTableInput{
        ProvisionedThroughput: provisionedThroughput,
        TableName:             &table,
    })

    if err := client.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: aws.String(table),
    }); err != nil {
		return err
	}

    if err != nil {
        fmt.Println("Got error calling UpdateTable:")
        fmt.Println(err)
    }

    return nil
}

func main() {
    fmt.Println("Updating Provisioned Capacity ...")
    updateTable()
    fmt.Println("Finished ...")
}
