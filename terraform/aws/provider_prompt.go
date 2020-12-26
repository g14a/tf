package aws

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func ProviderPrompt() {

	color.Green("\nSelect AWS Provider Info:\n\n", "text")

	_, region, err := RegionPrompt().Run()

	if err != nil {
		fmt.Println("came in")
		fmt.Println(err)
	}

	var promptOrder []string
	prompts := map[string]types.TfPrompt{}

	color.Red("\nWe recommend not providing us your access information.\n"+
		"We however assure that we use your information only to create your terraform configuration\n")

	prompts["access_key"] = types.TfPrompt{
		Label: "Enter your access_key: \nThis is the AWS access key. It must be provided, but it can also be sourced from the AWS_ACCESS_KEY_ID environment variable, or via a shared credentials file if profile is specified",
		Prompt: promptui.Prompt{
			Label: "",
			Mask: '*',
		},
	}
	promptOrder = append(promptOrder, "access_key")

	prompts["secret_key"] = types.TfPrompt{
		Label: "Enter your secret_key: \nThis is the AWS secret key. It must be provided, but it can also be sourced from the AWS_SECRET_ACCESS_KEY environment variable, or via a shared credentials file if profile is specified",
		Prompt: promptui.Prompt{
			Label: "",
			Mask: '*',
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
		Label:    "Enter max_retries:\nThis is the maximum number of times an API call is retried, in the case where requests are being throttled or experiencing transient failures. The delay between the subsequent API calls increases exponentially. If omitted, the default value is 25.",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "max_retries")

	prompts["allowed_account_ids"] = types.TfPrompt{
		Label: "Enter allowed_account_ids, eg:[a,b,c]:\nList of allowed AWS account IDs to prevent you from mistakenly using an incorrect one (and potentially end up destroying a live environment). Conflicts with forbidden_account_ids.",
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

	providerInfo := map[string]interface{}{}
	providerInfo["region"] = region

	for _, v := range promptOrder {
		p := prompts[v]
		value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		providerInfo[v] = value
	}

	var selectOrder []string
	selects := map[string]types.TfSelect{}

	selects["insecure"] = types.TfSelect{
		Label: "Enter insecure:\n Explicitly allow the provider to perform \"insecure\" SSL requests(bool).If omitted, the default value is false",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "insecure")

	selects["skip_credentials_validation"] = types.TfSelect{
		Label: "Enter skip_credentials_validation:\nSkip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or implemented.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}

	selectOrder = append(selectOrder, "skip_credentials_validation")

	selects["skip_get_ec2_platforms"] = types.TfSelect{
		Label: "Enter skip_get_ec2_platforms:\nSkip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "skip_get_ec2_platforms")

	selects["skip_metadata_api_check"] = types.TfSelect{
		Label: "Enter skip_metadata_api_check:\nSkip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Terraform from authenticating via the Metadata API.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "skip_metadata_api_check")

	selects["skip_requesting_account_id"] = types.TfSelect{
		Label: "Enter skip_requesting_account_id:\nSkip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "skip_requesting_account_id")

	selects["skip_region_validation"] = types.TfSelect{
		Label: "Enter skip_region_validation:\nSkip validation of provided region name. Useful for AWS-like implementations that use their own region names or to bypass the validation for regions that aren't publicly available yet.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "skip_region_validation")

	for _, v := range selectOrder {
		p := selects[v]
		value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		providerInfo[v] = value
	}

	providerInfo["assume_role"] = map[string]interface{}{
		"external_id": "sample id",
	}

	builder.ProviderBuilder("aws", promptOrder, selectOrder, providerInfo)
}
