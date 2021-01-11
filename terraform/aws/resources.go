package aws

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"strings"
	boilerplate2 "tf/boilerplate"
	"tf/terraform/aws/resourceprompts"
)

func resources() []string {
	return []string{
		"aws_instance",
		"aws_vpc",
		"aws_s3_bucket",
		"aws_s3_access_point",
		"aws_s3_account_public_access_block",
		"aws_s3_bucket_analytics_configuration",
		"aws_s3_bucket_inventory",
		"aws_s3_bucket_metric",
		"aws_s3_bucket_notification",
		"aws_s3_bucket_object",
		"aws_s3_bucket_ownership_controls",
		"aws_s3_bucket_policy",
		"aws_s3_bucket_public_access_block",
		"aws_access_analyzer_analyzer",
		"aws_acm_certificate",
		"aws_acm_certificate_validation",
		"aws_acmpca_certificate_authority",
		"aws_alb",
		"aws_alb_listener",
		"aws_alb_listener_certificate",
		"aws_alb_listener_rule",
		"aws_alb_target_group",
		"aws_alb_target_group_attachment",
		"aws_ami",
		"aws_ami_copy",
		"aws_ami_from_instance",
		"aws_ami_launch_permission",
		"aws_api_gateway_account",
		"aws_api_gateway_api_key",
		"aws_api_gateway_authorizer",
		"aws_api_gateway_base_path_mapping",
		"aws_api_gateway_client_certificate",
		"aws_api_gateway_deployment",
		"aws_api_gateway_documentation_part",
		"aws_api_gateway_documentation_version",
		"aws_api_gateway_domain_name",
		"aws_api_gateway_gateway_response",
		"aws_api_gateway_integration",
		"aws_api_gateway_integration_response",
		"aws_api_gateway_method",
		"aws_appautoscaling_scheduled_action",
		"aws_appautoscaling_target",
		"aws_appmesh_mesh",
		"aws_appmesh_route",
		"aws_appmesh_virtual_node",
		"aws_appmesh_virtual_router",
		"aws_appmesh_virtual_service",
		"aws_appsync_api_key",
		"aws_appsync_datasource",
		"aws_appsync_function",
		"aws_appsync_graphql_api",
		"aws_appsync_resolver",
		"aws_athena_database",
		"aws_athena_named_query",
		"aws_athena_workgroup",
		"aws_autoscaling_attachment",
		"aws_autoscaling_group",
		"aws_autoscaling_lifecycle_hook",
		"aws_autoscaling_notification",
		"aws_autoscaling_policy",
		"aws_autoscaling_schedule",
		"aws_backup_plan",
		"aws_backup_selection",
		"aws_backup_vault",
		"aws_batch_compute_environment",
		"aws_batch_job_definition",
		"aws_batch_job_queue",
		"aws_budgets_budget",
		"aws_cloud9_environment_ec2",
		"aws_cloudformation_stack",
		"aws_cloudformation_stack_set",
		"aws_api_gateway_method_response",
		"aws_api_gateway_method_settings",
		"aws_api_gateway_model",
		"aws_api_gateway_request_validator",
		"aws_api_gateway_resource",
		"aws_api_gateway_rest_api",
		"aws_api_gateway_stage",
		"aws_api_gateway_usage_plan",
		"aws_api_gateway_usage_plan_key",
		"aws_api_gateway_vpc_link",
		"aws_apigatewayv2_api",
		"aws_apigatewayv2_api_mapping",
		"aws_apigatewayv2_authorizer",
		"aws_apigatewayv2_deployment",
		"aws_apigatewayv2_domain_name",
		"aws_apigatewayv2_integration",
		"aws_apigatewayv2_integration_response",
		"aws_apigatewayv2_model",
		"aws_apigatewayv2_route",
		"aws_apigatewayv2_route_response",
		"aws_apigatewayv2_stage",
		"aws_apigatewayv2_vpc_link",
		"aws_app_cookie_stickiness_policy",
		"aws_cloud9_environment_ec2",
		"aws_cloudformation_stack",
		"aws_cloudformation_stack_set",
		"aws_cloudformation_stack_set_instance",
		"aws_cloudfront_distribution",
		"aws_cloudfront_origin_access_identity",
		"aws_cloudfront_public_key",
		"aws_cloudhsm_v2_cluster",
		"aws_cloudhsm_v2_hsm",
		"aws_cloudtrail", "aws_cloudwatch_dashboard",
		"aws_cloudwatch_event_permission",
		"aws_cloudwatch_event_rule",
		"aws_cloudwatch_event_target",
		"aws_cloudwatch_log_destination",
		"aws_cloudwatch_log_destination_policy",
		"aws_cloudwatch_log_group",
		"aws_cloudwatch_log_metric_filter",
		"aws_cloudwatch_log_resource_policy",
		"aws_cloudwatch_log_stream",
		"aws_cloudwatch_log_subscription_filter",
		"aws_cloudwatch_metric_alarm",
		"aws_codebuild_project",
		"aws_codebuild_source_credential",
		"aws_codebuild_webhook",
		"aws_codecommit_repository",
		"aws_codecommit_trigger",
		"aws_codedeploy_app",
		"aws_codedeploy_deployment_config",
		"aws_codedeploy_deployment_group",
		"aws_codepipeline",
		"aws_codepipeline_webhook",
		"aws_codestarnotifications_notification",
		"aws_cognito_identity_pool",
		"aws_cognito_identity_pool_roles_attachment",
		"aws_cognito_identity_provider",
		"aws_cognito_resource_server",
		"aws_cognito_user_group",
		"aws_cognito_user_pool",
		"aws_cognito_user_pool_client",
		"aws_cognito_user_pool_domain",
		"aws_config_aggregate_authorization",
		"aws_config_config_rule",
		"aws_config_configuration_aggregator",
		"aws_config_configuration_recorder",
		"aws_config_configuration_recorder_status",
		"aws_config_delivery_channel",
		"aws_config_organization_custom_rule",
		"aws_config_organization_managed_rule",
		"aws_cur_report_definition",
		"aws_customer_gateway",
		"aws_datapipeline_pipeline",
		"aws_datasync_agent",
		"aws_datasync_location_efs",
		"aws_datasync_location_nfs",
		"aws_datasync_location_s3",
		"aws_datasync_location_smb",
		"aws_datasync_task",
		"aws_dax_cluster",
		"aws_dax_parameter_group",
		"aws_dax_subnet_group",
		"aws_db_cluster_snapshot",
		"aws_db_event_subscription",
		"aws_db_instance",
		"aws_db_instance_role_association",
		"aws_db_option_group",
		"aws_db_parameter_group",
		"aws_db_proxy",
		"aws_db_proxy_default_target_group",
		"aws_db_proxy_target",
		"aws_db_security_group",
		"aws_db_snapshot",
		"aws_db_subnet_group",
		"aws_default_network_acl",
		"aws_default_route_table",
		"aws_default_security_group",
		"aws_default_subnet",
		"aws_default_vpc",
		"aws_default_vpc_dhcp_options",
		"aws_devicefarm_project",
		"aws_directory_service_conditional_forwarder",
		"aws_directory_service_directory",
		"aws_directory_service_log_subscription",
		"aws_dlm_lifecycle_policy",
		"aws_dms_certificate",
		"aws_dms_endpoint",
		"aws_dms_event_subscription",
		"aws_dms_replication_instance",
		"aws_dms_replication_subnet_group",
		"aws_dms_replication_task",
		"aws_docdb_cluster",
		"aws_docdb_cluster_instance",
		"aws_docdb_cluster_parameter_group",
		"aws_docdb_cluster_snapshot",
		"aws_docdb_subnet_group",
		"aws_dx_bgp_peer",
		"aws_dx_connection",
		"aws_dx_connection_association",
		"aws_dx_gateway",
		"aws_dx_gateway_association",
		"aws_dx_gateway_association_proposal",
		"aws_dx_hosted_private_virtual_interface",
		"aws_dx_hosted_private_virtual_interface_accepter",
		"aws_dx_hosted_public_virtual_interface",
		"aws_dx_hosted_public_virtual_interface_acceptor",
		"aws_dx_hosted_transit_virtual_interface",
		"aws_dx_hosted_transit_virtual_interface_acceptor",
		"aws_dx_lag",
		"aws_dx_private_virtual_interface",
		"aws_dx_public_virtual_interface",
		"aws_dynamodb_global_table",
		"aws_dynamodb_table",
		"aws_dynamodb_table_item",
		"aws_ebs_default_kms_key",
		"aws_ebs_encryption_by_default",
		"aws_ebs_snapshot",
		"aws_ebs_snapshot_copy",
		"aws_ebs_volume",
		"aws_ec2_availability_zone_group",
		"aws_ec2_capacity_reservation",
		"aws_ec2_client_vpn_authorization_rule",
		"aws_ec2_client_vpn_endpoint",
		"aws_ec2_client_vpn_network_association",
		"aws_ec2_client_vpn_route",
		"aws_ec2_fleet",
		"aws_ec2_local_gateway_route",
		"aws_ec2_local_gateway_route_table_vpc_association",
		"aws_ec2_tag",
		"aws_ec2_traffic_mirror_filter",
		"aws_ec2_traffic_mirror_filter_rule",
		"aws_ec2_traffic_mirror_session",
		"aws_ec2_traffic_mirror_target",
		"aws_ec2_transit_gateway",
		"aws_ec2_transit_gateway_peering_attachment",
		"aws_ec2_transit_gateway_peering_attachment_accepter",
		"aws_ec2_transit_gateway_route",
		"aws_ec2_transit_gateway_route_table",
		"aws_ec2_transit_gateway_route_table_association",
		"aws_ec2_transit_gateway_route_table_propagation",
		"aws_ec2_transit_gateway_vpc_attachment",
		"aws_ec2_transit_gateway_vpc_attachment_accepter",
		"aws_ecr_lifecycle_policy",
		"aws_ecr_repository",
		"aws_ecr_repository_policy",
		"aws_ecs_capacity_provider",
		"aws_ecs_cluster",
		"aws_ecs_service",
		"aws_ecs_task_definition",
		"aws_efs_access_point",
		"aws_efs_file_system",
		"aws_efs_file_system_policy",
		"aws_efs_mount_target",
		"aws_egress_only_internet_gateway",
		"aws_eip",
		"aws_eip_association",
		"aws_eks_cluster",
		"aws_eks_fargate_profile",
		"aws_eks_node_group",
		"aws_elastic_beanstalk_application",
		"aws_elastic_beanstalk_application_version",
		"aws_elastic_beanstalk_configuration_template",
		"aws_elastic_beanstalk_environment",
		"aws_elasticache_cluster",
		"aws_elasticache_parameter_group",
		"aws_elasticache_replication_group",
		"aws_elasticache_security_group",
		"aws_elasticache_subnet_group",
		"aws_elasticsearch_domain",
		"aws_elasticsearch_domain_policy",
		"aws_elastictranscoder_pipeline",
		"aws_elastictranscoder_preset",
		"aws_elb",
		"aws_elb_attachment",
		"aws_lambda_alias",
		"aws_lambda_code_signing_config",
		"aws_lambda_event_source_mapping",
		"aws_lambda_function",
		"aws_lambda_function_event_invoke_config",
		"aws_lambda_layer_version",
		"aws_lambda_permission",
		"aws_lambda_provisioned_concurrency_config",
		"aws_spot_datafeed_subscription",
		"aws_spot_fleet_request",
		"aws_spot_instance_request",
	}
}

func ResourcePrompt(resource string, boilerplate bool) {

	if resource == "" {
		color.Green("\nSelect aws Resources(e.g. aws_instance, aws_vpc):\n\n", "text")

		resourcePrompt := promptui.Select{
			Label:             "",
			Size:              20,
			Items:             resources(),
			StartInSearchMode: true,
			Searcher: func(input string, index int) bool {
				provider := resources()[index]
				name := strings.Replace(strings.ToLower(provider), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			},
		}

		var err error
		_, resource, err = resourcePrompt.Run()
		if err != nil {
			fmt.Println(err)
		}

		if boilerplate {
			boilerplate2.SelectResourceBP("aws", resource)
			return
		}
	}

	switch resource {
	case "aws_vpc":
		resourceprompts.AWSVPCPrompt()
	case "aws_default_network_acl":
		resourceprompts.AWSDefaultNetworkACLPrompt()
	case "aws_default_security_group":
		resourceprompts.AWSDefaultSecurityGroupPrompt()
	case "aws_default_subnet":
		resourceprompts.AWSDefaultSubnetPrompt()
	case "aws_default_vpc_dhcp_options":
		resourceprompts.AWSDefaultVPCDHCPOptionsPrompt()
	case "aws_default_vpc":
		resourceprompts.AWSDefaultVPCPrompt()
	case "aws_ec2_managed_prefix_list":
		resourceprompts.AWSEC2ManagedPrefixListPrompt()
	case "aws_egress_only_internet_gateway":
		resourceprompts.AWSEgressOnlyInternetGatewayPrompt()
	case "aws_customer_gateway":
		resourceprompts.AWSCustomerGatewayPrompt()
	case "aws_flow_log":
		resourceprompts.AWSFlowLogPrompt()
	case "aws_nat_gateway":
		resourceprompts.AWSNatGatewayPrompt()
	case "aws_network_acl":
		resourceprompts.AWSNetworkACLPrompt()
	case "aws_internet_gateway":
		resourceprompts.AWSInternetGatewayPrompt()
	case "aws_main_route_table_association":
		resourceprompts.AWSMainRouteTableAssociationPrompt()
	case "aws_default_route_table":
		resourceprompts.AWSDefaultRouteTablePrompt()
	case "aws_s3_bucket":
		resourceprompts.AWSS3BucketPrompt()
	case "aws_s3_access_point":
		resourceprompts.AWSS3AccessPointPrompt()
	case "aws_s3_account_public_access_block":
		resourceprompts.AWSS3AccountPublicAccessBlockPrompt()
	case "aws_s3_bucket_analytics_configuration":
		resourceprompts.AWSS3BucketAnalyticsConfigurationPrompt()
	case "aws_s3_bucket_metric":
		resourceprompts.AWSS3BucketMetricPrompt()
	case "aws_s3_bucket_notification":
		resourceprompts.AWSS3BucketNotificationPrompt()
	case "aws_s3_bucket_object":
		resourceprompts.AWSS3BucketObjectPrompt()
	case "aws_s3_bucket_ownership_controls":
		resourceprompts.AWSS3BucketOwnershipControlsPrompt()
	case "aws_s3_bucket_public_access_block":
		resourceprompts.AWSS3BucketPublicAccessBlockPrompt()
	case "aws_ami":
		resourceprompts.AWSAMIPrompt()
	case "aws_ami_copy":
		resourceprompts.AWSAMICopyPrompt()
	case "aws_ebs_snapshot":
		resourceprompts.AWSEBSSnapshotPrompt()
	case "aws_ebs_snapshot_copy":
		resourceprompts.AWSEBSSnapshotCopyPrompt()
	case "aws_ebs_default_kms_key":
		resourceprompts.AWSEBSDefaultKMSKeyPrompt()
	case "aws_ebs_encryption_by_default":
		resourceprompts.AWSEBSEncryptionByDefaultPrompt()
	case "aws_ami_from_instance":
		resourceprompts.AWSAMIFromInstancePrompt()
	case "aws_ami_launch_permission":
		resourceprompts.AWSAMILaunchPermissionPrompt()
	case "aws_elb":
		resourceprompts.AWSELBPrompt()
	case "aws_db_instance":
		resourceprompts.AWSDBInstancePrompt()
	case "aws_db_cluster_snapshot":
		resourceprompts.AWSDBClusterSnapshotPrompt()
	case "aws_db_event_subscription":
		resourceprompts.AWSDBEventSubscriptionPrompt()
	case "aws_db_instance_role_association":
		resourceprompts.AWSDBInstanceRoleAssociationPrompt()
	case "aws_db_option_group":
		resourceprompts.AWSDBOptionGroupPrompt()
	case "aws_ebs_volume":
		resourceprompts.AWSEBSVolumePrompt()
	case "aws_ec2_availability_zone_group":
		resourceprompts.AWSEC2AvailabilityZoneGroupPrompt()
	case "aws_ec2_capacity_reservation":
		resourceprompts.AWSEC2CapacityReservationPrompt()
	case "aws_ec2_carrier_gateway":
		resourceprompts.AWSEC2CarrierGatewayPrompt()
	case "aws_ec2_client_vpn_authorization_rule":
		resourceprompts.AWSEC2ClientVPNAuthorizationRulePrompt()
	case "aws_ec2_client_vpn_route":
		resourceprompts.AWSEC2ClientVPNRoutePrompt()
	case "aws_ec2_local_gateway_route_table_vpc_association":
		resourceprompts.AWSEC2LocalGatewayRouteTableVPCAssociationPrompt()
	case "aws_ec2_tag":
		resourceprompts.AWSEC2TagPrompt()
	case "aws_ec2_fleet":
		resourceprompts.AWSEC2FleetPrompt()
	case "aws_ec2_local_gateway_route":
		resourceprompts.AWSEC2LocalGatewayRoutePrompt()
	case "aws_ec2_client_vpn_endpoint":
		resourceprompts.AWSEC2ClientVPNEndpointPrompt()
	case "aws_ec2_client_vpn_network_association":
		resourceprompts.AWSEC2ClientVPNNetworkAssociationPrompt()
	case "aws_ec2_traffic_mirror_filter":
		resourceprompts.AWSEC2TrafficMirrorFilterPrompt()
	case "aws_ec2_traffic_mirror_filter_rule":
		resourceprompts.AWSEC2TrafficMirrorFilterRulePrompt()
	case "aws_ec2_traffic_mirror_session":
		resourceprompts.AWSEC2TrafficMirrorSessionPrompt()
	case "aws_ec2_traffic_mirror_target":
		resourceprompts.AWSEC2TrafficMirrorTargetPrompt()
	case "aws_ec2_transit_gateway":
		resourceprompts.AWSEC2TransitGatewayPrompt()
	case "aws_ec2_transit_gateway_peering_attachment":
		resourceprompts.AWSEC2TransitGatewayPeeringAttachmentPrompt()
	case "aws_ec2_transit_gateway_peering_attachment_accepter":
		resourceprompts.AWSEC2TransitGatewayPeeringAttachmentAccepterPrompt()
	case "aws_ec2_transit_gateway_route":
		resourceprompts.AWSEC2TransitGatewayRoutePrompt()
	case "aws_ec2_transit_gateway_route_table":
		resourceprompts.AWSEC2TransitGatewayRouteTablePrompt()
	case "aws_ec2_transit_gateway_route_table_association":
		resourceprompts.AWSEC2TransitGatewayRouteTableAssociationPrompt()
	case "aws_ec2_transit_gateway_route_table_propagation":
		resourceprompts.AWSEC2TransitGatewayRouteTablePropagationPrompt()
	case "aws_db_parameter_group":
		resourceprompts.AWSDBParameterGroupPrompt()
	case "aws_ec2_transit_gateway_vpc_attachment_accepter":
		resourceprompts.AWSEC2TransitGatewayVPCAttachmentAccepterPrompt()
	case "aws_eip":
		resourceprompts.AWSEIPPrompt()
	case "aws_eip_association":
		resourceprompts.AWSEIPAssociationPrompt()
	case "aws_ec2_transit_gateway_vpc_attachment":
		resourceprompts.AWSEC2TransitGatewayVPCAttachmentPrompt()
	case "aws_instance":
		resourceprompts.AWSInstancePrompt()
	case "aws_key_pair":
		resourceprompts.AWSKeyPairPrompt()
	case "aws_placement_group":
		resourceprompts.AWSPlacementGroupPrompt()
	case "aws_snapshot_create_volume_permission":
		resourceprompts.AWSSnapshotCreateVolumePermissionPrompt()
	case "aws_spot_datafeed_subscription":
		resourceprompts.AWSSpotDatafeedSubscriptionPrompt()
	case "aws_spot_fleet_request":
		resourceprompts.AWSSpotFleetRequestPrompt()
	case "aws_spot_instance_request":
		resourceprompts.AWSSpotInstanceRequestPrompt()
	case "aws_db_proxy":
		resourceprompts.AWSDbProxyPrompt()
	case "aws_db_proxy_default_target_group":
		resourceprompts.AWSDBProxyDefaultTargetGroupPrompt()
	case "aws_db_proxy_target":
		resourceprompts.AWSDBProxyTargetPrompt()
	case "aws_db_security_group":
		resourceprompts.AWSDBSecurityGroupPrompt()
	case "aws_db_snapshot":
		resourceprompts.AWSDBSnapshotPrompt()
	case "aws_db_subnet_group":
		resourceprompts.AWSDBSubnetGroupPrompt()
	case "aws_acm_certificate":
		resourceprompts.AWSACMCertificatePrompt()
	case "aws_acmpca_certificate_authority":
		resourceprompts.AWSACMPCACertificatePrompt()
	case "aws_api_gateway_account":
		resourceprompts.AWSAPIGatewayAccountPrompt()
	case "aws_api_gateway_api_key":
		resourceprompts.AWSAPIGatewayApiKeyPrompt()
	case "aws_api_gateway_authorizer":
		resourceprompts.AWSAPiGatewayAuthorizer()
	case "aws_api_gateway_base_path_mapping":
		resourceprompts.AWSAPIGatewayBasePathMappingPrompt()
	case "aws_api_gateway_client_certificate":
		resourceprompts.AWSAPIGatewayClientCertificatePrompt()
	case "aws_api_gateway_deployment":
		resourceprompts.AWSAPIGatewayDeploymentPrompt()
	case "aws_api_gateway_documentation_part":
		resourceprompts.AWSAPIGatewayDocumentationPartPrompt()
	case "aws_api_gateway_documentation_version":
		resourceprompts.AWSAPIGatewayDocumentationVersionPrompt()
	case "aws_api_gateway_domain_name":
		resourceprompts.AWSAPIGatewayDomainNamePrompt()
	case "aws_api_gateway_gateway_response":
		resourceprompts.AWSAPIGatewayGatewayResponsePrompt()
	case "aws_lambda_alias":
		resourceprompts.AWSLambdaAliasPrompt()
	case "aws_lambda_code_signing_config":
		resourceprompts.AWSLambdaCodeSigningConfigPrompt()
	case "aws_lambda_event_source_mapping":
		resourceprompts.AWSLambdaEventSourceMappingPrompt()
	case "aws_lambda_function":
		resourceprompts.AWSLambdaFunctionPrompt()
	case "aws_lambda_layer_version":
		resourceprompts.AWSLambdaLayerVersionPrompt()
	case "aws_lambda_permission":
		resourceprompts.AWSLambdaPermissionPrompt()
	case "aws_lambda_provisioned_concurrency_config":
		resourceprompts.AWSLambdaProvisionedConcurrencyConfigPrompt()
	case "aws_elastic_beanstalk_application":
		resourceprompts.AWSElasticBeanstalkApplicationPrompt()
	case "aws_elastic_beanstalk_application_version":
		resourceprompts.AWSElasticBeanstalkApplicationVersionPrompt()
	case "aws_elastic_beanstalk_configuration_template":
		resourceprompts.AWSElasticBeanstalkConfigurationTemplatePrompt()
	default:
		color.Red("No support added yet for your resource! Coming soon...")
		color.Yellow("\nMeanwhile try getting the boilerplate version by running `tf resource -p aws -r " + resource + " -b`")
	}
}
