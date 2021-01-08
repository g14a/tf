package aws

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"strings"
	boilerplate2 "tf/boilerplate"
	"tf/terraform/aws/resource_prompts"
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
		resource_prompts.AWSVPCPrompt()
	case "aws_customer_gateway":
		resource_prompts.AWSCustomerGatewayPrompt()
	case "aws_default_route_table":
		resource_prompts.AWSDefaultRouteTablePrompt()
	case "aws_s3_bucket":
		resource_prompts.AWSS3BucketPrompt()
	case "aws_s3_access_point":
		resource_prompts.AWSS3AccessPointPrompt()
	case "aws_s3_account_public_access_block":
		resource_prompts.AWSS3AccountPublicAccessBlockPrompt()
	case "aws_s3_bucket_analytics_configuration":
		resource_prompts.AWSS3BucketAnalyticsConfigurationPrompt()
	case "aws_s3_bucket_metric":
		resource_prompts.AWSS3BucketMetricPrompt()
	case "aws_s3_bucket_notification":
		resource_prompts.AWSS3BucketNotificationPrompt()
	case "aws_s3_bucket_object":
		resource_prompts.AWSS3BucketObjectPrompt()
	case "aws_s3_bucket_ownership_controls":
		resource_prompts.AWSS3BucketOwnershipControlsPrompt()
	case "aws_s3_bucket_public_access_block":
		resource_prompts.AWSS3BucketPublicAccessBlockPrompt()
	case "aws_ami":
		resource_prompts.AWSAMIPrompt()
	case "aws_ami_copy":
		resource_prompts.AWSAMICopyPrompt()
	case "aws_ebs_snapshot":
		resource_prompts.AWSEBSSnapshotPrompt()
	case "aws_ebs_snapshot_copy":
		resource_prompts.AWSEBSSnapshotCopyPrompt()
	case "aws_ebs_default_kms_key":
		resource_prompts.AWSEBSDefaultKMSKeyPrompt()
	case "aws_ebs_encryption_by_default":
		resource_prompts.AWSEBSEncryptionByDefaultPrompt()
	case "aws_ami_from_instance":
		resource_prompts.AWSAMIFromInstancePrompt()
	case "aws_ami_launch_permission":
		resource_prompts.AWSAMILaunchPermissionPrompt()
	case "aws_elb":
		resource_prompts.AWSELBPrompt()
	case "aws_db_instance":
		resource_prompts.AWSDBInstancePrompt()
	case "aws_db_cluster_snapshot":
		resource_prompts.AWSDBClusterSnapshotPrompt()
	case "aws_db_event_subscription":
		resource_prompts.AWSDBEventSubscriptionPrompt()
	case "aws_db_instance_role_association":
		resource_prompts.AWSDBInstanceRoleAssociationPrompt()
	case "aws_db_option_group":
		resource_prompts.AWSDBOptionGroupPrompt()
	case "aws_ebs_volume":
		resource_prompts.AWSEBSVolumePrompt()
	case "aws_ec2_availability_zone_group":
		resource_prompts.AWSEC2AvailabilityZoneGroupPrompt()
	case "aws_ec2_capacity_reservation":
		resource_prompts.AWSEC2CapacityReservationPrompt()
	case "aws_ec2_carrier_gateway":
		resource_prompts.AWSEC2CarrierGatewayPrompt()
	case "aws_ec2_client_vpn_authorization_rule":
		resource_prompts.AWSEC2ClientVPNAuthorizationRulePrompt()
	case "aws_ec2_client_vpn_route":
		resource_prompts.AWSEC2ClientVPNRoutePrompt()
	case "aws_ec2_local_gateway_route_table_vpc_association":
		resource_prompts.AWSEC2LocalGatewayRouteTableVPCAssociationPrompt()
	case "aws_ec2_tag":
		resource_prompts.AWSEC2TagPrompt()
	case "aws_ec2_fleet":
		resource_prompts.AWSEC2FleetPrompt()
	case "aws_ec2_local_gateway_route":
		resource_prompts.AWSEC2LocalGatewayRoutePrompt()
	case "aws_ec2_client_vpn_endpoint":
		resource_prompts.AWSEC2ClientVPNEndpointPrompt()
	case "aws_ec2_client_vpn_network_association":
		resource_prompts.AWSEC2ClientVPNNetworkAssociationPrompt()
	case "aws_ec2_traffic_mirror_filter":
		resource_prompts.AWSEC2TrafficMirrorFilterPrompt()
	case "aws_ec2_traffic_mirror_filter_rule":
		resource_prompts.AWSEC2TrafficMirrorFilterRulePrompt()
	case "aws_ec2_traffic_mirror_session":
		resource_prompts.AWSEC2TrafficMirrorSessionPrompt()
	case "aws_ec2_traffic_mirror_target":
		resource_prompts.AWSEC2TrafficMirrorTargetPrompt()
	case "aws_ec2_transit_gateway":
		resource_prompts.AWSEC2TransitGatewayPrompt()
	case "aws_ec2_transit_gateway_peering_attachment":
		resource_prompts.AWSEC2TransitGatewayPeeringAttachmentPrompt()
	case "aws_ec2_transit_gateway_peering_attachment_accepter":
		resource_prompts.AWSEC2TransitGatewayPeeringAttachmentAccepterPrompt()
	case "aws_ec2_transit_gateway_route":
		resource_prompts.AWSEC2TransitGatewayRoutePrompt()
	case "aws_ec2_transit_gateway_route_table":
		resource_prompts.AWSEC2TransitGatewayRouteTablePrompt()
	case "aws_ec2_transit_gateway_route_table_association":
		resource_prompts.AWSEC2TransitGatewayRouteTableAssociationPrompt()
	case "aws_ec2_transit_gateway_route_table_propagation":
		resource_prompts.AWSEC2TransitGatewayRouteTablePropagationPrompt()
	case "aws_db_parameter_group":
		resource_prompts.AWSDBParameterGroupPrompt()
	case "aws_ec2_transit_gateway_vpc_attachment_accepter":
		resource_prompts.AWSEC2TransitGatewayVPCAttachmentAccepterPrompt()
	case "aws_eip":
		resource_prompts.AWSEIPPrompt()
	case "aws_eip_association":
		resource_prompts.AWSEIPAssociationPrompt()
	case "aws_ec2_transit_gateway_vpc_attachment":
		resource_prompts.AWSEC2TransitGatewayVPCAttachmentPrompt()
	case "aws_instance":
		resource_prompts.AWSInstancePrompt()
	case "aws_key_pair":
		resource_prompts.AWSKeyPairPrompt()
	case "aws_placement_group":
		resource_prompts.AWSPlacementGroupPrompt()
	case "aws_snapshot_create_volume_permission":
		resource_prompts.AWSSnapshotCreateVolumePermissionPrompt()
	case "aws_spot_datafeed_subscription":
		resource_prompts.AWSSpotDatafeedSubscriptionPrompt()
	case "aws_db_proxy":
		resource_prompts.AWSDbProxyPrompt()
	case "aws_db_proxy_default_target_group":
		resource_prompts.AWSDBProxyDefaultTargetGroupPrompt()
	case "aws_db_proxy_target":
		resource_prompts.AWSDBProxyTargetPrompt()
	case "aws_db_security_group":
		resource_prompts.AWSDBSecurityGroupPrompt()
	case "aws_db_snapshot":
		resource_prompts.AWSDBSnapshotPrompt()
	case "aws_db_subnet_group":
		resource_prompts.AWSDBSubnetGroupPrompt()
	case "aws_acm_certificate":
		resource_prompts.AWSACMCertificatePrompt()
	case "aws_acmpca_certificate_authority":
		resource_prompts.AWSACMPCACertificatePrompt()
	case "aws_api_gateway_account":
		resource_prompts.AWSAPIGatewayAccountPrompt()
	case "aws_api_gateway_api_key":
		resource_prompts.AWSAPIGatewayApiKeyPrompt()
	case "aws_api_gateway_authorizer":
		resource_prompts.AWSAPiGatewayAuthorizer()
	case "aws_api_gateway_base_path_mapping":
		resource_prompts.AWSAPIGatewayBasePathMappingPrompt()
	case "aws_api_gateway_client_certificate":
		resource_prompts.AWSAPIGatewayClientCertificatePrompt()
	case "aws_api_gateway_deployment":
		resource_prompts.AWSAPIGatewayDeploymentPrompt()
	case "aws_api_gateway_documentation_part":
		resource_prompts.AWSAPIGatewayDocumentationPartPrompt()
	case "aws_api_gateway_documentation_version":
		resource_prompts.AWSAPIGatewayDocumentationVersionPrompt()
	case "aws_api_gateway_domain_name":
		resource_prompts.AWSAPIGatewayDomainNamePrompt()
	case "aws_api_gateway_gateway_response":
		resource_prompts.AWSAPIGatewayGatewayResponsePrompt()
	case "aws_lambda_alias":
		resource_prompts.AWSLambdaAliasPrompt()
	case "aws_lambda_code_signing_config":
		resource_prompts.AWSLambdaCodeSigningConfigPrompt()
	case "aws_lambda_event_source_mapping":
		resource_prompts.AWSLambdaEventSourceMappingPrompt()
	case "aws_lambda_function":
		resource_prompts.AWSLambdaFunctionPrompt()
	case "aws_lambda_layer_version":
		resource_prompts.AWSLambdaLayerVersionPrompt()
	case "aws_lambda_permission":
		resource_prompts.AWSLambdaPermissionPrompt()
	case "aws_lambda_provisioned_concurrency_config":
		resource_prompts.AWSLambdaProvisionedConcurrencyConfigPrompt()
	case "aws_elastic_beanstalk_application":
		resource_prompts.AWSElasticBeanstalkApplicationPrompt()
	case "aws_elastic_beanstalk_application_version":
		resource_prompts.AWSElasticBeanstalkApplicationVersionPrompt()
	case "aws_elastic_beanstalk_configuration_template":
		resource_prompts.AWSElasticBeanstalkConfigurationTemplatePrompt()
	default:
		color.Red("No support added yet for your resource! Coming soon...")
		color.Yellow("\nMeanwhile try getting the boilerplate version by running `tf resource -p aws -r " + resource + " -b`")
	}
}
