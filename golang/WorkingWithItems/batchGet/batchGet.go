package main

import (
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "fmt"
)

var awsRegion = "us-west-2"

func getSession() (*session.Session) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        // Provide SDK Config options, such as Region and Endpoint
        Config: aws.Config{
            Region: aws.String(awsRegion),
        },
    }))

    return sess
}

func getItems() error {
	dynamoDBClient := dynamodb.New(getSession())

	params := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
            "Reply": {
                Keys: []map[string]*dynamodb.AttributeValue{
                    {
                        "pk": &dynamodb.AttributeValue{
                            S: aws.String("vikram.johnson@somewhere.com"),
                        },
                        "sk": &dynamodb.AttributeValue{
                            S: aws.String("metadata"),
                        },
                    },
                    {
                        "pk": &dynamodb.AttributeValue{
                            S: aws.String("jose.schneller@somewhere.com"),
                        },
                        "sk": &dynamodb.AttributeValue{
                            S: aws.String("metadata"),
                        },
                    },
                },
            },
		},
	}

	items, err := dynamoDBClient.BatchGetItem(params)

	if err != nil {
		fmt.Println("Error getting batch of items ...", err)
		return err
    }

    fmt.Println("Items ...")
    fmt.Println(items)

	return nil
}

func main() {
    fmt.Println("Getting Batch of Items ...")
    getItems()
    fmt.Println("Finished ...")
}
