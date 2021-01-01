package resource_bps

import (
	"github.com/fatih/color"
)

func AWSDBInstanceBP()  {
	color.Green("\nresource \"aws_db_instance\" \"foo\" {\n  allocated_storage    = 20\n  storage_type         = \"gp2\"\n  engine               = \"mysql\"\n  engine_version       = \"5.7\"\n  instance_class       = \"db.t2.micro\"\n  name                 = \"mydb\"\n  username             = \"foo\"\n  password             = \"foobarbaz\"\n  parameter_group_name = \"default.mysql5.7\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance\n\n")
}