package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) (err error) {

		//SQS
		queue, err := sqs.NewQueue(ctx, "consumer", &sqs.QueueArgs{
			Name: pulumi.StringPtr("consumer"),
		})
		if err != nil {
			return err
		}

		//DynamoDB
		_, err = dynamodb.NewTable(ctx, "users", &dynamodb.TableArgs{
			Name: pulumi.StringPtr("users"),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("id"),
					Type: pulumi.String("S"),
				},
			},
			HashKey:       pulumi.String("id"),
			ReadCapacity:  pulumi.Int(10),
			WriteCapacity: pulumi.Int(10),
		})
		if err != nil {
			return err
		}

		//Lambda
		fileArchive := os.Getenv("LAMBDA_HANDLER")
		if len(fileArchive) == 0 {
			panic("empty lambda zip")
		}

		function, err := lambda.NewFunction(ctx, "user_gen_password", &lambda.FunctionArgs{
			Runtime:    lambda.RuntimeGo1dx,
			MemorySize: pulumi.IntPtr(128),
			Role:       pulumi.String("arn:aws:iam::000000000000:role/irrelevant"),
			Timeout:    pulumi.IntPtr(10),
			Handler:    pulumi.StringPtr("bin/main"),
			Code:       pulumi.NewFileArchive(fileArchive),
			Environment: lambda.FunctionEnvironmentArgs{
				Variables: pulumi.ToStringMap(map[string]string{
					"SERVER_ENVIRONMENT":    "development",
					"DEFAULT_REGION":        "us-east-1",
					"DYNAMODB_AWS_ENDPOINT": "http://172.16.0.11:4566",
					"TABLE_NAME":            "users",
				}),
			},
		})

		if err != nil {
			return err
		}

		_, err = lambda.NewEventSourceMapping(ctx, "trigger_consumer", &lambda.EventSourceMappingArgs{
			EventSourceArn:   queue.Arn,
			FunctionName:     function.Arn,
			StartingPosition: pulumi.String("LATEST"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
