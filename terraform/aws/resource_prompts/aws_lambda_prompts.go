package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSLambdaAliasPrompt() {
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

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) Name for the alias you are creating. Pattern: (?!^[0-9]+$)([a-zA-Z0-9-_]+)",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Description of the alias.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["function_name"] = types.TfPrompt{
		Label: "Enter function_name:\n(Required) Lambda Function name or ARN.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "function_name")

	prompts["function_version"] = types.TfPrompt{
		Label: "Enter function_version:\n(Required) Lambda function version for which you are creating the alias. Pattern: (\\$LATEST|[0-9]+)",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "function_version")

	builder.ResourceBuilder("aws_lambda_alias", blockName, promptOrder, nil, builder.PSOrder(promptOrder, nil, prompts, nil))
}

func AWSLambdaCodeSigningConfigPrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder, selectOrder, nestedPromptOrder, nestedSelectOrder []string

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Descriptive name for this code signing configuration.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")
	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	policiesSelect := map[string]types.TfSelect{}

	color.Green("\nEnter policies:\n(Optional) A configuration block of code signing policies that define the " +
		"\nactions to take if the validation checks fail. Detailed below." +
		"\nThe policies block supports the following argument:" +
		"\n1.untrusted_artifact_on_deployment")

	policiesSelect["untrusted_artifact_on_deployment"] = types.TfSelect{
		Label: "Enter untrusted_artifact_on_deployment:\n (Required) Code signing configuration policy for deployment validation failure. " +
			"\nIf you set the policy to Enforce, Lambda blocks the deployment request if " +
			"\ncode-signing validation checks fail. If you set the policy to Warn, " +
			"\nLambda allows the deployment and creates a CloudWatch log. \n" +
			"Valid values: Warn, Enforce. Default value: Warn",
		Select: promptui.Select{
			Label: "",
			Items: []string{"Warn", "Enforce"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "untrusted_artifact_on_deployment")
	selectOrder = append(selectOrder, "policies")

	resourceBlock["policies"] = builder.NestedPSOrder(nil, nestedSelectOrder, nil, policiesSelect)

	allowedPublishersPrompt := map[string]types.TfPrompt{}

	color.Green("\nEnter allowed_publishers:\n(Required) A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.")

	allowedPublishersPrompt["signing_profile_version_arns"] = types.TfPrompt{
		Label: "Enter signing_profile_version_arns:\n(Required) The Amazon Resource Name (ARN) for each of the signing profiles. " +
			"\nA signing profile defines a trusted user who can sign a code package.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "signing_profile_version_arns")
	selectOrder = append(selectOrder, "allowed_publishers")

	resourceBlock["allowed_publishers"] = builder.NestedPSOrder(nestedPromptOrder, nil, allowedPublishersPrompt, nil)

	builder.ResourceBuilder("aws_lambda_code_signing_config", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSLambdaEventSourceMappingPrompt() {
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

	prompts["batch_size"] = types.TfPrompt{
		Label: "Enter batch_size:\n(Optional) The largest number of records that Lambda will retrieve from " +
			"\nyour event source at the time of invocation. Defaults to 100 for DynamoDB and Kinesis, 10 for SQS.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "batch_size")

	prompts["maximum_batching_window_in_seconds"] = types.TfPrompt{
		Label: "Enter maximum_batching_window_in_seconds:\n(Optional) The maximum amount of time to gather records before invoking the function, " +
			"\nin seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case " +
			"\nof an SQS queue event source) until either maximum_batching_window_in_seconds expires or " +
			"\nbatch_size has been met. For streaming event sources, defaults to as soon as records are " +
			"\navailable in the stream. If the batch it reads from the stream/queue only has one record " +
			"\nin it, Lambda only sends one record to the function.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}

	promptOrder = append(promptOrder, "maximum_batching_window_in_seconds")

	prompts["event_source_arn"] = types.TfPrompt{
		Label: "Enter event_source_arn:\n(Required) The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}

	promptOrder = append(promptOrder, "event_source_arn")

	prompts["enabled"] = types.TfPrompt{
		Label: "Enter enabled:\n(Optional) Determines if the mapping will be enabled on creation. Defaults to true",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "enabled")

	prompts["function_name"] = types.TfPrompt{
		Label: "Enter function_name:\n(Required) The name or the ARN of the Lambda function that will be subscribing to events.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "function_name")

	prompts["parallelization_factor"] = types.TfPrompt{
		Label: "Enter parallelization_factor:\n(Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "parallelization_factor")

	prompts["maximum_retry_attempts"] = types.TfPrompt{
		Label: "Enter maximum_retry_attempts:\n(Optional) The maximum number of times to retry when the function returns an error. " +
			"\nOnly available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and " +
			"\ndefault of 10000.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "maximum_retry_attempts")

	prompts["maximum_record_age_in_seconds"] = types.TfPrompt{
		Label: "Enter maximum_record_age_in_seconds:\n(Optional) The maximum age of a record that Lambda sends to a function " +
			"\nfor processing. Only available for stream sources (DynamoDB and Kinesis). " +
			"\nMinimum of 60, maximum and default of 604800.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "maximum_record_age_in_seconds")

	prompts["bisect_batch_on_function_error"] = types.TfPrompt{
		Label: "Enter bisect_batch_on_function_error:\n(Optional) If the function returns an error, split the batch in two and retry. " +
			"\nOnly available for stream sources (DynamoDB and Kinesis). Defaults to false",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bisect_batch_on_function_error")

	prompts["starting_position_timestamp"] = types.TfPrompt{
		Label: "Enter starting_position_timestamp:\n (Optional) A timestamp in RFC3339 format of the data record which to start " +
			"\nreading when using starting_position set to AT_TIMESTAMP. If a record with this exact " +
			"\ntimestamp does not exist, the next later record is chosen. If the timestamp is older " +
			"\nthan the current trim horizon, the oldest available record is chosen.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "starting_position_timestamp")

	selects := map[string]types.TfSelect{}

	selects["starting_position"] = types.TfSelect{
		Label: "Enter starting_position:\n(Optional) The position in the stream where AWS Lambda should start reading. " +
			"\nMust be one of AT_TIMESTAMP (Kinesis only), LATEST or TRIM_HORIZON if getting events " +
			"\nfrom Kinesis or DynamoDB. Must not be provided if getting events from SQS.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"AT_TIMESTAMP", "LATEST", "TRIM_HORIZON"},
		},
	}
	selectOrder = append(selectOrder, "starting_position")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_lambda_event_source_mapping", blockName, promptOrder, selectOrder, resourceBlock)

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
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "publish")

	prompts["reserved_concurrent_executions"] = types.TfPrompt{
		Label: "Enter reserved_concurrent_executions(true/false):\n(Optional) The amount of reserved concurrent executions for this lambda function. " +
			"\nA value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. " +
			"\nDefaults to Unreserved Concurrency Limits -1",
		Prompt: promptui.Prompt{
			Label: "",
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
			Label: "",
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

	color.Green("Would you like to configure nested settings like vpc_config/file_system_config etc:")

	color.Yellow("\nConfigure nested settings like assume_role/ignore_tags [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_lambda_function", blockName, promptOrder, nil, resourceBlock)
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
	selectOrder = append(selectOrder, "vpc_config")

	resourceBlock["vpc_config"] = builder.NestedPSOrder(nestedPromptOrder, nil, vpcConfigPrompt, nil)

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
	selectOrder = append(selectOrder, "file_system_config")

	resourceBlock["file_system_config"] = builder.NestedPSOrder(nestedPromptOrder[len(nestedPromptOrder)-2:], nil, fileSystemConfigPrompt, nil)

	timeoutPrompt := map[string]types.TfPrompt{}
	timeoutPrompt["create"] = types.TfPrompt{
		Label: "Enter create:\n(Default 10m) How long to wait for slow uploads or EC2 throttling errors.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "create")
	selectOrder = append(selectOrder, "timeout")

	resourceBlock["timeout"] = builder.NestedPSOrder(nestedPromptOrder[len(nestedPromptOrder)-1:], nil, fileSystemConfigPrompt, nil)

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
	selectOrder = append(selectOrder, "tracing_config")

	resourceBlock["tracing_config"] = builder.NestedPSOrder(nestedPromptOrder[len(nestedPromptOrder)-1:], nil, tracingConfigPrompt, nil)

	deadLetterConfigPrompt := map[string]types.TfPrompt{}

	deadLetterConfigPrompt["target_arn"] = types.TfPrompt{
		Label: "Enter target_arn:\n(Required) The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM role must be granted suitable access to write to the target object, which means allowing either the sns:Publish or sqs:SendMessage action on this ARN, depending on which service is targeted.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "target_arn")
	selectOrder = append(selectOrder, "dead_letter_config")

	resourceBlock["dead_letter_config"] = builder.NestedPSOrder(nestedPromptOrder[len(nestedPromptOrder)-1:], nil, deadLetterConfigPrompt, nil)

	builder.ResourceBuilder("aws_lambda_function", blockName, promptOrder, selectOrder, resourceBlock)
}
