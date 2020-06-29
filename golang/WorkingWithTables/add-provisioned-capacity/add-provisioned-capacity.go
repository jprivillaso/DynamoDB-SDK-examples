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
            Endpoint: aws.String("http://localhost:8000"),
	    },
    }))

    return sess
}

func updateTable() error {
    client := dynamodb.New(getSession())

    table := "Music"

    provisionedThroughput := &dynamodb.ProvisionedThroughput{
        ReadCapacityUnits:  aws.Int64(20),
        WriteCapacityUnits: aws.Int64(5),
    }

    _, err := client.UpdateTable(&dynamodb.UpdateTableInput{
        ProvisionedThroughput: provisionedThroughput,
        TableName:             &table,
    })

    if err != nil {
        return err
    }

	err = client.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: aws.String(table),
    });

    if err != nil {
        fmt.Println("Got error calling CreateTable:")
		return err
	}

    return nil
}

func main() {
    fmt.Println("Updating Provisioned Capacity ...")
    updateTable()
    fmt.Println("Finished ...")
}
