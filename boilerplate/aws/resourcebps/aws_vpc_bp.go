package resourcebps

import "github.com/fatih/color"

func AWSVPCBP() {
	color.Green("\nresource \"aws_vpc\" \"foo\" {\n  cidr_block       = \"10.0.0.0/16\"\n  instance_tenancy = \"default\"\n\n  tags = {\n    Name = \"main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc\n\n")
}

func AWSCustomerGatewayBP() {
	color.Green("\nresource \"aws_customer_gateway\" \"main\" {\n  bgp_asn    = 65000\n  ip_address = \"172.83.124.10\"\n  type       = \"ipsec.1\"\n\n  tags = {\n    Name = \"main-customer-gateway\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/customer_gateway\n\n")
}

func AWSDefaultNetworkACLBP() {
	color.Green("\nresource \"aws_default_network_acl\" \"default\" {\n  default_network_acl_id = aws_vpc.mainvpc.default_network_acl_id\n\n  ingress {\n    protocol   = -1\n    rule_no    = 100\n    action     = \"allow\"\n    cidr_block = aws_vpc.mainvpc.cidr_block\n    from_port  = 0\n    to_port    = 0\n  }\n\n  egress {\n    protocol   = -1\n    rule_no    = 100\n    action     = \"allow\"\n    cidr_block = \"0.0.0.0/0\"\n    from_port  = 0\n    to_port    = 0\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_network_acl\n\n")
}

func AWSDefaultRouteTableBP() {
	color.Green("\nresource \"aws_default_route_table\" \"r\" {\n  default_route_table_id = aws_vpc.foo.default_route_table_id\n\n  route {\n    # ...\n  }\n\n  tags = {\n    Name = \"default table\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_route_table\n\n")
}

func AWSDefaultSecurityGroupBP() {
	color.Green("\nresource \"aws_vpc\" \"mainvpc\" {\n  cidr_block = \"10.1.0.0/16\"\n}\n\nresource \"aws_default_security_group\" \"default\" {\n  vpc_id = aws_vpc.mainvpc.id\n\n  ingress {\n    protocol  = -1\n    self      = true\n    from_port = 0\n    to_port   = 0\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_security_group\n\n")
}

func AWSDefaultSubnetBP() {
	color.Green("\nresource \"aws_default_subnet\" \"default_az1\" {\n  availability_zone = \"us-west-2a\"\n\n  tags = {\n    Name = \"Default subnet for us-west-2a\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_subnet\n\n")
}

func AWSDefaultVPCBP() {
	color.Green("\nresource \"aws_default_vpc\" \"default\" {\n  tags = {\n    Name = \"Default VPC\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_vpc\n\n")
}

func AWSDefaultVPCDHCPOptionsBP() {
	color.Green("\nresource \"aws_default_vpc_dhcp_options\" \"default\" {\n  tags = {\n    Name = \"Default DHCP Option Set\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_vpc_dhcp_options\n\n")
}

func AWSEC2ManagedPrefixListBP() {
	color.Green("\nresource \"aws_ec2_managed_prefix_list\" \"example\" {\n  name           = \"All VPC CIDR-s\"\n  address_family = \"IPv4\"\n  max_entries    = 5\n\n  entry {\n    cidr        = aws_vpc.example.cidr_block\n    description = \"Primary\"\n  }\n\n  entry {\n    cidr        = aws_vpc_ipv4_cidr_block_association.example.cidr_block\n    description = \"Secondary\"\n  }\n\n  tags = {\n    Env = \"live\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_managed_prefix_list\n\n")
}

func AWSEgressOnlyInternetGatewayBP() {
	color.Green("\nresource \"aws_vpc\" \"example\" {\n  cidr_block                       = \"10.1.0.0/16\"\n  assign_generated_ipv6_cidr_block = true\n}\n\nresource \"aws_egress_only_internet_gateway\" \"example\" {\n  vpc_id = aws_vpc.example.id\n\n  tags = {\n    Name = \"main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/egress_only_internet_gateway\n\n")
}

func AWSFlowLogBP() {
	color.Green("\nresource \"aws_flow_log\" \"example\" {\n  iam_role_arn    = aws_iam_role.example.arn\n  log_destination = aws_cloudwatch_log_group.example.arn\n  traffic_type    = \"ALL\"\n  vpc_id          = aws_vpc.example.id\n}\n\nresource \"aws_cloudwatch_log_group\" \"example\" {\n  name = \"example\"\n}\n\nresource \"aws_iam_role\" \"example\" {\n  name = \"example\"\n\n  assume_role_policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"\",\n      \"Effect\": \"Allow\",\n      \"Principal\": {\n        \"Service\": \"vpc-flow-logs.amazonaws.com\"\n      },\n      \"Action\": \"sts:AssumeRole\"\n    }\n  ]\n}\nEOF\n}\n\nresource \"aws_iam_role_policy\" \"example\" {\n  name = \"example\"\n  role = aws_iam_role.example.id\n\n  policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Action\": [\n        \"logs:CreateLogGroup\",\n        \"logs:CreateLogStream\",\n        \"logs:PutLogEvents\",\n        \"logs:DescribeLogGroups\",\n        \"logs:DescribeLogStreams\"\n      ],\n      \"Effect\": \"Allow\",\n      \"Resource\": \"*\"\n    }\n  ]\n}\nEOF\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/flow_log\n\n")
}

func AWSInternetGatewayBP() {
	color.Green("\nresource \"aws_internet_gateway\" \"gw\" {\n  vpc_id = aws_vpc.main.id\n\n  tags = {\n    Name = \"main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway\n\n")
}

func AWSMainRouteTableAssociationBP() {
	color.Green("\nresource \"aws_main_route_table_association\" \"a\" {\n  vpc_id         = aws_vpc.foo.id\n  route_table_id = aws_route_table.bar.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/main_route_table_association\n\n")
}

func AWSNatGatewayBP() {
	color.Green("\nresource \"aws_nat_gateway\" \"gw\" {\n  allocation_id = aws_eip.nat.id\n  subnet_id     = aws_subnet.example.id\n\n  tags = {\n    Name = \"gw NAT\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/nat_gateway\n\n")
}

func AWSNetworkACLBP() {
	color.Green("\nresource \"aws_network_acl\" \"main\" {\n  vpc_id = aws_vpc.main.id\n\n  egress {\n    protocol   = \"tcp\"\n    rule_no    = 200\n    action     = \"allow\"\n    cidr_block = \"10.3.0.0/18\"\n    from_port  = 443\n    to_port    = 443\n  }\n\n  ingress {\n    protocol   = \"tcp\"\n    rule_no    = 100\n    action     = \"allow\"\n    cidr_block = \"10.3.0.0/18\"\n    from_port  = 80\n    to_port    = 80\n  }\n\n  tags = {\n    Name = \"main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_acl\n\n")
}

func AWSNetworkACLRuleBP() {
	color.Green("\nresource \"aws_network_acl\" \"bar\" {\n  vpc_id = aws_vpc.foo.id\n}\n\nresource \"aws_network_acl_rule\" \"bar\" {\n  network_acl_id = aws_network_acl.bar.id\n  rule_number    = 200\n  egress         = false\n  protocol       = \"tcp\"\n  rule_action    = \"allow\"\n  cidr_block     = aws_vpc.foo.cidr_block\n  from_port      = 22\n  to_port        = 22\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_acl_rule\n\n")
}

func AWSNetworkInterfaceBP() {
	color.Green("\nresource \"aws_network_interface\" \"test\" {\n  subnet_id       = aws_subnet.public_a.id\n  private_ips     = [\"10.0.0.50\"]\n  security_groups = [aws_security_group.web.id]\n\n  attachment {\n    instance     = aws_instance.test.id\n    device_index = 1\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface\n\n")
}

func AWSNetworkInterfaceAttachmentBP() {
	color.Green("\nresource \"aws_network_interface_attachment\" \"test\" {\n  instance_id          = aws_instance.test.id\n  network_interface_id = aws_network_interface.test.id\n  device_index         = 0\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface_attachment\n\n")
}

func AWSNetworkInterfaceSGAttachmentBP() {
	color.Green("\ndata \"aws_ami\" \"ami\" {\n  most_recent = true\n\n  filter {\n    name   = \"name\"\n    values = [\"amzn-ami-hvm-*\"]\n  }\n\n  owners = [\"amazon\"]\n}\n\nresource \"aws_instance\" \"instance\" {\n  instance_type = \"t2.micro\"\n  ami           = data.aws_ami.ami.id\n\n  tags = {\n    \"type\" = \"terraform-test-instance\"\n  }\n}\n\nresource \"aws_security_group\" \"sg\" {\n  tags = {\n    \"type\" = \"terraform-test-security-group\"\n  }\n}\n\nresource \"aws_network_interface_sg_attachment\" \"sg_attachment\" {\n  security_group_id    = aws_security_group.sg.id\n  network_interface_id = aws_instance.instance.primary_network_interface_id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface_sg_attachment\n\n")
}

func AWSRouteBP() {
	color.Green("\nresource \"aws_route\" \"r\" {\n  route_table_id            = \"rtb-4fbb3ac4\"\n  destination_cidr_block    = \"10.0.1.0/22\"\n  vpc_peering_connection_id = \"pcx-45ff3dc1\"\n  depends_on                = [aws_route_table.testing]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route\n\n")
}

func AWSRouteTableBP() {
	color.Green("\nresource \"aws_route_table\" \"r\" {\n  vpc_id = aws_vpc.default.id\n\n  route {\n    cidr_block = \"10.0.1.0/24\"\n    gateway_id = aws_internet_gateway.main.id\n  }\n\n  route {\n    ipv6_cidr_block        = \"::/0\"\n    egress_only_gateway_id = aws_egress_only_internet_gateway.foo.id\n  }\n\n  tags = {\n    Name = \"main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table\n\n")
}

func AWSRouteTableAssociationBP() {
	color.Green("\nresource \"aws_route_table_association\" \"a\" {\n  subnet_id      = aws_subnet.foo.id\n  route_table_id = aws_route_table.bar.id\n}\n\nresource \"aws_route_table_association\" \"b\" {\n  gateway_id     = aws_internet_gateway.foo.id\n  route_table_id = aws_route_table.bar.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route_table_association\n\n")
}

func AWSSecurityGroupBP() {
	color.Green("\nresource \"aws_security_group\" \"allow_tls\" {\n  name        = \"allow_tls\"\n  description = \"Allow TLS inbound traffic\"\n  vpc_id      = aws_vpc.main.id\n\n  ingress {\n    description = \"TLS from VPC\"\n    from_port   = 443\n    to_port     = 443\n    protocol    = \"tcp\"\n    cidr_blocks = [aws_vpc.main.cidr_block]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name = \"allow_tls\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group\n\n")
}

func AWSSecurityGroupRuleBP() {
	color.Green("\nresource \"aws_security_group_rule\" \"example\" {\n  type              = \"ingress\"\n  from_port         = 0\n  to_port           = 65535\n  protocol          = \"tcp\"\n  cidr_blocks       = [aws_vpc.example.cidr_block]\n  security_group_id = \"sg-123456\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule\n\n")
}

func AWSSubnetBP() {
	color.Green("\nresource \"aws_subnet\" \"main\" {\n  vpc_id     = aws_vpc.main.id\n  cidr_block = \"10.0.1.0/24\"\n\n  tags = {\n    Name = \"Main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet\n\n")
}

func AWSVPCDHCPOptionsBP() {
	color.Green("\nresource \"aws_vpc_dhcp_options\" \"foo\" {\n  domain_name          = \"service.consul\"\n  domain_name_servers  = [\"127.0.0.1\", \"10.0.0.2\"]\n  ntp_servers          = [\"127.0.0.1\"]\n  netbios_name_servers = [\"127.0.0.1\"]\n  netbios_node_type    = 2\n\n  tags = {\n    Name = \"foo-name\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_dhcp_options\n\n")
}

func AWSVPCDHCPOptionsAssociationBP() {
	color.Green("\nresource \"aws_vpc_dhcp_options_association\" \"dns_resolver\" {\n  vpc_id          = aws_vpc.foo.id\n  dhcp_options_id = aws_vpc_dhcp_options.foo.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_dhcp_options_association\n\n")
}

func AWSVPCEndpointBP() {
	color.Green("\nresource \"aws_vpc_endpoint\" \"s3\" {\n  vpc_id       = aws_vpc.main.id\n  service_name = \"com.amazonaws.us-west-2.s3\"\n\n  tags = {\n    Environment = \"test\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint\n\n")
}

func AWSVPCEndpointConnectionNotificationBP() {
	color.Green("\nresource \"aws_sns_topic\" \"topic\" {\n  name = \"vpce-notification-topic\"\n\n  policy = <<POLICY\n{\n    \"Version\":\"2012-10-17\",\n    \"Statement\":[{\n        \"Effect\": \"Allow\",\n        \"Principal\": {\n            \"Service\": \"vpce.amazonaws.com\"\n        },\n        \"Action\": \"SNS:Publish\",\n        \"Resource\": \"arn:aws:sns:*:*:vpce-notification-topic\"\n    }]\n}\nPOLICY\n}\n\nresource \"aws_vpc_endpoint_service\" \"foo\" {\n  acceptance_required        = false\n  network_load_balancer_arns = [aws_lb.test.arn]\n}\n\nresource \"aws_vpc_endpoint_connection_notification\" \"foo\" {\n  vpc_endpoint_service_id     = aws_vpc_endpoint_service.foo.id\n  connection_notification_arn = aws_sns_topic.topic.arn\n  connection_events           = [\"Accept\", \"Reject\"]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_connection_notification\n\n")
}

func AWSVPCEndpointRouteTableAssociationBP() {
	color.Green("\nresource \"aws_vpc_endpoint_route_table_association\" \"example\" {\n  route_table_id  = aws_route_table.example.id\n  vpc_endpoint_id = aws_vpc_endpoint.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_route_table_association\n\n")
}

func AWSVPCEndpointServiceBP() {
	color.Green("\nresource \"aws_vpc_endpoint_service\" \"example\" {\n  acceptance_required        = false\n  network_load_balancer_arns = [aws_lb.example.arn]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_service\n\n")
}

func AWSVPCEndpointServiceAllowedPrincipalBP() {
	color.Green("\ndata \"aws_caller_identity\" \"current\" {}\n\nresource \"aws_vpc_endpoint_service_allowed_principal\" \"allow_me_to_foo\" {\n  vpc_endpoint_service_id = aws_vpc_endpoint_service.foo.id\n  principal_arn           = data.aws_caller_identity.current.arn\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_service_allowed_principal\n\n")
}

func AWSVPCEndpointSubnetAssociationBP() {
	color.Green("\nresource \"aws_vpc_endpoint_subnet_association\" \"sn_ec2\" {\n  vpc_endpoint_id = aws_vpc_endpoint.ec2.id\n  subnet_id       = aws_subnet.sn.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint_subnet_association\n\n")
}

func AWSVPCIPV4CIDRBlockAssociationBP() {
	color.Green("\nresource \"aws_vpc\" \"main\" {\n  cidr_block = \"10.0.0.0/16\"\n}\n\nresource \"aws_vpc_ipv4_cidr_block_association\" \"secondary_cidr\" {\n  vpc_id     = aws_vpc.main.id\n  cidr_block = \"172.2.0.0/16\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_ipv4_cidr_block_association\n\n")
}

func AWSVPCPeeringConnectionBP() {
	color.Green("\nresource \"aws_vpc_peering_connection\" \"foo\" {\n  peer_owner_id = var.peer_owner_id\n  peer_vpc_id   = aws_vpc.bar.id\n  vpc_id        = aws_vpc.foo.id\n\n  accepter {\n    allow_remote_vpc_dns_resolution = true\n  }\n\n  requester {\n    allow_remote_vpc_dns_resolution = true\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_peering_connection\n\n")
}

func AWSVPCPeeringConnectionAccepterBP() {
	color.Green("\nprovider \"aws\" {\n  region = \"us-east-1\"\n\n  # Requester's credentials.\n}\n\nprovider \"aws\" {\n  alias  = \"peer\"\n  region = \"us-west-2\"\n\n  # Accepter's credentials.\n}\n\nresource \"aws_vpc\" \"main\" {\n  cidr_block = \"10.0.0.0/16\"\n}\n\nresource \"aws_vpc\" \"peer\" {\n  provider   = aws.peer\n  cidr_block = \"10.1.0.0/16\"\n}\n\ndata \"aws_caller_identity\" \"peer\" {\n  provider = aws.peer\n}\n\n# Requester's side of the connection.\nresource \"aws_vpc_peering_connection\" \"peer\" {\n  vpc_id        = aws_vpc.main.id\n  peer_vpc_id   = aws_vpc.peer.id\n  peer_owner_id = data.aws_caller_identity.peer.account_id\n  peer_region   = \"us-west-2\"\n  auto_accept   = false\n\n  tags = {\n    Side = \"Requester\"\n  }\n}\n\n# Accepter's side of the connection.\nresource \"aws_vpc_peering_connection_accepter\" \"peer\" {\n  provider                  = aws.peer\n  vpc_peering_connection_id = aws_vpc_peering_connection.peer.id\n  auto_accept               = true\n\n  tags = {\n    Side = \"Accepter\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_peering_connection_accepter\n\n")
}

func AWSVPCPeeringConnectionOptionsBP() {
	color.Green("\nresource \"aws_vpc\" \"foo\" {\n  cidr_block = \"10.0.0.0/16\"\n}\n\nresource \"aws_vpc\" \"bar\" {\n  cidr_block = \"10.1.0.0/16\"\n}\n\nresource \"aws_vpc_peering_connection\" \"foo\" {\n  vpc_id      = aws_vpc.foo.id\n  peer_vpc_id = aws_vpc.bar.id\n  auto_accept = true\n}\n\nresource \"aws_vpc_peering_connection_options\" \"foo\" {\n  vpc_peering_connection_id = aws_vpc_peering_connection.foo.id\n\n  accepter {\n    allow_remote_vpc_dns_resolution = true\n  }\n\n  requester {\n    allow_vpc_to_remote_classic_link = true\n    allow_classic_link_to_remote_vpc = true\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_peering_connection_options\n\n")
}

func AWSVPNConnectionBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway\" \"example\" {}\n\nresource \"aws_customer_gateway\" \"example\" {\n  bgp_asn    = 65000\n  ip_address = \"172.0.0.1\"\n  type       = \"ipsec.1\"\n}\n\nresource \"aws_vpn_connection\" \"example\" {\n  customer_gateway_id = aws_customer_gateway.example.id\n  transit_gateway_id  = aws_ec2_transit_gateway.example.id\n  type                = aws_customer_gateway.example.type\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpn_connection\n\n")
}

func AWSVPNConnectionRouteBP() {
	color.Green("\nresource \"aws_vpc\" \"vpc\" {\n  cidr_block = \"10.0.0.0/16\"\n}\n\nresource \"aws_vpn_gateway\" \"vpn_gateway\" {\n  vpc_id = aws_vpc.vpc.id\n}\n\nresource \"aws_customer_gateway\" \"customer_gateway\" {\n  bgp_asn    = 65000\n  ip_address = \"172.0.0.1\"\n  type       = \"ipsec.1\"\n}\n\nresource \"aws_vpn_connection\" \"main\" {\n  vpn_gateway_id      = aws_vpn_gateway.vpn_gateway.id\n  customer_gateway_id = aws_customer_gateway.customer_gateway.id\n  type                = \"ipsec.1\"\n  static_routes_only  = true\n}\n\nresource \"aws_vpn_connection_route\" \"office\" {\n  destination_cidr_block = \"192.168.10.0/24\"\n  vpn_connection_id      = aws_vpn_connection.main.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpn_connection_route\n\n")
}

func AWSVPNGatewayBP() {
	color.Green("\nresource \"aws_vpn_gateway\" \"vpn_gw\" {\n  vpc_id = aws_vpc.main.id\n\n  tags = {\n    Name = \"main\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpn_gateway\n\n")
}

func AWSVPNGatewayAttachmentBP() {
	color.Green("\nresource \"aws_vpc\" \"network\" {\n  cidr_block = \"10.0.0.0/16\"\n}\n\nresource \"aws_vpn_gateway\" \"vpn\" {\n  tags = {\n    Name = \"example-vpn-gateway\"\n  }\n}\n\nresource \"aws_vpn_gateway_attachment\" \"vpn_attachment\" {\n  vpc_id         = aws_vpc.network.id\n  vpn_gateway_id = aws_vpn_gateway.vpn.id\n}\n\n")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpn_gateway_attachment\n\n")
}

func AWSVPNGatewayRoutePropagationBP() {
	color.Green("\nresource \"aws_vpn_gateway_route_propagation\" \"example\" {\n  vpn_gateway_id = aws_vpn_gateway.example.id\n  route_table_id = aws_route_table.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpn_gateway_route_propagation\n\n")
}
