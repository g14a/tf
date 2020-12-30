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

	color.Green("\nEnter block name(required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts["ami"] = types.TfPrompt{
		Label: "Enter ami(required):\nThe AMI to use for the instance",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ami")

	prompts["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type(required) e.g. t2.micro\nThe type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.",
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
		Label: "Enter cpu_threads_per_core(number):\n(Optional - has no effect unless cpu_core_count is also set) " +
			"If set to to 1, hyperthreading is disabled on the launched instance. " +
			"Defaults to 2 if not set. See Optimizing CPU Options for more information.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "cpu_threads_per_core")

	prompts["ebs_optimized"] = types.TfPrompt{
		Label: "Select true/false for EBS-optimized(bool):\n(Optional) If true, the launched EC2 instance will be EBS-optimized. " +
			"Note that if this is not set on an instance type that is optimized by default then this will show " +
			"as disabled but if the instance type is optimized by default then there is no " +
			"need to set this and there is no effect to disabling it. See the " +
			"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html of AWS User Guide for more information.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ebs_optimized")

	prompts["monitoring"] = types.TfPrompt{
		Label: "Select true/false for monitoring:\n(Optional) " +
			"If true, the launched EC2 instance will have detailed monitoring enabled",
		Prompt: promptui.Prompt{
			Label: "",
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

	selects := map[string]types.TfSelect{}
	var selectOrder []string

	selects["associate_public_ip_address"] = types.TfSelect{
		Label: "Enter associate_public_ip_address.(Optional)Associate a public ip address with an instance in a VPC.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "associate_public_ip_address")

	selects["placement_group"] = types.TfSelect{
		Label: "Enter placement_group:\nThe Placement Group to start the instance in",
		Select: promptui.Select{
			Label: "",
			Items: []string{"cluster", "partition", "spread"},
		},
	}
	selectOrder = append(selectOrder, "placement_group")

	selects["hibernation"] = types.TfSelect{
		Label: "Enter hibernation.\n(Optional)If true, the launched EC2 instance will support hibernation.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "hibernation")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Yellow("\nConfigure nested settings like tags [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_instance", blockName, promptOrder, selectOrder, resourceBlock)
		return
	}

	tagPrompt := map[string]types.TfPrompt{}
	var nestedOrder []string

	color.Green("\nEnter tags (Optional) A map of tags to assign to the resource:\n\n")

	tagPrompt["Name"] = types.TfPrompt{
		Label: "Enter Name: ",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "Name")
	selectOrder = append(selectOrder, "tags")

	resourceBlock["tags"] = builder.NestedPSOrder(nestedOrder, tagPrompt, nil)
	builder.ResourceBuilder("aws_instance", blockName, promptOrder, selectOrder, resourceBlock)

}

func AWSVPCPrompt() {

	color.Green("\nEnter block name(required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder []string

	prompts["cidr_block"] = types.TfPrompt{
		Label: "Enter cidr_block:\n(Required) The CIDR block for the VPC",
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
		Label: "Enter instance_tenancy:\bTenancy of instances spin up within VPC. Default is `default`",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"dedicated", "host"},
		},
	}
	selectOrder = append(selectOrder, "instance_tenancy")

	selects["enable_classiclink"] = types.TfSelect{
		Label: "Enter enable_classiclink:\nWhether or not the VPC has Classiclink enabled",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_classiclink")

	selects["enable_dns_hostnames"] = types.TfSelect{
		Label: "Enter enable_dns_hostnames:\nWhether or not the VPC has DNS hostname support",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_dns_hostnames")

	selects["enable_dns_support"] = types.TfSelect{
		Label: "Enter enable_dns_hostnames:\nWhether or not the VPC has DNS support",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_dns_support")

	selects["enable_classiclink_dns_support"] = types.TfSelect{
		Label: "Enter enable_classiclink_dns_support:\n(Optional) A boolean flag to enable/disable ClassicLink DNS Support for the VPC." +
			" Only valid in regions and accounts that support EC2 Classic.",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "enable_classiclink_dns_support")

	selects["assign_generated_ipv6_cidr_block"] = types.TfSelect{
		Label: "Enter assign_generated_ipv6_cidr_block:\nEnter (Optional) Requests an Amazon-provided IPv6 CIDR block with a /56 prefix " +
			"length for the VPC. You cannot specify the range of IP addresses, " +
			"or the size of the CIDR block. Default is false",
		Select: promptui.Select{
			Label: ",",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "assign_generated_ipv6_cidr_block")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Yellow("\nConfigure nested settings like tags [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_vpc", blockName, promptOrder, selectOrder, resourceBlock)
		return
	}

	tagPrompt := map[string]types.TfPrompt{}
	var nestedOrder []string

	color.Green("\nEnter tags (Optional) A map of tags to assign to the resource:\n\n")

	tagPrompt["Name"] = types.TfPrompt{
		Label: "Enter Name: ",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "Name")
	selectOrder = append(selectOrder, "tags")

	resourceBlock["tags"] = builder.NestedPSOrder(nestedOrder, tagPrompt, nil)

	builder.ResourceBuilder("aws_vpc", blockName, promptOrder, selectOrder, resourceBlock)

}

func AWSS3BucketPrompt() {

	color.Green("\nEnter block name(required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
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
		Label: "Enter acl:\nThe canned ACL to apply",
		Select: promptui.Select{
			Label: "",
			Items: []string{"private", "public-read", "public-read-write", "aws-exec-read", "authenticated-read", "log-delivery-write"},
		},
	}
	selectOrder = append(selectOrder, "acl")

	selects["force_destroy"] = types.TfSelect{
		Label: "Enter force_destroy:\n(Optional, Default:false) A boolean that indicates all objects \n" +
			"(including any locked objects) should be deleted from the bucket \n" +
			"so that the bucket can be destroyed without error. These objects are not recoverable.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "force_destroy")

	selects["acceleration_status"] = types.TfSelect{
		Label: "Enter acceleration_status:\n(Optional) Sets the accelerate " +
			"configuration of an existing bucket. Can be Enabled or Suspended",
		Select: promptui.Select{
			Label: "",
			Items: []string{"Enabled", "Suspended"},
		},
	}
	selectOrder = append(selectOrder, "acceleration_status")
	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Yellow("\nConfigure nested settings like cors_rule/versioning/website etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket", blockName, promptOrder, selectOrder, resourceBlock)
		return
	}

	corsRulePrompt := map[string]types.TfPrompt{}
	var nestedOrder []string

	color.Green("\nEnter cors_rule (Optional) A rule of Cross-Origin Resource Sharing :\n\n")

	corsRulePrompt["allowed_headers"] = types.TfPrompt{
		Label: "Enter allowed_headers:\n(Optional) Specifies which headers are allowed",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "allowed_headers")

	corsRulePrompt["allowed_methods"] = types.TfPrompt{
		Label: "Enter allowed_methods:\nRequired) Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "allowed_methods")

	corsRulePrompt["allowed_origins"] = types.TfPrompt{
		Label: "Enter allowed_origins:\n(Required) Specifies which origins are allowed.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "allowed_origins")

	corsRulePrompt["exposed_headers"] = types.TfPrompt{
		Label: "Enter exposed_headers:\n(Optional) Specifies expose header in the response.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "exposed_headers")

	corsRulePrompt["max_age_seconds"] = types.TfPrompt{
		Label: "Enter max_age_seconds:\n(Optional) Specifies time in seconds " +
			"that browser can cache the response for a preflight request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "max_age_seconds")
	selectOrder = append(selectOrder, "cors_rule")

	resourceBlock["cors_rule"] = builder.NestedPSOrder(nestedOrder, corsRulePrompt, nil)

	color.Green("\nEnter website:\nThe website object supports the following:" +
		"\n1.index_document\n2.error_document\n3.redirect_all_requests_to\n4.routing_rules\n\n")

	websitePrompt := map[string]types.TfPrompt{}

	websitePrompt["index_document"] = types.TfPrompt{
		Label: "Enter index_document:\n(Required), unless using redirect_all_requests_to) Amazon S3\n" +
			" returns this index document when requests are made to the" +
			" root domain or any of the subfolders.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "index_document")

	websitePrompt["error_document"] = types.TfPrompt{
		Label: "Enter error_document:\n(Optional) An absolute path to the document to return in case of a 4XX error.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "error_document")

	websitePrompt["redirect_all_requests_to"] = types.TfPrompt{
		Label: "Enter redirect_all_requests_to:\nOptional) A hostname to redirect all website requests for \n" +
			"this bucket to. Hostname can optionally be prefixed with a \n" +
			"protocol (http:// or https://) to use when redirecting requests. \n" +
			"The default is the protocol that is used in the original request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "redirect_all_requests_to")

	websitePrompt["routing_rules"] = types.TfPrompt{
		Label: "Enter routing_rules:\n(Optional) A json array containing routing rules describing redirect behavior and when redirects are applied.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "routing_rules")
	selectOrder = append(selectOrder, "website")
	
	resourceBlock["website"] = builder.NestedPSOrder(nestedOrder[len(nestedOrder)-4:], websitePrompt, nil)

	builder.ResourceBuilder("aws_s3_bucket", blockName, promptOrder, selectOrder, resourceBlock)
}
