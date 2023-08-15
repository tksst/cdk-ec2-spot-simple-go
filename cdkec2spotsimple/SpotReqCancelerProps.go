package cdkec2spotsimple

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Options related to Lambda functions to cancel spot requests.
type SpotReqCancelerProps struct {
	// Internal Lambda functions execution role.
	// Default: - Create a new Role that can do ec2:DescribeInstances and ec2:CancelSpotInstanceRequests and has "service-role/AWSLambdaBasicExecutionRole".
	//
	LambdaExcecutionRole awsiam.IRole `field:"optional" json:"lambdaExcecutionRole" yaml:"lambdaExcecutionRole"`
	// Log retention period for internal Lambda functions logs kept in CloudWatch Logs.
	// Default: - Three months.
	//
	LambdaLogRetention awslogs.RetentionDays `field:"optional" json:"lambdaLogRetention" yaml:"lambdaLogRetention"`
	// Runtime environment for the internal Lambda function.
	//
	// If anything other than Node.js is specified, an error will occur.
	// Default: - Node.js 16
	//
	LambdaRuntime awslambda.Runtime `field:"optional" json:"lambdaRuntime" yaml:"lambdaRuntime"`
}

