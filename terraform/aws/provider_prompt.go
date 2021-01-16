package aws

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/utils"
	"github.com/manifoldco/promptui"
)

func ProviderPrompt() {

	var promptOrder []string
	prompts := map[string]types.TfPrompt{}

	color.Red("\nWe recommend not providing us your access information.\n" +
		"We however assure that we use your information only to create your terraform configuration\n")

	prompts["access_key"] = types.TfPrompt{
		Label: "Enter your access_key: \nThis is the AWS access key. It must be provided, \n" +
			"but it can also be sourced from the AWS_ACCESS_KEY_ID \n" +
			"environment variable, or via a shared \ncredentials file if profile is specified",
		Prompt: promptui.Prompt{
			Label: "",
			Mask:  '*',
		},
	}
	promptOrder = append(promptOrder, "access_key")

	prompts["secret_key"] = types.TfPrompt{
		Label: "Enter your secret_key: \nThis is the AWS secret key. It must be provided, \n" +
			"but it can also be sourced from the AWS_SECRET_ACCESS_KEY \n" +
			"environment variable, or via a shared \ncredentials file if profile is specified",
		Prompt: promptui.Prompt{
			Label: "",
			Mask:  '*',
		},
	}
	promptOrder = append(promptOrder, "secret_key")

	prompts["profile"] = types.TfPrompt{
		Label: "Enter profile:\nThis is the AWS profile name as set in the shared credentials file",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "profile")

	prompts["max_retries"] = types.TfPrompt{
		Label: "Enter max_retries:\nThis is the maximum number of times an API call is retried, \n" +
			"in the case where requests are being throttled or experiencing \n" +
			"transient failures. The delay between the subsequent API calls \n" +
			"increases exponentially. If omitted, the default value is 25.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "max_retries")

	prompts["allowed_account_ids"] = types.TfPrompt{
		Label: "Enter allowed_account_ids, eg:[\"a\",\"b\",\"c\"]:\nList of allowed AWS account IDs to prevent you from \n" +
			"mistakenly using an incorrect one (and potentially end up destroying \n" +
			"a live environment). Conflicts with forbidden_account_ids.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "allowed_account_ids")

	prompts["token"] = types.TfPrompt{
		Label: "Enter token: Session token for validating temporary credentials",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "token")

	prompts["shared_credentials_file"] = types.TfPrompt{
		Label: "Enter shared_credentials_file :\n(Optional) This is the path to the shared " +
			"credentials file. If this is not set and a profile is specified," +
			" ~/.aws/credentials will be used.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "shared_credentials_file")

	prompts["insecure"] = types.TfPrompt{
		Label: "Enter insecure(true/false):\nExplicitly allow the provider to perform \"insecure\" SSL \n" +
			"requests(bool).If omitted, the default value is false",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "insecure")

	prompts["skip_credentials_validation"] = types.TfPrompt{
		Label: "Enter skip_credentials_validation(true/false):\nSkip the credentials validation via the STS API.\n" +
			"Useful for AWS API implementations that do not have STS available or implemented.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}

	promptOrder = append(promptOrder, "skip_credentials_validation")

	prompts["skip_get_ec2_platforms"] = types.TfPrompt{
		Label: "Enter skip_get_ec2_platforms(true/false):\nSkip getting the supported EC2 platforms. \n" +
			"Used by users that don't have ec2:DescribeAccountAttributes permissions.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "skip_get_ec2_platforms")

	prompts["skip_metadata_api_check"] = types.TfPrompt{
		Label: "Enter skip_metadata_api_check(true/false):\nSkip the AWS Metadata API check. \n" +
			"Useful for AWS API implementations that do not have a metadata \n" +
			"API endpoint. Setting to true prevents Terraform from authenticating via the Metadata API.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "skip_metadata_api_check")

	prompts["skip_requesting_account_id"] = types.TfPrompt{
		Label: "Enter skip_requesting_account_id(true/false):\nSkip requesting the account ID. \n" +
			"Useful for AWS API implementations that do not \n" +
			"have the IAM, STS API, or metadata API.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "skip_requesting_account_id")

	prompts["skip_region_validation"] = types.TfPrompt{
		Label: "Enter skip_region_validation(true/false):\nSkip validation of provided region name. \n" +
			"Useful for AWS-like implementations that use their own \n" +
			"region names or to bypass the validation for regions \nthat aren't publicly available yet.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "skip_region_validation")

	var selectOrder []string
	selects := map[string]types.TfSelect{}

	selects["region"] = RegionPrompt()

	selectOrder = append(selectOrder, "region")

	providerBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

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
		"3.policy\n4.policy_arns\n5.role_arn\n6.session_name\n7.tags(not supported by this cli yet)\n")

	assumeRolePrompts := map[string]types.TfPrompt{}
	var nestedOrder []string

	assumeRolePrompts["duration_seconds"] = types.TfPrompt{
		Label: "\nEnter duration_seconds:\n(Optional) Number of seconds to restrict the assume role session duration.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedOrder = append(nestedOrder, "duration_seconds")

	assumeRolePrompts["external_id"] = types.TfPrompt{
		Label: "Enter external_id:\n(Optional) External identifier to use when assuming the role.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "external_id")

	assumeRolePrompts["session_name"] = types.TfPrompt{
		Label: "Enter session_name:\n(Optional) Session name to use when assuming the role.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "session_name")

	assumeRolePrompts["role_arn"] = types.TfPrompt{
		Label: "Enter role_arn:\n(Optional) Amazon Resource Name (ARN) of the IAM Role to assume.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "role_arn")
	selectOrder = append(selectOrder, "assume_role")

	providerBlock["assume_role"] = builder.PSOrder(nestedOrder, nil, assumeRolePrompts, nil)

	ignoreTagsPrompts := map[string]types.TfPrompt{}

	color.Green("\nEnter ignore_tags:\nThe ignore_tags configuration block supports " +
		"\n1.keys\n2.key_prefixes\n")

	ignoreTagsPrompts["keys"] = types.TfPrompt{
		Label: "Enter keys([\"a\",\"b\",\"c\"]):(Optional) List of exact resource tag keys to ignore \nacross all resources handled by this provider." +
			"\nCheck https://registry.terraform.io/providers/hashicorp/aws/latest/docs#ignore_tags-configuration-block for more info.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "keys")

	ignoreTagsPrompts["key_prefixes"] = types.TfPrompt{
		Label: "Enter key_prefixes([\"a\",\"b\",\"c\"])(Optional) List of resource tag key prefixes to ignore across all resources handled by this provider." +
			"\nCheck https://registry.terraform.io/providers/hashicorp/aws/latest/docs#key_prefixes for more info",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "key_prefixes")
	selectOrder = append(selectOrder, "ignore_tags")

	providerBlock["ignore_tags"] = builder.PSOrder(nestedOrder[len(nestedOrder)-2:], nil, ignoreTagsPrompts, nil)

	builder.ProviderBuilder("aws", providerBlock)
}
