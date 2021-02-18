package resourcebps

import "github.com/fatih/color"

func AWSAPIGatewayAccountBP()  {
	color.Green("\nresource \"aws_api_gateway_account\" \"demo\" {\n  cloudwatch_role_arn = aws_iam_role.cloudwatch.arn\n}\n\nresource \"aws_iam_role\" \"cloudwatch\" {\n  name = \"api_gateway_cloudwatch_global\"\n\n  assume_role_policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"\",\n      \"Effect\": \"Allow\",\n      \"Principal\": {\n        \"Service\": \"apigateway.amazonaws.com\"\n      },\n      \"Action\": \"sts:AssumeRole\"\n    }\n  ]\n}\nEOF\n}\n\nresource \"aws_iam_role_policy\" \"cloudwatch\" {\n  name = \"default\"\n  role = aws_iam_role.cloudwatch.id\n\n  policy = <<EOF\n{\n    \"Version\": \"2012-10-17\",\n    \"Statement\": [\n        {\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"logs:CreateLogGroup\",\n                \"logs:CreateLogStream\",\n                \"logs:DescribeLogGroups\",\n                \"logs:DescribeLogStreams\",\n                \"logs:PutLogEvents\",\n                \"logs:GetLogEvents\",\n                \"logs:FilterLogEvents\"\n            ],\n            \"Resource\": \"*\"\n        }\n    ]\n}\nEOF\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_account\n\n")
}

func AWSAPIGatewayAPIKeyBP() {
	color.Green("\nresource \"aws_api_gateway_api_key\" \"MyDemoApiKey\" {\n  name = \"demo\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_api_key\n\n")
}

func AWSAPIGatewayAuthorizerBP()  {
	color.Green("\nresource \"aws_api_gateway_authorizer\" \"demo\" {\n  name                   = \"demo\"\n  rest_api_id            = aws_api_gateway_rest_api.demo.id\n  authorizer_uri         = aws_lambda_function.authorizer.invoke_arn\n  authorizer_credentials = aws_iam_role.invocation_role.arn\n}\n\nresource \"aws_api_gateway_rest_api\" \"demo\" {\n  name = \"auth-demo\"\n}\n\nresource \"aws_iam_role\" \"invocation_role\" {\n  name = \"api_gateway_auth_invocation\"\n  path = \"/\"\n\n  assume_role_policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Action\": \"sts:AssumeRole\",\n      \"Principal\": {\n        \"Service\": \"apigateway.amazonaws.com\"\n      },\n      \"Effect\": \"Allow\",\n      \"Sid\": \"\"\n    }\n  ]\n}\nEOF\n}\n\nresource \"aws_iam_role_policy\" \"invocation_policy\" {\n  name = \"default\"\n  role = aws_iam_role.invocation_role.id\n\n  policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Action\": \"lambda:InvokeFunction\",\n      \"Effect\": \"Allow\",\n      \"Resource\": \"${aws_lambda_function.authorizer.arn}\"\n    }\n  ]\n}\nEOF\n}\n\nresource \"aws_iam_role\" \"lambda\" {\n  name = \"demo-lambda\"\n\n  assume_role_policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Action\": \"sts:AssumeRole\",\n      \"Principal\": {\n        \"Service\": \"lambda.amazonaws.com\"\n      },\n      \"Effect\": \"Allow\",\n      \"Sid\": \"\"\n    }\n  ]\n}\nEOF\n}\n\nresource \"aws_lambda_function\" \"authorizer\" {\n  filename      = \"lambda-function.zip\"\n  function_name = \"api_gateway_authorizer\"\n  role          = aws_iam_role.lambda.arn\n  handler       = \"exports.example\"\n\n  source_code_hash = filebase64sha256(\"lambda-function.zip\")\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_authorizer\n\n")
}
