package resourcebps

import "github.com/fatih/color"

func AWSRoute53DelegationSetBP() {
	color.Green("\nresource \"aws_route53_delegation_set\" \"main\" {\n  reference_name = \"DynDNS\"\n}\n\nresource \"aws_route53_zone\" \"primary\" {\n  name              = \"hashicorp.com\"\n  delegation_set_id = aws_route53_delegation_set.main.id\n}\n\nresource \"aws_route53_zone\" \"secondary\" {\n  name              = \"terraform.io\"\n  delegation_set_id = aws_route53_delegation_set.main.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_delegation_set\n\n")
}

func AWSRoute53HealthCheckBP() {
	color.Green("\nresource \"aws_route53_health_check\" \"example\" {\n  fqdn              = \"example.com\"\n  port              = 80\n  type              = \"HTTP\"\n  resource_path     = \"/\"\n  failure_threshold = \"5\"\n  request_interval  = \"30\"\n\n  tags = {\n    Name = \"tf-test-health-check\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_health_check\n\n")
}

func AWSRoute53QueryLogBP() {
	color.Green("\nprovider \"aws\" {\n  alias  = \"us-east-1\"\n  region = \"us-east-1\"\n}\n\nresource \"aws_cloudwatch_log_group\" \"aws_route53_example_com\" {\n  provider = aws.us-east-1\n\n  name              = \"/aws/route53/${aws_route53_zone.example_com.name}\"\n  retention_in_days = 30\n}\n\n# Example CloudWatch log resource policy to allow Route53 to write logs\n# to any log group under /aws/route53/*\n\ndata \"aws_iam_policy_document\" \"route53-query-logging-policy\" {\n  statement {\n    actions = [\n      \"logs:CreateLogStream\",\n      \"logs:PutLogEvents\",\n    ]\n\n    resources = [\"arn:aws:logs:*:*:log-group:/aws/route53/*\"]\n\n    principals {\n      identifiers = [\"route53.amazonaws.com\"]\n      type        = \"Service\"\n    }\n  }\n}\n\nresource \"aws_cloudwatch_log_resource_policy\" \"route53-query-logging-policy\" {\n  provider = aws.us-east-1\n\n  policy_document = data.aws_iam_policy_document.route53-query-logging-policy.json\n  policy_name     = \"route53-query-logging-policy\"\n}\n\n# Example Route53 zone with query logging\n\nresource \"aws_route53_zone\" \"example_com\" {\n  name = \"example.com\"\n}\n\nresource \"aws_route53_query_log\" \"example_com\" {\n  depends_on = [aws_cloudwatch_log_resource_policy.route53-query-logging-policy]\n\n  cloudwatch_log_group_arn = aws_cloudwatch_log_group.aws_route53_example_com.arn\n  zone_id                  = aws_route53_zone.example_com.zone_id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_query_log\n\n")
}

func AWSRoute53RecordBP() {
	color.Green("\nresource \"aws_route53_record\" \"www\" {\n  zone_id = aws_route53_zone.primary.zone_id\n  name    = \"www.example.com\"\n  type    = \"A\"\n  ttl     = \"300\"\n  records = [aws_eip.lb.public_ip]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record\n\n")
}

func AWSRoute53VPCAssociationAuthorizationBP() {
	color.Green("\nresource \"aws_vpc\" \"example\" {\n  cidr_block           = \"10.6.0.0/16\"\n  enable_dns_hostnames = true\n  enable_dns_support   = true\n}\n\nresource \"aws_route53_zone\" \"example\" {\n  name = \"example.com\"\n\n  vpc {\n    vpc_id = aws_vpc.example.id\n  }\n}\n\nresource \"aws_vpc\" \"alternate\" {\n  provider = \"aws.alternate\"\n\n  cidr_block           = \"10.7.0.0/16\"\n  enable_dns_hostnames = true\n  enable_dns_support   = true\n}\n\nresource \"aws_route53_vpc_association_authorization\" \"example\" {\n  vpc_id  = aws_vpc.alternate.id\n  zone_id = aws_route53_zone.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_vpc_association_authorization\n\n")
}

func AWSRoute53ZoneBP() {
	color.Green("\nresource \"aws_route53_zone\" \"main\" {\n  name = \"example.com\"\n}\n\nresource \"aws_route53_zone\" \"dev\" {\n  name = \"dev.example.com\"\n\n  tags = {\n    Environment = \"dev\"\n  }\n}\n\nresource \"aws_route53_record\" \"dev-ns\" {\n  zone_id = aws_route53_zone.main.zone_id\n  name    = \"dev.example.com\"\n  type    = \"NS\"\n  ttl     = \"30\"\n  records = aws_route53_zone.dev.name_servers\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_zone\n\n")
}

func AWSRoute53ZoneAssociationBP() {
	color.Green("\nresource \"aws_vpc\" \"primary\" {\n  cidr_block           = \"10.6.0.0/16\"\n  enable_dns_hostnames = true\n  enable_dns_support   = true\n}\n\nresource \"aws_vpc\" \"secondary\" {\n  cidr_block           = \"10.7.0.0/16\"\n  enable_dns_hostnames = true\n  enable_dns_support   = true\n}\n\nresource \"aws_route53_zone\" \"example\" {\n  name = \"example.com\"\n\n  # NOTE: The aws_route53_zone vpc argument accepts multiple configuration\n  #       blocks. The below usage of the single vpc configuration, the\n  #       lifecycle configuration, and the aws_route53_zone_association\n  #       resource is for illustrative purposes (e.g. for a separate\n  #       cross-account authorization process, which is not shown here).\n  vpc {\n    vpc_id = aws_vpc.primary.id\n  }\n\n  lifecycle {\n    ignore_changes = [vpc]\n  }\n}\n\nresource \"aws_route53_zone_association\" \"secondary\" {\n  zone_id = aws_route53_zone.example.zone_id\n  vpc_id  = aws_vpc.secondary.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_zone_association\n\n")
}
