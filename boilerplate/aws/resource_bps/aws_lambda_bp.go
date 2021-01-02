package resource_bps

import "github.com/fatih/color"

func AWSLambdaFunctionBP() {
	color.Green("\nresource \"aws_lambda_function\" \"test_lambda\" {\n  filename      = \"lambda_function_payload.zip\"\n  function_name = \"lambda_function_name\"\n  role          = aws_iam_role.iam_for_lambda.arn\n  handler       = \"exports.test\"\n\n  # The filebase64sha256() function is available in Terraform 0.11.12 and later\n  # For Terraform 0.11.11 and earlier, use the base64sha256() function and the file() function:\n  # source_code_hash = \"${base64sha256(file(\"lambda_function_payload.zip\"))}\"\n  source_code_hash = filebase64sha256(\"lambda_function_payload.zip\")\n\n  runtime = \"nodejs12.x\"\n\n  environment {\n    variables = {\n      foo = \"bar\"\n    }\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function\n\n")
}

func AWSLambdaCodeSigningConfigBP()  {
	color.Green("\nresource \"aws_lambda_code_signing_config\" \"foo\" {\n  allowed_publishers {\n    signing_profile_version_arns = [\n      aws_signer_signing_profile.example1.arn,\n      aws_signer_signing_profile.example2.arn,\n    ]\n  }\n\n  policies {\n    untrusted_artifact_on_deployment = \"Warn\"\n  }\n\n  description = \"My awesome code signing config.\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_code_signing_config\n\n")
}

func AWSLambdaEventSourceMappingBP()  {
	color.Green("\nresource \"aws_lambda_event_source_mapping\" \"foo\" {\n  event_source_arn  = aws_dynamodb_table.example.stream_arn\n  function_name     = aws_lambda_function.example.arn\n  starting_position = \"LATEST\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_event_source_mapping\n\n")
}

func AWSLambdaAliasBP()  {
	color.Green("\nresource \"aws_lambda_alias\" \"foo\" {\n  name             = \"my_alias\"\n  description      = \"a sample description\"\n  function_name    = aws_lambda_function.lambda_function_test.arn\n  function_version = \"1\"\n\n  routing_config {\n    additional_version_weights = {\n      \"2\" = 0.5\n    }\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_alias\n\n")
}

func AWSLambdaFunctionEventInvokeConfigBP()  {
	color.Green("\nresource \"aws_lambda_function_event_invoke_config\" \"example\" {\n  function_name = aws_lambda_alias.example.function_name\n\n  destination_config {\n    on_failure {\n      destination = aws_sqs_queue.example.arn\n    }\n\n    on_success {\n      destination = aws_sns_topic.example.arn\n    }\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function_event_invoke_config\n\n")
}

func AWSLambdaLayerVersionBP()  {
	color.Green("\nresource \"aws_lambda_layer_version\" \"lambda_layer\" {\n  filename   = \"lambda_layer_payload.zip\"\n  layer_name = \"lambda_layer_name\"\n\n  compatible_runtimes = [\"nodejs12.x\"]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_layer_version\n\n")
}

func AWSLambdaPermissionBP()  {
	color.Green("\nresource \"aws_lambda_permission\" \"allow_cloudwatch\" {\n  statement_id  = \"AllowExecutionFromCloudWatch\"\n  action        = \"lambda:InvokeFunction\"\n  function_name = aws_lambda_function.test_lambda.function_name\n  principal     = \"events.amazonaws.com\"\n  source_arn    = \"arn:aws:events:eu-west-1:111122223333:rule/RunDaily\"\n  qualifier     = aws_lambda_alias.test_alias.name\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission")
}

func AWSLambdaProvisionedConcurrencyConfig()  {
	color.Green("\nresource \"aws_lambda_provisioned_concurrency_config\" \"example\" {\n  function_name                     = aws_lambda_alias.example.function_name\n  provisioned_concurrent_executions = 1\n  qualifier                         = aws_lambda_alias.example.name\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_provisioned_concurrency_config\n\n")
}