package resourcebps

import "github.com/fatih/color"

func AWSCloudFrontDistributionPrompt() {
	color.Green("\nresource \"aws_cloudfront_distribution\" \"foo\" {\n  enabled = false\n  default_cache_behavior {\n    allowed_methods = []\n    cached_methods = []\n    target_origin_id = \"\"\n    viewer_protocol_policy = \"\"\n    forwarded_values {\n      query_string = false\n      cookies {\n        forward = \"\"\n      }\n    }\n  }\n  origin {\n    domain_name = \"\"\n    origin_id = \"\"\n  }\n  restrictions {\n    geo_restriction {\n      restriction_type = \"\"\n    }\n  }\n  viewer_certificate {}\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution\n\n")
}
