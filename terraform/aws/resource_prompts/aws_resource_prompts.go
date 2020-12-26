package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSInstanceBuilderPrompt() {
	prompts := map[string]types.TfPrompt{}

	blockPrompt := promptui.Prompt{
		Label: "Enter block name(required) e.g. web",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts["ami"] = types.TfPrompt{
		Label: "Enter ami(required)",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ami")

	prompts["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type(required) e.g. t2.micro",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_type")

	prompts["cpu_core_count"] = types.TfPrompt{
		Label: "Enter cpu_core_count(number)",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "cpu_core_count")

	prompts["cpu_threads_per_core"] = types.TfPrompt{
		Label: "Enter cpu_threads_per_core(number)",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "cpu_threads_per_core")

	prompts["ebs_optimized"] = types.TfPrompt{
		Label: "Select true/false for EBS-optimized(bool)",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ebs_optimized")

	prompts["monitoring"] = types.TfPrompt{
		Label: "Select true/false for detailed monitoring in EC2 instance",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "monitoring")

	prompts["subnet_id"] = types.TfPrompt{
		Label: "The VPC Subnet ID to launch in",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["private_ip"] = types.TfPrompt{
		Label: "Private IP address to associate with the instance in a VPC",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "private_ip")

	prompts["iam_instance_profile"] = types.TfPrompt{
		Label: "The IAM Instance Profile to launch the instance with",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "iam_instance_profile")

	prompts["security_groups"] = types.TfPrompt{
		Label: "A list of security group names (EC2-Classic) or IDs (default VPC) to associate with\ne.g.[a,b,c]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "security_groups")

	prompts["vpc_security_group_ids"] = types.TfPrompt{
		Label: "A list of security group IDs to associate with(Only VPC) e.g. [a,b,c]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_security_group_ids")

	selects := map[string]types.TfSelect{}
	var selectOrder []string

	selects["associate_public_ip_address"] = types.TfSelect{
		Label: "Associate a public ip address with an instance in a VPC.(bool)",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "vpc_security_group_ids")

	selects["placement_group"] = types.TfSelect{
		Label: "The Placement Group to start the instance in",
		Select: promptui.Select{
			Label: "",
			Items: []string{"cluster", "partition", "spread"},
		},
	}
	selectOrder = append(selectOrder, "placement_group")

	selects["hibernation"] = types.TfSelect{
		Label: "If true, the launched EC2 instance will support hibernation.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "hibernation")

	builder.ResourceBuilder("aws_instance", blockName, promptOrder, selectOrder, builder.PSOrder(promptOrder, selectOrder, prompts, selects))

	color.Green("\nDo you want to fill nested blocks like connection, timeouts: [y/n]\n\n", "text")
	ynPrompt := promptui.Prompt{}
	yn, err := ynPrompt.Run()

	if yn == "n" {
		return
	}

}

func AWSVPCPrompt() {
	prompts := map[string]types.TfPrompt{}

	blockPrompt := promptui.Prompt{
		Label: "Enter block name(required) e.g. web",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts["cidr_block"] = types.TfPrompt{
		Label: "Enter cidr_block(required)",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "cidr_block")

	prompts["owner_id"] = types.TfPrompt{
		Label: "The ID of the AWS account that owns the VPC.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "owner_id")

	var selectOrder []string
	selects := map[string]types.TfSelect{}

	selects["instance_tenancy"] = types.TfSelect{
		Label: "A tenancy option for instances launched into the VPC. Default is `default`",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"dedicated", "host"},
		},
	}
	selectOrder = append(selectOrder, "instance_tenancy")

	selects["enable_classiclink"] = types.TfSelect{
		Label: "A boolean flag to enable/disable ClassicLink for the VPC.",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_classiclink")

	selects["enable_dns_hostnames"] = types.TfSelect{
		Label: "A boolean flag to enable/disable DNS hostnames in the VPC.",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_dns_hostnames")

	selects["enable_dns_support"] = types.TfSelect{
		Label: "A boolean flag to enable/disable DNS hostnames in the VPC.",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_dns_support")

	selects["enable_classiclink_dns_support"] = types.TfSelect{
		Label: "A boolean flag to enable/disable ClassicLink DNS Support for the VPC. \nOnly valid in regions and accounts that support EC2 Classic.",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_classiclink_dns_support")

	selects["assign_generated_ipv6_cidr_block"] = types.TfSelect{
		Label: "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "assign_generated_ipv6_cidr_block")

	builder.ResourceBuilder("aws_vpc", blockName, promptOrder, selectOrder, builder.PSOrder(promptOrder, selectOrder, prompts, selects))

}

func AWSS3BucketPrompt() {
	blockPrompt := promptui.Prompt{
		Label: "Enter block name(required) e.g. web",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder []string

	prompts["bucket"] = types.TfPrompt{
		Label: "The name of the bucket. If omitted, Terraform will assign a random, unique name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	selects := map[string]types.TfSelect{}
	var selectOrder []string

	selects["acl"] = types.TfSelect{
		Label: "The canned ACL to apply",
		Select: promptui.Select{
			Label: "",
			Items: []string{"private", "public-read", "public-read-write", "aws-exec-read", "authenticated-read", "log-delivery-write"},
		},
	}
	selectOrder = append(selectOrder, "acl")

	selects["force_destroy"] = types.TfSelect{
		Label: "indicates all objects (including any locked objects) should be deleted from the bucket so that the bucket can be destroyed without error",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "force_destroy")

	selects["acceleration_status"] = types.TfSelect{
		Label: "Sets the accelerate configuration of an existing bucket.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"Enabled", "Suspended"},
		},
	}
	selectOrder = append(selectOrder, "acceleration_status")

	builder.ResourceBuilder("aws_s3_bucket", blockName, promptOrder, selectOrder, builder.PSOrder(promptOrder, selectOrder, prompts, selects))
}
