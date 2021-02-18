package resourcebps

import "github.com/fatih/color"

func AWSElasticBeanstalkApplicationBP() {
	color.Green("\nresource \"aws_elastic_beanstalk_application\" \"tftest\" {\n  name        = \"tf-test-name\"\n  description = \"tf-test-desc\"\n\n  appversion_lifecycle {\n    service_role          = aws_iam_role.beanstalk_service.arn\n    max_count             = 128\n    delete_source_from_s3 = true\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elastic_beanstalk_application\n\n")
}

func AWSElasticBeanstalkApplicationVersionBP()  {
	color.Green("\nresource \"aws_s3_bucket\" \"default\" {\n  bucket = \"tftest.applicationversion.bucket\"\n}\n\nresource \"aws_s3_bucket_object\" \"default\" {\n  bucket = aws_s3_bucket.default.id\n  key    = \"beanstalk/go-v1.zip\"\n  source = \"go-v1.zip\"\n}\n\nresource \"aws_elastic_beanstalk_application\" \"default\" {\n  name        = \"tf-test-name\"\n  description = \"tf-test-desc\"\n}\n\nresource \"aws_elastic_beanstalk_application_version\" \"default\" {\n  name        = \"tf-test-version-label\"\n  application = \"tf-test-name\"\n  description = \"application version created by terraform\"\n  bucket      = aws_s3_bucket.default.id\n  key         = aws_s3_bucket_object.default.id\n}\n\n")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elastic_beanstalk_application_version\n\n")
}

func AWSElasticBeanstalkApplicationConfigurationTemplateBP()  {
	color.Green("\nresource \"aws_elastic_beanstalk_application\" \"tftest\" {\n  name        = \"tf-test-name\"\n  description = \"tf-test-desc\"\n}\n\nresource \"aws_elastic_beanstalk_configuration_template\" \"tf_template\" {\n  name                = \"tf-test-template-config\"\n  application         = aws_elastic_beanstalk_application.tftest.name\n  solution_stack_name = \"64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4\"\n}")
	color.Yellow("\n\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elastic_beanstalk_configuration_template\n\n")
}

func AWSElasticBeanstalkEnvironmentBP()  {
	color.Green("\nresource \"aws_elastic_beanstalk_application\" \"tftest\" {\n  name        = \"tf-test-name\"\n  description = \"tf-test-desc\"\n}\n\nresource \"aws_elastic_beanstalk_environment\" \"tfenvtest\" {\n  name                = \"tf-test-name\"\n  application         = aws_elastic_beanstalk_application.tftest.name\n  solution_stack_name = \"64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4\"\n}")
	color.Yellow("\n\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elastic_beanstalk_environment\n\n")
}