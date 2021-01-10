package resourcebps

import "github.com/fatih/color"

func AWSElasticBeanstalkApplication() {
	color.Green("\nresource \"aws_elastic_beanstalk_application\" \"tftest\" {\n  name        = \"tf-test-name\"\n  description = \"tf-test-desc\"\n\n  appversion_lifecycle {\n    service_role          = aws_iam_role.beanstalk_service.arn\n    max_count             = 128\n    delete_source_from_s3 = true\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elastic_beanstalk_application\n\n")
}
