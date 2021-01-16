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

	schema := []types.Schema{
		{

			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The ID of the associated VPC.",
		},
		{

			Field: "subnet_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc:   "(Optional) A list of Subnet IDs to apply the ACL to",
		},
		{

			Field: "tags",
			Ex:    "k1=v1,k2=v2",
			Doc:   "(Optional) A list of Subnet IDs to apply the ACL to",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	ingressEgressSchema := []types.Schema{
		{
			Field:     "from_port",
			Ex:        "",
			Doc:       "(Required) The from port to match.",
			Validator: utils.IntValidator,
		},
		{
			Field:     "to_port",
			Ex:        "",
			Doc:       "(Required) The to port to match.",
			Validator: utils.IntValidator,
		},
		{
			Field:     "rule_no",
			Ex:        "",
			Doc:       "(Required) The rule number. Used for ordering.",
			Validator: utils.IntValidator,
		},
		{

			Field: "action",
			Ex:    "",
			Doc:   "(Required) The action to take.",
		},
		{

			Field: "protocol",
			Ex:    "-1",
			Doc:   "(Required) The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.",
		},
		{

			Field: "cidr_block",
			Ex:    "172.16.0.0/24",
			Doc:   "(Optional) The CIDR block to match. This must be a valid network mask.",
		},
		{

			Field: "ipv6_cidr_block",
			Ex:    "172.16.0.0/24",
			Doc:   "(Optional) The IPv6 CIDR block.",
		},
		{
			Field:     "icmp_type",
			Ex:        "",
			Doc:       "(Optional) The ICMP type to be used. Default 0.",
			Validator: utils.IntValidator,
		},
		{
			Field:     "icmp_code",
			Ex:        "",
			Doc:       "(Optional) The ICMP type code to be used. Default 0.",
			Validator: utils.IntValidator,
		},
	}

	resourceBlock["ingress"] = builder.PSOrder(types.ProvidePS(ingressEgressSchema))

	color.Green("\nEnter egress:\n(Optional) Specifies an egress rule." +
		"\n1.from_port\n2.to_port\n3.rule_no\n4.action\n5.protocol\n6.cidr_block\n7.ipv6_cidr_block\n8.icmp_type\n9.icmp_code\n")

	resourceBlock["egress"] = builder.PSOrder(types.ProvidePS(ingressEgressSchema))

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

	schema := []types.Schema{
		{
			Field: "network_acl_id",
			Ex:    "",
			Doc:   "(Required) The ID of the network ACL.",
		},
		{
			Field:     "role_number",
			Ex:        "",
			Doc:       "(Required) The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.",
			Validator: utils.IntValidator,
		},
		{
			Field:     "egress",
			Ex:        "(true/false)",
			Doc:       "(Optional, bool) Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default false",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "protocol",
			Ex:        "10",
			Doc:       "(Required) The protocol. A value of -1 means all protocols.",
			Validator: utils.IntValidator,
		},
		{
			Field: "cidr_block",
			Ex:    "172.16.0.0/24",
			Doc:   "(Optional) The network range to allow or deny, in CIDR notation.",
		},
		{
			Field: "ipv6_cidr_block",
			Ex:    "2001:db8:1234:1a00::/56",
			Doc:   "(Optional) The IPv6 CIDR block to allow or deny.",
		},
		{
			Field:     "from_port",
			Ex:        "443",
			Doc:       "(Optional) The from port to match.",
			Validator: utils.IntValidator,
		},
		{
			Field:     "to_port",
			Ex:        "443",
			Doc:       "(Optional) The to port to match.",
			Validator: utils.IntValidator,
		},
		{
			Field:     "icmp_type",
			Ex:        "-1",
			Doc:       "(Optional) ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1",
			Validator: utils.IntValidator,
		},
		{
			Field:     "icmp_code",
			Ex:        "-1",
			Doc:       "(Optional) ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1",
			Validator: utils.IntValidator,
		},
		{
			Field: "rule_action",
			Doc:   "(Required) Indicates whether to allow or deny the traffic that matches the rule.",
			Items: []string{"allow", "deny"},
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "subnet_id",
			Ex:    "",
			Doc:   "(Required) Subnet ID to create the ENI in.",
		},
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) A description for the network interface.",
		},
		{
			Field: "private_ips",
			Ex:    "[\"ip1\",\"ip2\"]",
			Doc:   "(Optional) List of private IPs to assign to the ENI.",
		},
		{
			Field: "private_ips_count",
			Ex:    "10",
			Doc: "(Optional) Number of secondary private IPs to assign to the ENI. " +
				"\nThe total number nof private IPs will be 1 + private_ips_count, " +
				"\nas a primary private IP will be assiged to an ENI by default.",
			Validator: utils.IntValidator,
		},
		{
			Field: "ipv6_addresses",
			Ex:    "",
			Doc: "(Optional) One or more specific IPv6 addresses from the IPv6 CIDR block " +
				"\nrange of your subnet. You can't use this option if you're specifying ipv6_address_count",
		},
		{
			Field: "ipv6_addresses_count",
			Ex:    "10",
			Doc: "(Optional) One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. " +
				"\nYou can't use this option if you're specifying ipv6_address_count",
			Validator: utils.IntValidator,
		},
		{
			Field: "security_groups",
			Ex:    "",
			Doc:   "(Optional) List of security group IDs to assign to the ENI.",
		},
		{
			Field: "source_dest_checks",
			Ex:    "(true/false)",
			Doc:   "(Optional) Whether to enable source destination checking for the ENI.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	attachmentSchema := []types.Schema{
		{
			Field: "instance",
			Ex:    "",
			Doc:   "(Required) ID of the instance to attach to.",
		},
		{
			Field:     "device_index",
			Ex:        "1",
			Doc:       "(Required) Integer to define the devices index.",
			Validator: utils.IntValidator,
		},
	}

	resourceBlock["attachment"] = builder.PSOrder(types.ProvidePS(attachmentSchema))

	builder.ResourceBuilder("aws_network_interface", blockName, resourceBlock)
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

	schema := []types.Schema{
		{
			Field: "instance_id",
			Ex:    "instance-123",
			Doc:   "(Required) Instance ID to attach.",
		},
		{
			Field: "network_interface_id",
			Ex:    "nid-123",
			Doc:   "(Required) ENI ID to attach.",
		},
		{
			Field:     "device_index",
			Ex:        "1",
			Doc:       "(Required) Network interface index.",
			Validator: utils.IntValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "security_group_id",
			Ex:    "sec-gr-123",
			Doc:   "(Required) The ID of the security group.",
		},
		{
			Field: "network_interface_id",
			Ex:    "net-id-123",
			Doc:   "(Required) The ID of the network interface to attach to.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "route_table_id",
			Ex:    "",
			Doc:   "(Required) The ID of the routing table.",
		},
		{
			Field: "destination_cidr_block",
			Ex:    "172.2.0.0/16",
			Doc:   "(Optional) The destination CIDR block.",
		},
		{
			Field: "destination_ipv6_cidr_block",
			Ex:    "2001:db8:1234:1a00::/56",
			Doc:   "(Optional) The destination IPv6 CIDR block.",
		},
		{
			Field: "egress_only_gateway_id",
			Ex:    "gateway-123",
			Doc:   "(Optional) Identifier of a VPC Egress Only Internet Gateway.",
		},
		{
			Field: "gateway_id",
			Ex:    "gateway-123",
			Doc:   "(Optional) Identifier of a VPC internet gateway or a virtual private gateway.",
		},
		{
			Field: "instance_id",
			Ex:    "instance-123",
			Doc:   "(Optional) Identifier of an EC2 instance.",
		},
		{
			Field: "nat_gateway_id",
			Ex:    "gateway-nat-123",
			Doc:   "(Optional) Identifier of a VPC NAT gateway.",
		},
		{
			Field: "local_gateway_id",
			Ex:    "gateway-local-123",
			Doc:   "(Optional) Identifier of a Outpost local gateway.",
		},
		{
			Field: "network_interface_id",
			Ex:    "ni-123",
			Doc:   "(Optional) Identifier of an EC2 network interface.",
		},
		{
			Field: "transit_gateway_id",
			Ex:    "tra-gateway-123",
			Doc:   "(Optional) Identifier of an EC2 Transit Gateway.",
		},
		{
			Field: "vpc_endpoint_id",
			Ex:    "vpc-ei-123",
			Doc:   "(Optional) Identifier of a VPC Endpoint.",
		},
		{
			Field: "vpc_peering_connection_id",
			Ex:    "vpc-pci-123",
			Doc:   "(Optional) Identifier of a VPC peering connection.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "vpc-2f09a348",
			Doc:   "(Required) The VPC ID.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
		{
			Field: "propagating_vgws",
			Ex:    "[\"g1\",\"g2\"]",
			Doc:   "(Optional) A list of virtual gateways for propagation.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	routeSchema := []types.Schema{
		{
			Field: "cidr_block",
			Ex:    "172.2.0.0/16",
			Doc:   "(Required) The CIDR block of the route.",
		},
		{
			Field: "ipv6_cidr_block",
			Ex:    "2001:db8:1234:1a00::/56",
			Doc:   "(Optional) The Ipv6 CIDR block of the route.",
		},
		{
			Field: "egress_only_gateway_id",
			Ex:    "gateway-123",
			Doc:   "(Optional) Identifier of a VPC Egress Only Internet Gateway.",
		},
		{
			Field: "gateway_id",
			Ex:    "gateway-123",
			Doc:   "(Optional) Identifier of a VPC internet gateway or a virtual private gateway.",
		},
		{
			Field: "instance_id",
			Ex:    "instance-123",
			Doc:   "(Optional) Identifier of an EC2 instance.",
		},
		{
			Field: "nat_gateway_id",
			Ex:    "gateway-nat-123",
			Doc:   "(Optional) Identifier of a VPC NAT gateway.",
		},
		{
			Field: "local_gateway_id",
			Ex:    "gateway-local-123",
			Doc:   "(Optional) Identifier of a Outpost local gateway.",
		},
		{
			Field: "network_interface_id",
			Ex:    "nid-123",
			Doc:   "(Optional) Identifier of an EC2 network interface.",
		},
		{
			Field: "transit_gateway_id",
			Ex:    "tra-gateway-123",
			Doc:   "(Optional) Identifier of an EC2 Transit Gateway.",
		},
		{
			Field: "vpc_endpoint_id",
			Ex:    "vpd-eid-123",
			Doc:   "(Optional) Identifier of a VPC Endpoint.",
		},
		{
			Field: "vpc_peering_connection_id",
			Ex:    "vpc-pci-123",
			Doc:   "(Optional) Identifier of a VPC peering connection.",
		},
	}

	resourceBlock["route"] = builder.PSOrder(types.ProvidePS(routeSchema))

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

	schema := []types.Schema{
		{
			Field: "subnet_id",
			Ex:    "subnet-0bb1c79de3EXAMPLE",
			Doc:   "(Optional) The subnet ID to create an association. Conflicts with gateway_id",
		},
		{
			Field: "gateway_id",
			Ex:    "gateway-0bb1c79de3EXAMPLE",
			Doc:   "(Optional) The gateway ID to create an association. Conflicts with subnet_id",
		},
		{
			Field: "route_table_id",
			Ex:    "rtb-09ba434c1bEXAMPLE",
			Doc:   "(Required) The ID of the routing table to associate with.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "name",
			Ex:    "new group",
			Doc:   "(Optional, Forces new resource) The name of the security group. If omitted, Terraform will assign a random, unique name",
		},
		{
			Field: "name_prefix",
			Ex:    "new prefix",
			Doc:   "(Optional, Forces new resource) Creates a unique name beginning with the specified prefix. Conflicts with name",
		},
		{Field: "description",
			Ex: "example description",
			Doc: "(Optional, Forces new resource) The security group description. " +
				"\nDefaults to \"Managed by Terraform\". Cannot be \"\". " +
				"\nNOTE: This field maps to the AWS GroupDescription attribute, " +
				"\nfor which there is no Update API. If you'd like to classify your " +
				"\nsecurity groups in a way that can be updated, use tags",
		},
		{
			Field: "revoke_rules_on_delete",
			Ex:    "(true/false)",
			Doc: "(Optional) Instruct Terraform to revoke all of the Security Groups attached ingress and " +
				"\negress rules before deleting the rule itself. This is normally not needed, however " +
				"\ncertain AWS services such as Elastic Map Reduce may automatically add required rules " +
				"\nto security groups used with the service, and those rules may contain a cyclic " +
				"\ndependency that prevent the security groups from being destroyed without removing " +
				"\nthe dependency first. Default false",
			Validator: utils.BoolValidator,
		},
		{
			Field: "vpc_id",
			Ex:    "vpc-2f09a348",
			Doc:   "(Optional, Forces new resource) The VPC ID.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	ingressEgressSchema := []types.Schema{
		{
			Field: "cidr_blocks",
			Ex:    "[\"172.2.0.0/16\",\"173.0.0.1/24\"]",
			Doc:   "(Optional) List of CIDR blocks.",
		},
		{
			Field: "ipv6_cidr_blocks",
			Ex:    "[\"2001:db8:1234:1a00::/56\"]",
			Doc:   "(Optional) List of IPv6 CIDR blocks.",
		},
		{
			Field: "prefix_list_ids",
			Ex:    "[\"pl-63a5400a\"]",
			Doc:   "(Optional) List of IPv6 CIDR blocks.",
		},
		{
			Field:     "from_port",
			Ex:        "443",
			Doc:       "(Required) The start port (or ICMP type number if protocol is \"icmp\" or \"icmpv6\")",
			Validator: utils.IntValidator,
		},
		{
			Field: "protocol",
			Ex:    "-1",
			Doc: "(Required) The protocol. If you select a protocol of \"-1\" (semantically equivalent to \"all\", " +
				"\nwhich is not a valid value here), you must specify a \"from_port\" and \"to_port\" " +
				"\nequal to 0. The supported values are defined in the \"IpProtocol\" argument on the " +
				"\nIpPermission API reference. This argument is normalized to a lowercase value to match " +
				"\nthe AWS API requirement when using with Terraform 0.12.x and above, please make sure " +
				"\nthat the value of the protocol is specified as lowercase when using with older " +
				"\nversion of Terraform to avoid an issue during upgrade.",
			Validator: utils.IntValidator,
		},
		{
			Field: "security_groups",
			Ex:    "[\"g1\",\"g2\"]",
			Doc:   "(Optional) List of security group Group Names if using EC2-Classic, or Group IDs if using a VPC.",
		},
		{
			Field:     "self",
			Ex:        "(true/false)",
			Doc:       "(Optional) If true, the security group itself will be added as a source to this ingress/egress rule.",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "to_port",
			Ex:        "443",
			Doc:       "(Required) The end range port (or ICMP code if protocol is \"icmp\").",
			Validator: utils.IntValidator,
		},
		{
			Field: "description",
			Ex:    "example description",
			Doc:   "(Optional) Description of this ingress/egress rule.",
		},
	}

	resourceBlock["ingress"] = builder.PSOrder(types.ProvidePS(ingressEgressSchema))

	color.Green("\nEnter egress:\n(Optional) Specifies an egress rule." +
		"\n1.cidr_blocks\n2.ipv6_cidr_blocks\n3.prefix_list_ids\n4.from_port\n5.protocol\n6.security_groups\n7.self\n8.to_port\n9.description\n")

	resourceBlock["egress"] = builder.PSOrder(types.ProvidePS(ingressEgressSchema))

	builder.ResourceBuilder("aws_security_group", blockName, resourceBlock)
}

func AWSSecurityGroupRulePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "cidr_blocks",
			Ex:    "[\"172.2.0.0/16\",\"173.0.0.1/24\"]",
			Doc:   "(Optional) List of CIDR blocks. Cannot be specified with source_security_group_id",
		},
		{
			Field: "ipv6_cidr_blocks",
			Ex:    "[\"2001:db8:1234:1a00::/56\"]",
			Doc:   "(Optional) List of IPv6 CIDR blocks.",
		},
		{
			Field: "prefix_list_ids",
			Ex:    "[\"pl-63a5400a\"]",
			Doc:   "(Optional) List of Prefix List IDs.",
		},
		{
			Field:     "from_port",
			Ex:        "443",
			Doc:       "(Required) The start port (or ICMP type number if protocol is \"icmp\" or \"icmpv6\").",
			Validator: utils.IntValidator,
		},
		{
			Field: "protocol",
			Ex:    "tcp",
			Doc: "(Required) The protocol. If not icmp, icmpv6, tcp, udp, or all use the protocol number" +
				"\nCheckout https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml",
		},
		{
			Field: "security_group_id",
			Ex:    "tcp",
			Doc:   "(Required) The security group to apply this rule to.",
		},
		{
			Field: "source_security_group_id",
			Ex:    "sg-903004f8",
			Doc: "(Optional) The security group id to allow access to/from, depending on the type. " +
				"\nCannot be specified with cidr_blocks and self.",
		},
		{
			Field: "self",
			Ex:    "(true/false)",
			Doc: "(Optional) If true, the security group itself will be added as a source to this ingress rule. " +
				"\nCannot be specified with source_security_group_id",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "to_port",
			Ex:        "443",
			Doc:       "(Required) The end port (or ICMP code if protocol is \"icmp\").",
			Validator: utils.IntValidator,
		},
		{
			Field: "description",
			Ex:    "an example description",
			Doc:   "(Optional) Description of the rule.",
		},
		{
			Type:  "select",
			Field: "type",
			Ex:    "",
			Doc:   "(Required) The type of rule being created. Valid options are ingress (inbound) or egress (outbound).",
			Items: []string{"ingress", "egress"},
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_security_group_rule", blockName, resourceBlock)
}

func AWSSubnetPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "availability_zone",
			Ex:    "us-east-1a",
			Doc:   "(Optional) The AZ for the subnet.",
		},
		{
			Field: "availability_zone_id",
			Ex:    "use-az2",
			Doc:   "(Optional) The AZ ID of the subnet.",
		},
		{
			Field: "cidr_block",
			Ex:    "172.2.0.0/16",
			Doc:   "(Required) The CIDR block for the subnet.",
		},
		{Field: "ipv6_cidr_block",
			Ex: "2001:db8:1234:1a00::/56",
			Doc: "(Optional) The IPv6 network range for the subnet, in CIDR notation. " +
				"\nThe subnet size must use a /64 prefix length.",
		},
		{Field: "map_public_ip_on_launch",
			Ex: "(true/false)",
			Doc: "(Optional) Specify true to indicate that instances launched " +
				"\ninto the subnet should be assigned a public IP address. Default is false",
			Validator: utils.BoolValidator,
		},
		{
			Field: "outpost_arn",
			Ex:    "arn:aws:outposts:us-west-2:123456789012:outpost/op-xxxxxxxxxxxxxxxx",
			Doc:   "(Optional) The Amazon Resource Name (ARN) of the Outpost.",
		},
		{
			Field: "assign_ipv6_address_on_creation",
			Ex:    "(true/false)",
			Doc: "(Optional) Specify true to indicate that network interfaces created " +
				"\nin the specified subnet should be assigned an IPv6 address. Default is false",
			Validator: utils.BoolValidator,
		},
		{
			Field: "vpc_id",
			Ex:    "vpc-2f09a348",
			Doc:   "(Required) The VPC ID.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k1=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_subnet", blockName, resourceBlock)
}

func AWSVPCDHCPOptionsPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "domain_name",
			Ex:    "service.consul",
			Doc:   "(Optional) the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the search value in the /etc/resolv.conf",
		},
		{
			Field: "domain_name_servers",
			Ex:    "[\"s1\", \"s2\"]",
			Doc:   "(Optional) List of name servers to configure in /etc/resolv.conf. If you want to use the default AWS nameservers you should set this to AmazonProvidedDNS",
		},
		{
			Field: "ntp_servers",
			Ex:    "[\"s1\", \"s2\"]",
			Doc:   "(Optional) List of NTP servers to configure.",
		},
		{
			Field: "netbios_name_servers",
			Ex:    "[\"s1\", \"s2\"]",
			Doc:   "(Optional) List of NETBIOS name servers.",
		},
		{
			Field: "netbios_node_type",
			Ex:    "1",
			Doc:   "(Optional) The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see RFC 2132.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpc_dhcp_options", blockName, resourceBlock)
}

func AWSVPCDHCPOptionsAssociationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The ID of the VPC to which we would like to associate a DHCP Options Set.",
		},
		{
			Field: "dhcp_options_id",
			Ex:    "dhcp-id-123",
			Doc:   "(Required) The ID of the DHCP Options Set to associate to the VPC.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpc_dhcp_options_association", blockName, resourceBlock)
}

func AWSVPCEndpointPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "service_name",
			Ex:    "srv-123",
			Doc:   "(Required) The service name. For AWS services the service name is usually in the form com.amazonaws.<region>.<service> (the SageMaker Notebook service is an exception to this rule, the service name is in the form aws.sagemaker.<region>.notebook",
		},
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The ID of the VPC in which the endpoint will be used.",
		},
		{
			Field:     "auto_accept",
			Ex:        "(true/false)",
			Doc:       "(Optional) Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).",
			Validator: utils.BoolValidator,
		},
		{
			Field: "policy",
			Ex:    "",
			Doc: "(Optional) A policy to attach to the endpoint that controls access to the service. " +
				"\nDefaults to full access. All Gateway and some Interface endpoints support " +
				"\npolicies - see the relevant AWS documentation for more details. For more " +
				"\ninformation about building AWS IAM policy documents with Terraform, see the AWS IAM Policy Document Guide.",
		},
		{
			Field: "private_dns_enabled",
			Ex:    "(true/false)",
			Doc: "(Optional) AWS services and AWS Marketplace partner services only) Whether " +
				"\nor not to associate a private hosted zone with the specified VPC. " +
				"\nApplicable for endpoints of type Interface. Defaults to false.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "route_table_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc:   "(Optional) One or more route table IDs. Applicable for endpoints of type Gateway",
		},
		{
			Field: "subnet_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc: "(Optional) The ID of one or more subnets in which to create a network interface " +
				"\nfor the endpoint. Applicable for endpoints of type GatewayLoadBalancer and Interface",
		},
		{
			Field: "security_group_ids",
			Ex:    "[\"id1\",\"id2\"]",
			Doc: "(Optional) The ID of one or more security groups to associate with the network interface. " +
				"\nRequired for endpoints of type Interface.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
		{
			Type:  "select",
			Field: "vpd_endpoint_type",
			Doc:   "(Optional) The VPC endpoint type",
			Items: []string{"Gateway", "GatewayLoadBalancer", "Interface"},
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	timeoutSchema := []types.Schema{
		{
			Field: "create",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for creating a VPC endpoint",
		},
		{
			Field: "update",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for VPC endpoint modifications",
		},
		{
			Field: "delete",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for destroying VPC endpoints",
		},
	}

	resourceBlock["timeout"] = builder.PSOrder(types.ProvidePS(timeoutSchema))

	builder.ResourceBuilder("aws_vpc_endpoint", blockName, resourceBlock)
}

func AWSVPCEndpointConnectionNotificationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	color.Yellow("\nOne of vpc_endpoint_service_id or vpc_endpoint_id must be specified.\n")

	schema := []types.Schema{
		{
			Field: "vpc_endpoint_service_id",
			Ex:    "vpc-srv-id-123",
			Doc:   "(Optional) The ID of the VPC Endpoint Service to receive notifications for.",
		},
		{
			Field: "vpc_endpoint_id",
			Ex:    "vpc-ei-123",
			Doc:   "(Optional) The ID of the VPC Endpoint to receive notifications for.",
		},
		{
			Field: "connection_notification_arn",
			Ex:    "",
			Doc:   "(Required) The ARN of the SNS topic for the notifications.",
		},
		{
			Field: "connection_events",
			Ex:    "[\"e1\",\"e2\"]",
			Doc:   "(Required) One or more endpoint events for which to receive notifications.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpc_endpoint_connection_notification", blockName, resourceBlock)
}

func AWSVPCEndpointRouteTableAssociationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "route_table_id",
			Ex:    "rti-123",
			Doc:   "(Required) Identifier of the EC2 Route Table to be associated with the VPC Endpoint.",
		},
		{
			Field: "vpc_endpoint_id",
			Ex:    "vpc-ei-123",
			Doc:   "(Required) Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpc_endpoint_route_table_association", blockName, resourceBlock)
}

func AWSVPCEndpointServicePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field:     "acceptance_required",
			Ex:        "(true/false)",
			Doc:       "(Required) Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - true or false.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "allowed_principals",
			Ex:    "[\"p1\",\"p2\"]",
			Doc:   "(Optional) The ARNs of one or more principals allowed to discover the endpoint service.",
		},
		{
			Field: "gateway_load_balancer_arns",
			Ex:    "[\"a1\",\"a2\"]",
			Doc:   "(Optional) Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.",
		},
		{
			Field: "network_load_balancer_arns",
			Ex:    "[\"a1\",\"a2\"]",
			Doc:   "(Optional) Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
		{
			Field: "private_dns_name",
			Ex:    "dns-123",
			Doc:   "(Optional) The private DNS name for the service.",
		},
	}
	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpc_endpoint_service", blockName, resourceBlock)
}

func AWSCallerIdentityPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "vpc_endpoint_service_id",
			Ex:    "vpc-epsrv-123",
			Doc:   "(Required) The ID of the VPC endpoint service to allow permission.",
		},
		{
			Field: "principal_arn",
			Ex:    "",
			Doc:   "(Required) The ARN of the principal to allow permissions.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_caller_identity", blockName, resourceBlock)
}

func AWSVPCEndpointSubnetAssociationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "vpc_endpoint_id",
			Ex:    "vpc-ei-123",
			Doc:   "(Required) The ID of the VPC endpoint with which the subnet will be associated.",
		},
		{
			Field: "subnet_id",
			Ex:    "subnet-123",
			Doc:   "(Required) The ID of the subnet to be associated with the VPC endpoint.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	timeoutSchema := []types.Schema{
		{
			Field: "create",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for creating a VPC endpoint",
		},
		{
			Field: "delete",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for destroying VPC endpoints",
		},
	}

	resourceBlock["timeout"] = builder.PSOrder(types.ProvidePS(timeoutSchema))

	builder.ResourceBuilder("aws_vpc_endpoint_subnet_association", blockName, resourceBlock)
}

func AWSVPCIPV4CIDRBlockAssociationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "cidr_block",
			Ex:    "172.16.0.0/24",
			Doc:   "(Required) The additional IPv4 CIDR block to associate with the VPC.",
		},
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The ID of the VPC to make the association with.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like timeout [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_vpc_ipv4_cidr_block_association", blockName, resourceBlock)
		return
	}

	timeoutSchema := []types.Schema{
		{
			Field: "create",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for creating a VPC endpoint",
		},
		{
			Field: "delete",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for destroying VPC endpoints",
		},
	}

	resourceBlock["timeout"] = builder.PSOrder(types.ProvidePS(timeoutSchema))

	builder.ResourceBuilder("aws_vpc_ipv4_cidr_block_association", blockName, resourceBlock)
}

func AWSVPCPeeringConnectionPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "peer_owner_id",
			Ex:    "peer-owner-123",
			Doc:   "(Optional) The AWS account ID of the owner of the peer VPC. Defaults to the account ID the AWS provider is currently connected to.",
		},
		{
			Field: "peer_vpc_id",
			Ex:    "peer-vpc-123",
			Doc:   "(Required) The ID of the VPC with which you are creating the VPC Peering Connection.",
		},
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The ID of the requester VPC.",
		},
		{
			Field:     "auto_accept",
			Ex:        "(true/false)",
			Doc:       "(Optional) Accept the peering (both VPCs need to be in the same AWS account).",
			Validator: utils.BoolValidator,
		},
		{
			Field: "peer_region",
			Ex:    "",
			Doc: "(Optional) The region of the accepter VPC of the [VPC Peering Connection]. " +
				"\nauto_accept must be false, and use the aws_vpc_peering_connection_accepter " +
				"\nto manage the accepter side.",
		},
		{
			Field: "tags",
			Ex:    "k1=v1,k2=v2",
			Doc:   "(Optional) A map of tags to assign to the resource.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like accepter/requester/timeout [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_vpc_peering_connection", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter accepter:\n(Optional) - An optional configuration block that allows for [VPC Peering Connection] " +
		"\n(https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the " +
		"\nVPC that accepts the peering connection (a maximum of one)." +
		"\n\nThe accepter block supports the following arguments:" +
		"\n1.allow_remote_vpc_dns_resolution\n2.allow_classic_link_to_remote_vpc\n3.allow_vpc_to_remote_classic_link")

	accepterRequesterSchema := []types.Schema{
		{
			Field: "allow_remote_vpc_dns_resolution",
			Ex:    "(true/false)",
			Doc: "(Optional) Allow a local VPC to resolve public DNS hostnames to private IP " +
				"\naddresses when queried from instances in the peer VPC. This is not supported " +
				"\nfor inter-region VPC peering.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "allow_classic_link_to_remote_vpc",
			Ex:    "(true/false)",
			Doc: "(Optional) Allow a local linked EC2-Classic instance to communicate with " +
				"\ninstances in a peer VPC. This enables an outbound communication from " +
				"\nthe local ClassicLink connection to the remote VPC.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "allow_vpc_to_remote_classic_link",
			Ex:    "(true/false)",
			Doc: "(Optional) Allow a local VPC to communicate with a linked EC2-Classic instance in a " +
				"\bpeer VPC. This enables an outbound communication from the local VPC to the remote " +
				"\nClassicLink connection.",
			Validator: utils.BoolValidator,
		},
	}

	resourceBlock["accepter"] = builder.PSOrder(types.ProvidePS(accepterRequesterSchema))

	color.Green("\nEnter requester:\n(Optional) - An optional configuration block that allows for [VPC Peering Connection] " +
		"\n(https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the " +
		"\nVPC that accepts the peering connection (a maximum of one)." +
		"\nThe requester block supports the following arguments:" +
		"\n1.allow_remote_vpc_dns_resolution\n2.allow_classic_link_to_remote_vpc\n3.allow_vpc_to_remote_classic_link")

	resourceBlock["requester"] = builder.PSOrder(types.ProvidePS(accepterRequesterSchema))

	color.Green("Enter timeout configuration:\n")

	timeoutSchema := []types.Schema{
		{
			Field: "create",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for creating a VPC endpoint",
		},
		{
			Field: "update",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for VPC endpoint modifications",
		},
		{
			Field: "delete",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for destroying VPC endpoints",
		},
	}

	resourceBlock["timeout"] = builder.PSOrder(types.ProvidePS(timeoutSchema))

	builder.ResourceBuilder("aws_vpc_peering_connection", blockName, resourceBlock)
}

func AWSVPCPeeringConnectionAccepterPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "vpc_peering_connection_id",
			Ex:    "vpc-pci-123",
			Doc:   "(Required) The VPC Peering Connection ID to manage.",
		},
		{
			Field:     "auto_accept",
			Ex:        "(true/false)",
			Doc:       "(Optional) Whether or not to accept the peering request. Defaults to false.",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpc_peering_connection_accepter", blockName, resourceBlock)
}

func AWSVPCPeeringConnectionOptionsPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "vpc_peering_connection_id",
			Ex:    "vpc-pci-123",
			Doc:   "(Required) The ID of the requester VPC peering connection.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like accepter/requester [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_vpc_peering_connection", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter accepter:\n(Optional) - An optional configuration block that allows for [VPC Peering Connection] " +
		"\n(https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the " +
		"\nVPC that accepts the peering connection (a maximum of one)." +
		"\n\nThe accepter block supports the following arguments:" +
		"\n1.allow_remote_vpc_dns_resolution\n2.allow_classic_link_to_remote_vpc\n3.allow_vpc_to_remote_classic_link")

	accepterRequesterSchema := []types.Schema{
		{
			Field: "allow_remote_vpc_dns_resolution",
			Ex:    "(true/false)",
			Doc: "(Optional) Allow a local VPC to resolve public DNS hostnames to " +
				"\nprivate IP addresses when queried from instances in the peer VPC.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "allow_classic_link_to_remote_vpc",
			Ex:    "(true/false)",
			Doc: "(Optional) Allow a local linked EC2-Classic instance to communicate " +
				"\nwith instances in a peer VPC. This enables an outbound communication " +
				"\nfrom the local ClassicLink connection to the remote VPC. This option " +
				"\nis not supported for inter-region VPC peering.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "allow_vpc_to_remote_classic_link",
			Ex:    "(true/false)",
			Doc: "(Optional) Allow a local VPC to communicate with a linked EC2-Classic " +
				"\ninstance in a peer VPC. This enables an outbound communication from " +
				"\nthe local VPC to the remote ClassicLink connection. This option is not " +
				"\nsupported for inter-region VPC peering.",
			Validator: utils.BoolValidator,
		},
	}

	resourceBlock["accepter"] = builder.PSOrder(types.ProvidePS(accepterRequesterSchema))

	color.Green("\nEnter requester:\n(Optional) - A optional configuration block that allows for [VPC Peering Connection] " +
		"\n(https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests " +
		"\nthe peering connection (a maximum of one)." +
		"\nThe requester block supports the following arguments:" +
		"\n1.allow_remote_vpc_dns_resolution\n2.allow_classic_link_to_remote_vpc\n3.allow_vpc_to_remote_classic_link")

	resourceBlock["requester"] = builder.PSOrder(types.ProvidePS(accepterRequesterSchema))

	builder.ResourceBuilder("aws_vpc_peering_connection_options", blockName, resourceBlock)
}

func AWSVPCConnectionPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	color.Yellow("\nOnly one of the transit_gateway_id and vpn_gateway_id is required\n")

	schema := []types.Schema{
		{
			Field: "customer_gateway_id",
			Ex:    "cg-123",
			Doc:   "(Required) The ID of the customer gateway.",
		},
		{
			Type:  "select",
			Field: "type",
			Doc:   "(Required) The type of VPN connection. The only type AWS supports at this time is \"ipsec.1\"",
			Items: []string{"ipsec.1"},
		},
		{
			Field: "transit_gateway_id",
			Ex:    "tg-123",
			Doc:   "(Optional) The ID of the EC2 Transit Gateway.",
		},
		{
			Field: "vpn_gateway_id",
			Ex:    "vgi-123",
			Doc:   "(Optional) The ID of the Virtual Private Gateway.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpn_connection", blockName, resourceBlock)
}

func AWSVPNConnectionRoutePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "destination_cidr_block",
			Ex:    "172.16.0.0/24",
			Doc:   "(Required) The CIDR block associated with the local subnet of the customer network.",
		},
		{
			Field: "vpn_connection_id",
			Ex:    "vpc-c123",
			Doc:   "(Required) The ID of the VPN connection.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpn_connection_route", blockName, resourceBlock)
}

func AWSVPNGatewayPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Optional) The VPC ID to create in.",
		},
		{
			Field: "availability_zone",
			Ex:    "az-123",
			Doc:   "(Optional) The Availability Zone for the virtual private gateway.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
		{
			Field: "amazon_side_asn",
			Ex:    "",
			Doc: "(Optional) The Autonomous System Number (ASN) for the Amazon side of the gateway. " +
				"\nIf you don't specify an ASN, the virtual private gateway is created with the default ASN.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpn_gateway", blockName, resourceBlock)
}

func AWSVPNGatewayAttachmentPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex: "vpc-123",
			Doc: "(Required) The ID of the VPC.",
		},
		{
			Field: "vpn_gateway_id",
			Ex: "vpn-gi-123",
			Doc: "(Required) The ID of the Virtual Private Gateway.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpn_gateway_attachment", blockName, resourceBlock)
}






















