package resourceprompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/utils"
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
			Doc:   "(Required) Code signing configuration policy for deployment validation failure. " +
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
			Doc:   "(Required) The Amazon Resource Name (ARN) for each of the signing profiles. " +
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
			Validator: utils.IntValidator,
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
			Validator: utils.IntValidator,
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
			Validator: utils.BoolValidator,
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
			Validator: utils.MinMaxIntValidator(1, 10),
		},
		{
			Field: "maximum_retry_attempts",
			Ex:    "",
			Doc: "(Optional) The maximum number of times to retry when the function returns an error. " +
				"\nOnly available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.",
			Validator: utils.MinMaxIntValidator(0, 10000),
		},
		{
			Field: "maximum_record_age_in_seconds",
			Ex:    "",
			Doc: "(Optional) The maximum age of a record that Lambda sends to a function for processing. " +
				"\nOnly available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.",
			Validator: utils.MinMaxIntValidator(60, 604800),
		},
		{
			Field:     "bisect_batch_on_function_error",
			Ex:        "",
			Doc:       "(Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to false.",
			Validator: utils.BoolValidator,
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
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder, selectOrder []string

	prompts["function_name"] = types.TfPrompt{
		Label: "Enter function_name:\n(Required) A unique name for your Lambda Function.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "function_name")

	prompts["handler"] = types.TfPrompt{
		Label: "Enter handler:\n(Required) The function entrypoint in your code." +
			"\nCheckout https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "handler")

	prompts["role"] = types.TfPrompt{
		Label: "Enter role:\n(Required) IAM role attached to the Lambda Function. This governs both " +
			"\nwho / what can invoke your Lambda Function, as well as what resources our Lambda " +
			"\nFunction has access to. See Lambda Permission Model for more details.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "role")

	prompts["runtime"] = types.TfPrompt{
		Label: "Enter runtime:\n(Optional) See Runtimes for valid values." +
			"Checkout https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "runtime")

	prompts["filename"] = types.TfPrompt{
		Label: "Enter filename:\n(Optional) The path to the function's deployment package within the local " +
			"\nfilesystem. If defined, The s3_-prefixed options and image_uri cannot be used.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "filename")

	prompts["s3_bucket"] = types.TfPrompt{
		Label: "Enter s3_bucket:\n(Optional) The S3 bucket location containing the function's deployment package. " +
			"\nConflicts with filename and image_uri. This bucket must reside in the same " +
			"\nAWS region where you are creating the Lambda function.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "s3_bucket")

	prompts["s3_key"] = types.TfPrompt{
		Label: "Enter s3_key:\n(Optional) The S3 key of an object containing the function's deployment package. Conflicts with filename and image_uri.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "s3_key")

	prompts["s3_object_version"] = types.TfPrompt{
		Label: "Enter s3_object_version:\n(Optional) The object version containing the function's deployment package. Conflicts with filename and image_uri.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "s3_object_version")

	prompts["image_uri"] = types.TfPrompt{
		Label: "Enter image_uri:\n(Optional) The ECR image URI containing the function's deployment package. Conflicts with filename, s3_bucket, s3_key, and s3_object_version",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "image_uri")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Description of what your Lambda Function does.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["layers"] = types.TfPrompt{
		Label: "Enter layers: e.g. [\"\"]\n(Optional) List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See Lambda Layers",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "layers")

	prompts["memory_size"] = types.TfPrompt{
		Label: "Enter memory_size:\n(Optional) Amount of memory in MB your Lambda Function can use at runtime. Defaults to 128. See Limits",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "memory_size")

	prompts["publish"] = types.TfPrompt{
		Label: "Enter publish(true/false):\n(Optional) Whether to publish creation/change as new Lambda Function Version. Defaults to false.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "publish")

	prompts["reserved_concurrent_executions"] = types.TfPrompt{
		Label: "Enter reserved_concurrent_executions(true/false):\n(Optional) The amount of reserved concurrent executions for this lambda function. " +
			"\nA value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. " +
			"\nDefaults to Unreserved Concurrency Limits -1",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "reserved_concurrent_executions")

	prompts["kms_key_arn"] = types.TfPrompt{
		Label: "Enter kms_key_arn:\n(Optional) Amazon Resource Name (ARN) of the AWS Key Management Service " +
			"\n(KMS) key that is used to encrypt environment variables. If this configuration is not " +
			"\nprovided when environment variables are in use, AWS Lambda uses a default service key. " +
			"\nIf this configuration is provided when environment variables are not in use, " +
			"\nthe AWS Lambda API does not save this configuration and Terraform will show a " +
			"\nperpetual difference of adding the key. To fix the perpetual difference, " +
			"\nremove this configuration.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "kms_key_arn")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: For e.g. k1=v1,k2=v2\n(Optional) A map of tags to assign to the object.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}
	selects["package_type"] = types.TfSelect{
		Label: "Enter package_type:\n(Optional) The Lambda deployment package type. Valid values are Zip and Image. Defaults to Zip.",
		Select: promptui.Select{
			Label: "",
		},
	}
	selectOrder = append(selectOrder, "package_type")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

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

	vpcConfigPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	vpcConfigPrompt["security_group_ids"] = types.TfPrompt{
		Label: "Enter security_group_ids  e.g.[\"id1\",\"id2\"]:\n(Required) A list of security group IDs associated with the Lambda function.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "security_group_ids")

	vpcConfigPrompt["subnet_ids"] = types.TfPrompt{
		Label: "Enter subnet_ids  e.g.[\"id1\",\"id2\"]:\n(Required) A list of subnet IDs associated with the Lambda function.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "subnet_ids")

	resourceBlock["vpc_config"] = builder.PSOrder(nestedPromptOrder, nil, vpcConfigPrompt, nil)

	fileSystemConfigPrompt := map[string]types.TfPrompt{}

	fileSystemConfigPrompt["arn"] = types.TfPrompt{
		Label: "Enter arn:\n(Required) The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "arn")

	fileSystemConfigPrompt["local_mount_path"] = types.TfPrompt{
		Label: "Enter local_mount_path:\n(Required) The path where the function can access the file system, starting with /mnt/.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "local_mount_path")

	resourceBlock["file_system_config"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-2:], nil, fileSystemConfigPrompt, nil)

	timeoutPrompt := map[string]types.TfPrompt{}
	timeoutPrompt["create"] = types.TfPrompt{
		Label: "Enter create:\n(Default 10m) How long to wait for slow uploads or EC2 throttling errors.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "create")

	resourceBlock["timeout"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-1:], nil, fileSystemConfigPrompt, nil)

	tracingConfigPrompt := map[string]types.TfPrompt{}
	tracingConfigPrompt["mode"] = types.TfPrompt{
		Label: "Enter mode:\n(Required) Can be either PassThrough or Active. If PassThrough, Lambda will only trace " +
			"\nthe request from an upstream service if it contains a tracing header with \"sampled=1\". " +
			"\nIf Active, Lambda will respect any tracing header it receives from an upstream service. " +
			"\nIf no tracing header is received, Lambda will call X-Ray for a tracing decision.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "mode")

	resourceBlock["tracing_config"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-1:], nil, tracingConfigPrompt, nil)

	deadLetterConfigPrompt := map[string]types.TfPrompt{}

	deadLetterConfigPrompt["target_arn"] = types.TfPrompt{
		Label: "Enter target_arn:\n(Required) The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM role must be granted suitable access to write to the target object, which means allowing either the sns:Publish or sqs:SendMessage action on this ARN, depending on which service is targeted.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "target_arn")

	resourceBlock["dead_letter_config"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-1:], nil, deadLetterConfigPrompt, nil)

	builder.ResourceBuilder("aws_lambda_function", blockName, resourceBlock)
}

func AWSLambdaLayerVersionPrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string

	prompts["layer_name"] = types.TfPrompt{
		Label: "Enter layer_name:\n(Required) A unique name for your Lambda Layer",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "layer_name")

	prompts["filename"] = types.TfPrompt{
		Label: "Enter filename:\n(Optional) The path to the function's deployment package within the local filesystem. " +
			"\nIf defined, The s3_-prefixed options cannot be used.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "filename")

	prompts["s3_bucket"] = types.TfPrompt{
		Label: "Enter s3_bucket:\n(Optional) The S3 bucket location containing the function's deployment package. " +
			"\nConflicts with filename. This bucket must reside in the same AWS region where you are " +
			"\ncreating the Lambda function.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "s3_bucket")

	prompts["s3_key"] = types.TfPrompt{
		Label: "Enter s3_key:\n(Optional) The S3 key of an object containing the function's deployment package. Conflicts with filename",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "s3_key")

	prompts["s3_object_version"] = types.TfPrompt{
		Label: "Enter s3_object_version:\n(Optional) The object version containing the function's deployment package. Conflicts with filename",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "s3_object_version")

	prompts["compatible_runtimes"] = types.TfPrompt{
		Label: "Enter compatible_runtimes: e.g. [\"nodejs12.x\"]\n(Optional) A list of Runtimes this layer is compatible with. Up to 5 runtimes can be specified.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "compatible_runtimes")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Description of what your Lambda Layer does.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["license_info"] = types.TfPrompt{
		Label: "Enter license_info:\n(Optional) License info for your Lambda Layer. " +
			"\nCheckout https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "license_info")

	prompts["source_code_hash"] = types.TfPrompt{
		Label: "Enter source_code_hash:\n(Optional) Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package " +
			"\nfile specified with either filename or s3_key. The usual way to set this is ${filebase64sha256(\"file.zip\")}",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_code_hash")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	lifecyclePrompt := map[string]types.TfPrompt{}
	var nestedSelectOrder []string

	color.Green("Enter lifecycle block:\nThe lifecycle block supports" +
		"\n1.create_before_destroy\n2.prevent_destroy\n3.ignore_changes\n")

	lifecyclePrompt["create_before_destroy"] = types.TfPrompt{
		Label: "Enter create_before_destroy:(true/false)\nBy default, when Terraform must change a resource argument \n" +
			"that cannot be updated in-place due to remote API limitations, \n" +
			"Terraform will instead destroy the existing object and then \n" +
			"create a new replacement object with the new configured arguments.\n" +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#create_before_destroy",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "create_before_destroy")

	lifecyclePrompt["prevent_destroy"] = types.TfPrompt{
		Label: "Enter prevent_destroy:(true/false)\nThis meta-argument, when set to true, will cause Terraform to \n" +
			"reject with an error any plan that would destroy the infrastructure \n" +
			"object associated with the resource, as long as the argument \n" +
			"remains present in the configuration.\n" +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#prevent_destroy",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "prevent_destroy")

	lifecyclePrompt["ignore_changes"] = types.TfPrompt{
		Label: "Enter ignore_changes: e.g.[\"c1\",\"c2\"]\nBy default, Terraform detects any difference in the " +
			"current settings of a real infrastructure object and plans to " +
			"update the remote object to match configuration." +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#ignore_changes",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "ignore_changes")

	resourceBlock["lifecycle"] = builder.PSOrder(nestedSelectOrder, nil, lifecyclePrompt, nil)

	builder.ResourceBuilder("aws_lambda_layer_version", blockName, resourceBlock)
}

func AWSLambdaPermissionPrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string

	prompts["action"] = types.TfPrompt{
		Label: "Enter action:\n(Required) The AWS Lambda action you want to allow in this statement. (e.g. lambda:InvokeFunction)",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "action")

	prompts["event_source_token"] = types.TfPrompt{
		Label: "Enter event_source_token:\n(Optional) The Event Source Token to validate. Used with Alexa Skills.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "event_source_token")

	prompts["function_name"] = types.TfPrompt{
		Label: "Enter function_name:\n(Required) Name of the Lambda function whose resource policy you are updating",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "function_name")

	prompts["principal"] = types.TfPrompt{
		Label: "Enter principal:\n(Required) The principal who is getting this permission. e.g. s3.amazonaws.com, an AWS account ID, " +
			"\nor any valid AWS service principal such as events.amazonaws.com or sns.amazonaws.com",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "principal")

	prompts["qualifier"] = types.TfPrompt{
		Label: "Enter qualifier:\n(Optional) Query parameter to specify function version or alias name. " +
			"\nThe permission will then apply to the specific qualified ARN. " +
			"\ne.g. arn:aws:lambda:aws-region:acct-id:function:function-name:2",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "qualifier")

	prompts["source_account"] = types.TfPrompt{
		Label: "Enter source_account:\n(Optional) This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_account")

	prompts["source_arn"] = types.TfPrompt{
		Label: "Enter source_arn:\n(Optional) When the principal is an AWS service, the ARN of the specific resource " +
			"\nwithin that service to grant permission to. Without this, any resource from principal will be granted " +
			"\npermission – even if that resource is from another account. For S3, this should be the ARN " +
			"\nof the S3 Bucket. For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule. " +
			"\nFor API Gateway, this should be the ARN of the API, as described in" +
			"\nhttps://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_arn")

	prompts["statement_id"] = types.TfPrompt{
		Label: "Enter statement_id:\n(Optional) A unique statement identifier. By default generated by Terraform.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "statement_id")

	prompts["statement_id_prefix"] = types.TfPrompt{
		Label: "Enter statement_id_prefix:\n(Optional) A statement identifier prefix. Terraform will generate a unique suffix. Conflicts with statement_id.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "statement_id_prefix")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_lambda_permission", blockName, resourceBlock)

}

func AWSLambdaProvisionedConcurrencyConfigPrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string

	prompts["function_name"] = types.TfPrompt{
		Label: "Enter function_name:\n(Required) Name or Amazon Resource Name (ARN) of the Lambda Function.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "function_name")

	prompts["provisioned_concurrent_executions"] = types.TfPrompt{
		Label: "Enter provisioned_concurrent_executions:\n(Required) Amount of capacity to allocate. Must be greater than or equal to 1.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "provisioned_concurrent_executions")

	prompts["qualifier"] = types.TfPrompt{
		Label: "Enter qualifier:\n(Required) Lambda Function version or Lambda Alias name.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "qualifier")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)
	builder.ResourceBuilder("aws_lambda_provisioned_concurrency_config", blockName, resourceBlock)

}
