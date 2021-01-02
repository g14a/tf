package aws

import (
	"github.com/fatih/color"
	"tf/boilerplate/aws/resource_bps"
)

func ResourceBP(resource string) {

	switch resource {
	case "aws_instance":
		resource_bps.AWSEC2BP()
	case "aws_db_instance":
		resource_bps.AWSDBInstanceBP()
	case "aws_s3_bucket":
		resource_bps.AWSS3BucketBP()
	case "aws_vpc":
		resource_bps.AWSVPCBP()
	case "aws_sns_platform_application":
		resource_bps.AWSSNSPlatformApplicationBP()
	case "aws_elastic_beanstalk_application":
		resource_bps.AWSElasticBeanstalkApplication()
	case "aws_cloudfront_distribution":
		resource_bps.AWSCloudFrontDistributionPrompt()
	case "aws_lambda_function":
		resource_bps.AWSLambdaFunctionBP()
	case "aws_lambda_code_signing_config":
		resource_bps.AWSLambdaCodeSigningConfigBP()
	case "aws_lambda_layer_version":
		resource_bps.AWSLambdaLayerVersionBP()
	case "aws_lambda_permission":
		resource_bps.AWSLambdaPermissionBP()
	case "aws_lambda_function_event_invoke_config":
		resource_bps.AWSLambdaFunctionEventInvokeConfigBP()
	case "aws_lambda_provisioned_concurrency_config":
		resource_bps.AWSLambdaProvisionedConcurrencyConfig()
	default:
		color.Red("No such resource present in AWS")
	}
}
