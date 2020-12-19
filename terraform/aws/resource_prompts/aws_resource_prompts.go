package resource_prompts

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/utils"
)

func AWSInstanceBuilderPrompt() {
	prompts := map[string]promptui.Prompt{}

	blockPrompt := promptui.Prompt{
		Label:  "Enter block name(required) e.g. web",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts["ami"] = promptui.Prompt{
		Label: "Enter ami(required)",
	}
	promptOrder = append(promptOrder, "ami")

	prompts["instance_type"] = promptui.Prompt{
		Label: "Enter instance_type(required) e.g. t2.micro",
	}
	promptOrder = append(promptOrder, "instance_type")

	prompts["cpu_core_count"] = promptui.Prompt{
		Label: "Enter cpu_core_count(number)",
		Validate: utils.IntValidator,
	}
	promptOrder = append(promptOrder, "cpu_core_count")

	prompts["cpu_threads_per_core"] = promptui.Prompt{
		Label: "Enter cpu_threads_per_core(number)",
		Validate: utils.IntValidator,
	}
	promptOrder = append(promptOrder, "cpu_threads_per_core")

	prompts["ebs_optimized"] = promptui.Prompt{
		Label: "Select true/false for EBS-optimized(bool)",
	}
	promptOrder = append(promptOrder, "ebs_optimized")

	prompts["monitoring"] = promptui.Prompt{
		Label: "Select true/false for detailed monitoring in EC2 instance",
	}
	promptOrder = append(promptOrder, "monitoring")

	prompts["subnet_id"] = promptui.Prompt{
		Label: "The VPC Subnet ID to launch in",
		Validate: utils.StringValidator,
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["private_ip"] = promptui.Prompt{
		Label: "Private IP address to associate with the instance in a VPC",
		Validate: utils.StringValidator,
	}
	promptOrder = append(promptOrder, "private_ip")

	prompts["iam_instance_profile"] = promptui.Prompt{
		Label: "The IAM Instance Profile to launch the instance with",
		Validate: utils.StringValidator,
	}
	promptOrder = append(promptOrder, "iam_instance_profile")

	prompts["security_groups"] = promptui.Prompt{
		Label: "A list of security group names (EC2-Classic) or IDs (default VPC) to associate with\ne.g.[a,b,c]",
	}
	promptOrder = append(promptOrder, "security_groups")

	prompts["vpc_security_group_ids"] = promptui.Prompt{
		Label: "A list of security group IDs to associate with(Only VPC) e.g. [a,b,c]",
	}
	promptOrder = append(promptOrder, "vpc_security_group_ids")

	selects := map[string]promptui.Select{}
	var selectOrder []string

	selects["associate_public_ip_address"] = promptui.Select{
		Label: "Associate a public ip address with an instance in a VPC.(bool)",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "vpc_security_group_ids")

	selects["placement_group"] = promptui.Select{
		Label: "The Placement Group to start the instance in",
		Items: []string{"cluster","partition","spread"},
	}

	selects["hibernation"] = promptui.Select{
		Label: "If true, the launched EC2 instance will support hibernation.",
		Items: []string{"true", "false"},
	}

	resourceBlock := map[string]interface{}{}

	for _, v := range promptOrder {
		p := prompts[v]
		value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		resourceBlock[v] = value
	}

	for _, v := range selectOrder {
		p := selects[v]
		_, value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		resourceBlock[v] = value
	}

	builder.ResourceBuilder("aws_instance", blockName, resourceBlock)
}

func AWSVPCPrompt()  {
	prompts := map[string]promptui.Prompt{}

	blockPrompt := promptui.Prompt{
		Label:  "Enter block name(required) e.g. web",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts["cidr_block"] = promptui.Prompt{
		Label: "Enter cidr_block(required)",
	}
	promptOrder = append(promptOrder, "cidr_block")

	prompts["owner_id"] = promptui.Prompt{
		Label: "The ID of the AWS account that owns the VPC.",
	}
	promptOrder = append(promptOrder, "owner_id")

	var selectOrder []string
	selects := map[string]promptui.Select{}

	selects["instance_tenancy"] = promptui.Select{
		Label: "A tenancy option for instances launched into the VPC. Default is `default`",
		Items: []string{"dedicated","host"},
	}
	selectOrder = append(selectOrder, "instance_tenancy")

	selects["enable_classiclink"] = promptui.Select{
		Label: "A boolean flag to enable/disable ClassicLink for the VPC.",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "enable_classiclink")

	selects["enable_dns_hostnames"] = promptui.Select{
		Label: "A boolean flag to enable/disable DNS hostnames in the VPC.",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "enable_dns_hostnames")

	selects["enable_dns_support"] = promptui.Select{
		Label: "A boolean flag to enable/disable DNS hostnames in the VPC.",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "enable_dns_support")

	selects["enable_classiclink_dns_support"] = promptui.Select{
		Label: "A boolean flag to enable/disable ClassicLink DNS Support for the VPC. \nOnly valid in regions and accounts that support EC2 Classic.",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "enable_classiclink_dns_support")

	selects["assign_generated_ipv6_cidr_block"] = promptui.Select{
		Label: "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC",
		Items: []string{"true", "false"},
	}
	selectOrder = append(selectOrder, "assign_generated_ipv6_cidr_block")

	resourceBlock := map[string]interface{}{}

	for _, v := range promptOrder {
		p := prompts[v]
		value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		resourceBlock[v] = value
	}

	for _, v := range selectOrder {
		p := selects[v]
		_, value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		resourceBlock[v] = value
	}

	builder.ResourceBuilder("aws_vpc", blockName, resourceBlock)

}

func AWSS3BucketPrompt() {
	prompts := map[string]promptui.Prompt{}
	var promptOrder []string

	prompts["bucket"] = promptui.Prompt{
		Label: "The name of the bucket. If omitted, Terraform will assign a random, unique name",
	}
	promptOrder = append(promptOrder, "bucket")

	selects := map[string]promptui.Select{}
	var selectOrder []string

	selects["acl"] = promptui.Select{
		Label: "The canned ACL to apply",
		Items: []string{"private","public-read","public-read-write","aws-exec-read","authenticated-read","log-delivery-write"},
	}
	selectOrder = append(selectOrder, "acl")

	selects["force_destroy"] = promptui.Select{
		Label: "indicates all objects (including any locked objects) should be deleted from the bucket so that the bucket can be destroyed without error",
		Items: []string{"true","false"},
	}
	selectOrder = append(selectOrder, "force_destroy")

	selects["acceleration_status"] = promptui.Select{
		Label: "Sets the accelerate configuration of an existing bucket.",
		Items: []string{"Enabled","Suspended"},
	}
	selectOrder = append(selectOrder, "force_destroy")

}