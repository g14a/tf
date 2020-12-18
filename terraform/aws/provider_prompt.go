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

	var order []string
	prompts := map[string]promptui.Prompt{}

	color.Red("\nWe recommend not providing us your access information.\n" +
		"We however assure that we use your information only to create your terraform configuration\n", "text")

	prompts["access_key"] = promptui.Prompt{
		Label: "Enter your access_key",
		Mask:  '*',
	}
	order = append(order, "access_key")

	prompts["secret_key"] = promptui.Prompt{
		Label: "Enter your secret_key",
		Mask:  '*',
	}
	order = append(order, "secret_key")

	prompts["profile"] = promptui.Prompt{
		Label: "Enter profile - This is the AWS profile name as set in the shared credentials file",
	}
	order = append(order, "profile")

	prompts["max_retries"] = promptui.Prompt{
		Label:    "Enter max_retries",
		Validate: utils.IntValidator,
	}
	order = append(order, "max_retries")

	prompts["allowed_account_ids"] = promptui.Prompt{
		Label: "Enter allowed_account_ids, eg:[a,b,c]",
	}
	order = append(order, "allowed_account_ids")

	prompts["insecure"] = promptui.Prompt{
		Label: "Enter insecure - Explicitly allow the provider to perform \"insecure\" SSL requests(bool)",
	}
	order = append(order, "insecure")

	prompts["token"] = promptui.Prompt{
		Label: "Enter token - Session token for validating temporary credentials",
	}
	order = append(order, "token")

	providerInfo := map[string]interface{}{}
	providerInfo["region"] = region

	for k, v := range prompts {
		value, err := v.Run()
		if err != nil {
			fmt.Println(err)
		}
		providerInfo[k] = value
	}

	selects := map[string]promptui.Select{}
	selects["skip_credentials_validation"] = promptui.Select{
		Label: "Enter skip_credentials_validation",
		Items: []string{"true","false"},
	}

	selects["skip_get_ec2_platforms"] = promptui.Select{
		Label: "Enter skip_get_ec2_platforms",
		Items: []string{"true","false"},
	}

	selects["skip_metadata_api_check"] = promptui.Select{
		Label: "Enter skip_metadata_api_check",
		Items: []string{"true","false"},
	}

	selects["skip_requesting_account_id"] = promptui.Select{
		Label: "Enter skip_requesting_account_id",
		Items: []string{"true","false"},
	}

	selects["skip_region_validation"] = promptui.Select{
		Label: "Enter skip_region_validation",
		Items: []string{"true","false"},
	}

	for k, v := range selects {
		_, value, err := v.Run()
		if err != nil {
			fmt.Println(err)
		}
		providerInfo[k] = value
	}

	builder.ProviderBuilder("aws", providerInfo)
}
