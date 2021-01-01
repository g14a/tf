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
	default:
		color.Red("No such resource present in AWS")
	}
}
