package aws

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/utils"
)

func ProviderPrompt() {
	var prompt promptui.Prompt

	fmt.Println("AWS Provider Prompt...")

	_, region, err := AWSRegionPrompt().Run()

	prompt = promptui.Prompt{
		Label: "Enter your access_key",
		Mask:  '*',
	}

	accessKey, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	prompt = promptui.Prompt{
		Label: "Enter your secret_key",
		Mask:  '*',
	}

	secretKey, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	prompt = promptui.Prompt{
		Label: "Enter profile - This is the AWS profile name as set in the shared credentials file",
	}

	profile, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	prompt = promptui.Prompt{
		Label:    "Enter max_retries",
		Validate: utils.IntValidator,
	}

	maxRetries, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Enter allowed_account_ids, eg:[a,b,c]",
	}

	allowedAccountIds, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Enter insecure - Explicitly allow the provider to perform \"insecure\" SSL requests",
	}

	insecure, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Enter token - Session token for validating temporary credentials",
	}

	token, err := prompt.Run()

	providerInfo := map[string]interface{}{
		"region":      region,
		"access_key":  accessKey,
		"secret_key":  secretKey,
		"profile":     profile,
		"max_retries": maxRetries,
		"allowed_account_ids": allowedAccountIds,
		"insecure": insecure,
		"token": token,
	}

	builder.ProviderBuilder("aws", providerInfo)
}
