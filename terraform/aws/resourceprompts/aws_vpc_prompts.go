package resourceprompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/utils"
	"github.com/manifoldco/promptui"
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

	schema := []types.Schema{
		{
			Field:     "bgp_asn",
			Ex:        "10000",
			Doc:       "(Required) The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).",
			Validator: utils.IntValidator,
		},
		{
			Field: "ip_address",
			Ex:    "172.83.124.10",
			Doc:   "(Required) The IP address of the gateway's Internet-routable external interface.",
		},
		{
			Type:  "select",
			Field: "type",
			Ex:    "ipsec.1",
			Doc:   "(Required) The type of customer gateway.",
			Items: []string{"ipsec.1"},
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) Tags to apply to the gateway.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "default_network_acl_id",
			Ex:    "",
			Doc:   "(Required) The Network ACL ID to manage. This attribute is exported from aws_vpc, or manually found via the AWS Console.",
		},
		{
			Field: "subnet_ids",
			Ex:    "",
			Doc:   "(Optional) A list of Subnet IDs to apply the ACL to. See the notes below on managing Subnets in the Default Network ACL",
		},
		{
			Field:     "tags",
			Ex:        "",
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
		builder.ResourceBuilder("aws_default_network_acl", blockName, resourceBlock)
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

	schema := []types.Schema{
		{
			Field: "default_route_table_id",
			Ex:    "",
			Doc:   "(Required) The ID of the Default Routing Table.",
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
		builder.ResourceBuilder("aws_default_route_table", blockName, resourceBlock)
		return
	}

	routeSchema := []types.Schema{
		{
			Field: "cidr_block",
			Ex:    "172.16.0.0/24",
			Doc:   "(Required) The CIDR block of the route.",
		},
		{
			Field: "ipv6_cidr_block",
			Ex:    "2001:db8:1234:1a00::/56",
			Doc:   "(Optional) The Ipv6 CIDR block of the route",
		},
		{
			Field: "egress_only_gateway_id",
			Ex:    "eg-123",
			Doc:   "(Optional) Identifier of a VPC Egress Only Internet Gateway.",
		},
		{
			Field: "gateway_id",
			Ex:    "",
			Doc:   "(Optional) Identifier of a VPC internet gateway or a virtual private gateway.",
		},
		{
			Field: "instance_id",
			Ex:    "",
			Doc:   "(Optional) Identifier of an EC2 instance.",
		},
		{
			Field: "nat_gateway_id",
			Ex:    "",
			Doc:   "(Optional) Identifier of a VPC NAT gateway.",
		},
		{
			Field: "network_interface_id",
			Ex:    "",
			Doc:   "(Optional) Identifier of an EC2 network interface.",
		},
		{
			Field: "transit_gateway_id",
			Ex:    "",
			Doc:   "(Optional) Identifier of an EC2 Transit Gateway.",
		},
		{
			Field: "vpc_endpoint_id",
			Ex:    "",
			Doc:   "(Optional) Identifier of a VPC Endpoint. This route must be removed prior to VPC Endpoint deletion.",
		},
		{
			Field: "vpc_peering_connection_id",
			Ex:    "",
			Doc:   "(Optional) Identifier of a VPC peering connection.",
		},
	}

	resourceBlock["route"] = builder.PSOrder(types.ProvidePS(routeSchema))

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

	schema := []types.Schema{
		{
			Field: "cidr_block",
			Ex:    "",
			Doc:   "(Required) The CIDR block for the VPC.",
		},
		{
			Field:     "enable_dns_support",
			Ex:        "(true/false)",
			Doc:       "(Optional) A boolean flag to enable/disable DNS support in the VPC. Defaults true.",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "enable_dns_hostnames",
			Ex:        "(true/false)",
			Doc:       "(Optional) A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "enable_classiclink",
			Ex:    "(true/false)",
			Doc: "(Optional) A boolean flag to enable/disable ClassicLink for the VPC. Only valid in regions and accounts that support EC2 Classic. Defaults false." +
				"\nCheckout https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "enable_classiclink_dns_support",
			Ex:        "(true/false)",
			Doc:       "(Optional) A boolean flag to enable/disable ClassicLink DNS Support for the VPC. Only valid in regions and accounts that support EC2 Classic.",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "assign_generated_ipv6_cidr_block",
			Ex:        "(true/false)",
			Doc:       "(Optional) Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is false.",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
		{
			Type:  "select",
			Field: "instance_tenancy",
			Doc:   "(Optional) A tenancy option for instances launched into the VPC. Default is default, which makes your instances shared on the host. Using either of the other options (dedicated or host) costs at least $2/hr.",
			Items: []string{"dedicated", "host"},
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "",
			Doc:   "(Optional, Forces new resource) The VPC ID. Note that changing the vpc_id will not restore any default security group rules that were modified, added, or removed. It will be left in its current state",
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
		builder.ResourceBuilder("aws_default_security_group", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter ingress:\n(Optional) Specifies an ingress rule." +
		"\n1.cidr_blocks\n2.description\n3.from_port\n4.ipv6_cidr_blocks\n5.prefix_list_ids\n6.protocol\n7.security_groups\n8.self\n9.to_port\n")

	ingressEgressSchema := []types.Schema{
		{
			Field: "cidr_blocks",
			Ex:    "[\"172.2.0.0/16\",\"173.0.0.1/24\"]",
			Doc:   "(Optional) List of CIDR blocks.",
		},
		{
			Field: "description",
			Ex:    "example description",
			Doc:   "(Optional) Description of this ingress/egress rule.",
		},
		{
			Field:     "from_port",
			Ex:        "443",
			Doc:       "(Required) The start port (or ICMP type number if protocol is \"icmp\" or \"icmpv6\")",
			Validator: utils.IntValidator,
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
			Field:     "protocol",
			Ex:        "-1",
			Doc:       "(Required) The protocol. If you select a protocol of \"-1\" (semantically equivalent to \"all\", which is not a valid value here), you must specify a \"from_port\" and \"to_port\" equal to 0. If not icmp, icmpv6, tcp, udp, or \"-1\" use the protocol number",
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
	}

	resourceBlock["ingress"] = builder.PSOrder(types.ProvidePS(ingressEgressSchema))

	color.Green("\nEnter egress:\n(Optional) Specifies an egress rule." +
		"\n1.cidr_blocks\n2.ipv6_cidr_blocks\n3.prefix_list_ids\n4.from_port\n5.protocol\n6.security_groups\n7.self\n8.to_port\n9.description\n")

	resourceBlock["egress"] = builder.PSOrder(types.ProvidePS(ingressEgressSchema))

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

	schema := []types.Schema{
		{
			Field: "availability_zone",
			Ex:    "",
			Doc:   "The AZ for the subnet.",
		},
		{
			Field: "map_public_ip_on_launch",
			Ex:    "",
			Doc:   "(Optional) Specify true to indicate that instances launched into the subnet should be assigned a public IP address.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field:     "enable_dns_support",
			Ex:        "(true/false)",
			Doc:       "(Optional) A boolean flag to enable/disable DNS support in the VPC. Defaults true.",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "enable_dns_hostnames",
			Ex:        "(true/false)",
			Doc:       "(Optional) A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "enable_classiclink",
			Ex:    "(true/false)",
			Doc: "(Optional) A boolean flag to enable/disable ClassicLink for the VPC. Only valid in regions " +
				"\nand accounts that support EC2 Classic. Defaults false." +
				"\nCheckout https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html",
			Validator: utils.BoolValidator,
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.BoolValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "netbios_name_servers",
			Ex:    "[\"s1\",\"s2\"]",
			Doc:   "(Optional) List of NETBIOS name servers.",
		},
		{
			Field: "netbios_node_type",
			Ex:    "",
			Doc: "(Optional) The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast " +
				"\nand multicast are not supported in their network. For more information about these node types," +
				"\ncheckout http://www.ietf.org/rfc/rfc2132.txt",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	color.Yellow("\nWhen you reference a Prefix List in a resource, the maximum number of entries " +
		"\nfor the prefix lists counts as the same number of rules or entries for the resource. " +
		"\nFor example, if you create a prefix list with a maximum of 20 entries and you reference " +
		"\nthat prefix list in a security group rule, this counts as 20 rules for the security group.")

	schema := []types.Schema{
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) The name of this resource. The name must not start with com.amazonaws.",
		},
		{
			Field: "address_family",
			Ex:    "",
			Doc:   "(Required, Forces new resource) The address family (IPv4 or IPv6) of this prefix list.",
		},
		{
			Field:     "max_retries",
			Ex:        "",
			Doc:       "(Required, Forces new resource) The address family (IPv4 or IPv6) of this prefix list.",
			Validator: utils.IntValidator,
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to this resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	entrySchema := []types.Schema{
		{
			Field: "cidr",
			Ex:    "172.0.0.1/24",
			Doc:   "(Required) The CIDR block of this entry.",
		},
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) Description of this entry.",
		},
	}

	resourceBlock["entry"] = builder.PSOrder(types.ProvidePS(entrySchema))

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

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The VPC ID to create in.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	color.Yellow("\nOne of eni_id, subnet_id, or vpc_id must be specified.")

	schema := []types.Schema{
		{
			Type:  "select",
			Field: "traffic_type",
			Doc:   "(Required) The type of traffic to capture",
			Items: []string{"ACCEPT", "REJECT", "ALL"},
		},
		{
			Field: "eni_id",
			Ex:    "",
			Doc:   "(Optional) Elastic Network Interface ID to attach to",
		},
		{
			Field: "iam_role_arn",
			Ex:    "",
			Doc:   "(Optional) The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group",
		},
		{
			Type:  "select",
			Field: "log_destination",
			Doc:   "(Optional) The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group",
			Items: []string{"cloud-watch-logs", "s3"},
		},
		{
			Field: "subnet_id",
			Ex:    "",
			Doc:   "(Optional) Subnet ID to attach to",
		},
		{
			Field: "vpc_id",
			Ex:    "",
			Doc:   "(Optional) VPC ID to attach to",
		},
		{
			Field: "log_format",
			Ex:    "",
			Doc:   "(Optional) The fields to include in the flow log record, in the order in which they should appear.",
		},
		{
			Type:  "select",
			Field: "max_aggregation_interval",
			Doc: "(Optional) The maximum interval(in seconds) of time during which a flow of " +
				"\npackets is captured and aggregated into a flow log record.",
			Items: []string{"60", "600"},
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) Key-value map of resource tags",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The VPC ID to create in.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) The VPC ID to create in.",
		},
		{
			Field: "route_table_id",
			Ex:    "",
			Doc:   "(Required) The ID of the Route Table to set as the new main route table for the target VPC",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	schema := []types.Schema{
		{
			Field: "allocation_id",
			Ex:    "",
			Doc:   "(Required) The Allocation ID of the Elastic IP address for the gateway.",
		},
		{
			Field: "subnet_id",
			Ex:    "",
			Doc:   "(Required) The Subnet ID of the subnet in which to place the gateway.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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
			Ex:    "vpc-123",
			Doc:   "(Required) The ID of the VPC.",
		},
		{
			Field: "vpn_gateway_id",
			Ex:    "vpn-gi-123",
			Doc:   "(Required) The ID of the Virtual Private Gateway.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpn_gateway_attachment", blockName, resourceBlock)
}

func AWSVPNGatewayRoutePropagationPrompt() {
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
			Field: "vpn_gateway_id",
			Ex:    "vpn-gi-123",
			Doc:   "The id of the aws_vpn_gateway to propagate routes from.",
		},
		{
			Field: "route_table_id",
			Ex:    "rti-123",
			Doc:   "The id of the aws_route_table to propagate routes into.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_vpn_gateway_route_propagation", blockName, resourceBlock)
}
