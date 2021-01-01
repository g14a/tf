package resource_bps

import (
	"github.com/fatih/color"
)

func AWSEC2BP() {
	color.Green("\nresource \"aws_instance\" \"foo\" {\n  ami           = data.aws_ami.ubuntu.id\n  instance_type = \"t3.micro\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance\n\n")
}
