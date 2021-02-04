package aws

import (
	"fmt"
	"github.com/g14a/tf/boilerplate/aws/resourcebps"
)

func ResourceBP(resource string) {

	fmt.Println(resource, "=========resources")
	if resource != "" {
		switch resource {
		case "aws_instance":
			resourcebps.AWSEC2InstanceBP()
		case "aws_ami":
			resourcebps.AWSAMIBP()
		case "aws_ami_copy":
			resourcebps.AWSAMICopyBP()
		case "aws_ami_from_instance":
			resourcebps.AWSAMIFromInstanceBP()
		case "aws_ami_launch_permission":
			resourcebps.AWSAMIFromLaunchPermissionBP()
		case "aws_ebs_default_kms_key":
			resourcebps.AWSEBSDefaultKMSKeyBP()
		case "aws_ebs_encryption_by_default":
			resourcebps.AWSEBSEncryptionByDefaultBP()
		case "aws_ebs_snapshot":
			resourcebps.AWSEBSSnapshotBP()
		case "aws_ebs_snapshot_copy":
			resourcebps.AWSEBSSnapshotCopyBP()
		case "aws_db_instance":
			resourcebps.AWSDBInstanceBP()
		case "aws_db_cluster_snapshot":
			resourcebps.AWSDBClusterSnapshotBP()
		case "aws_db_event_subscription":
			resourcebps.AWSDBEventSubscriptionBP()
		case "aws_db_instance_role_association":
			resourcebps.AWSDBInstanceRoleAssociationBP()
		case "aws_db_option_group":
			resourcebps.AWSDBOptionGroupBP()
		case "aws_ebs_volume":
			resourcebps.AWSEBSVolumeBP()
		case "aws_ec2_availability_zone_group":
			resourcebps.AWSEC2AvailabilityZoneGroupBP()
		case "aws_ec2_capacity_reservation":
			resourcebps.AWSEC2CapacityReservationBP()
		case "aws_ec2_carrier_gateway":
			resourcebps.AWSEC2CarrierGatewayPrompt()
		case "aws_ec2_client_vpn_endpoint":
			resourcebps.AWSEC2ClientVPNEndpointBP()
		case "aws_ec2_client_vpn_authorization_rule":
			resourcebps.AWSEC2ClientVPNAuthorizationRuleBP()
		case "aws_ec2_client_vpn_network_association":
			resourcebps.AWSEC2ClientVPNNetworkAssociationBP()
		case "aws_ec2_client_vpn_route":
			resourcebps.AWSEC2ClientVPNRouteBP()
		case "aws_ec2_fleet":
			resourcebps.AWSEC2FleetBP()
		case "aws_ec2_local_gateway_route":
			resourcebps.AWSEC2LocalGatewayRouteBP()
		case "aws_ec2_local_gateway_route_table_vpc_association":
			resourcebps.AWSEC2LocalGatewayRouteTableVPCAssociationBP()
		case "aws_ec2_transit_gateway_peering_attachment_accepter":
			resourcebps.AWSEC2TransitGatewayPeeringAttachmentAccepterBP()
		case "aws_ec2_transit_gateway_route":
			resourcebps.AWSEC2TransitGatewayRouteBP()
		case "aws_ec2_transit_gateway_route_table":
			resourcebps.AWSEC2TransitGatewayRouteTableBP()
		case "aws_ec2_transit_gateway_route_table_association":
			resourcebps.AWSEC2TransitGatewayRouteTableAssociationBP()
		case "aws_ec2_transit_gateway_route_table_propagation":
			resourcebps.AWSEC2TransitGatewayRouteTablePropagationBP()
		case "aws_ec2_transit_gateway_vpc_attachment":
			resourcebps.AWSEC2TransitGatewayVPCAttachmentBP()
		case "aws_ec2_transit_gateway_vpc_attachment_accepter":
			resourcebps.AWSEC2TransitGatewayVPCAttachmentAccepterBP()
		case "aws_eip":
			resourcebps.AWSEIPBP()
		case "aws_eip_association":
			resourcebps.AWSEIPAssociationBP()
		case "aws_key_pair":
			resourcebps.AWSKeyPairBP()
		case "aws_launch_configuration":
			resourcebps.AWSLaunchConfigurationBP()
		case "aws_launch_template":
			resourcebps.AWSLaunchTemplateBP()
		case "aws_placement_group":
			resourcebps.AWSPlacementGroupBP()
		case "aws_snapshot_create_volume_permission":
			resourcebps.AWSSnapshotCreateVolumePermissionBP()
		case "aws_spot_datafeed_subscription":
			resourcebps.AWSSpotDatafeedSubscriptionBP()
		case "aws_spot_fleet_request":
			resourcebps.AWSSpotFleetRequestBP()
		case "aws_spot_instance_request":
			resourcebps.AWSSpotInstanceRequestBP()
		case "aws_volume_attachment":
			resourcebps.AWSVolumeAttachmentBP()
		case "aws_ec2_tag":
			resourcebps.AWSEC2TagBP()
		case "aws_ec2_traffic_mirror_filter":
			resourcebps.AWSEC2TrafficMirrorFilterBP()
		case "aws_ec2_traffic_mirror_filter_rule":
			resourcebps.AWSEC2TrafficMirrorFilterRuleBP()
		case "aws_ec2_traffic_mirror_session":
			resourcebps.AWSEC2TrafficMirrorSessionBP()
		case "aws_ec2_transit_gateway_peering_attachment":
			resourcebps.AWSEC2TransitGatewayPeeringAttachmentBP()
		case "aws_ec2_transit_gateway":
			resourcebps.AWSEC2TransitGatewayBP()
		case "aws_ec2_traffic_mirror_target":
			resourcebps.AWSEC2TrafficMirrorTargetBP()
		case "aws_db_parameter_option":
			resourcebps.AWSDBParameterOptionBP()
		case "aws_db_proxy":
			resourcebps.AWSDBProxyBP()
		case "aws_db_proxy_default_target_group":
			resourcebps.AWSDBProxyDefaultTargetGroupBP()
		case "aws_db_proxy_target_group":
			resourcebps.AWSDBProxyTargetGroupBP()
		case "aws_db_security_group":
			resourcebps.AWSDBSecurityGroupBP()
		case "aws_db_snapshot":
			resourcebps.AWSDBSnapshotBP()
		case "aws_db_subnet_group":
			resourcebps.AWSDBSubnetGroupBP()
		case "aws_s3_bucket":
			resourcebps.AWSS3BucketBP()
		case "aws_s3_access_point":
			resourcebps.AWSS3AccessPointBP()
		case "aws_s3_account_public_access_block":
			resourcebps.AWSS3AccountPublicAccessBlockBP()
		case "aws_s3_bucket_analytics_configuration":
			resourcebps.AWSS3BucketAnalyticsConfigurationBP()
		case "aws_s3_bucket_inventory":
			resourcebps.AWSS3BucketInventoryBP()
		case "aws_s3_bucket_metric":
			resourcebps.AWSS3BucketMetricBP()
		case "aws_s3_bucket_notification":
			resourcebps.AWSS3BucketNotificationBP()
		case "aws_s3_bucket_object":
			resourcebps.AWSS3BucketObjectBP()
		case "aws_s3_bucket_ownership_controls":
			resourcebps.AWSS3BucketOwnershipControlsBP()
		case "aws_s3_bucket_policy":
			resourcebps.AWSS3BucketPolicyBP()
		case "aws_s3_bucket_public_access_block":
			resourcebps.AWSS3BucketPolicyAccessBlockBP()
		case "aws_lambda_event_source_mapping":
			resourcebps.AWSLambdaEventSourceMappingBP()
		case "aws_lambda_alias":
			resourcebps.AWSLambdaAliasBP()
		case "aws_vpc":
			resourcebps.AWSVPCBP()
		case "aws_customer_gateway":
			resourcebps.AWSCustomerGatewayBP()
		case "aws_default_network_acl":
			resourcebps.AWSDefaultNetworkACLBP()
		case "aws_default_route_table":
			resourcebps.AWSDefaultRouteTableBP()
		case "aws_default_security_group":
			resourcebps.AWSDefaultSecurityGroupBP()
		case "aws_default_subnet":
			resourcebps.AWSDefaultSubnetBP()
		case "aws_default_vpc":
			resourcebps.AWSDefaultVPCBP()
		case "aws_default_vpc_dhcp_options":
			resourcebps.AWSDefaultVPCDHCPOptionsBP()
		case "aws_ec2_managed_prefix_list":
			resourcebps.AWSEC2ManagedPrefixListBP()
		case "aws_egress_only_internet_gateway":
			resourcebps.AWSEgressOnlyInternetGatewayBP()
		case "aws_flow_log":
			resourcebps.AWSFlowLogBP()
		case "aws_internet_gateway":
			resourcebps.AWSInternetGatewayBP()
		case "aws_main_route_table_association":
			resourcebps.AWSMainRouteTableAssociationBP()
		case "aws_nat_gateway":
			resourcebps.AWSNatGatewayBP()
		case "aws_network_acl":
			resourcebps.AWSNetworkACLBP()
		case "aws_network_acl_rule":
			resourcebps.AWSNetworkACLRuleBP()
		case "aws_network_interface":
			resourcebps.AWSNetworkInterfaceBP()
		case "aws_network_interface_attachment":
			resourcebps.AWSNetworkInterfaceAttachmentBP()
		case "aws_network_interface_sg_attachment":
			resourcebps.AWSNetworkInterfaceSGAttachmentBP()
		case "aws_route":
			resourcebps.AWSRouteBP()
		case "aws_route_table":
			resourcebps.AWSRouteTableBP()
		case "aws_route_table_association":
			resourcebps.AWSRouteTableAssociationBP()
		case "aws_security_group":
			resourcebps.AWSSecurityGroupBP()
		case "aws_security_group_rule":
			resourcebps.AWSSecurityGroupRuleBP()
		case "aws_subnet":
			resourcebps.AWSSubnetBP()
		case "aws_vpc_dhcp_options":
			resourcebps.AWSVPCDHCPOptionsBP()
		case "aws_vpc_dhcp_options_association":
			resourcebps.AWSVPCDHCPOptionsAssociationBP()
		case "aws_vpc_endpoint":
			resourcebps.AWSVPCEndpointBP()
		case "aws_vpc_endpoint_connection_notification":
			resourcebps.AWSVPCEndpointConnectionNotificationBP()
		case "aws_vpc_endpoint_route_table_association":
			resourcebps.AWSVPCEndpointRouteTableAssociationBP()
		case "aws_vpc_endpoint_service":
			resourcebps.AWSVPCEndpointServiceBP()
		case "aws_vpc_endpoint_service_allowed_principal":
			resourcebps.AWSVPCEndpointServiceAllowedPrincipalBP()
		case "aws_vpc_endpoint_subnet_association":
			resourcebps.AWSVPCEndpointSubnetAssociationBP()
		case "aws_vpc_ipv4_cidr_block_association":
			resourcebps.AWSVPCIPV4CIDRBlockAssociationBP()
		case "aws_vpc_peering_connection":
			resourcebps.AWSVPCPeeringConnectionBP()
		case "aws_vpc_peering_connection_accepter":
			resourcebps.AWSVPCPeeringConnectionAccepterBP()
		case "aws_vpc_peering_connection_options":
			resourcebps.AWSVPCPeeringConnectionOptionsBP()
		case "aws_vpn_connection":
			resourcebps.AWSVPNConnectionBP()
		case "aws_vpn_connection_route":
			resourcebps.AWSVPNConnectionRouteBP()
		case "aws_vpn_gateway":
			resourcebps.AWSVPNGatewayBP()
		case "aws_vpn_gateway_attachment":
			resourcebps.AWSVPNGatewayAttachmentBP()
		case "aws_vpn_gateway_route_propagation":
			resourcebps.AWSVPNGatewayRoutePropagationBP()
		case "aws_sns_platform_application":
			resourcebps.AWSSNSPlatformApplicationBP()
		case "aws_elastic_beanstalk_application":
			resourcebps.AWSElasticBeanstalkApplication()
		case "aws_cloudfront_distribution":
			resourcebps.AWSCloudFrontDistributionBP()
		case "aws_lambda_function":
			resourcebps.AWSLambdaFunctionBP()
		case "aws_lambda_code_signing_config":
			resourcebps.AWSLambdaCodeSigningConfigBP()
		case "aws_lambda_layer_version":
			resourcebps.AWSLambdaLayerVersionBP()
		case "aws_lambda_permission":
			resourcebps.AWSLambdaPermissionBP()
		case "aws_lambda_function_event_invoke_config":
			resourcebps.AWSLambdaFunctionEventInvokeConfigBP()
		case "aws_lambda_provisioned_concurrency_config":
			resourcebps.AWSLambdaProvisionedConcurrencyConfigBP()
		case "aws_route53_delegation_set":
			resourcebps.AWSRoute53DelegationSetBP()
		case "aws_route53_health_check":
			resourcebps.AWSRoute53HealthCheckBP()
		case "aws_route53_query_log":
			resourcebps.AWSRoute53QueryLogBP()
		case "aws_route53_record":
			resourcebps.AWSRoute53RecordBP()
		case "aws_route53_vpc_association_authorization":
			resourcebps.AWSRoute53VPCAssociationAuthorizationBP()
		case "aws_route53_zone":
			resourcebps.AWSRoute53ZoneBP()
		case "aws_route53_zone_association":
			resourcebps.AWSRoute53ZoneAssociationBP()
		}
	}
}
