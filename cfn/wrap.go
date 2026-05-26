// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package cfn

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// CustomResourceLambdaFunction is a standard form Lambda for a Custom Resource.
type CustomResourceLambdaFunction func(context.Context, Event) (reason string, err error)

// SNSCustomResourceLambdaFunction is a standard form Lambda for a Custom Resource
// that is triggered via a SNS topic.
type SNSCustomResourceLambdaFunction func(context.Context, events.SNSEvent) (reason string, err error)

// CustomResourceFunction is a representation of the customer's Custom Resource function.
// LambdaWrap will take the returned values and turn them into a response to be sent
// to CloudFormation.
type CustomResourceFunction func(context.Context, Event) (physicalResourceID string, data map[string]interface{}, err error)

func lambdaWrapWithClient(lambdaFunction CustomResourceFunction, client httpClient) (fn CustomResourceLambdaFunction) {
	_ = "STUB: not implemented"
	return *new(CustomResourceLambdaFunction)
}

// A previous physical resource id exists unless this is a create request.

// If this is a create request, the fallback should be the request ID

// FIXME: something should be done if an error is returned here

// LambdaWrap returns a CustomResourceLambdaFunction which is something lambda.Start()
// will understand. The purpose of doing this is so that Response Handling boiler
// plate is taken away from the customer and it makes writing a Custom Resource
// simpler.
//
//	func myLambda(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
//		physicalResourceID = "arn:...."
//		return
//	}
//
//	func main() {
//		lambda.Start(cfn.LambdaWrap(myLambda))
//	}
func LambdaWrap(lambdaFunction CustomResourceFunction) (fn CustomResourceLambdaFunction) {
	_ = "STUB: not implemented"
	return *new(CustomResourceLambdaFunction)
}

// LambdaWrapSNS wraps a Lambda handler with support for SNS-based custom
// resources. Usage and purpose otherwise same as LambdaWrap().
func LambdaWrapSNS(lambdaFunction CustomResourceFunction) SNSCustomResourceLambdaFunction {
	_ = "STUB: not implemented"
	return *new(SNSCustomResourceLambdaFunction)
}
