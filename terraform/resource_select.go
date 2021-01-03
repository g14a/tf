package terraform

import "tf/terraform/aws"

func SelectResourceTree(provider string, resource string) {
	switch provider {
	case "aws":
		aws.ResourcePrompt(resource)
	}
}
