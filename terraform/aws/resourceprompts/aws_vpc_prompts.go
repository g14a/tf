package resourceprompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSCustomerGatewayPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder []string

	prompts["bgp_asn"] = types.TfPrompt{
		Label: "Enter bgp_asn:\n(Required) The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bgp_asn")

	prompts["ip_address"] = types.TfPrompt{
		Label: "Enter ip_address:\n(Required) The IP address of the gateway's Internet-routable external interface.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ip_address")

	prompts["type"] = types.TfPrompt{
		Label: "Enter type:\n(Required) The type of customer gateway. The only type AWS supports at this time is \"ipsec.1\"",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "type")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) Tags to apply to the gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_customer_gateway", blockName, resourceBlock)
}

func AWSDefaultRouteTablePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["default_route_table_id"] = types.TfPrompt{
		Label: "Enter default_route_table_id:\n(Required) The ID of the Default Routing Table.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "default_route_table_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g.k1=v1,k2=v2\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["propagating_vgws"] = types.TfPrompt{
		Label: "Enter propagating_vgws:\n(Optional) A list of virtual gateways for propagation.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "propagating_vgws")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like route [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_default_route_table", blockName, resourceBlock)
		return
	}

	routePrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	routePrompt["cidr_block"] = types.TfPrompt{
		Label: "Enter cidr_block:\n(Required) The CIDR block of the route.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cidr_block")

	routePrompt["ipv6_cidr_block"] = types.TfPrompt{
		Label: "Enter ipv6_cidr_block:\n(Optional) The Ipv6 CIDR block of the route.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ipv6_cidr_block")

	routePrompt["egress_only_gateway_id"] = types.TfPrompt{
		Label: "Enter egress_only_gateway_id:\n(Optional) Identifier of a VPC Egress Only Internet Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "egress_only_gateway_id")

	routePrompt["gateway_id"] = types.TfPrompt{
		Label: "Enter gateway_id:\n(Optional) Identifier of a VPC internet gateway or a virtual private gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "gateway_id")

	routePrompt["instance_id"] = types.TfPrompt{
		Label: "Enter instance_id:\n(Optional) Identifier of an EC2 instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "instance_id")

	routePrompt["nat_gateway_id"] = types.TfPrompt{
		Label: "Enter nat_gateway_id:\n(Optional) Identifier of a VPC NAT gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "nat_gateway_id")

	routePrompt["network_interface_id"] = types.TfPrompt{
		Label: "Enter network_interface_id:\n(Optional) Identifier of an EC2 network interface.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "network_interface_id")

	routePrompt["transit_gateway_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_id:\n(Optional) Identifier of an EC2 Transit Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "transit_gateway_id")

	routePrompt["vpc_endpoint_id"] = types.TfPrompt{
		Label: "Enter vpc_endpoint_id:\n(Optional) Identifier of a VPC Endpoint. This route must be removed prior to VPC Endpoint deletion.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "vpc_endpoint_id")

	routePrompt["vpc_peering_connection_id"] = types.TfPrompt{
		Label: "Enter vpc_peering_connection_id:\n(Optional) Identifier of a VPC Endpoint. This route must be removed prior to VPC Endpoint deletion.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "vpc_peering_connection_id")
	selectOrder = append(selectOrder, "route")

	resourceBlock["route"] = builder.PSOrder(nestedPromptOrder, nil, routePrompt, nil)

	builder.ResourceBuilder("aws_default_route_table", blockName, resourceBlock)
}

func AWSVPCPrompt() {

	color.Green("\nEnter block name(Required) e.g. web\n\n")
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

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n For e.g. k1=v1,k2=v2",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["enable_classiclink"] = types.TfPrompt{
		Label: "Enter enable_classiclink(true/false):\nWhether or not the VPC has Classiclink enabled",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enable_classiclink")

	prompts["enable_dns_hostnames"] = types.TfPrompt{
		Label: "Enter enable_dns_hostnames(true/false):\nWhether or not the VPC has DNS hostname support",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enable_dns_hostnames")

	prompts["enable_dns_support"] = types.TfPrompt{
		Label: "Enter enable_dns_hostnames(true/false):\nWhether or not the VPC has DNS support",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enable_dns_support")

	prompts["enable_classiclink_dns_support"] = types.TfPrompt{
		Label: "Enter enable_classiclink_dns_support(true/false):\n(Optional) A boolean flag to enable/disable ClassicLink DNS Support for the VPC." +
			" Only valid in regions and accounts that support EC2 Classic.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enable_classiclink_dns_support")

	prompts["assign_generated_ipv6_cidr_block"] = types.TfPrompt{
		Label: "Enter assign_generated_ipv6_cidr_block(true/false):\nEnter (Optional) Requests an Amazon-provided IPv6 CIDR block with a /56 prefix " +
			"length for the VPC. You cannot specify the range of IP addresses, " +
			"or the size of the CIDR block. Default is false",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "assign_generated_ipv6_cidr_block")

	var selectOrder []string
	selects := map[string]types.TfSelect{}

	selects["instance_tenancy"] = types.TfSelect{
		Label: "Enter instance_tenancy:\bTenancy of instances spin up within VPC. Default is `default`",
		Select: promptui.Select{
			Label: "",
			Items: []string{"dedicated", "host"},
		},
	}
	selectOrder = append(selectOrder, "instance_tenancy")

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
		builder.ResourceBuilder("aws_vpc", blockName, resourceBlock)
		return
	}

	lifecyclePrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	color.Green("Enter lifecycle block:\n")

	lifecyclePrompt["create_before_destroy"] = types.TfPrompt{
		Label: "Enter create_before_destroy:(true/false)\nBy default, when Terraform must change a resource argument \n" +
			"that cannot be updated in-place due to remote API limitations, \n" +
			"Terraform will instead destroy the existing object and then \n" +
			"create a new replacement object with the new configured arguments.\n" +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#create_before_destroy",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "create_before_destroy")

	lifecyclePrompt["prevent_destroy"] = types.TfPrompt{
		Label: "Enter prevent_destroy:(true/false)\nThis meta-argument, when set to true, will cause Terraform to \n" +
			"reject with an error any plan that would destroy the infrastructure \n" +
			"object associated with the resource, as long as the argument \n" +
			"remains present in the configuration.\n" +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#prevent_destroy",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "prevent_destroy")

	lifecyclePrompt["ignore_changes"] = types.TfPrompt{
		Label: "Enter ignore_changes: e.g.[\"c1\",\"c2\"]\nBy default, Terraform detects any difference in the " +
			"current settings of a real infrastructure object and plans to " +
			"update the remote object to match configuration." +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#ignore_changes",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ignore_changes")
	selectOrder = append(selectOrder, "lifecycle")

	resourceBlock["lifecycle"] = builder.PSOrder(nestedPromptOrder, nil, lifecyclePrompt, nil)

	builder.ResourceBuilder("aws_vpc", blockName, resourceBlock)
}
