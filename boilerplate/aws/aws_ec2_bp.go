package aws

import "fmt"

func AWSEC2BP() {
	fmt.Println("resource \"aws_instance\" \"foo\" {\n  ami = \"\"\n  instance_type = \"\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
}
