package terraform

import "github.com/g14a/tf/terraform/aws"

func SelectResourceTree(provider string, resource string, boilerplate bool) {
	switch provider {
	case "aws":
		aws.ResourcePrompt(resource, boilerplate)
	}
}
