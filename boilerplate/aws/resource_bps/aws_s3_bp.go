package resource_bps

import "github.com/fatih/color"

func AWSS3BucketBP() {
	color.Green("\nresource \"aws_s3_bucket\" \"foo\" {\n  bucket = \"my-tf-test-bucket\"\n  acl    = \"private\"\n\n  tags = {\n    Name        = \"My bucket\"\n    Environment = \"Dev\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket\n\n")
}
