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
