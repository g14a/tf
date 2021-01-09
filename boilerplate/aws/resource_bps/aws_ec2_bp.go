package resource_bps

import (
	"github.com/fatih/color"
)

func AWSEC2InstanceBP() {
	color.Green("\nresource \"aws_instance\" \"foo\" {\n  ami           = data.aws_ami.ubuntu.id\n  instance_type = \"t3.micro\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance\n\n")
}

func AWSAMIBP()  {
	color.Green("\n# Create an AMI that will start a machine whose root device is backed by\n# an EBS volume populated from a snapshot. It is assumed that such a snapshot\n# already exists with the id \"snap-xxxxxxxx\".\nresource \"aws_ami\" \"example\" {\n  name                = \"terraform-example\"\n  virtualization_type = \"hvm\"\n  root_device_name    = \"/dev/xvda\"\n\n  ebs_block_device {\n    device_name = \"/dev/xvda\"\n    snapshot_id = \"snap-xxxxxxxx\"\n    volume_size = 8\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami\n\n")
}

func AWSAMICopyBP()  {
	color.Green("\nresource \"aws_ami_copy\" \"example\" {\n  name              = \"terraform-example\"\n  description       = \"A copy of ami-xxxxxxxx\"\n  source_ami_id     = \"ami-xxxxxxxx\"\n  source_ami_region = \"us-west-1\"\n\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami_copy\n\n")
}

func AWSAMIFromInstanceBP()  {
	color.Green("\nresource \"aws_ami_from_instance\" \"example\" {\n  name               = \"terraform-example\"\n  source_instance_id = \"i-xxxxxxxx\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami_from_instance\n\n")
}

func AWSAMIFromLaunchPermissionBP() {
	color.Green("\nresource \"aws_ami_launch_permission\" \"example\" {\n  image_id   = \"ami-12345678\"\n  account_id = \"123456789012\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ami_launch_permission\n\n")
}

func AWSEBSDefaultKMSKeyBP()  {
	color.Green("\nresource \"aws_ebs_default_kms_key\" \"example\" {\n  key_arn = aws_kms_key.example.arn\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_default_kms_key\n\n")
}

func AWSEBSEncryptionByDefaultBP()  {
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

func AWSEC2AvailabilityZoneGroupBP()  {
	color.Green("\nresource \"aws_ec2_availability_zone_group\" \"example\" {\n  group_name    = \"us-west-2-lax-1\"\n  opt_in_status = \"opted-in\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_availability_zone_group\n\n")
}

func AWSEC2CapacityReservationBP()  {
	color.Green("\nresource \"aws_ec2_capacity_reservation\" \"default\" {\n  instance_type     = \"t2.micro\"\n  instance_platform = \"Linux/UNIX\"\n  availability_zone = \"eu-west-1a\"\n  instance_count    = 1\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_capacity_reservation\n\n")
}

func AWSEC2CarrierGatewayPrompt() {
	color.Green("\nresource \"aws_ec2_carrier_gateway\" \"example\" {\n  vpc_id = aws_vpc.example.id\n\n  tags = {\n    Name = \"example-carrier-gateway\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_carrier_gateway\n\n")
}

func AWSEC2ClientVPNAuthorizationRuleBP()  {
	color.Green("\nresource \"aws_ec2_client_vpn_authorization_rule\" \"example\" {\n  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.example.id\n  target_network_cidr    = aws_subnet.example.cidr_block\n  authorize_all_groups   = true\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_client_vpn_authorization_rule\n\n")
}

func AWSEC2ClientVPNEndpointBP() {
	color.Green("\nresource \"aws_ec2_client_vpn_endpoint\" \"example\" {\n  description            = \"terraform-clientvpn-example\"\n  server_certificate_arn = aws_acm_certificate.cert.arn\n  client_cidr_block      = \"10.0.0.0/16\"\n\n  authentication_options {\n    type                       = \"certificate-authentication\"\n    root_certificate_chain_arn = aws_acm_certificate.root_cert.arn\n  }\n\n  connection_log_options {\n    enabled               = true\n    cloudwatch_log_group  = aws_cloudwatch_log_group.lg.name\n    cloudwatch_log_stream = aws_cloudwatch_log_stream.ls.name\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_client_vpn_endpoint\n\n")
}