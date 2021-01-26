package resourceprompts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
)

func AWSRoute53DelegationSetPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "reference_name",
			Ex:    "",
			Doc:   "(Optional) This is a reference name used in Caller Reference (helpful for identifying single delegation set amongst others)",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_route53_delegation_set", blockName, resourceBlock)
}

func AWSRoute53HealthCheckPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "reference_name",
			Ex:    "",
			Doc:   "(Optional) This is a reference name used in Caller Reference (helpful for identifying single health_check set amongst others)",
		},
		{
			Field: "fqdn",
			Ex:    "",
			Doc:   "(Optional) The fully qualified domain name of the endpoint to be checked.",
		},
		{
			Field: "ip_address",
			Ex:    "",
			Doc:   "(Optional) The IP address of the endpoint to be checked.",
		},
		{
			Field: "port",
			Ex:    "",
			Doc:   "(Optional) The port of the endpoint to be checked.",
		},
		{
			Type:  "select",
			Field: "type",
			Doc:   "(Required) The protocol to use when performing health checks.",
			Items: []string{"HTTP", "HTTPS", "HTTP_STR_MATCH", "HTTPS_STR_MATCH", "TCP", "CALCULATED", "CLOUDWATCH_METRIC"},
		},
		{
			Field: "failure_threshold",
			Ex:    "",
			Doc:   "(Required) The number of consecutive health checks that an endpoint must pass or fail.",
		},
		{
			Field: "request_interval",
			Ex:    "",
			Doc:   "(Required) The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.",
		},
		{
			Field: "resource_path",
			Ex:    "",
			Doc:   "(Optional) The path that you want Amazon Route 53 to request when performing health checks.",
		},
		{
			Field: "search_string",
			Ex:    "",
			Doc:   "(Optional) String searched in the first 5120 bytes of the response body for check to be considered healthy",
		},
		{
			Field: "measure_latency",
			Ex:    "",
			Doc:   "(Optional) A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.",
		},
		{
			Field:     "invert_healthcheck",
			Ex:        "(true/false)",
			Doc:       "(Optional) A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "disabled",
			Ex:    "(true/false)",
			Doc: "(Optional) A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:" +
				"\n 	For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource." +
				"\n		For calculated health checks, Route 53 stops aggregating the status of the referenced health checks." +
				"\n 	For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.",
		},
		{
			Field: "enable_sni",
			Ex:    "(true/false)",
			Doc:   "(Optional) A boolean value that indicates whether Route53 should send the fqdn to the endpoint when performing the health check. This defaults to AWS' defaults: when the type is \"HTTPS\" enable_sni defaults to true, when type is anything else enable_sni defaults to false.",
		},
		{
			Field: "child_healthchecks",
			Ex:    "",
			Doc:   "(Optional) For a specified parent health check, a list of HealthCheckId values for the associated child health checks.",
		},
		{
			Field:     "child_health_threshold",
			Ex:        "",
			Doc:       "(Optional) The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive",
			Validator: validators.MinMaxIntValidator(0, 256),
		},
		{
			Field: "cloudwatch_alarm_name",
			Ex:    "",
			Doc:   "(Optional) The name of the CloudWatch alarm.",
		},
		{
			Field: "cloudwatch_alarm_region",
			Ex:    "",
			Doc:   "(Optional) The CloudWatchRegion that the CloudWatch alarm was created in.",
		},
		{
			Type:  "select",
			Field: "insufficient_data_health_status",
			Doc:   "(Optional) The status of the health check when CloudWatch has insufficient data about the state of associated alarm.",
			Items: []string{"Healthy", "Unhealthy", "LastKnownStatus"},
		},
		{
			Field: "regions",
			Ex:    "",
			Doc:   "(Optional) A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the health check.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_route53_health_check", blockName, resourceBlock)
}

func AWSRoute53QueryLogPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "cloudwatch_log_group_arn",
			Ex:    "",
			Doc:   "(Required) CloudWatch log group ARN to send query logs.",
		},
		{
			Field: "zone_id",
			Ex:    "",
			Doc:   "(Required) Route53 hosted zone ID to enable query logs.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_route53_query_log", blockName, resourceBlock)
}

func AWSRoute53RecordPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "zone_id",
			Ex:    "",
			Doc:   "(Required) The ID of the hosted zone to contain this record.",
		},
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) The name of the record.",
		},
		{
			Type:  "select",
			Field: "type",
			Doc:   "(Required) The record type.",
			Items: []string{"A", "AAAA", "CAA", "CNAME", "MX", "NAPTR", "NS", "PTR", "SOA", "SPF", "SRV", "TXT"},
		},
		{
			Field:     "ttl",
			Ex:        "10",
			Doc:       "(Required for non-alias records) The TTL of the record.",
			Validator: validators.IntValidator,
		},
		{
			Field: "records",
			Ex:    "",
			Doc: "(Required for non-alias records) A string list of records. To specify a single record value " +
				"\nlonger than 255 characters such as a TXT record for DKIM, add \\\"\\\" inside the " +
				"\nTerraform configuration string (e.g. \"first255characters\\\"\\\"morecharacters\")",
		},
		{
			Field: "set_identifier",
			Ex:    "",
			Doc: "(Optional) Unique identifier to differentiate records with routing policies from " +
				"\none another. Required if using failover, geolocation, latency, or weighted routing " +
				"\npolicies documented below.",
		},
		{
			Field: "health_check_id",
			Ex:    "",
			Doc:   "(Optional) The health check the record should be associated with.",
		},
		{
			Field: "allow_overwrite",
			Ex:    "(true/false)",
			Doc: "(Optional) Allow creation of this record in Terraform to overwrite an existing " +
				"\nrecord, if any. This does not affect the ability to update the record in " +
				"\nTerraform and does not prevent other resources within Terraform or manual " +
				"\nRoute 53 changes outside Terraform from overwriting this record. false by default. " +
				"\nThis configuration is not recommended for most environments.",
			Validator: validators.BoolValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like alias/failover_routing_policies/geolocation? etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_route53_record", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter alias: (Optional) An alias block. Conflicts with ttl & records:" +
		"\nThe alias block supports the following arguments:" +
		"\n1.name\n2.zone_id\n3.evaluate_target_health")
	aliasSchema := []types.Schema{
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) DNS domain name for a CloudFront distribution, S3 bucket, ELB, or another resource record set in this hosted zone.",
		},
		{
			Field: "zone_id",
			Ex:    "",
			Doc:   "(Required) Hosted zone ID for a CloudFront distribution, S3 bucket, ELB, or Route 53 hosted zone",
		},
		{
			Field: "evaluate_target_health",
			Ex:    "",
			Doc: "(Required) Set to true if you want Route 53 to determine whether to respond to DNS " +
				"\nqueries using this resource record set by checking the health of the resource record set." +
				"\nCheckout https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values.html#rrsets-values-alias-evaluate-target-health",
			Validator: validators.BoolValidator,
		},
	}

	resourceBlock["alias"] = builder.PSOrder(types.ProvidePS(aliasSchema))

	failoverRoutingPolicySchema := []types.Schema{
		{
			Type:  "select",
			Field: "type",
			Doc:   "(Required) PRIMARY or SECONDARY. A PRIMARY record will be served if its healthcheck is passing, otherwise the SECONDARY will be served.",
			Items: []string{"PRIMARY", "SECONDARY"},
		},
	}
	resourceBlock["failover_routing_policy"] = builder.PSOrder(types.ProvidePS(failoverRoutingPolicySchema))

	geolocationRoutingPolicySchema := []types.Schema{
		{
			Field: "continent",
			Ex:    "",
			Doc: "A two-letter continent code. Either continent or country must be specified." +
				"\nCheckout http://docs.aws.amazon.com/Route53/latest/APIReference/API_GetGeoLocation.html",
		},
		{
			Field: "country",
			Ex:    "",
			Doc:   "A two-character country code or * to indicate a default resource record set.",
		},
		{
			Field: "subdivision",
			Ex:    "",
			Doc:   "(Optional) A subdivision code for a country.",
		},
	}

	resourceBlock["geolocation_routing_policy"] = builder.PSOrder(types.ProvidePS(geolocationRoutingPolicySchema))

	latencyRoutingPolicySchema := []types.Schema{
		{
			Field: "region",
			Ex:    "",
			Doc:   "(Required) An AWS region from which to measure latency.",
		},
	}

	resourceBlock["latency_routing_policy"] = builder.PSOrder(types.ProvidePS(latencyRoutingPolicySchema))

	weightedRoutingPolicySchema := []types.Schema{
		{
			Field:     "weight",
			Ex:        "",
			Doc:       "(Required) A numeric value indicating the relative weight of the record.",
			Validator: validators.IntValidator,
		},
	}

	resourceBlock["latency_routing_policy"] = builder.PSOrder(types.ProvidePS(weightedRoutingPolicySchema))

	builder.ResourceBuilder("aws_route53_record", blockName, resourceBlock)
}

func AWSRoute53VPCAssociationAuthorizationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "zone_id",
			Ex:    "",
			Doc:   "(Required) The ID of the private hosted zone that you want to authorize associating a VPC with.",
		},
		{
			Field: "vpc_id",
			Ex:    "",
			Doc:   "(Required) The VPC to authorize for association with the private hosted zone.",
		},
		{
			Field: "vpc_region",
			Ex:    "",
			Doc:   "(Optional) The VPC's region. Defaults to the region of the AWS provider.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_route53_vpc_association_authorization", blockName, resourceBlock)
}

func AWSRoute53ZonePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) This is the name of the hosted zone.",
		},
		{
			Field: "comment",
			Ex:    "",
			Doc:   "(Optional) A comment for the hosted zone. Defaults to 'Managed by Terraform'.",
		},
		{
			Field: "delegation_set_id",
			Ex:    "",
			Doc: "(Optional) The ID of the reusable delegation set whose NS records you " +
				"\nwant to assign to the hosted zone. Conflicts with vpc as delegation " +
				"\nsets can only be used for public zones.",
		},
		{
			Field:     "force_destroy",
			Ex:        "(true/false)",
			Doc:       "(Optional) Whether to destroy all records (possibly managed outside of Terraform) in the zone when destroying the zone.",
			Validator: validators.BoolValidator,
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the zone.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Green("\nEnter vpc:\n(Optional) Configuration block(s) specifying VPC(s) to associate with a private hosted zone. " +
		"\nConflicts with the delegation_set_id argument in this resource" +
		"\nThe vpc block supports the following arguments:" +
		"\n1.vpc_id\n2.vpc_region")

	vpcSchema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "",
			Doc:   "(Required) ID of the VPC to associate.",
		},
		{
			Field: "vpc_region",
			Ex:    "",
			Doc:   "(Optional) Region of the VPC to associate. Defaults to AWS provider region.",
		},
	}

	resourceBlock["vpc"] = builder.PSOrder(types.ProvidePS(vpcSchema))

	builder.ResourceBuilder("aws_route53_zone", blockName, resourceBlock)
}

func AWSRoute53ZoneAssociationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	color.Yellow("\nUnless explicit association ordering is required (e.g. a separate cross-account association " +
		"\nauthorization), usage of this resource is not recommended. Use the vpc configuration blocks available within " +
		"\nthe aws_route53_zone resource instead.")

	schema := []types.Schema{
		{
			Field: "zone_id",
			Ex:    "",
			Doc:   "(Required) The private hosted zone to associate.",
		},
		{
			Field: "vpc_id",
			Ex:    "",
			Doc:   "(Required) The VPC to associate with the private hosted zone.",
		},
		{
			Field: "vpc_region",
			Ex:    "",
			Doc:   "(Optional) The VPC's region. Defaults to the region of the AWS provider.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_route53_zone_association", blockName, resourceBlock)
}
