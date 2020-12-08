package aws

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"reflect"
	"tf/builder"
	"tf/utils"
)

func AWSProviderPrompt() {
	fmt.Println("AWS Provider Prompt...")

	_, region, err := AWSRegionPrompt().Run()

	accessKeyPrompt := promptui.Prompt{
		Label: "Enter your access_key",
		Mask:  '*',
	}

	accessKey, err := accessKeyPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	secretKeyPrompt := promptui.Prompt{
		Label: "Enter your secret_key",
		Mask:  '*',
	}

	secretKey, err := secretKeyPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	profilePrompt := promptui.Prompt{
		Label: "Enter profile - This is the AWS profile name as set in the shared credentials file",
	}

	profile, err := profilePrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	maxRetriesPrompt := promptui.Prompt{
		Label:    "Enter Max Retries",
		Validate: utils.IntValidator,
	}

	maxRetries, err := maxRetriesPrompt.Run()

	fmt.Println(reflect.TypeOf(maxRetries), "============")

	providerInfo := map[string]interface{}{
		"region":      region,
		"access_key":  accessKey,
		"secret_key":  secretKey,
		"profile":     profile,
		"max_retries": maxRetries,
	}

	builder.ProviderBuilder("aws", providerInfo)
}
