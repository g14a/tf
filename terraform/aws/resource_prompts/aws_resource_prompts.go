package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSInstancePrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts["ami"] = types.TfPrompt{
		Label: "Enter ami(Required):\nThe AMI to use for the instance",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ami")

	prompts["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type(Required) e.g. t2.micro\nThe type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_type")

	prompts["cpu_core_count"] = types.TfPrompt{
		Label: "Enter cpu_core_count(number):\n(Optional)Sets the number of CPU cores for an instance. " +
			"This option is only supported on creation of instance type that support CPU Options - " +
			"specifying this option for unsupported instance types will return an error from the EC2 API. Checkout " +
			"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values for more info.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "cpu_core_count")

	prompts["cpu_threads_per_core"] = types.TfPrompt{
		Label: "Enter cpu_threads_per_core(number):\n(Optional) - has no effect unless cpu_core_count is also set) " +
			"If set to to 1, hyperthreading is disabled on the launched instance. " +
			"Defaults to 2 if not set. See Optimizing CPU Options for more information.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "cpu_threads_per_core")

	prompts["ebs_optimized"] = types.TfPrompt{
		Label: "Enter EBS-optimized(true/false):\n(Optional) If true, the launched EC2 instance will be EBS-optimized. " +
			"\nNote that if this is not set on an instance type that is optimized by default then this will show " +
			"\nas disabled but if the instance type is optimized by default then there is no " +
			"\nneed to set this and there is no effect to disabling it. " +
			"\nCheckout https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html of AWS User Guide for more information.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ebs_optimized")

	prompts["monitoring"] = types.TfPrompt{
		Label: "Select true/false for monitoring:\n(Optional) " +
			"If true, the launched EC2 instance will have detailed monitoring enabled",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "monitoring")

	prompts["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Optional) The VPC Subnet ID to launch in.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["private_ip"] = types.TfPrompt{
		Label: "Enter private_ip:\n(Optional) Private IP address to associate with the instance in a VPC.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "private_ip")

	prompts["iam_instance_profile"] = types.TfPrompt{
		Label: "Enter iam_instance_profile:\n(Optional) The IAM Instance Profile to launch the " +
			"instance with. Specified as the name of the Instance Profile. " +
			"Ensure your credentials have the correct permission to assign " +
			"the instance profile according to the EC2 documentation, notably iam:PassRole",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "iam_instance_profile")

	prompts["security_groups"] = types.TfPrompt{
		Label: "A list of security group names (EC2-Classic) or IDs (default VPC) to associate with\ne.g.[\"a\",\"b\",\"c\"]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "security_groups")

	prompts["vpc_security_group_ids"] = types.TfPrompt{
		Label: "A list of security group IDs to associate with(Only VPC) e.g. [\"a\",\"b\",\"c\"]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_security_group_ids")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["associate_public_ip_address"] = types.TfPrompt{
		Label: "Enter associate_public_ip_address(true/false):\n(Optional)Associate a public ip address with an instance in a VPC.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "associate_public_ip_address")

	prompts["hibernation"] = types.TfPrompt{
		Label: "Enter hibernation(true/false).\n(Optional)If true, the launched EC2 instance will support hibernation.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "hibernation")

	selects := map[string]types.TfSelect{}
	var selectOrder []string

	selects["placement_group"] = types.TfSelect{
		Label: "Enter placement_group:\nThe Placement Group to start the instance in",
		Select: promptui.Select{
			Label: "",
			Items: []string{"cluster", "partition", "spread"},
		},
	}
	selectOrder = append(selectOrder, "placement_group")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_instance", blockName, resourceBlock)
}

func AWSELBPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder []string
	var selectOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Optional) The name of the ELB. By default generated by Terraform.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["instances"] = types.TfPrompt{
		Label: "Enter instances:\n(Optional) A list of instance ids to place in the ELB pool. e.g. [\"id1\",\"id2\"]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instances")

	prompts["cross_zone_load_balancing"] = types.TfPrompt{
		Label: "Enter cross_zone_load_balancing(true/false):\n(Optional) Enable cross-zone load balancing. Default: true",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "cross_zone_load_balancing")

	prompts["idle_timeout"] = types.TfPrompt{
		Label: "Enter idle_timeout:\n(Optional) The time in seconds that the connection is allowed to be idle. Default: 60",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "idle_timeout")

	prompts["connection_draining"] = types.TfPrompt{
		Label: "Enter connection_draining(true/false):\n(Optional) Boolean to enable connection draining. Default: false",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "connection_draining")

	prompts["connection_draining_timeout"] = types.TfPrompt{
		Label: "Enter connection_draining_timeout(true/false):\n(Optional) The time in seconds to allow for connections to drain. Default: 300",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "connection_draining_timeout")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like listener/access_logs/website etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket", blockName, resourceBlock)
		return
	}

	listenerPrompt := map[string]types.TfPrompt{}
	listenerSelect := map[string]types.TfSelect{}

	var nestedPromptOrder, nestedSelectOrder []string

	listenerPrompt["instance_port"] = types.TfPrompt{
		Label: "Enter instance_port:\n(Required) The port on the instance to route to",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "instance_port")

	listenerPrompt["lb_port"] = types.TfPrompt{
		Label: "Enter lb_port:\n(Required) The port to listen on for the load balancer",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "lb_port")

	listenerPrompt["ssl_certificate_id"] = types.TfPrompt{
		Label: "Enter ssl_certificate_id:\n(Optional) The ARN of an SSL certificate you have uploaded to AWS IAM. Note ECDSA-specific restrictions below. Only valid when lb_protocol is either HTTPS or SSL",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ssl_certificate_id")

	listenerSelect["instance_protocol"] = types.TfSelect{
		Label: "Enter instance_protocol:\n(Required) The protocol to use to the instance.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"HTTP", "HTTPS", "TCP", "SSL"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "instance_protocol")

	listenerSelect["lb_protocol"] = types.TfSelect{
		Label: "Enter lb_protocol:\n(Required) The protocol to listen on.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"HTTP", "HTTPS", "TCP", "SSL"},
		},
	}

	nestedSelectOrder = append(nestedSelectOrder, "lb_protocol")
	selectOrder = append(selectOrder, "listener")

	resourceBlock["listener"] = builder.PSOrder(nestedPromptOrder, nestedSelectOrder, listenerPrompt, listenerSelect)

	accessLogsPrompt := map[string]types.TfPrompt{}
	accessLogsPrompt["bucket"] = types.TfPrompt{
		Label: "Enter bucket:\n(Required) The S3 bucket name to store the logs in.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "bucket")

	accessLogsPrompt["bucket_prefix"] = types.TfPrompt{
		Label: "Enter bucket_prefix:\n(Optional) The S3 bucket prefix. Logs are stored in the root if not configured.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "bucket_prefix")

	accessLogsPrompt["interval"] = types.TfPrompt{
		Label: "Enter interval:\n(Optional) The publishing interval in minutes. Default: 60 minutes.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "interval")

	accessLogsPrompt["enabled"] = types.TfPrompt{
		Label: "Enter enabled:(true/false)\n(Optional) Boolean to enable / disable access_logs. Default is true",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "enabled")
	selectOrder = append(selectOrder, "access_logs")

	resourceBlock["access_logs"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-4:], nil, accessLogsPrompt, nil)

	healthCheckPrompt := map[string]types.TfPrompt{}
	healthCheckPrompt["healthy_threshold"] = types.TfPrompt{
		Label: "Enter healthy_threshold:\n(Required) The number of checks before the instance is declared healthy.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "healthy_threshold")

	healthCheckPrompt["unhealthy_threshold"] = types.TfPrompt{
		Label: "Enter unhealthy_threshold:\n(Required) The number of checks before the instance is declared unhealthy.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "unhealthy_threshold")

	healthCheckPrompt["target"] = types.TfPrompt{
		Label: "Enter target:\n(Required) The target of the check. Valid pattern is \"${PROTOCOL}:${PORT}${PATH}\" e.g. HTTP:8000/",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "target")

	healthCheckPrompt["timeout"] = types.TfPrompt{
		Label: "Enter timeout:\n(Required) The length of time before the check times out.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "timeout")

	healthCheckPrompt["interval"] = types.TfPrompt{
		Label: "Enter interval:\n(Required) The interval between checks.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "interval")
	selectOrder = append(selectOrder, "health_check")

	resourceBlock["health_check"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-5:], nil, healthCheckPrompt, nil)

	builder.ResourceBuilder("aws_elb", blockName, resourceBlock)
}
