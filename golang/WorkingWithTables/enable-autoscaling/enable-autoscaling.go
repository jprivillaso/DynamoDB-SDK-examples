package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/iam"
    "encoding/json"
    "fmt"
)

// StatementEntry will dictate what this policy allows or doesn't allow.
type StatementEntry struct {
    Effect   string
    Action   []string
    Resource string
}

// PrincipalEntry will dictate what this policy allows or doesn't allow.
type PrincipalEntry struct {
    Service []string
}

// PrinciplaStatementEntry will dictate what this policy allows or doesn't allow.
type PrinciplaStatementEntry struct {
    Effect    string
    Action    []string
    Principal PrincipalEntry
}

// AssumePolicyDocument is our definition of our policies to be uploaded to IAM.
type AssumePolicyDocument struct {
    Version   string
    Statement []StatementEntry
}

// PolicyDocument is our definition of our policies to be uploaded to IAM.
type PolicyDocument struct {
    Version   string
    Statement []PrinciplaStatementEntry
}

var tableName = "Music"

func getAssumeRolePolicyDocument() (string, error) {
    policy := AssumePolicyDocument{
        Version: "2012-10-17",
        Statement: [] StatementEntry{
            StatementEntry{
                Effect: "Allow",
                Action: []string{
                    "dynamodb:DescribeTable",
                    "dynamodb:UpdateTable",
                },
                Resource: "*",
            },
            StatementEntry{
                Effect: "Allow",
                Action: []string{
                    "cloudwatch:PutMetricAlarm",
					"cloudwatch:DescribeAlarms",
					"cloudwatch:GetMetricStatistics",
					"cloudwatch:SetAlarmState",
					"cloudwatch:DeleteAlarms",
                },
                Resource: "*",
            },
        },
    }

    marshalledPolicy, err := json.Marshal(&policy)

    if err != nil {
        return "", err
    }

    return string(marshalledPolicy), nil
}

func getPolicyDocument() (string, error) {
    policy := PolicyDocument{
        Version: "2012-10-17",
        Statement: []PrinciplaStatementEntry {
            PrinciplaStatementEntry{
                Effect: "Allow",
                Action: []string{
                    "sts:AssumeRole",
                },
                Principal: PrincipalEntry{
                    Service: []string{
                        "ec2.amazonaws.com",
                    },
                },
            },
        },
    }

    marshalledPolicy, err := json.Marshal(&policy)

    if err != nil {
        return "", err
    }

    return string(marshalledPolicy), nil
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
    policy, err := getAssumeRolePolicyDocument()

    if err != nil {
        fmt.Println("Error creating Assume Role Policy Document")
        return err
    }

	_, err = iamClient.CreateRole(&iam.CreateRoleInput{
        AssumeRolePolicyDocument: aws.String(policy),
        Path:                     aws.String("/"),
        RoleName: 				  &roleName,
    })

    if err != nil {
        fmt.Println("Error creating IAM Role")
        return err
    }

	return nil
}

func createPolicy(iamClient *iam.IAM) (string, error) {
    policyName := fmt.Sprintf("%s_%s", tableName, "_TableScalingPolicy")
    policyConfig, err := getPolicyDocument()

    if err != nil {
        fmt.Println("Error creating Policy Document")
        return "", err
    }

	policy, err := iamClient.CreatePolicy(&iam.CreatePolicyInput{
		PolicyDocument: aws.String(policyConfig),
		PolicyName:     aws.String(policyName),
    })

    if err != nil {
        fmt.Println("Error creating Policy Document")
        return "", err
    }

    policyArn := *policy.Policy.Arn

    return policyArn, nil
}

func attachPolicy(policyArn string) {

}

func createPreRequisites() error {
	iamClient := iam.New(session.New())

	// Perform IAM requirements before being able to alter the table
	createRole(iamClient)
    policyArn, err := createPolicy(iamClient)

    if err != nil {
        fmt.Println("Error creating Policy Document")
        return err
    }

    attachPolicy(policyArn)

    return nil
}

func updateTable() error {
    client := configureSession()

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
    fmt.Println("Updating table to enable autoscaling ...")
    updateTable()
    fmt.Println("Finished ...")
}
