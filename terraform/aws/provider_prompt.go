package aws

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/utils"
)

func ProviderPrompt() {

	color.Green("\nSelect AWS Provider Info:\n\n", "text")

	_, region, err := RegionPrompt().Run()

	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts := map[string]promptui.Prompt{}

	color.Red("\nWe recommend not providing us your access information.\n"+
		"We however assure that we use your information only to create your terraform configuration\n", "text")

	prompts["access_key"] = promptui.Prompt{
		Label: "Enter your access_key",
		Mask:  '*',
	}
	promptOrder = append(promptOrder, "access_key")

	prompts["secret_key"] = promptui.Prompt{
		Label: "Enter your secret_key",
		Mask:  '*',
	}
	promptOrder = append(promptOrder, "secret_key")

	prompts["profile"] = promptui.Prompt{
		Label: "Enter profile - This is the AWS profile name as set in the shared credentials file",
	}
	promptOrder = append(promptOrder, "profile")

	prompts["max_retries"] = promptui.Prompt{
		Label:    "Enter max_retries",
		Validate: utils.IntValidator,
	}
	promptOrder = append(promptOrder, "max_retries")

	prompts["allowed_account_ids"] = promptui.Prompt{
		Label: "Enter allowed_account_ids, eg:[a,b,c]",
	}
	promptOrder = append(promptOrder, "allowed_account_ids")

	prompts["token"] = promptui.Prompt{
		Label: "Enter token - Session token for validating temporary credentials",
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
	selects := map[string]promptui.Select{}

	selects["insecure"] = promptui.Select{
		Label: "Enter insecure - Explicitly allow the provider to perform \"insecure\" SSL requests(bool)",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "insecure")

	selects["skip_credentials_validation"] = promptui.Select{
		Label: "Enter skip_credentials_validation",
		Items: []string{"true", "false"},
	}

	selectOrder = append(selectOrder, "skip_credentials_validation")

	selects["skip_get_ec2_platforms"] = promptui.Select{
		Label: "Enter skip_get_ec2_platforms",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "skip_get_ec2_platforms")

	selects["skip_metadata_api_check"] = promptui.Select{
		Label: "Enter skip_metadata_api_check",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "skip_metadata_api_check")

	selects["skip_requesting_account_id"] = promptui.Select{
		Label: "Enter skip_requesting_account_id",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "skip_requesting_account_id")

	selects["skip_region_validation"] = promptui.Select{
		Label: "Enter skip_region_validation",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "skip_region_validation")

	for _, v := range selectOrder {
		p := selects[v]
		_, value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		providerInfo[v] = value
	}

	providerInfo["assume_role"] = map[string]interface{}{
		"external_id": "sample id",
	}

	builder.ProviderBuilder("aws", providerInfo)
}
