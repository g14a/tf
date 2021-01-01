package resource_bps

import "github.com/fatih/color"

func AWSVPCBP() {
	color.Green("\nresource \"aws_vpc\" \"foo\" {\n  cidr_block       = \"10.0.0.0/16\"\n  instance_tenancy = \"default\"\n\n  tags = {\n    Name = \"main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc\n\n")
}
