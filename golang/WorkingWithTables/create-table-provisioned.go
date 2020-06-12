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
        // Provide SDK Config options, such as Region.
        Config: aws.Config{
            Region: aws.String("us-west-2"),
            Endpoint: aws.String("http://localhost:8000"),
	    },
    }))

    svc := dynamodb.New(sess)

    return svc
}

func createTable(svc *dynamodb.DynamoDB) {
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

    provisionedThroughput := &dynamodb.ProvisionedThroughput{
        ReadCapacityUnits:  aws.Int64(10),
        WriteCapacityUnits: aws.Int64(10),
    }

    _, err := svc.CreateTable(&dynamodb.CreateTableInput{
        AttributeDefinitions:  attributeDefinitions,
        KeySchema:             keySchema,
        ProvisionedThroughput: provisionedThroughput,
        TableName:             &table,
    })

    if err != nil {
        fmt.Println("Got error calling CreateTable:")
        fmt.Println(err)
    }
}

func main() {
    fmt.Println("Creating Provisioned Table ...")

    svc := configureSession()
    createTable(svc)

    fmt.Println("Finished ...")
}
