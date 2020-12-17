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

	prompts := map[string]promptui.Prompt{}

	prompts["access_key"] = promptui.Prompt{
		Label: "Enter your access_key",
		Mask:  '*',
	}

	prompts["secret_key"] = promptui.Prompt{
		Label: "Enter your secret_key",
		Mask:  '*',
	}

	prompts["profile"] = promptui.Prompt{
		Label: "Enter profile - This is the AWS profile name as set in the shared credentials file",
	}

	prompts["max_retries"] = promptui.Prompt{
		Label:    "Enter max_retries",
		Validate: utils.IntValidator,
	}

	prompts["allowed_account_ids"] = promptui.Prompt{
		Label: "Enter allowed_account_ids, eg:[a,b,c]",
	}

	prompts["insecure"] = promptui.Prompt{
		Label: "Enter insecure - Explicitly allow the provider to perform \"insecure\" SSL requests",
	}

	prompts["token"] = promptui.Prompt{
		Label: "Enter token - Session token for validating temporary credentials",
	}

	providerInfo := map[string]interface{}{}
	providerInfo["region"] = region

	for k, v := range prompts {
		value, err := v.Run()
		if err != nil {
			fmt.Println(err)
		}
		providerInfo[k] = value
	}

	builder.ProviderBuilder("aws", providerInfo)
}
