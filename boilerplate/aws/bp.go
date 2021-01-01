package aws

import (
	"github.com/fatih/color"
)

func ResourceBP(resource string) {

	switch resource {
	case "aws_instance":
		AWSEC2BP()
	default:
		color.Red("No such resource present in AWS")
	}
}
