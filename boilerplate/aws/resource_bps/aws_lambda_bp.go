package resource_bps

import "github.com/fatih/color"

func AWSLambdaFunctionBP() {
	color.Green("\nresource \"aws_lambda_function\" \"test_lambda\" {\n  filename      = \"lambda_function_payload.zip\"\n  function_name = \"lambda_function_name\"\n  role          = aws_iam_role.iam_for_lambda.arn\n  handler       = \"exports.test\"\n\n  # The filebase64sha256() function is available in Terraform 0.11.12 and later\n  # For Terraform 0.11.11 and earlier, use the base64sha256() function and the file() function:\n  # source_code_hash = \"${base64sha256(file(\"lambda_function_payload.zip\"))}\"\n  source_code_hash = filebase64sha256(\"lambda_function_payload.zip\")\n\n  runtime = \"nodejs12.x\"\n\n  environment {\n    variables = {\n      foo = \"bar\"\n    }\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function\n\n")
}
