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

	prompts["ami"] = promptui.Prompt{
		Label: "Enter ami(required)",
	}

	prompts["instance_type"] = promptui.Prompt{
		Label: "Enter instance_type(required) e.g. t2.micro",
	}

	prompts["cpu_core_count"] = promptui.Prompt{
		Label: "Enter cpu_core_count(number)",
		Validate: utils.IntValidator,
	}

	prompts["cpu_threads_per_core"] = promptui.Prompt{
		Label: "Enter cpu_threads_per_core(number)",
		Validate: utils.IntValidator,
	}

	prompts["ebs_optimized"] = promptui.Prompt{
		Label: "Select true/false for EBS-optimized(bool)",
	}

	prompts["monitoring"] = promptui.Prompt{
		Label: "Select true/false for detailed monitoring in EC2 instance",
	}

	prompts["subnet_id"] = promptui.Prompt{
		Label: "The VPC Subnet ID to launch in",
		Validate: utils.StringValidator,
	}

	prompts["private_ip"] = promptui.Prompt{
		Label: "Private IP address to associate with the instance in a VPC",
		Validate: utils.StringValidator,
	}

	prompts["iam_instance_profile"] = promptui.Prompt{
		Label: "The IAM Instance Profile to launch the instance with",
		Validate: utils.StringValidator,
	}

	placementGroupSelect := promptui.Select{
		Label: "The Placement Group to start the instance in",
		Items: []string{"cluster","partition","spread"},
	}

	_, pg, err := placementGroupSelect.Run()

	resourceBlock := map[string]interface{}{}
	resourceBlock["placement_group"] = pg

	for k, v := range prompts {
		value, err := v.Run()
		if err != nil {
			fmt.Println(err)
		}
		resourceBlock[k] = value
	}

	builder.ResourceBuilder("aws_instance", blockName, resourceBlock)
}

func AWSVPCPrompt()  {
	
}