package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/applicationautoscaling"
    "fmt"
)

var tableName = "Music"

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

func registerScalableTarget(
    autoscalingClient *applicationautoscaling.ApplicationAutoScaling,
    dimension string,
    resourceID string,
    roleARN string,
) {
    input := &applicationautoscaling.RegisterScalableTargetInput{
        MaxCapacity:       aws.Int64(500),
        MinCapacity:       aws.Int64(1),
        ResourceId:        aws.String(resourceID),
        RoleARN:           aws.String(roleARN),
        ScalableDimension: aws.String(dimension),
        ServiceNamespace:  aws.String("dynamodb"),
    }
    autoscalingClient.RegisterScalableTarget(input)
}

func registerAutoscaling(roleARN string) {
    autoscalingClient := applicationautoscaling.New(getSession())
    resourceID := fmt.Sprintf("%s%s", "table/", tableName)

    registerScalableTarget(autoscalingClient, "dynamodb:table:ReadCapacityUnits", resourceID, roleARN)
    fmt.Println("Read scalable target registered ...")

    registerScalableTarget(autoscalingClient, "dynamodb:table:WriteCapacityUnits", resourceID, roleARN)
    fmt.Println("Write scalable target registered ...")
}

func main() {
    fmt.Println("Updating table to enable autoscaling ...")

    roleARN := "arn:aws:iam::618326157558:role/Music_TableScalingRole"
    registerAutoscaling(roleARN)

    fmt.Println("Finished ...")
}
