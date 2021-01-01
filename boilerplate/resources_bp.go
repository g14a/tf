package boilerplate

import "tf/boilerplate/aws"

func SelectResourceBP(provider string, resource string) {
	switch provider {
	case "aws":
		aws.ResourceBP(resource)
	}
}
