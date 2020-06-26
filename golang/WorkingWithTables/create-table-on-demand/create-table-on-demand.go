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

func createTable() error {
    client := configureSession()

    table := "Music"

    attributeDefinitions := []*dynamodb.AttributeDefinition{
        {
            AttributeName: aws.String("Artist"),
            AttributeType: aws.String("S"),
        },
        {
            AttributeName: aws.String("SongTitle"),
            AttributeType: aws.String("S"),
        },
    }

    keySchema := []*dynamodb.KeySchemaElement{
        {
            AttributeName: aws.String("Artist"),
            KeyType:       aws.String("HASH"), // Partition Key
        },
        {
            AttributeName: aws.String("SongTitle"),
            KeyType:       aws.String("RANGE"), // Sort Key
        },
    }

    billingMode := aws.String("PAY_PER_REQUEST")

    _, err := client.CreateTable(&dynamodb.CreateTableInput{
        AttributeDefinitions:  attributeDefinitions,
        KeySchema:             keySchema,
        BillingMode:           billingMode,
        TableName:             &table,
    })

	if err := client.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: aws.String(table),
    }); err != nil {
		return err
	}

    if err != nil {
        fmt.Println("Got error calling CreateTable:")
        fmt.Println(err)
    }

    return nil
}

func main() {
    fmt.Println("Creating Provisioned Table ...")
    createTable()
    fmt.Println("Finished ...")
}
