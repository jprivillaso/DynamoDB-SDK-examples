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

func writeItems() error {
	dynamoDBClient := dynamodb.New(getSession())

	params := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			// Here it comes the table name
			"Reply": {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: aws.String("vikram.johnson@somewhere.com"),
							},
							"sk": {
								S: aws.String("metadata"),
							},
							"username": {
								S: aws.String("vikj"),
							},
							"first_name": {
								S: aws.String("Vikram"),
							},
							"last_name": {
								S: aws.String("Johnson"),
							},
							"name": {
								S: aws.String("Vikram Johnson"),
							},
							"age": {
								N: aws.String("31"),
							},
							"address": {
								M: map[string]*dynamodb.AttributeValue{
									"road": {
										S: aws.String("89105 Bakken Rd"),
									},
									"city": {
										S: aws.String("Greenbank"),
									},
									"pcode": {
										N: aws.String("98253"),
									},
									"state": {
										S: aws.String("WA"),
									},
									"country": {
										S: aws.String("USA"),
									},
								},
						  	},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: aws.String("jose.schneller@somewhere.com"),
						  	},
						  	"sk": {
								S: aws.String("metadata"),
							},
						  	"username": {
								S: aws.String("joses"),
							},
						  	"first_name": {
								S: aws.String("Jose"),
							},
						  	"last_name": {
								S: aws.String("Schneller"),
							},
						  	"name": {
								S: aws.String("Jose Schneller"),
							},
						  	"age": {
								N: aws.String("27"),
							},
						  	"address": {
								M: map[string]*dynamodb.AttributeValue{
									"road": {
										S: aws.String("12341 Fish Rd"),
									},
									"city": {
										S: aws.String("Freeland"),
									},
									"pcode": {
										N: aws.String("98249"),
									},
									"state": {
										S: aws.String("WA"),
									},
									"country": {
										S: aws.String("USA"),
									},
								},
						  	},
						},
					},
				},
			},
		},
	}

	_, err := dynamoDBClient.BatchWriteItem(params)

	if err != nil {
		fmt.Println("Error writing batch of items ...", err)
		return err
	}

	return nil
}

func main() {
    fmt.Println("Writing Batch of Items ...")
    writeItems()
    fmt.Println("Finished ...")
}
