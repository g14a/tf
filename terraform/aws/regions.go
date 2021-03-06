package aws

import (
	"strings"

	"github.com/fatih/color"

	"github.com/g14a/tf/types"
	"github.com/manifoldco/promptui"
)

func GetRegions() []string {
	return []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2",
		"af-south-1", "ap-east-1", "ap-northeast-2", "ap-southeast-1",
		"ap-southeast-2", "ap-northeast-1", "ca-central-1", "eu-central-1",
		"eu-west-1", "eu-west-2", "eu-south-1", "eu-west-3", "eu-north-1",
		"me-south-1", "sa-east-1",
	}
}

func RegionPrompt() types.TfSelect {

	color.Green("Select one of the AWS regions:")

	return types.TfSelect{
		Label: "",
		Select: promptui.Select{
			Label:             "",
			Size:              20,
			Items:             GetRegions(),
			StartInSearchMode: true,
			Searcher: func(input string, index int) bool {
				provider := GetRegions()[index]
				name := strings.Replace(strings.ToLower(provider), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			},
		},
	}
}
