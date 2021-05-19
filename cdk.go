package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type DemoFunctionStackProps struct {
	awscdk.StackProps
}

func NewGoCDKDemoFunctionStack(scope constructs.Construct, id string, props *DemoFunctionStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	awslambda.NewFunction(stack, jsii.String("GoCDKDemoFunction"), &awslambda.FunctionProps{
		Code:         awslambda.NewAssetCode(jsii.String("hello.zip"), nil),
		Handler:      jsii.String("hello"),
		Description:  jsii.String("Demo Lambda with Go CDK"),
		FunctionName: jsii.String("GoCDKDemoFunction"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(30)),
		Tracing:      awslambda.Tracing_ACTIVE,
		Environment: &map[string]*string{
			"REGION": jsii.String("us-east-1"),
			"TZ":     jsii.String("Asia/Tokyo"),
		},
	})
	return stack
}

func main() {
	app := awscdk.NewApp(nil)
	NewGoCDKDemoFunctionStack(app, "GoCDKDemoFunctionStack", &DemoFunctionStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})
	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
