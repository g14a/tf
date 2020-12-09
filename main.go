package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
	"tf/file"
	"tf/terraform"
	"tf/terraform/aws"
)

func main() {

	file.FilePrompt()

	provider := promptui.Select{
		Label:             "Select Provider",
		Size:              20,
		StartInSearchMode: true,
		Items:             terraform.GetProviders(),
		Searcher: func(input string, index int) bool {
			provider := terraform.GetProviders()[index]
			name := strings.Replace(strings.ToLower(provider), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		},
	}

	_, tfProvider, err := provider.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Println(tfProvider)

	aws.ProviderPrompt()
}
