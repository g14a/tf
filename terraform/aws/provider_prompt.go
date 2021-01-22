package aws

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
)

func ProviderPrompt() {

	fmt.Println()

	region, err := RegionPrompt().Run()
	if err != nil {
		fmt.Println(err)
	}

	color.Red("\nWe recommend not providing us your access information.\n" +
		"We however assure that we use your information only to create your terraform configuration\n")

	schema := []types.Schema{
		{
			Field: "access_key",
			Ex:    "",
			Doc:   "(Optional) This is the AWS access key. It must be provided, but it can also be sourced from the AWS_ACCESS_KEY_ID environment variable, or via a shared credentials file if profile is specified.",
		},
		{
			Field: "secret_key",
			Ex:    "",
			Doc:   "(Optional) This is the AWS secret key. It must be provided, but it can also be sourced from the AWS_SECRET_ACCESS_KEY environment variable, or via a shared credentials file if profile is specified.",
		},
		{
			Field: "profile",
			Ex:    "",
			Doc:   "(Optional) This is the AWS profile name as set in the shared credentials file.",
		},
		{
			Field: "shared_credentials_file",
			Ex:    "",
			Doc:   "(Optional) This is the path to the shared credentials file. If this is not set and a profile is specified, ~/.aws/credentials will be used.",
		},
		{
			Field:     "max_retries",
			Ex:        "25",
			Doc:       "(Optional) This is the maximum number of times an API call is retried, in the case where requests are being throttled or experiencing transient failures. The delay between the subsequent API calls increases exponentially. If omitted, the default value is 25",
			Validator: validators.IntValidator,
		},
		{
			Field: "allowed_account_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc:   "(Optional) List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one (and potentially end up destroying a live environment). Conflicts with forbidden_account_ids",
		},
		{
			Field: "forbidden_account_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc:   "(Optional) List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with allowed_account_ids",
		},
		{
			Field: "token",
			Ex:    "",
			Doc:   "(Optional) Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials. It can also be sourced from the AWS_SESSION_TOKEN environment variable.",
		},
		{
			Field:     "insecure",
			Ex:        "(true/false)",
			Doc:       "(Optional) Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted, the default value is false.",
			Validator: validators.BoolValidator,
		},
		{
			Field:     "skip_credentials_validation",
			Ex:        "(true/false)",
			Doc:       "(Optional) Skip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or implemented.",
			Validator: validators.BoolValidator,
		},
		{
			Field:     "skip_get_ec2_platforms",
			Ex:        "(true/false)",
			Doc:       "(Optional) Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.",
			Validator: validators.BoolValidator,
		},
		{
			Field:     "skip_region_validation",
			Ex:        "(true/false)",
			Doc:       "(Optional) Skip validation of provided region name. Useful for AWS-like implementations that use their own region names or to bypass the validation for regions that aren't publicly available yet.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "skip_requesting_account_id",
			Ex:    "(true/false)",
			Doc: "(Optional) Skip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API." +
				"\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs#skip_requesting_account_id for extra information.",
			Validator: validators.BoolValidator,
		},
		{
			Field:     "skip_metadata_api_check",
			Ex:        "(true/false)",
			Doc:       "(Optional) Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
			Validator: validators.BoolValidator,
		},
	}

	providerBlock := builder.PSOrder(types.ProvidePS(schema))
	providerBlock["region"] = region

	color.Yellow("\nConfigure nested settings like assume_role/ignore_tags [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ProviderBuilder("aws", providerBlock)
		return
	}

	color.Green("\nEnter assume_role:\nThe assume_role configuration block supports " +
		"the following optional arguments:\n1.duration_seconds\n2.external_id\n" +
		"3.policy\n4.policy_arns\n5.role_arn\n6.session_name\n\n")

	assumeRoleSchema := []types.Schema{
		{
			Field:     "duration_seconds",
			Ex:        "10",
			Doc:       "(Optional) Number of seconds to restrict the assume role session duration.",
			Validator: validators.IntValidator,
		},
		{
			Field: "external_id",
			Ex:    "",
			Doc:   "(Optional) External identifier to use when assuming the role.",
		},
		{
			Field: "policy_arns",
			Ex:    "",
			Doc:   "(Optional) Set of Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.",
		},
		{
			Field: "session_name",
			Ex:    "",
			Doc:   "(Optional) Session name to use when assuming the role.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) Map of assume role session tags.",
			Validator: validators.RCValidator,
		},
		{
			Field: "transitive_tag_keys",
			Ex:    "",
			Doc:   "(Optional) Set of assume role session tag keys to pass to any subsequent sessions.",
		},
	}

	providerBlock["assume_role"] = builder.PSOrder(types.ProvidePS(assumeRoleSchema))

	color.Green("\nEnter ignore_tags:\nThe ignore_tags configuration block supports " +
		"\n1.keys\n2.key_prefixes\n")

	ignoreTagsSchema := []types.Schema{
		{
			Field: "keys",
			Ex:    "[\"a\",\"b\",\"c\"]",
			Doc:   "(Optional) List of exact resource tag keys to ignore across all resources handled by this provider. This configuration prevents Terraform from returning the tag in any tags attributes and displaying any configuration difference for the tag value. If any resource configuration still has this tag key configured in the tags argument, it will display a perpetual difference until the tag is removed from the argument or ignore_changes is also used.",
		},
		{
			Field: "key_prefixes",
			Ex:    "",
			Doc:   "(Optional) List of resource tag key prefixes to ignore across all resources handled by this provider. This configuration prevents Terraform from returning any tag key matching the prefixes in any tags attributes and displaying any configuration difference for those tag values. If any resource configuration still has a tag matching one of the prefixes configured in the tags argument, it will display a perpetual difference until the tag is removed from the argument or ignore_changes is also used.",
		},
	}

	providerBlock["ignore_tags"] = builder.PSOrder(types.ProvidePS(ignoreTagsSchema))

	builder.ProviderBuilder("aws", providerBlock)
}
