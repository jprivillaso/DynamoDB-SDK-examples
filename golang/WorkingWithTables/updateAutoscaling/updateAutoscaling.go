package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/applicationautoscaling"
    "fmt"
)

var tableName      = "Music"
var readDimension  = "dynamodb:table:ReadCapacityUnits"
var writeDimension = "dynamodb:table:WriteCapacityUnits"
var resourceID     = fmt.Sprintf("%s%s", "table/", tableName)
var roleARN        = "PUT_YOUR_ROLE_ARN_HERE"

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

func registerAutoscaling() {
    autoscalingClient := applicationautoscaling.New(getSession())

    registerScalableTarget(autoscalingClient, readDimension)
    fmt.Println("Read scalable target registered ...")

    registerScalableTarget(autoscalingClient, writeDimension)
    fmt.Println("Write scalable target registered ...")
}

func main() {
    fmt.Println("Updating autoscaling settings ...")
    registerAutoscaling()
    fmt.Println("Finished ...")
}
