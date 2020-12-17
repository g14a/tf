package terraform

import "tf/terraform/aws"

func SelectResourceTree(provider string)  {
	switch provider {
	case "aws":
		aws.ResourcePrompt()
	}
}