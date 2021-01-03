package resource_prompts

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
	
	builder.ResourceBuilder("aws_customer_gateway", blockName, promptOrder, nil, resourceBlock)
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
			Label: "",
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
		builder.ResourceBuilder("aws_default_route_table", blockName, promptOrder, nil, resourceBlock)
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

	resourceBlock["route"] = builder.NestedPSOrder(nestedPromptOrder, nil, routePrompt, nil)

	builder.ResourceBuilder("aws_default_route_table", blockName, promptOrder, selectOrder, resourceBlock)
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

	builder.ResourceBuilder("aws_vpc", blockName, promptOrder, selectOrder, resourceBlock)
}

