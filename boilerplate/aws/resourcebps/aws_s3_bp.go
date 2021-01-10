package resourcebps

import "github.com/fatih/color"

func AWSS3BucketBP() {
	color.Green("\nresource \"aws_s3_bucket\" \"foo\" {\n  bucket = \"my-tf-test-bucket\"\n  acl    = \"private\"\n\n  tags = {\n    Name        = \"My bucket\"\n    Environment = \"Dev\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket\n\n")
}

func AWSS3AccessPointBP()  {
	color.Green("\nresource \"aws_s3_bucket\" \"example\" {\n  bucket = \"example\"\n}\n\nresource \"aws_s3_access_point\" \"example\" {\n  bucket = aws_s3_bucket.example.id\n  name   = \"example\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_access_point\n\n")
}

func AWSS3AccountPublicAccessBlockBP()  {
	color.Green("\nresource \"aws_s3_account_public_access_block\" \"example\" {\n  block_public_acls   = true\n  block_public_policy = true\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_account_public_access_block\n\n")
}

func AWSS3BucketAnalyticsConfigurationBP()  {
	color.Green("\nresource \"aws_s3_bucket_analytics_configuration\" \"example-entire-bucket\" {\n  bucket = aws_s3_bucket.example.bucket\n  name   = \"EntireBucket\"\n\n  storage_class_analysis {\n    data_export {\n      destination {\n        s3_bucket_destination {\n          bucket_arn = aws_s3_bucket.analytics.arn\n        }\n      }\n    }\n  }\n}\n\nresource \"aws_s3_bucket\" \"example\" {\n  bucket = \"example\"\n}\n\nresource \"aws_s3_bucket\" \"analytics\" {\n  bucket = \"analytics destination\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_analytics_configuration\n\n")
}

func AWSS3BucketInventoryBP() {
	color.Green("\nresource \"aws_s3_bucket\" \"test\" {\n  bucket = \"my-tf-test-bucket\"\n}\n\nresource \"aws_s3_bucket\" \"inventory\" {\n  bucket = \"my-tf-inventory-bucket\"\n}\n\nresource \"aws_s3_bucket_inventory\" \"test\" {\n  bucket = aws_s3_bucket.test.id\n  name   = \"EntireBucketDaily\"\n\n  included_object_versions = \"All\"\n\n  schedule {\n    frequency = \"Daily\"\n  }\n\n  destination {\n    bucket {\n      format     = \"ORC\"\n      bucket_arn = aws_s3_bucket.inventory.arn\n    }\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_inventory\n\n")
}

func AWSS3BucketMetricBP()  {
	color.Green("\nresource \"aws_s3_bucket\" \"example\" {\n  bucket = \"example\"\n}\n\nresource \"aws_s3_bucket_metric\" \"example-filtered\" {\n  bucket = aws_s3_bucket.example.bucket\n  name   = \"ImportantBlueDocuments\"\n\n  filter {\n    prefix = \"documents/\"\n\n    tags = {\n      priority = \"high\"\n      class    = \"blue\"\n    }\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_metric\n\n")
}

func AWSS3BucketNotificationBP()  {
	color.Green("\nresource \"aws_sns_topic\" \"topic\" {\n  name = \"s3-event-notification-topic\"\n\n  policy = <<POLICY\n{\n    \"Version\":\"2012-10-17\",\n    \"Statement\":[{\n        \"Effect\": \"Allow\",\n        \"Principal\": {\"AWS\":\"*\"},\n        \"Action\": \"SNS:Publish\",\n        \"Resource\": \"arn:aws:sns:*:*:s3-event-notification-topic\",\n        \"Condition\":{\n            \"ArnLike\":{\"aws:SourceArn\":\"${aws_s3_bucket.bucket.arn}\"}\n        }\n    }]\n}\nPOLICY\n}\n\nresource \"aws_s3_bucket\" \"bucket\" {\n  bucket = \"your_bucket_name\"\n}\n\nresource \"aws_s3_bucket_notification\" \"bucket_notification\" {\n  bucket = aws_s3_bucket.bucket.id\n\n  topic {\n    topic_arn     = aws_sns_topic.topic.arn\n    events        = [\"s3:ObjectCreated:*\"]\n    filter_suffix = \".log\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_notification\n\n")
}

func AWSS3BucketObjectBP()  {
	color.Green("\nresource \"aws_s3_bucket_object\" \"object\" {\n  bucket = \"your_bucket_name\"\n  key    = \"new_object_key\"\n  source = \"path/to/file\"\n\n  # The filemd5() function is available in Terraform 0.11.12 and later\n  # For Terraform 0.11.11 and earlier, use the md5() function and the file() function:\n  # etag = \"${md5(file(\"path/to/file\"))}\"\n  etag = filemd5(\"path/to/file\")\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_object\n\n")
}

func AWSS3BucketOwnershipControlsBP()  {
	color.Green("\nresource \"aws_s3_bucket\" \"example\" {\n  bucket = \"example\"\n}\n\nresource \"aws_s3_bucket_ownership_controls\" \"example\" {\n  bucket = aws_s3_bucket.example.id\n\n  rule {\n    object_ownership = \"BucketOwnerPreferred\"\n  }\n}\n\n")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_ownership_controls\n\n")
}

func AWSS3BucketPolicyBP()  {
	color.Green("\nresource \"aws_s3_bucket\" \"b\" {\n  bucket = \"my_tf_test_bucket\"\n}\n\nresource \"aws_s3_bucket_policy\" \"b\" {\n  bucket = aws_s3_bucket.b.id\n\n  policy = <<POLICY\n{\n  \"Version\": \"2012-10-17\",\n  \"Id\": \"MYBUCKETPOLICY\",\n  \"Statement\": [\n    {\n      \"Sid\": \"IPAllow\",\n      \"Effect\": \"Deny\",\n      \"Principal\": \"*\",\n      \"Action\": \"s3:*\",\n      \"Resource\": \"arn:aws:s3:::my_tf_test_bucket/*\",\n      \"Condition\": {\n         \"IpAddress\": {\"aws:SourceIp\": \"8.8.8.8/32\"}\n      }\n    }\n  ]\n}\nPOLICY\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_policy\n\n")
}

func AWSS3BucketPolicyAccessBlockBP()  {
	color.Green("\nresource \"aws_s3_bucket\" \"example\" {\n  bucket = \"example\"\n}\n\nresource \"aws_s3_bucket_public_access_block\" \"example\" {\n  bucket = aws_s3_bucket.example.id\n\n  block_public_acls   = true\n  block_public_policy = true\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_public_access_block\n\n")
}