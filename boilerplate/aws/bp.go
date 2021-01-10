package aws

import (
	"tf/boilerplate/aws/resource_bps"
)

func ResourceBP(resource string) {

	switch resource {
	case "aws_instance":
		resource_bps.AWSEC2InstanceBP()
	case "aws_ami":
		resource_bps.AWSAMIBP()
	case "aws_ami_copy":
		resource_bps.AWSAMICopyBP()
	case "aws_ami_from_instance":
		resource_bps.AWSAMIFromInstanceBP()
	case "aws_ami_launch_permission":
		resource_bps.AWSAMIFromLaunchPermissionBP()
	case "aws_ebs_default_kms_key":
		resource_bps.AWSEBSDefaultKMSKeyBP()
	case "aws_ebs_encryption_by_default":
		resource_bps.AWSEBSEncryptionByDefaultBP()
	case "aws_ebs_snapshot":
		resource_bps.AWSEBSSnapshotBP()
	case "aws_ebs_snapshot_copy":
		resource_bps.AWSEBSSnapshotCopyBP()
	case "aws_db_instance":
		resource_bps.AWSDBInstanceBP()
	case "aws_db_cluster_snapshot":
		resource_bps.AWSDBClusterSnapshotBP()
	case "aws_db_event_subscription":
		resource_bps.AWSDBEventSubscriptionBP()
	case "aws_db_instance_role_association":
		resource_bps.AWSDBInstanceRoleAssociationBP()
	case "aws_db_option_group":
		resource_bps.AWSDBOptionGroupBP()
	case "aws_ebs_volume":
		resource_bps.AWSEBSVolumeBP()
	case "aws_ec2_availability_zone_group":
		resource_bps.AWSEC2AvailabilityZoneGroupBP()
	case "aws_ec2_capacity_reservation":
		resource_bps.AWSEC2CapacityReservationBP()
	case "aws_ec2_carrier_gateway":
		resource_bps.AWSEC2CarrierGatewayPrompt()
	case "aws_ec2_client_vpn_endpoint":
		resource_bps.AWSEC2ClientVPNEndpointBP()
	case "aws_ec2_client_vpn_authorization_rule":
		resource_bps.AWSEC2ClientVPNAuthorizationRuleBP()
	case "aws_ec2_client_vpn_network_association":
		resource_bps.AWSEC2ClientVPNNetworkAssociationBP()
	case "aws_ec2_client_vpn_route":
		resource_bps.AWSEC2ClientVPNRouteBP()
	case "aws_ec2_fleet":
		resource_bps.AWSEC2FleetBP()
	case "aws_ec2_local_gateway_route":
		resource_bps.AWSEC2LocalGatewayRouteBP()
	case "aws_ec2_local_gateway_route_table_vpc_association":
		resource_bps.AWSEC2LocalGatewayRouteTableVPCAssociationBP()
	case "aws_ec2_transit_gateway_peering_attachment_accepter":
		resource_bps.AWSEC2TransitGatewayPeeringAttachmentAccepterBP()
	case "aws_ec2_transit_gateway_route":
		resource_bps.AWSEC2TransitGatewayRouteBP()
	case "aws_ec2_transit_gateway_route_table":
		resource_bps.AWSEC2TransitGatewayRouteTableBP()
	case "aws_ec2_transit_gateway_route_table_association":
		resource_bps.AWSEC2TransitGatewayRouteTableAssociationBP()
	case "aws_ec2_transit_gateway_route_table_propagation":
		resource_bps.AWSEC2TransitGatewayRouteTablePropagationBP()
	case "aws_ec2_transit_gateway_vpc_attachment":
		resource_bps.AWSEC2TransitGatewayVPCAttachmentBP()
	case "aws_ec2_transit_gateway_vpc_attachment_accepter":
		resource_bps.AWSEC2TransitGatewayVPCAttachmentAccepterBP()
	case "aws_eip":
		resource_bps.AWSEIPBP()
	case "aws_eip_association":
		resource_bps.AWSEIPAssociationBP()
	case "aws_key_pair":
		resource_bps.AWSKeyPairBP()
	case "aws_launch_configuration":
		resource_bps.AWSLaunchConfigurationBP()
	case "aws_launch_template":
		resource_bps.AWSLaunchTemplateBP()
	case "aws_placement_group":
		resource_bps.AWSPlacementGroupBP()
	case "aws_snapshot_create_volume_permission":
		resource_bps.AWSSnapshotCreateVolumePermissionBP()
	case "aws_spot_datafeed_subscription":
		resource_bps.AWSSpotDatafeedSubscriptionBP()
	case "aws_spot_fleet_request":
		resource_bps.AWSSpotFleetRequestBP()
	case "aws_spot_instance_request":
		resource_bps.AWSSpotInstanceRequestBP()
	case "aws_volume_attachment":
		resource_bps.AWSVolumeAttachmentBP()
	case "aws_ec2_tag":
		resource_bps.AWSEC2TagBP()
	case "aws_ec2_traffic_mirror_filter":
		resource_bps.AWSEC2TrafficMirrorFilterBP()
	case "aws_ec2_traffic_mirror_filter_rule":
		resource_bps.AWSEC2TrafficMirrorFilterRuleBP()
	case "aws_ec2_traffic_mirror_session":
		resource_bps.AWSEC2TrafficMirrorSessionBP()
	case "aws_ec2_transit_gateway_peering_attachment":
		resource_bps.AWSEC2TransitGatewayPeeringAttachmentBP()
	case "aws_ec2_transit_gateway":
		resource_bps.AWSEC2TransitGatewayBP()
	case "aws_ec2_traffic_mirror_target":
		resource_bps.AWSEC2TrafficMirrorTargetBP()
	case "aws_db_parameter_option":
		resource_bps.AWSDBParameterOptionBP()
	case "aws_db_proxy":
		resource_bps.AWSDBProxyBP()
	case "aws_db_proxy_default_target_group":
		resource_bps.AWSDBProxyDefaultTargetGroupBP()
	case "aws_db_proxy_target_group":
		resource_bps.AWSDBProxyTargetGroupBP()
	case "aws_db_security_group":
		resource_bps.AWSDBSecurityGroupBP()
	case "aws_db_snapshot":
		resource_bps.AWSDBSnapshotBP()
	case "aws_db_subnet_group":
		resource_bps.AWSDBSubnetGroupBP()
	case "aws_s3_bucket":
		resource_bps.AWSS3BucketBP()
	case "aws_vpc":
		resource_bps.AWSVPCBP()
	case "aws_sns_platform_application":
		resource_bps.AWSSNSPlatformApplicationBP()
	case "aws_elastic_beanstalk_application":
		resource_bps.AWSElasticBeanstalkApplication()
	case "aws_cloudfront_distribution":
		resource_bps.AWSCloudFrontDistributionPrompt()
	case "aws_lambda_function":
		resource_bps.AWSLambdaFunctionBP()
	case "aws_lambda_code_signing_config":
		resource_bps.AWSLambdaCodeSigningConfigBP()
	case "aws_lambda_layer_version":
		resource_bps.AWSLambdaLayerVersionBP()
	case "aws_lambda_permission":
		resource_bps.AWSLambdaPermissionBP()
	case "aws_lambda_function_event_invoke_config":
		resource_bps.AWSLambdaFunctionEventInvokeConfigBP()
	case "aws_lambda_provisioned_concurrency_config":
		resource_bps.AWSLambdaProvisionedConcurrencyConfig()
	}
}
