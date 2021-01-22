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
