package resourceprompts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
)

func AWSLambdaAliasPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) Name for the alias you are creating. Pattern: (?!^[0-9]+$)([a-zA-Z0-9-_]+)",
		},
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) Description of the alias.",
		},
		{
			Field: "function_name",
			Ex:    "",
			Doc:   "(Required) Lambda Function name or ARN.",
		},
		{
			Field: "function_version",
			Ex:    "",
			Doc:   "(Required) Lambda function version for which you are creating the alias. Pattern: (\\$LATEST|[0-9]+).",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_lambda_alias", blockName, resourceBlock)
}

func AWSLambdaCodeSigningConfigPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) Descriptive name for this code signing configuration.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Green("\nEnter policies:\n(Optional) A configuration block of code signing policies that define the " +
		"\nactions to take if the validation checks fail. Detailed below." +
		"\nThe policies block supports the following argument:" +
		"\n1.untrusted_artifact_on_deployment")

	policiesSchema := []types.Schema{
		{
			Type:  "select",
			Field: "untrusted_artifact_on_deployment",
			Doc: "(Required) Code signing configuration policy for deployment validation failure. " +
				"\nIf you set the policy to Enforce, Lambda blocks the deployment request if code-signing " +
				"\nvalidation checks fail. If you set the policy to Warn, Lambda allows the deployment and " +
				"\ncreates a CloudWatch log.",
			Items: []string{"Warn", "Enforce"},
		},
	}

	resourceBlock["policies"] = builder.PSOrder(types.ProvidePS(policiesSchema))

	color.Green("\nEnter allowed_publishers:\n(Required) A configuration block of allowed publishers as " +
		"\nsigning profiles for this code signing configuration. Detailed below.")

	allowedPublishersSchema := []types.Schema{
		{
			Field: "signing_profile_version_arns",
			Ex:    "",
			Doc: "(Required) The Amazon Resource Name (ARN) for each of the signing profiles. " +
				"\nA signing profile defines a trusted user who can sign a code package.",
		},
	}

	resourceBlock["allowed_publishers"] = builder.PSOrder(types.ProvidePS(allowedPublishersSchema))

	builder.ResourceBuilder("aws_lambda_code_signing_config", blockName, resourceBlock)
}

func AWSLambdaEventSourceMappingPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field:     "batch_size",
			Ex:        "",
			Doc:       "(Optional) The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to 100 for DynamoDB and Kinesis, 10 for SQS.",
			Validator: validators.IntValidator,
		},
		{
			Field: "maximum_batching_window_in_seconds",
			Ex:    "",
			Doc: "(Optional) The maximum amount of time to gather records before invoking the function, " +
				"\nin seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case " +
				"\nof an SQS queue event source) until either maximum_batching_window_in_seconds expires or " +
				"\nbatch_size has been met. For streaming event sources, defaults to as soon as records are " +
				"\navailable in the stream. If the batch it reads from the stream/queue only has one record in " +
				"\nit, Lambda only sends one record to the function.",
			Validator: validators.IntValidator,
		},
		{
			Field: "event_source_arn",
			Ex:    "",
			Doc:   "(Required) The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.",
		},
		{
			Field:     "enabled",
			Ex:        "",
			Doc:       "(Optional) Determines if the mapping will be enabled on creation. Defaults to true",
			Validator: validators.BoolValidator,
		},
		{
			Field: "function_name",
			Ex:    "",
			Doc:   "(Required) The name or the ARN of the Lambda function that will be subscribing to events.",
		},
		{
			Field: "parallelization_factor",
			Ex:    "",
			Doc: "(Optional) The number of batches to process from each shard concurrently. " +
				"\nOnly available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.",
			Validator: validators.MinMaxIntValidator(1, 10),
		},
		{
			Field: "maximum_retry_attempts",
			Ex:    "",
			Doc: "(Optional) The maximum number of times to retry when the function returns an error. " +
				"\nOnly available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.",
			Validator: validators.MinMaxIntValidator(0, 10000),
		},
		{
			Field: "maximum_record_age_in_seconds",
			Ex:    "",
			Doc: "(Optional) The maximum age of a record that Lambda sends to a function for processing. " +
				"\nOnly available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.",
			Validator: validators.MinMaxIntValidator(60, 604800),
		},
		{
			Field:     "bisect_batch_on_function_error",
			Ex:        "",
			Doc:       "(Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to false.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "starting_position_timestamp",
			Ex:    "",
			Doc: "(Optional) A timestamp in RFC3339 format of the data record which to start " +
				"\nreading when using starting_position set to AT_TIMESTAMP. If a record with this " +
				"\nexact timestamp does not exist, the next later record is chosen. " +
				"\nIf the timestamp is older than the current trim horizon, the oldest available record is chosen.",
		},
		{
			Type:  "select",
			Field: "starting_position",
			Doc: "(Optional) The position in the stream where AWS Lambda should start reading. Must be one of " +
				"\nAT_TIMESTAMP (Kinesis only), LATEST or TRIM_HORIZON if getting events from Kinesis or DynamoDB. " +
				"\nMust not be provided if getting events from SQS",
			Items: []string{"AT_TIMESTAMP", "LATEST", "TRIM_HORIZON"},
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_lambda_event_source_mapping", blockName, resourceBlock)

}

func AWSLambdaFunctionPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "filename",
			Ex:    "",
			Doc:   "(Optional) The path to the function's deployment package within the local filesystem. If defined, The s3_-prefixed options and image_uri cannot be used.",
		},
		{
			Field: "s3_bucket",
			Ex:    "",
			Doc:   "(Optional) The S3 bucket location containing the function's deployment package. Conflicts with filename and image_uri. This bucket must reside in the same AWS region where you are creating the Lambda function.",
		},
		{
			Field: "s3_key",
			Ex:    "",
			Doc:   "(Optional) The S3 key of an object containing the function's deployment package. Conflicts with filename and image_uri.",
		},
		{
			Field: "s3_object_version",
			Ex:    "",
			Doc:   "(Optional) The object version containing the function's deployment package. Conflicts with filename and image_uri",
		},
		{
			Field: "image_uri",
			Ex:    "",
			Doc:   "(Optional) The ECR image URI containing the function's deployment package. Conflicts with filename, s3_bucket, s3_key, and s3_object_version",
		},
		{
			Type:  "select",
			Field: "package_type",
			Doc:   "(Optional) The Lambda deployment package type. Valid values are Zip and Image. Defaults to Zip",
			Items: []string{"Zip", "Image"},
		},
		{
			Field: "function_name",
			Ex:    "",
			Doc:   "(Required) A unique name for your Lambda Function.",
		},
		{
			Field: "handler",
			Ex:    "",
			Doc:   "(Required) The function entrypoint in your code.",
		},
		{
			Field: "role",
			Ex:    "",
			Doc: "(Required) IAM role attached to the Lambda Function. This governs both who / what " +
				"\ncan invoke your Lambda Function, as well as what resources our Lambda Function has access to. " +
				"\nCheckout https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html",
		},
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) Description of what your Lambda Function does.",
		},
		{
			Field: "layers",
			Ex:    "",
			Doc: "(Optional) List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function." +
				"\nCheckout https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html",
		},
		{
			Field:     "memory_size",
			Ex:        "",
			Doc:       "(Optional) Amount of memory in MB your Lambda Function can use at runtime. Defaults to 128",
			Validator: validators.IntValidator,
		},
		{
			Field: "runtime",
			Ex:    "",
			Doc:   "(Optional) Checkout https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime",
		},
		{
			Field:     "timeout",
			Ex:        "3",
			Doc:       "(Optional) The amount of time your Lambda Function has to run in seconds. Defaults to 3",
			Validator: validators.IntValidator,
		},
		{
			Field: "reserved_concurrent_executions",
			Ex:    "",
			Doc:   "(Optional) The amount of reserved concurrent executions for this lambda function. A value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. Defaults to Unreserved Concurrency Limits -1",
		},
		{
			Field:     "publish",
			Ex:        "(true/false)",
			Doc:       "(Optional) Whether to publish creation/change as new Lambda Function Version. Defaults to false",
			Validator: validators.BoolValidator,
		},
		{
			Field: "kms_key_arn",
			Ex:    "",
			Doc: "(Optional) Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) " +
				"\nkey that is used to encrypt environment variables. If this configuration is not " +
				"\nprovided when environment variables are in use, AWS Lambda uses a default service key. " +
				"\nIf this configuration is provided when environment variables are not in use, " +
				"\nthe AWS Lambda API does not save this configuration and Terraform will show a " +
				"\nperpetual difference of adding the key. To fix the perpetual difference, remove this configuration.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the object.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nWould you like to configure nested settings like vpc_config/file_system_config etc: [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_lambda_function", blockName, resourceBlock)
		return
	}

	vpcConfigSchema := []types.Schema{
		{
			Field: "subnet_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc:   "(Required) A list of subnet IDs associated with the Lambda function.",
		},
		{
			Field: "security_group_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc:   "(Required) A list of security group IDs associated with the Lambda function.",
		},
	}

	resourceBlock["vpc_config"] = builder.PSOrder(types.ProvidePS(vpcConfigSchema))

	fileSystemConfigSchema := []types.Schema{
		{
			Field: "arn",
			Ex:    "",
			Doc:   "(Required) The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.",
		},
		{
			Field: "local_mount_path",
			Ex:    "",
			Doc:   "(Required) The path where the function can access the file system, starting with /mnt/",
		},
	}

	resourceBlock["file_system_config"] = builder.PSOrder(types.ProvidePS(fileSystemConfigSchema))

	timeoutSchema := []types.Schema{
		{
			Field: "create",
			Ex:    "60s | 10m | 2h",
			Doc:   "(Default 10m) How long to wait for slow uploads or EC2 throttling errors.",
		},
	}

	resourceBlock["timeout"] = builder.PSOrder(types.ProvidePS(timeoutSchema))

	tracingConfigSchema := []types.Schema{
		{
			Type:  "select",
			Field: "mode",
			Ex:    "",
			Doc:   "(Required) Can be either PassThrough or Active. If PassThrough, Lambda will only trace the request from an upstream service if it contains a tracing header with \"sampled=1\". If Active, Lambda will respect any tracing header it receives from an upstream service. If no tracing header is received, Lambda will call X-Ray for a tracing decision.",
			Items: []string{"PassThrough", "Active"},
		},
	}

	resourceBlock["tracing_config"] = builder.PSOrder(types.ProvidePS(tracingConfigSchema))

	deadLetterConfigSchema := []types.Schema{
		{
			Field: "target_arn",
			Ex:    "",
			Doc:   "(Required) The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM role must be granted suitable access to write to the target object, which means allowing either the sns:Publish or sqs:SendMessage action on this ARN, depending on which service is targeted.",
		},
	}

	resourceBlock["dead_letter_config"] = builder.PSOrder(types.ProvidePS(deadLetterConfigSchema))

	builder.ResourceBuilder("aws_lambda_function", blockName, resourceBlock)
}

func AWSLambdaFunctionEventInvokeConfigPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "function_name",
			Ex:    "",
			Doc:   "(Required) Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.",
		},
		{
			Field:     "maximum_event_age_in_seconds",
			Ex:        "100",
			Doc:       "(Optional) Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.",
			Validator: validators.MinMaxIntValidator(60, 21600),
		},
		{
			Field:     "maximum_retry_attempts",
			Ex:        "1",
			Doc:       "(Optional) Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.",
			Validator: validators.MinMaxIntValidator(0, 2),
		},
		{
			Field: "qualifier",
			Ex:    "",
			Doc:   "(Optional) Lambda Function published version, $LATEST, or Lambda Alias name.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Green("\nEnter destination_config:\n(Optional) Configuration block with destination configuration." +
		"\nThe destination_config block supports the following arguments:" +
		"\n1.on_failure\n2.on_success")

	color.Green("\nEnter on_failure block:\n(Optional) Configuration block with destination configuration for failed asynchronous invocations." +
		"\nThe on_failure block supports the following arguments:" +
		"\n1.destination")

	onFailureSchema := []types.Schema{
		{
			Field: "destination",
			Ex:    "",
			Doc: "(Required) Amazon Resource Name (ARN) of the destination resource. " +
				"\nCheckout https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations for acceptable resource types and associated IAM permissions.",
		},
	}

	onFailureBlock := builder.PSOrder(types.ProvidePS(onFailureSchema))

	color.Green("\nEnter on_success block:\n(Optional) Configuration block with destination configuration for failed asynchronous invocations." +
		"\nThe on_success block supports the following arguments:" +
		"\n1.destination")

	onSuccessSchema := []types.Schema{
		{
			Field: "destination",
			Ex:    "",
			Doc: "(Required) Amazon Resource Name (ARN) of the destination resource. " +
				"\nCheckout https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations for acceptable resource types and associated IAM permissions.",
		},
	}

	onSuccessBlock := builder.PSOrder(types.ProvidePS(onSuccessSchema))

	destinationConfig := map[string]interface{}{
		"on_failure": onFailureBlock,
		"on_success": onSuccessBlock,
	}

	resourceBlock["destination_config"] = destinationConfig

	builder.ResourceBuilder("aws_lambda_function_event_invoke_config", blockName, resourceBlock)
}

func AWSLambdaLayerVersionPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "layer_name",
			Ex:    "",
			Doc:   "(Required) A unique name for your Lambda Layer",
		},
		{
			Field: "filename",
			Ex:    "",
			Doc:   "(Optional) The path to the function's deployment package within the local filesystem. If defined, The s3_-prefixed options cannot be used.",
		},
		{
			Field: "s3_bucket",
			Ex:    "",
			Doc:   "(Optional) The S3 bucket location containing the function's deployment package. Conflicts with filename. This bucket must reside in the same AWS region where you are creating the Lambda function.",
		},
		{
			Field: "s3_key",
			Ex:    "",
			Doc:   "(Optional) The S3 key of an object containing the function's deployment package. Conflicts with filename",
		},
		{
			Field: "s3_object_version",
			Ex:    "",
			Doc:   "(Optional) The object version containing the function's deployment package. Conflicts with filename",
		},
		{
			Field: "compatible_runtimes",
			Ex:    "",
			Doc:   "(Optional) A list of Runtimes this layer is compatible with. Up to 5 runtimes can be specified.",
		},
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) Description of what your Lambda Layer does.",
		},
		{
			Field: "license_info",
			Ex:    "",
			Doc: "(Optional) License info for your Lambda Layer." +
				"\nCheckout https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo",
		},
		{
			Field: "source_code_hash",
			Ex:    "",
			Doc:   "(Optional) Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key. The usual way to set this is ${filebase64sha256(\"file.zip\")} (Terraform 0.11.12 or later) or ${base64sha256(file(\"file.zip\"))} (Terraform 0.11.11 and earlier), where \"file.zip\" is the local filename of the lambda layer source archive.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_lambda_layer_version", blockName, resourceBlock)
}

func AWSLambdaPermissionPrompt() {

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "action",
			Ex:    "lambda:InvokeFunction",
			Doc:   "(Required) The AWS Lambda action you want to allow in this statement.",
		},
		{
			Field: "event_source_token",
			Ex:    "",
			Doc:   "(Optional) The Event Source Token to validate.",
		},
		{
			Field: "function_name",
			Ex:    "",
			Doc:   "(Required) Name of the Lambda function whose resource policy you are updating",
		},
		{
			Field: "principal",
			Ex:    "s3.amazonaws.com",
			Doc: "(Required) The principal who is getting this permission. " +
				"\nAny valid AWS service principal such as events.amazonaws.com or sns.amazonaws.com",
		},
		{
			Field: "qualifier",
			Ex:    "arn:aws:lambda:aws-region:acct-id:function:function-name:2",
			Doc: "(Optional) Query parameter to specify function version or alias name. " +
				"\nThe permission will then apply to the specific qualified ARN.",
		},
		{
			Field: "source_account",
			Ex:    "",
			Doc: "(Optional) This parameter is used for S3 and SES. " +
				"\nThe AWS account ID (without a hyphen) of the source owner.",
		},
		{
			Field: "source_arn",
			Ex:    "",
			Doc: "(Optional) When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to. " +
				"\nWithout this, any resource from principal will be granted permission – even if that resource is from another account. " +
				"\nFor S3, this should be the ARN of the S3 Bucket. For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule. " +
				"\nFor API Gateway, this should be the ARN of the API, as described here.",
		},
		{
			Field: "statement_id",
			Ex:    "",
			Doc:   "(Optional) A unique statement identifier. By default generated by Terraform.",
		},
		{
			Field: "statement_id_prefix",
			Ex:    "",
			Doc: "(Optional) A statement identifier prefix. Terraform will generate a unique suffix. " +
				"\nConflicts with statement_id",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_lambda_permission", blockName, resourceBlock)

}

func AWSLambdaProvisionedConcurrencyConfigPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "function_name",
			Ex:    "",
			Doc:   "(Required) Name or Amazon Resource Name (ARN) of the Lambda Function.",
		},
		{
			Field:     "provisioned_concurrent_executions",
			Ex:        "",
			Doc:       "(Required) Amount of capacity to allocate. Must be greater than or equal to 1",
			Validator: validators.IntValidator,
		},
		{
			Field: "qualifier",
			Ex:    "",
			Doc:   "(Required) Lambda Function version or Lambda Alias name.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like timeouts [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_lambda_provisioned_concurrency_config", blockName, resourceBlock)
		return
	}

	timeoutSchema := []types.Schema{
		{
			Field: "create",
			Ex:    "60s | 30m | 2h",
			Doc:   "(Default 15 minutes) How long to wait for the Lambda Provisioned Concurrency Config to be ready on creation.",
		},
		{
			Field: "update",
			Ex:    "60s | 30m | 2h",
			Doc:   "(Default 15 minutes) How long to wait for the Lambda Provisioned Concurrency Config to be ready on update.",
		},
	}

	resourceBlock["timeouts"] = builder.PSOrder(types.ProvidePS(timeoutSchema))

	builder.ResourceBuilder("aws_lambda_provisioned_concurrency_config", blockName, resourceBlock)
}
