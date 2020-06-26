package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
		"github.com/aws/aws-sdk-go/service/dynamodb"
		"github.com/aws/aws-sdk-go/service/iam"
		"encoding/json"
    "fmt"
)

var tableName = "Music"

func getPolicyDocument() (string) {
	policyDocument := `
	{
		Version: "2012-10-17",
		Statement: [
			{
				Effect: "Allow",
				Action: [
					"dynamodb:DescribeTable",
					"dynamodb:UpdateTable",
					"cloudwatch:PutMetricAlarm",
					"cloudwatch:DescribeAlarms",
					"cloudwatch:GetMetricStatistics",
					"cloudwatch:SetAlarmState",
					"cloudwatch:DeleteAlarms",
				],
				Resource: "*",
			},
		],
	}
	`

	return policyDocument
}

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

func createRole(iamClient *iam.IAM) error {

	roleName := fmt.Sprintf("%s_%s", tableName, "_TableScalingRole")
	policyDocument, err := json.Marshal(getPolicyDocument())

	if err != nil {
		fmt.Println("Error serializing role's policy document:")
		return err
	}

	iamClient.CreateRole(&iam.CreateRoleInput{
    AssumeRolePolicyDocument: aws.String(string(policyDocument)),
    Path:                     aws.String("/"),
    RoleName: 							  &roleName,
	})

	return nil
}

func createTable() error {
    client := configureSession()

		iamClient := iam.New(session.New())

		// Perform IAM requirements before being able to alter the table
		createRole(iamClient)

    _, err := client.UpdateTable(&dynamodb.UpdateTableInput{
			TableName: &tableName,
    })

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
