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

func AWSDefaultNetworkACLPrompt() {
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

	prompts["default_network_acl_id"] = types.TfPrompt{
		Label: "Enter default_network_acl_id:\n(Required) The Network ACL ID to manage. This attribute is exported from aws_vpc, or manually found via the AWS Console.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "default_network_acl_id")

	prompts["subnet_ids"] = types.TfPrompt{
		Label: "Enter subnet_ids e.g.[\"id1\",\"id2\"]:\n(Optional) A list of Subnet IDs to apply the ACL to. See the notes below on managing Subnets in the Default Network ACL",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_ids")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like ingress/egress [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_default_network_acl", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter ingress:\n(Optional) Specifies an ingress rule." +
		"\n1.from_port\n2.to_port\n3.rule_no\n4.action\n5.protocol\n6.cidr_block\n7.ipv6_cidr_block\n8.icmp_type\n9.icmp_code\n")

	ingressEgressPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	ingressEgressPrompt["from_port"] = types.TfPrompt{
		Label: "Enter from_port:\n(Required) The from port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "from_port")

	ingressEgressPrompt["to_port"] = types.TfPrompt{
		Label: "Enter to_port:\n(Required) The to port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "to_port")

	ingressEgressPrompt["rule_no"] = types.TfPrompt{
		Label: "Enter rule_no:\n(Required) The to port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "rule_no")

	ingressEgressPrompt["action"] = types.TfPrompt{
		Label: "Enter action:\n(Required) The action to take.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "action")

	ingressEgressPrompt["protocol"] = types.TfPrompt{
		Label: "Enter protocol:\n(Required) The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "protocol")

	ingressEgressPrompt["cidr_block"] = types.TfPrompt{
		Label: "Enter cidr_block:\n(Optional) The CIDR block to match. This must be a valid network mask.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cidr_block")

	ingressEgressPrompt["ipv6_cidr_block"] = types.TfPrompt{
		Label: "Enter ipv6_cidr_block:\n(Optional) The IPv6 CIDR block.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ipv6_cidr_block")

	ingressEgressPrompt["icmp_type"] = types.TfPrompt{
		Label: "Enter icmp_type:\n(Optional) The ICMP type to be used. Default 0.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "icmp_type")

	ingressEgressPrompt["icmp_code"] = types.TfPrompt{
		Label: "Enter icmp_code:\n(Optional) The ICMP type code to be used. Default 0.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "icmp_code")

	resourceBlock["ingress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	color.Green("\nEnter egress:\n(Optional) Specifies an egress rule." +
		"\n1.from_port\n2.to_port\n3.rule_no\n4.action\n5.protocol\n6.cidr_block\n7.ipv6_cidr_block\n8.icmp_type\n9.icmp_code\n")

	resourceBlock["egress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	builder.ResourceBuilder("aws_default_network_acl", blockName, resourceBlock)
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

func AWSDefaultSecurityGroupPrompt() {
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

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Optional, Forces new resource) The VPC ID. Note that changing the vpc_id will not restore any default security group rules that were modified, added, or removed. It will be left in its current state",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like ingress/egress [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_default_security_group", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter ingress:\n(Optional) Specifies an ingress rule." +
		"\n1.cidr_blocks\n2.description\n3.from_port\n4.ipv6_cidr_blocks\n5.prefix_list_ids\n6.protocol\n7.security_groups\n8.self\n9.to_port\n")

	ingressEgressPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	ingressEgressPrompt["cidr_blocks"] = types.TfPrompt{
		Label: "Enter cidr_blocks:\n(Optional) List of CIDR blocks.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cidr_blocks")

	ingressEgressPrompt["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Description of this ingress/egress rule.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "description")

	ingressEgressPrompt["from_port"] = types.TfPrompt{
		Label: "Enter from_port:\n(Required) The start port (or ICMP type number if protocol is \"icmp\" or \"icmpv6\")",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "from_port")

	ingressEgressPrompt["ipv6_cidr_blocks"] = types.TfPrompt{
		Label: "Enter ipv6_cidr_blocks:\n(Optional) List of IPv6 CIDR blocks.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ipv6_cidr_blocks")

	ingressEgressPrompt["prefix_list_ids"] = types.TfPrompt{
		Label: "Enter prefix_list_ids:\n(Optional) List of prefix list IDs.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "prefix_list_ids")

	ingressEgressPrompt["protocol"] = types.TfPrompt{
		Label: "Enter protocol:\n(Required) The protocol. If you select a protocol of \"-1\" (semantically equivalent to \"all\", which is not a valid value here), you must specify a \"from_port\" and \"to_port\" equal to 0. If not icmp, icmpv6, tcp, udp, or \"-1\" use the protocol number",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "protocol")

	ingressEgressPrompt["security_groups"] = types.TfPrompt{
		Label: "Enter security_groups:\n(Optional) List of security group Group Names if using EC2-Classic, or Group IDs if using a VPC.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "security_groups")

	ingressEgressPrompt["self"] = types.TfPrompt{
		Label: "Enter self:\n(Optional) If true, the security group itself will be added as a source to this ingress/egress rule.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "self")

	ingressEgressPrompt["to_port"] = types.TfPrompt{
		Label: "Enter to_port:\n(Required) The end range port (or ICMP code if protocol is \"icmp\").",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "to_port")

	resourceBlock["ingress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	color.Green("\nEnter egress:\n(Optional) Specifies an egress rule." +
		"\n1.cidr_blocks\n2.description\n3.from_port\n4.ipv6_cidr_blocks\n5.prefix_list_ids\n6.protocol\n7.security_groups\n8.self\n9.to_port\n")

	resourceBlock["egress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	builder.ResourceBuilder("aws_default_security_group", blockName, resourceBlock)
}

func AWSDefaultSubnetPrompt() {
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

	prompts["availability_zone"] = types.TfPrompt{
		Label: "Enter availability_zone:\n(Optional) The AZ for the subnet.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "availability_zone")

	prompts["map_public_ip_on_launch"] = types.TfPrompt{
		Label: "Enter map_public_ip_on_launch:\n(Optional) Specify true to indicate that instances launched into the subnet " +
			"\nshould be assigned a public IP address.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "map_public_ip_on_launch")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_default_subnet", blockName, resourceBlock)
}

func AWSDefaultVPCPrompt() {
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

	prompts["enable_dns_support"] = types.TfPrompt{
		Label: "Enter enable_dns_support:\n(Optional) A boolean flag to enable/disable DNS support in the VPC. Defaults to true.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enable_dns_support")

	prompts["enable_dns_hostnames"] = types.TfPrompt{
		Label: "Enter enable_dns_hostnames:\n(Optional) A boolean flag to enable/disable DNS hostnames in the VPC. Defaults to false.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enable_dns_hostnames")

	prompts["enable_classiclink"] = types.TfPrompt{
		Label: "Enter enable_classiclink:\n(Optional) A boolean flag to enable/disable ClassicLink for the VPC. Only valid in " +
			"\nregions and accounts that support EC2 Classic. Defaults false." +
			"\nCheckout https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enable_classiclink")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_default_vpc", blockName, resourceBlock)
}

func AWSDefaultVPCDHCPOptionsPrompt() {
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

	prompts["netbios_name_servers"] = types.TfPrompt{
		Label: "Enter netbios_name_servers:\n(Optional) List of NETBIOS name servers.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "netbios_name_servers")

	prompts["netbios_node_type"] = types.TfPrompt{
		Label: "Enter netbios_node_type:\n(Optional) List of NETBIOS name servers.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "netbios_node_type")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) List of NETBIOS name servers.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_default_vpc_dhcp_options", blockName, resourceBlock)
}

func AWSEC2ManagedPrefixListPrompt() {
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

	color.Yellow("\nWhen you reference a Prefix List in a resource, the maximum number of entries " +
		"\nfor the prefix lists counts as the same number of rules or entries for the resource. " +
		"\nFor example, if you create a prefix list with a maximum of 20 entries and you reference " +
		"\nthat prefix list in a security group rule, this counts as 20 rules for the security group.")

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) The name of this resource. The name must not start with com.amazonaws",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["address_family"] = types.TfPrompt{
		Label: "Enter address_family:\n(Required, Forces new resource) The address family (IPv4 or IPv6) of this prefix list.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "address_family")

	prompts["max_retries"] = types.TfPrompt{
		Label: "Enter max_retries:\n(Required, Forces new resource) The maximum number of entries that this prefix list can contain.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "max_retries")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to this resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like entry [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_ec2_managed_prefix_list", blockName, resourceBlock)
		return
	}

	entryPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	entryPrompt["cidr"] = types.TfPrompt{
		Label: "Enter cidr:\n(Required) The CIDR block of this entry.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cidr")

	entryPrompt["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Description of this entry.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "description")

	resourceBlock["entry"] = builder.PSOrder(nestedPromptOrder, nil, entryPrompt, nil)

	builder.ResourceBuilder("aws_ec2_managed_prefix_list", blockName, resourceBlock)
}

func AWSEgressOnlyInternetGatewayPrompt() {
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

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) The VPC ID to create in.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_egress_only_internet_gateway", blockName, resourceBlock)
}

func AWSFlowLogPrompt() {
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

	color.Yellow("\nOne of eni_id, subnet_id, or vpc_id must be specified.")

	prompts["eni_id"] = types.TfPrompt{
		Label: "Enter eni_id:\n(Optional) Elastic Network Interface ID to attach to",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "eni_id")

	prompts["iam_role_arn"] = types.TfPrompt{
		Label: "Enter iam_role_arn:\n(Optional) The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "iam_role_arn")

	prompts["log_destination"] = types.TfPrompt{
		Label: "Enter log_destination:\n(Optional) The ARN of the logging destination.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "log_destination")

	prompts["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Optional) Subnet ID to attach to",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Optional) VPC ID to attach to",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["log_format"] = types.TfPrompt{
		Label: "Enter log_format:\n(Optional) The fields to include in the flow log record, in the order in which they should appear.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "log_format")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) Key-value map of resource tags",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}

	selects["traffic_type"] = types.TfSelect{
		Label: "Enter traffic_type:\n(Required) The type of traffic to capture.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"ACCEPT", "REJECT"},
		},
	}
	selectOrder = append(selectOrder, "traffic_type")

	selects["log_destination_type"] = types.TfSelect{
		Label: "Enter log_destination_type:\n(Optional) The type of the logging destination.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"cloud-watch-logs", "s3"},
		},
	}
	selectOrder = append(selectOrder, "log_destination_type")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_flow_log", blockName, resourceBlock)
}

func AWSInternetGatewayPrompt() {
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

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) The VPC ID to create in.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_internet_gateway", blockName, resourceBlock)
}

func AWSMainRouteTableAssociationPrompt() {
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

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) The ID of the VPC whose main route table should be set",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["route_table_id"] = types.TfPrompt{
		Label: "Enter route_table_id:\n(Required) The ID of the Route Table to set as the new main route table for the target VPC",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "route_table_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_main_route_table_association", blockName, resourceBlock)
}

func AWSNatGatewayPrompt() {
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

	prompts["allocation_id"] = types.TfPrompt{
		Label: "Enter allocation_id:\n(Required) The Allocation ID of the Elastic IP address for the gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "allocation_id")

	prompts["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Required) The Subnet ID of the subnet in which to place the gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_nat_gateway", blockName, resourceBlock)
}

func AWSNetworkACLPrompt() {
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

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) The ID of the associated VPC.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["subnet_ids"] = types.TfPrompt{
		Label: "Enter subnet_ids e.g.[\"id1\",\"id2\"]:\n(Optional) A list of Subnet IDs to apply the ACL to",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_ids")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like ingress/egress [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_network_acl", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter ingress:\n(Optional) Specifies an ingress rule." +
		"\n1.from_port\n2.to_port\n3.rule_no\n4.action\n5.protocol\n6.cidr_block\n7.ipv6_cidr_block\n8.icmp_type\n9.icmp_code\n")

	ingressEgressPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	ingressEgressPrompt["from_port"] = types.TfPrompt{
		Label: "Enter from_port:\n(Required) The from port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "from_port")

	ingressEgressPrompt["to_port"] = types.TfPrompt{
		Label: "Enter to_port:\n(Required) The to port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "to_port")

	ingressEgressPrompt["rule_no"] = types.TfPrompt{
		Label: "Enter rule_no:\n(Required) The to port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "rule_no")

	ingressEgressPrompt["action"] = types.TfPrompt{
		Label: "Enter action:\n(Required) The action to take.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "action")

	ingressEgressPrompt["protocol"] = types.TfPrompt{
		Label: "Enter protocol:\n(Required) The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "protocol")

	ingressEgressPrompt["cidr_block"] = types.TfPrompt{
		Label: "Enter cidr_block:\n(Optional) The CIDR block to match. This must be a valid network mask.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cidr_block")

	ingressEgressPrompt["ipv6_cidr_block"] = types.TfPrompt{
		Label: "Enter ipv6_cidr_block:\n(Optional) The IPv6 CIDR block.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ipv6_cidr_block")

	ingressEgressPrompt["icmp_type"] = types.TfPrompt{
		Label: "Enter icmp_type:\n(Optional) The ICMP type to be used. Default 0.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "icmp_type")

	ingressEgressPrompt["icmp_code"] = types.TfPrompt{
		Label: "Enter icmp_code:\n(Optional) The ICMP type code to be used. Default 0.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "icmp_code")

	resourceBlock["ingress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	color.Green("\nEnter egress:\n(Optional) Specifies an egress rule." +
		"\n1.from_port\n2.to_port\n3.rule_no\n4.action\n5.protocol\n6.cidr_block\n7.ipv6_cidr_block\n8.icmp_type\n9.icmp_code\n")

	resourceBlock["egress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	builder.ResourceBuilder("aws_network_acl", blockName, resourceBlock)
}

func AWSNetworkACLRulePrompt() {
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

	prompts["network_acl_id"] = types.TfPrompt{
		Label: "Enter network_acl_id:\n(Required) The ID of the network ACL.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_acl_id")

	prompts["role_number"] = types.TfPrompt{
		Label: "Enter role_number:\n(Required) The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "role_number")

	prompts["egress"] = types.TfPrompt{
		Label: "Enter egress(true/false):\n(Optional, bool) Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default false",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "egress")

	prompts["protocol"] = types.TfPrompt{
		Label: "Enter protocol:\n(Required) The protocol. A value of -1 means all protocols.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "protocol")

	prompts["cidr_block"] = types.TfPrompt{
		Label: "Enter cidr_block:\n(Optional) The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24).",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "cidr_block")

	prompts["ipv6_cidr_block"] = types.TfPrompt{
		Label: "Enter ipv6_cidr_block:\n(Optional) The IPv6 CIDR block to allow or deny.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ipv6_cidr_block")

	prompts["from_port"] = types.TfPrompt{
		Label: "Enter from_port:\n(Optional) The from port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "from_port")

	prompts["to_port"] = types.TfPrompt{
		Label: "Enter to_port:\n(Optional) The to port to match.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "to_port")

	prompts["icmp_type"] = types.TfPrompt{
		Label: "Enter icmp_type:\n(Optional) ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "icmp_type")

	prompts["icmp_code"] = types.TfPrompt{
		Label: "Enter icmp_code:\n(Optional) ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "icmp_code")

	selects := map[string]types.TfSelect{}

	selects["rule_action"] = types.TfSelect{
		Label: "Enter rule_action:\n(Required) Indicates whether to allow or deny the traffic that matches the rule.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"allow", "deny"},
		},
	}
	selectOrder = append(selectOrder, "rule_action")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_network_acl_rule", blockName, resourceBlock)
}

func AWSNetworkInterfacePrompt() {
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

	prompts["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Required) Subnet ID to create the ENI in.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A description for the network interface.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["private_ips"] = types.TfPrompt{
		Label: "Enter private_ips e.g.[\"ip1\",\"ip2\"]:\n(Optional) List of private IPs to assign to the ENI.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "private_ips")

	prompts["private_ips_count"] = types.TfPrompt{
		Label: "Enter private_ips_count:\n(Optional) Number of secondary private IPs to assign to the ENI. The total number " +
			"\nof private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "private_ips_count")

	prompts["ipv6_addresses"] = types.TfPrompt{
		Label: "Enter ipv6_addresses:\n(Optional) One or more specific IPv6 addresses from the IPv6 CIDR block " +
			"\nrange of your subnet. You can't use this option if you're specifying ipv6_address_count",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ipv6_addresses")

	prompts["ipv6_addresses_count"] = types.TfPrompt{
		Label: "Enter ipv6_addresses_count:\n(Optional) One or more specific IPv6 addresses from the IPv6 CIDR block " +
			"\nrange of your subnet. You can't use this option if you're specifying ipv6_address_count",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ipv6_addresses_count")

	prompts["ipv6_addresses_count"] = types.TfPrompt{
		Label: "Enter ipv6_addresses_count:\n(Optional) One or more specific IPv6 addresses from the IPv6 CIDR block " +
			"\nrange of your subnet. You can't use this option if you're specifying ipv6_address_count",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "ipv6_addresses_count")

	prompts["security_groups"] = types.TfPrompt{
		Label: "Enter security_groups:\n(Optional) List of security group IDs to assign to the ENI.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "security_groups")

	prompts["source_dest_checks"] = types.TfPrompt{
		Label: "Enter source_dest_checks:\n(Optional) Whether to enable source destination checking for the ENI.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "source_dest_checks")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like attachment [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_network_interface", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter attachment:\n(Optional) Block to define the attachment of the ENI." +
		"\n1.instance\n2.device_index")

	attachmentPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	attachmentPrompt["instance"] = types.TfPrompt{
		Label: "Enter instance:\n(Required) ID of the instance to attach to.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "instance")

	attachmentPrompt["device_index"] = types.TfPrompt{
		Label: "Enter device_index:\n(Required) Integer to define the devices index.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "device_index")

	resourceBlock["attachment"] = builder.PSOrder(nestedPromptOrder, nil, attachmentPrompt, nil)

	builder.ResourceBuilder("aws_network_inteface", blockName, resourceBlock)
}

func AWSNetworkInterfaceAttachmentPrompt() {
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

	prompts["instance_id"] = types.TfPrompt{
		Label: "Enter instance_id:\n(Required) Instance ID to attach.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_id")

	prompts["network_interface_id"] = types.TfPrompt{
		Label: "Enter network_interface_id:\n(Required) ENI ID to attach.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_interface_id")

	prompts["device_index"] = types.TfPrompt{
		Label: "Enter device_index:\n(Required) Network interface index (int).",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "device_index")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_network_interface_attachment", blockName, resourceBlock)
}

func AWSNetworkInterfaceSGAttachmentPrompt() {
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

	prompts["security_group_id"] = types.TfPrompt{
		Label: "Enter security_group_id:\n(Required) The ID of the security group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "security_group_id")

	prompts["network_interface_id"] = types.TfPrompt{
		Label: "Enter network_interface_id:\n(Required) The ID of the network interface to attach to.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_interface_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_network_interface_sg_attachment", blockName, resourceBlock)
}

func AWSRoutePrompt() {
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

	prompts["route_table_id"] = types.TfPrompt{
		Label: "Enter route_table_id:\n(Required) The ID of the routing table.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "route_table_id")

	prompts["destination_cidr_block"] = types.TfPrompt{
		Label: "Enter destination_cidr_block:\n(Optional) The destination CIDR block.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "destination_cidr_block")

	prompts["destination_ipv6_cidr_block"] = types.TfPrompt{
		Label: "Enter destination_ipv6_cidr_block:\n(Optional) The destination IPv6 CIDR block.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "destination_ipv6_cidr_block")

	prompts["egress_only_gateway_id"] = types.TfPrompt{
		Label: "Enter egress_only_gateway_id:\n(Optional) Identifier of a VPC Egress Only Internet Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "egress_only_gateway_id")

	prompts["gateway_id"] = types.TfPrompt{
		Label: "Enter gateway_id:\n(Optional) Identifier of a VPC internet gateway or a virtual private gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "gateway_id")

	prompts["instance_id"] = types.TfPrompt{
		Label: "Enter instance_id:\n(Optional) Identifier of an EC2 instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_id")

	prompts["nat_gateway_id"] = types.TfPrompt{
		Label: "Enter nat_gateway_id:\n(Optional) Identifier of a VPC NAT gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "nat_gateway_id")

	prompts["local_gateway_id"] = types.TfPrompt{
		Label: "Enter local_gateway_id:\n(Optional) Identifier of a Outpost local gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "local_gateway_id")

	prompts["network_interface_id"] = types.TfPrompt{
		Label: "Enter network_interface_id:\n(Optional) Identifier of an EC2 network interface.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_interface_id")

	prompts["transit_gateway_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_id:\n(Optional) Identifier of an EC2 Transit Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_id")

	prompts["vpc_endpoint_id"] = types.TfPrompt{
		Label: "Enter vpc_endpoint_id:\n(Optional) Identifier of a VPC Endpoint.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_endpoint_id")

	prompts["vpc_peering_connection_id"] = types.TfPrompt{
		Label: "Enter vpc_peering_connection_id:\n(Optional) Identifier of a VPC peering connection.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_peering_connection_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_route", blockName, resourceBlock)
}

func AWSRouteTablePrompt() {
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

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) The VPC ID.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["propagating_vgws"] = types.TfPrompt{
		Label: "Enter propagating_vgws e.g.[\"g1\",\"g2\"]:\n(Optional) A list of virtual gateways for propagation.",
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
		builder.ResourceBuilder("aws_route_table", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter route:\n(Optional) A list of route objects." +
		"\n1.cidr_block\n2.ipv6_cidr_block\n3.egress_only_gateway_id\n4.gateway_id\n5.instance_id\n6.nat_gateway_id\n7.local_gateway_id\n8.network_interface_id\n9.transit_gateway_id\n10.vpc_endpoint_id\n11.vpc_peering_connection_id\n")

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

	routePrompt["local_gateway_id"] = types.TfPrompt{
		Label: "Enter local_gateway_id:\n(Optional) Identifier of a Outpost local gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "local_gateway_id")

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
		Label: "Enter vpc_endpoint_id:\n(Optional) Identifier of a VPC Endpoint.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "vpc_endpoint_id")

	routePrompt["vpc_peering_connection_id"] = types.TfPrompt{
		Label: "Enter vpc_peering_connection_id:\n(Optional) Identifier of a VPC peering connection.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "vpc_peering_connection_id")

	resourceBlock["route"] = builder.PSOrder(nestedPromptOrder, nil, routePrompt, nil)

	builder.ResourceBuilder("aws_route_table", blockName, resourceBlock)
}

func AWSRouteTableAssociationPrompt() {
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

	prompts["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Optional) The subnet ID to create an association. Conflicts with gateway_id",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["gateway_id"] = types.TfPrompt{
		Label: "Enter gateway_id:\n(Optional) The gateway ID to create an association. Conflicts with subnet_id",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "gateway_id")

	prompts["route_table_id"] = types.TfPrompt{
		Label: "Enter route_table_id:\n(Required) The ID of the routing table to associate with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "route_table_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_route_table_association", blockName, resourceBlock)
}

func AWSSecurityGroupPrompt() {
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

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Optional, Forces new resource) The name of the security group. If omitted, Terraform will assign a random, unique name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["name_prefix"] = types.TfPrompt{
		Label: "Enter name_prefix:\n(Optional, Forces new resource) Creates a unique name beginning with the specified prefix. Conflicts with name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name_prefix")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional, Forces new resource) The security group description. Defaults to \"Managed by Terraform\". " +
			"\nCannot be \"\". NOTE: This field maps to the AWS GroupDescription attribute, " +
			"\nfor which there is no Update API. If you'd like to classify your security groups in a " +
			"\nway that can be updated, use tags",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["revoke_rules_on_delete"] = types.TfPrompt{
		Label: "Enter revoke_rules_on_delete(true/false):\n(Optional) Instruct Terraform to revoke all of the Security Groups attached ingress " +
			"\nand egress rules before deleting the rule itself. This is normally not needed, however certain AWS services " +
			"\nsuch as Elastic Map Reduce may automatically add required rules to security groups used with the service, " +
			"\nand those rules may contain a cyclic dependency that prevent the security groups from being destroyed " +
			"\nwithout removing the dependency first. Default false",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "revoke_rules_on_delete")

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Optional, Forces new resource) The VPC ID.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like ingress/egress [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_security_group", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter ingress:\n(Optional) Specifies an ingress rule." +
		"\n1.cidr_blocks\n2.ipv6_cidr_blocks\n3.prefix_list_ids\n4.from_port\n5.protocol\n6.security_groups\n7.self\n8.to_port\n9.description\n")

	ingressEgressPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	ingressEgressPrompt["cidr_blocks"] = types.TfPrompt{
		Label: "Enter cidr_blocks:\n(Optional) List of CIDR blocks.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cidr_blocks")

	ingressEgressPrompt["ipv6_cidr_blocks"] = types.TfPrompt{
		Label: "Enter ipv6_cidr_blocks:\n(Optional) List of IPv6 CIDR blocks.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ipv6_cidr_blocks")

	ingressEgressPrompt["prefix_list_ids"] = types.TfPrompt{
		Label: "Enter prefix_list_ids:\n(Optional) List of prefix list IDs.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "prefix_list_ids")

	ingressEgressPrompt["from_port"] = types.TfPrompt{
		Label: "Enter from_port:\n(Required) The start port (or ICMP type number if protocol is \"icmp\" or \"icmpv6\")",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "from_port")

	ingressEgressPrompt["protocol"] = types.TfPrompt{
		Label: "Enter protocol:\n(Required) The protocol. If you select a protocol of \"-1\" (semantically equivalent to \"all\", which is " +
			"\nnot a valid value here), you must specify a \"from_port\" and \"to_port\" equal to 0. The supported values are " +
			"\ndefined in the \"IpProtocol\" argument on the IpPermission API reference. This argument is normalized to a " +
			"\nlowercase value to match the AWS API requirement when using with Terraform 0.12.x and above, please make " +
			"\nsure that the value of the protocol is specified as lowercase when using with older version of Terraform " +
			"\nto avoid an issue during upgrade.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "protocol")

	ingressEgressPrompt["security_groups"] = types.TfPrompt{
		Label: "Enter security_groups:\n(Optional) List of security group Group Names if using EC2-Classic, or Group IDs if using a VPC.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "security_groups")

	ingressEgressPrompt["self"] = types.TfPrompt{
		Label: "Enter self:\n(Optional) If true, the security group itself will be added as a source to this ingress/egress rule.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "self")

	ingressEgressPrompt["to_port"] = types.TfPrompt{
		Label: "Enter to_port:\n(Required) The end range port (or ICMP code if protocol is \"icmp\").",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "to_port")

	ingressEgressPrompt["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Description of this ingress/egress rule.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "description")

	resourceBlock["ingress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	color.Green("\nEnter egress:\n(Optional) Specifies an egress rule." +
		"\n1.cidr_blocks\n2.ipv6_cidr_blocks\n3.prefix_list_ids\n4.from_port\n5.protocol\n6.security_groups\n7.self\n8.to_port\n9.description\n")

	resourceBlock["egress"] = builder.PSOrder(nestedPromptOrder, nil, ingressEgressPrompt, nil)

	builder.ResourceBuilder("aws_security_group", blockName, resourceBlock)
}