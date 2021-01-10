package resource_bps

import (
	"github.com/fatih/color"
)

func AWSEC2InstanceBP() {
	color.Green("\nresource \"aws_instance\" \"foo\" {\n  ami           = data.aws_ami.ubuntu.id\n  instance_type = \"t3.micro\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance\n\n")
}

func AWSAMIBP() {
	color.Green("\n# Create an AMI that will start a machine whose root device is backed by\n# an EBS volume populated from a snapshot. It is assumed that such a snapshot\n# already exists with the id \"snap-xxxxxxxx\".\nresource \"aws_ami\" \"example\" {\n  name                = \"terraform-example\"\n  virtualization_type = \"hvm\"\n  root_device_name    = \"/dev/xvda\"\n\n  ebs_block_device {\n    device_name = \"/dev/xvda\"\n    snapshot_id = \"snap-xxxxxxxx\"\n    volume_size = 8\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami\n\n")
}

func AWSAMICopyBP() {
	color.Green("\nresource \"aws_ami_copy\" \"example\" {\n  name              = \"terraform-example\"\n  description       = \"A copy of ami-xxxxxxxx\"\n  source_ami_id     = \"ami-xxxxxxxx\"\n  source_ami_region = \"us-west-1\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami_copy\n\n")
}

func AWSAMIFromInstanceBP() {
	color.Green("\nresource \"aws_ami_from_instance\" \"example\" {\n  name               = \"terraform-example\"\n  source_instance_id = \"i-xxxxxxxx\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami_from_instance\n\n")
}

func AWSAMIFromLaunchPermissionBP() {
	color.Green("\nresource \"aws_ami_launch_permission\" \"example\" {\n  image_id   = \"ami-12345678\"\n  account_id = \"123456789012\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami_launch_permission\n\n")
}

func AWSEBSDefaultKMSKeyBP() {
	color.Green("\nresource \"aws_ebs_default_kms_key\" \"example\" {\n  key_arn = aws_kms_key.example.arn\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_default_kms_key\n\n")
}

func AWSEBSEncryptionByDefaultBP() {
	color.Green("\nresource \"aws_ebs_encryption_by_default\" \"example\" {\n  enabled = true\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_encryption_by_default\n\n")
}

func AWSEBSSnapshotBP() {
	color.Green("\nresource \"aws_ebs_snapshot\" \"example_snapshot\" {\n  volume_id = aws_ebs_volume.example.id\n\n  tags = {\n    Name = \"HelloWorld_snap\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_snapshot\n\n")
}

func AWSEBSSnapshotCopyBP() {
	color.Green("\nresource \"aws_ebs_snapshot_copy\" \"example_copy\" {\n  source_snapshot_id = aws_ebs_snapshot.example_snapshot.id\n  source_region      = \"us-west-2\"\n\n  tags = {\n    Name = \"HelloWorld_copy_snap\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_snapshot_copy\n\n")
}

func AWSEBSVolumeBP() {
	color.Green("\nresource \"aws_ebs_volume\" \"example\" {\n  availability_zone = \"us-west-2a\"\n  size              = 40\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_volume\n\n")
}

func AWSEC2AvailabilityZoneGroupBP() {
	color.Green("\nresource \"aws_ec2_availability_zone_group\" \"example\" {\n  group_name    = \"us-west-2-lax-1\"\n  opt_in_status = \"opted-in\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_availability_zone_group\n\n")
}

func AWSEC2CapacityReservationBP() {
	color.Green("\nresource \"aws_ec2_capacity_reservation\" \"default\" {\n  instance_type     = \"t2.micro\"\n  instance_platform = \"Linux/UNIX\"\n  availability_zone = \"eu-west-1a\"\n  instance_count    = 1\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_capacity_reservation\n\n")
}

func AWSEC2CarrierGatewayPrompt() {
	color.Green("\nresource \"aws_ec2_carrier_gateway\" \"example\" {\n  vpc_id = aws_vpc.example.id\n\n  tags = {\n    Name = \"example-carrier-gateway\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_carrier_gateway\n\n")
}

func AWSEC2ClientVPNAuthorizationRuleBP() {
	color.Green("\nresource \"aws_ec2_client_vpn_authorization_rule\" \"example\" {\n  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.example.id\n  target_network_cidr    = aws_subnet.example.cidr_block\n  authorize_all_groups   = true\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_client_vpn_authorization_rule\n\n")
}

func AWSEC2ClientVPNEndpointBP() {
	color.Green("\nresource \"aws_ec2_client_vpn_endpoint\" \"example\" {\n  description            = \"terraform-clientvpn-example\"\n  server_certificate_arn = aws_acm_certificate.cert.arn\n  client_cidr_block      = \"10.0.0.0/16\"\n\n  authentication_options {\n    type                       = \"certificate-authentication\"\n    root_certificate_chain_arn = aws_acm_certificate.root_cert.arn\n  }\n\n  connection_log_options {\n    enabled               = true\n    cloudwatch_log_group  = aws_cloudwatch_log_group.lg.name\n    cloudwatch_log_stream = aws_cloudwatch_log_stream.ls.name\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_client_vpn_endpoint\n\n")
}

func AWSEC2ClientVPNNetworkAssociationBP() {
	color.Green("\nresource \"aws_ec2_client_vpn_network_association\" \"example\" {\n  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.example.id\n  subnet_id              = aws_subnet.example.id\n  security_groups        = [aws_security_group.example1.id, aws_security_group.example2.id]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_client_vpn_network_association\n\n")
}

func AWSEC2ClientVPNRouteBP() {
	color.Green("\nresource \"aws_ec2_client_vpn_route\" \"example\" {\n  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.example.id\n  destination_cidr_block = \"0.0.0.0/0\"\n  target_vpc_subnet_id   = aws_ec2_client_vpn_network_association.example.subnet_id\n}\n\nresource \"aws_ec2_client_vpn_network_association\" \"example\" {\n  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.example.id\n  subnet_id              = aws_subnet.example.id\n}\n\nresource \"aws_ec2_client_vpn_endpoint\" \"example\" {\n  description            = \"Example Client VPN endpoint\"\n  server_certificate_arn = aws_acm_certificate.example.arn\n  client_cidr_block      = \"10.0.0.0/16\"\n\n  authentication_options {\n    type                       = \"certificate-authentication\"\n    root_certificate_chain_arn = aws_acm_certificate.example.arn\n  }\n\n  connection_log_options {\n    enabled = false\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_client_vpn_route\n\n")
}

func AWSEC2FleetBP() {
	color.Green("\nresource \"aws_ec2_fleet\" \"example\" {\n  launch_template_config {\n    launch_template_specification {\n      launch_template_id = aws_launch_template.example.id\n      version            = aws_launch_template.example.latest_version\n    }\n  }\n\n  target_capacity_specification {\n    default_target_capacity_type = \"spot\"\n    total_target_capacity        = 5\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_fleet\n\n")
}

func AWSEC2LocalGatewayRouteBP() {
	color.Green("\nresource \"aws_ec2_local_gateway_route\" \"example\" {\n  destination_cidr_block                   = \"172.16.0.0/16\"\n  local_gateway_route_table_id             = data.aws_ec2_local_gateway_route_table.example.id\n  local_gateway_virtual_interface_group_id = data.aws_ec2_local_gateway_virtual_interface_group.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_local_gateway_route\n\n")
}

func AWSEC2LocalGatewayRouteTableVPCAssociationBP() {
	color.Green("\ndata \"aws_ec2_local_gateway_route_table\" \"example\" {\n  outpost_arn = \"arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef\"\n}\n\nresource \"aws_vpc\" \"example\" {\n  cidr_block = \"10.0.0.0/16\"\n}\n\nresource \"aws_ec2_local_gateway_route_table_vpc_association\" \"example\" {\n  local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.example.id\n  vpc_id                       = aws_vpc.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_local_gateway_route_table_vpc_association\n\n")
}

func AWSEC2TagBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway\" \"example\" {}\n\nresource \"aws_customer_gateway\" \"example\" {\n  bgp_asn    = 65000\n  ip_address = \"172.0.0.1\"\n  type       = \"ipsec.1\"\n}\n\nresource \"aws_vpn_connection\" \"example\" {\n  customer_gateway_id = aws_customer_gateway.example.id\n  transit_gateway_id  = aws_ec2_transit_gateway.example.id\n  type                = aws_customer_gateway.example.type\n}\n\nresource \"aws_ec2_tag\" \"example\" {\n  resource_id = aws_vpn_connection.example.transit_gateway_attachment_id\n  key         = \"Name\"\n  value       = \"Hello World\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_tag\n\n")
}

func AWSEC2TrafficMirrorFilterBP() {
	color.Green("\nresource \"aws_ec2_traffic_mirror_filter\" \"foo\" {\n  description      = \"traffic mirror filter - terraform example\"\n  network_services = [\"amazon-dns\"]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_traffic_mirror_filter\n\n")
}

func AWSEC2TrafficMirrorFilterRuleBP() {
	color.Green("\nresource \"aws_ec2_traffic_mirror_filter\" \"filter\" {\n  description      = \"traffic mirror filter - terraform example\"\n  network_services = [\"amazon-dns\"]\n}\n\nresource \"aws_ec2_traffic_mirror_filter_rule\" \"ruleout\" {\n  description              = \"test rule\"\n  traffic_mirror_filter_id = aws_ec2_traffic_mirror_filter.filter.id\n  destination_cidr_block   = \"10.0.0.0/8\"\n  source_cidr_block        = \"10.0.0.0/8\"\n  rule_number              = 1\n  rule_action              = \"accept\"\n  traffic_direction        = \"egress\"\n}\n\nresource \"aws_ec2_traffic_mirror_filter_rule\" \"rulein\" {\n  description              = \"test rule\"\n  traffic_mirror_filter_id = aws_ec2_traffic_mirror_filter.filter.id\n  destination_cidr_block   = \"10.0.0.0/8\"\n  source_cidr_block        = \"10.0.0.0/8\"\n  rule_number              = 1\n  rule_action              = \"accept\"\n  traffic_direction        = \"ingress\"\n  protocol                 = 6\n\n  destination_port_range {\n    from_port = 22\n    to_port   = 53\n  }\n\n  source_port_range {\n    from_port = 0\n    to_port   = 10\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_traffic_mirror_filter_rule\n\n")
}

func AWSEC2TrafficMirrorSessionBP() {
	color.Green("\nresource \"aws_ec2_traffic_mirror_filter\" \"filter\" {\n  description      = \"traffic mirror filter - terraform example\"\n  network_services = [\"amazon-dns\"]\n}\n\nresource \"aws_ec2_traffic_mirror_target\" \"target\" {\n  network_load_balancer_arn = aws_lb.lb.arn\n}\n\nresource \"aws_ec2_traffic_mirror_session\" \"session\" {\n  description              = \"traffic mirror session - terraform example\"\n  network_interface_id     = aws_instance.test.primary_network_interface_id\n  traffic_mirror_filter_id = aws_ec2_traffic_mirror_filter.filter.id\n  traffic_mirror_target_id = aws_ec2_traffic_mirror_target.target.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_traffic_mirror_session\n\n")
}

func AWSEC2TrafficMirrorTargetBP() {
	color.Green("\nresource \"aws_ec2_traffic_mirror_target\" \"nlb\" {\n  description               = \"NLB target\"\n  network_load_balancer_arn = aws_lb.lb.arn\n}\n\nresource \"aws_ec2_traffic_mirror_target\" \"eni\" {\n  description          = \"ENI target\"\n  network_interface_id = aws_instance.test.primary_network_interface_id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_traffic_mirror_target\n\n")
}

func AWSEC2TransitGatewayBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway\" \"example\" {\n  description = \"example\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway\n\n")
}

func AWSEC2TransitGatewayPeeringAttachmentBP() {
	color.Green("\nprovider \"aws\" {\n  alias  = \"local\"\n  region = \"us-east-1\"\n}\n\nprovider \"aws\" {\n  alias  = \"peer\"\n  region = \"us-west-2\"\n}\n\ndata \"aws_region\" \"peer\" {\n  provider = aws.peer\n}\n\nresource \"aws_ec2_transit_gateway\" \"local\" {\n  provider = aws.local\n\n  tags = {\n    Name = \"Local TGW\"\n  }\n}\n\nresource \"aws_ec2_transit_gateway\" \"peer\" {\n  provider = aws.peer\n\n  tags = {\n    Name = \"Peer TGW\"\n  }\n}\n\nresource \"aws_ec2_transit_gateway_peering_attachment\" \"example\" {\n  peer_account_id         = aws_ec2_transit_gateway.peer.owner_id\n  peer_region             = data.aws_region.peer.name\n  peer_transit_gateway_id = aws_ec2_transit_gateway.peer.id\n  transit_gateway_id      = aws_ec2_transit_gateway.local.id\n\n  tags = {\n    Name = \"TGW Peering Requestor\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_peering_attachment\n\n")
}

func AWSEC2TransitGatewayPeeringAttachmentAccepterBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway_peering_attachment_accepter\" \"example\" {\n  transit_gateway_attachment_id = aws_ec2_transit_gateway_peering_attachment.example.id\n\n  tags = {\n    Name = \"Example cross-account attachment\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_peering_attachment_accepter\n\n")
}

func AWSEC2TransitGatewayRouteBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway_route\" \"example\" {\n  destination_cidr_block         = \"0.0.0.0/0\"\n  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.example.id\n  transit_gateway_route_table_id = aws_ec2_transit_gateway.example.association_default_route_table_id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_route\n\n")
}

func AWSEC2TransitGatewayRouteTableBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway_route_table\" \"example\" {\n  transit_gateway_id = aws_ec2_transit_gateway.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_route_table\n\n")
}

func AWSEC2TransitGatewayRouteTableAssociationBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway_route_table_association\" \"example\" {\n  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.example.id\n  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_route_table_association\n\n")
}

func AWSEC2TransitGatewayRouteTablePropagationBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway_route_table_propagation\" \"example\" {\n  transit_gateway_attachment_id  = aws_ec2_transit_gateway_vpc_attachment.example.id\n  transit_gateway_route_table_id = aws_ec2_transit_gateway_route_table.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_route_table_propagation\n\n")
}

func AWSEC2TransitGatewayVPCAttachmentBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway_vpc_attachment\" \"example\" {\n  subnet_ids         = [aws_subnet.example.id]\n  transit_gateway_id = aws_ec2_transit_gateway.example.id\n  vpc_id             = aws_vpc.example.id\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_vpc_attachment\n\n")
}

func AWSEC2TransitGatewayVPCAttachmentAccepterBP() {
	color.Green("\nresource \"aws_ec2_transit_gateway_vpc_attachment_accepter\" \"example\" {\n  transit_gateway_attachment_id = aws_ec2_transit_gateway_vpc_attachment.example.id\n\n  tags = {\n    Name = \"Example cross-account attachment\"\n  }\n}")
	color.Yellow("\bCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_transit_gateway_vpc_attachment_accepter\n\n")
}

func AWSEIPBP() {
	color.Green("\nresource \"aws_eip\" \"lb\" {\n  instance = aws_instance.web.id\n  vpc      = true\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip\n\n")
}

func AWSEIPAssociationBP() {
	color.Green("\nresource \"aws_eip_association\" \"eip_assoc\" {\n  instance_id   = aws_instance.web.id\n  allocation_id = aws_eip.example.id\n}\n\nresource \"aws_instance\" \"web\" {\n  ami               = \"ami-21f78e11\"\n  availability_zone = \"us-west-2a\"\n  instance_type     = \"t2.micro\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}\n\nresource \"aws_eip\" \"example\" {\n  vpc = true\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip_association\n\n")
}

func AWSKeyPairBP() {
	color.Green("\nresource \"aws_key_pair\" \"deployer\" {\n  key_name   = \"deployer-key\"\n  public_key = \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD3F6tyPEFEzV0LX3X8BsXdMsQz1x2cEikKDEY0aIj41qgxMCP/iteneqXSIFZBp5vizPvaoIR3Um9xK7PGoW8giupGn+EPuxIA4cDM4vzOqOkiMPhz5XK0whEjkVzTo4+S0puvDZuwIsdiW9mxhJc7tgBNL0cYlWSYVkz4G/fslNfRPW5mYAM49f4fhtxPb5ok4Q2Lg9dPKVHO/Bgeu5woMc7RY0p1ej6D4CKFE6lymSDJpW0YHX/wqE9+cfEauh7xZcG0q9t2ta6F6fmX0agvpFyZo8aFbXeUBr7osSCJNgvavWbM/06niWrOvYX2xwWdhXmXSrbX8ZbabVohBK41 email@example.com\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/key_pair\n\n")
}

func AWSLaunchConfigurationBP() {
	color.Green("\nresource \"aws_launch_configuration\" \"as_conf\" {\n  name          = \"web_config\"\n  image_id      = data.aws_ami.ubuntu.id\n  instance_type = \"t2.micro\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/launch_configuration\n\n")
}

func AWSLaunchTemplateBP() {
	color.Green("\nresource \"aws_launch_template\" \"foo\" {\n  name = \"foo\"\n\n  block_device_mappings {\n    device_name = \"/dev/sda1\"\n\n    ebs {\n      volume_size = 20\n    }\n  }\n\n  capacity_reservation_specification {\n    capacity_reservation_preference = \"open\"\n  }\n\n  cpu_options {\n    core_count       = 4\n    threads_per_core = 2\n  }\n\n  credit_specification {\n    cpu_credits = \"standard\"\n  }\n\n  disable_api_termination = true\n\n  ebs_optimized = true\n\n  elastic_gpu_specifications {\n    type = \"test\"\n  }\n\n  elastic_inference_accelerator {\n    type = \"eia1.medium\"\n  }\n\n  iam_instance_profile {\n    name = \"test\"\n  }\n\n  image_id = \"ami-test\"\n\n  instance_initiated_shutdown_behavior = \"terminate\"\n\n  instance_market_options {\n    market_type = \"spot\"\n  }\n\n  instance_type = \"t2.micro\"\n\n  kernel_id = \"test\"\n\n  key_name = \"test\"\n\n  license_specification {\n    license_configuration_arn = \"arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef\"\n  }\n\n  metadata_options {\n    http_endpoint               = \"enabled\"\n    http_tokens                 = \"required\"\n    http_put_response_hop_limit = 1\n  }\n\n  monitoring {\n    enabled = true\n  }\n\n  network_interfaces {\n    associate_public_ip_address = true\n  }\n\n  placement {\n    availability_zone = \"us-west-2a\"\n  }\n\n  ram_disk_id = \"test\"\n\n  vpc_security_group_ids = [\"sg-12345678\"]\n\n  tag_specifications {\n    resource_type = \"instance\"\n\n    tags = {\n      Name = \"test\"\n    }\n  }\n\n  user_data = filebase64(\"${path.module}/example.sh\")\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/launch_template\n\n")
}

func AWSPlacementGroupBP() {
	color.Green("\nresource \"aws_placement_group\" \"web\" {\n  name     = \"hunky-dory-pg\"\n  strategy = \"cluster\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/placement_group\n\n")
}

func AWSSnapshotCreateVolumePermissionBP() {
	color.Green("\nresource \"aws_snapshot_create_volume_permission\" \"example_perm\" {\n  snapshot_id = aws_ebs_snapshot.example_snapshot.id\n  account_id  = \"12345678\"\n}\n")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/snapshot_create_volume_permission\n\n")
}

func AWSSpotDatafeedSubscriptionBP() {
	color.Green("\nresource \"aws_s3_bucket\" \"default\" {\n  bucket = \"tf-spot-datafeed\"\n}\n\nresource \"aws_spot_datafeed_subscription\" \"default\" {\n  bucket = aws_s3_bucket.default.bucket\n  prefix = \"my_subdirectory\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/spot_datafeed_subscription\n\n")
}

func AWSSpotFleetRequestBP() {
	color.Green("\n# Request a Spot fleet\nresource \"aws_spot_fleet_request\" \"cheap_compute\" {\n  iam_fleet_role      = \"arn:aws:iam::12345678:role/spot-fleet\"\n  spot_price          = \"0.03\"\n  allocation_strategy = \"diversified\"\n  target_capacity     = 6\n  valid_until         = \"2019-11-04T20:44:20Z\"\n\n  launch_specification {\n    instance_type            = \"m4.10xlarge\"\n    ami                      = \"ami-1234\"\n    spot_price               = \"2.793\"\n    placement_tenancy        = \"dedicated\"\n    iam_instance_profile_arn = aws_iam_instance_profile.example.arn\n  }\n\n  launch_specification {\n    instance_type            = \"m4.4xlarge\"\n    ami                      = \"ami-5678\"\n    key_name                 = \"my-key\"\n    spot_price               = \"1.117\"\n    iam_instance_profile_arn = aws_iam_instance_profile.example.arn\n    availability_zone        = \"us-west-1a\"\n    subnet_id                = \"subnet-1234\"\n    weighted_capacity        = 35\n\n    root_block_device {\n      volume_size = \"300\"\n      volume_type = \"gp2\"\n    }\n\n    tags = {\n      Name = \"spot-fleet-example\"\n    }\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/spot_fleet_request\n\n")
}

func AWSSpotInstanceRequestBP() {
	color.Green("\nresource \"aws_spot_instance_request\" \"cheap_worker\" {\n  ami           = \"ami-1234\"\n  spot_price    = \"0.03\"\n  instance_type = \"c4.xlarge\"\n\n  tags = {\n    Name = \"CheapWorker\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/spot_instance_request\n\n")
}

func AWSVolumeAttachmentBP() {
	color.Green("\nresource \"aws_volume_attachment\" \"ebs_att\" {\n  device_name = \"/dev/sdh\"\n  volume_id   = aws_ebs_volume.example.id\n  instance_id = aws_instance.web.id\n}\n\nresource \"aws_instance\" \"web\" {\n  ami               = \"ami-21f78e11\"\n  availability_zone = \"us-west-2a\"\n  instance_type     = \"t2.micro\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}\n\nresource \"aws_ebs_volume\" \"example\" {\n  availability_zone = \"us-west-2a\"\n  size              = 1\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/volume_attachment\n\n")
}
