package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSDBInstancePrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	selects := map[string]types.TfSelect{}

	var promptOrder, selectOrder []string

	prompts["allocated_storage"] = types.TfPrompt{
		Label: "Enter allocated_storage:\n(Required unless a snapshot_identifier or replicate_source_db is provided) \n" +
			"The allocated storage in gibibytes. If max_allocated_storage is configured, \n" +
			"this argument represents the initial storage allocation and differences \n" +
			"from the configuration will be ignored automatically when Storage Autoscaling occurs.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "allocated_storage")

	selects["storage_type"] = types.TfSelect{
		Label: "Enter storage_type:\n(Optional) One of \"standard\" (magnetic), \"gp2\" \n" +
			"(general purpose SSD), or \"io1\" (provisioned IOPS SSD). \n" +
			"The default is \"io1\" if iops is specified, \"gp2\" if not.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"standard", "gp2", "io1"},
		},
	}
	selectOrder = append(selectOrder, "storage_type")

	selects["engine"] = types.TfSelect{
		Label: "Enter engine:\n The database engine.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"aurora", "aurora-mysql", "aurora-postgresql", "mariadb", "mysql", "oracle-ee",
				"oracle-se2", "oracle-se1", "oracle-se", "postgres", "sqlserver-ee", "sqlserver-se", "sqlserver-ex", "sqlserver-web"},
		},
	}
	selectOrder = append(selectOrder, "engine")

	prompts["engine_version"] = types.TfPrompt{
		Label: "Enter engine_version:\nThe engine version to use.\n" +
			"Check https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance#engine_version",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "engine_version")

	prompts["instance_class"] = types.TfPrompt{
		Label: "Enter instance_class:\n(Required) The instance type of the RDS instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_class")

	prompts["name"] = types.TfPrompt{
		Label: "Enter password:\nOptional) The name of the database to create when the \n" +
			"DB instance is created. If this parameter is not specified, \n" +
			"no database is created in the DB instance. Note that this \n" +
			"does not apply for Oracle or SQL Server engines\n" +
			"Check https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance.html",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["username"] = types.TfPrompt{
		Label: "Enter username:\n(Required unless a snapshot_identifier or replicate_source_db is provided) \n" +
			"Username for the master DB user.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "username")

	prompts["password"] = types.TfPrompt{
		Label: "Enter password:\n(Required unless a snapshot_identifier or replicate_source_db is provided) \n" +
			"Password for the master DB user. Note that this may show up in logs, \n" +
			"and it will be stored in the state file.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "password")

	prompts["parameter_group_name"] = types.TfPrompt{
		Label: "Enter parameter_group_name:\n(Optional) Name of the DB parameter group to associate.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "parameter_group_name")

	prompts["availability_zone"] = types.TfPrompt{
		Label: "Enter availability_zone:\nThe AZ for the RDS instance",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "availability_zone")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_db_instance", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSDBClusterSnapshotPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder []string

	prompts["db_cluster_identifier"] = types.TfPrompt{
		Label: "Enter db_cluster_identifier:\n(Required) The DB Cluster Identifier from which to take the snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_cluster_identifier")

	prompts["db_cluster_snapshot_identifier"] = types.TfPrompt{
		Label: "Enter db_cluster_snapshot_identifier:\n(Required) The Identifier for the snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_cluster_snapshot_identifier")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A map of tags to assign to the DB cluster.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_db_cluster_snapshot", blockName, promptOrder, nil, resourceBlock)
}

func AWSDBEventSubscriptionPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	selects := map[string]types.TfSelect{}

	var promptOrder, selectOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Optional) The name of the DB event subscription. By default generated by Terraform.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["name_prefix"] = types.TfPrompt{
		Label: "Enter name_prefix:\n(Optional) The name of the DB event subscription. Conflicts with name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name_prefix")

	prompts["sns_topic"] = types.TfPrompt{
		Label: "Enter sns_topic:\n(Required) The SNS topic to send events to.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "sns_topic")

	prompts["source_ids"] = types.TfPrompt{
		Label: "Enter source_ids: e.g. [\"c1\",\"c2\"]\n(Optional) A list of identifiers of the event sources for which " +
			"\nevents will be returned. If not specified, then all sources are " +
			"\nincluded in the response. If specified, a source_type must also be specified.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_ids")

	prompts["event_categories"] = types.TfPrompt{
		Label: "Enter event_categories: e.g. [\"c1\",\"c2\"]\n(Optional) A list of event categories for a SourceType that you want to subscribe to. " +
			"\nSee http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run aws rds describe-event-categories.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "event_categories")

	prompts["enabled"] = types.TfPrompt{
		Label: "Enter enabled:\n(Optional) A boolean flag to enable/disable the subscription. Defaults to true.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "enabled")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g.k1=v1,k2=v2\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects["source_type"] = types.TfSelect{
		Label: "Enter source_type:\n(Optional) The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot. If not set, all sources will be subscribed to.",
		Select: promptui.Select{
			Label: "",
		},
	}
	selectOrder = append(selectOrder, "source_type")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Green("Enter timeouts block:\n" +
		"The timeout block supports the following arguments:" +
		"\n1.create\n2.delete\n3.update")

	timeoutsPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	timeoutsPrompt["create"] = types.TfPrompt{
		Label: "Enter create: e.g. 40m\n(Default 40m) How long to wait for an RDS event notification subscription to be ready.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "create")

	timeoutsPrompt["update"] = types.TfPrompt{
		Label: "Enter update: e.g. 40m\n(Default 40m) How long to wait for an RDS event notification subscription to be updated.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "update")

	timeoutsPrompt["delete"] = types.TfPrompt{
		Label: "Enter delete: e.g. 40m\n(Default 40m) How long to wait for an RDS event notification subscription to be deleted.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "delete")
	selectOrder = append(selectOrder, "timeouts")

	resourceBlock["timeouts"] = builder.NestedPSOrder(nestedPromptOrder, nil, timeoutsPrompt, nil)

	builder.ResourceBuilder("aws_db_event_subscription", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSDBInstanceRoleAssociationPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder []string

	prompts["db_instance_identifier"] = types.TfPrompt{
		Label: "Enter db_instance_identifier:\n(Required) DB Instance Identifier to associate with the IAM Role.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_instance_identifier")

	prompts["feature_name"] = types.TfPrompt{
		Label: "Enter feature_name:\n(Required) Name of the feature for association. This can be found in the " +
			"\nAWS documentation relevant to the integration or a full list is available in the SupportedFeatureNames list returned by AWS CLI rds describe-db-engine-versions." +
			"\nCheckout https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "feature_name")

	prompts["role_arn"] = types.TfPrompt{
		Label: "Enter role_arn:\n(Required) Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "role_arn")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_db_instance_role_association", blockName, promptOrder, nil, resourceBlock)

}

func AWSDBOptionGroupPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder, selectOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Optional, Forces new resource) The name of the option group. " +
			"\nIf omitted, Terraform will assign a random, unique name. Must be lowercase, " +
			"\nto match as it is stored in AWS.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["name_prefix"] = types.TfPrompt{
		Label: "Enter name_prefix:\n(Optional, Forces new resource) Creates a unique name beginning " +
			"\nwith the specified prefix. Conflicts with name. Must be lowercase, " +
			"\nto match as it is stored in AWS.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name_prefix")

	prompts["option_group_description"] = types.TfPrompt{
		Label: "Enter option_group_description:\n(Optional) The description of the option group. Defaults to \"Managed by Terraform\".",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "option_group_description")

	prompts["engine_name"] = types.TfPrompt{
		Label: "Enter engine_name:\n(Required) Specifies the name of the engine that this option group should be associated with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "engine_name")

	prompts["major_engine_version"] = types.TfPrompt{
		Label: "Enter major_engine_version:\n(Required) Specifies the major version of the engine that this option group should be associated with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "major_engine_version")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g. k1=v1,k2=v2\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like options etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_db_option_group", blockName, promptOrder, nil, resourceBlock)
		return
	}

	color.Green("\nEnter option:\nThe option block supports the following arguments:" +
		"\n1.option_name\n2.option_settings(not supported by this cli yet)\n3.port\n4.version\n5.db_security_group_memberships\n6.vpc_security_group_memberships\n")

	optionPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	optionPrompt["option_name"] = types.TfPrompt{
		Label: "Enter option_name:\n(Required) The Name of the Option (e.g. MEMCACHED).",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "option_name")

	optionPrompt["port"] = types.TfPrompt{
		Label: "Enter port:\n(Optional) The Port number when connecting to the Option (e.g. 11211).",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "port")

	optionPrompt["version"] = types.TfPrompt{
		Label: "Enter version:\n(Optional) The version of the option (e.g. 13.1.0.0).",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "version")

	optionPrompt["db_security_group_memberships"] = types.TfPrompt{
		Label: "Enter db_security_group_memberships: e.g.[\"g1\",\"g2\"]\n(Optional) A list of DB Security Groups for which the option is enabled.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "db_security_group_memberships")

	optionPrompt["vpc_security_group_memberships"] = types.TfPrompt{
		Label: "Enter vpc_security_group_memberships: e.g.[\"g1\",\"g2\"]\n(Optional) A list of VPC Security Groups for which the option is enabled.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "vpc_security_group_memberships")
	selectOrder = append(selectOrder, "option")

	resourceBlock["option"] = builder.NestedPSOrder(nestedPromptOrder, nil, optionPrompt, nil)

	builder.ResourceBuilder("aws_db_option_group", blockName, promptOrder, selectOrder, resourceBlock)

}

func AWSDBParameterGroupPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder, selectOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Optional, Forces new resource) The name of the DB parameter group. If omitted, " +
			"\nTerraform will assign a random, unique name.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["name_prefix"] = types.TfPrompt{
		Label: "Enter name_prefix:\n(Optional, Forces new resource) Creates a unique name beginning with the specified prefix. Conflicts with name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name_prefix")

	prompts["family"] = types.TfPrompt{
		Label: "Enter family:\n(Required, Forces new resource) The family of the DB parameter group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "family")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional, Forces new resource) The description of the DB parameter group. Defaults to \"Managed by Terraform\".",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n (Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like options etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_db_parameter_group", blockName, promptOrder, nil, resourceBlock)
		return
	}

	parameterPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	parameterPrompt["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) The name of the DB parameter.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "name")

	parameterPrompt["value"] = types.TfPrompt{
		Label: "Enter value:\n(Required) The value of the DB parameter.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "value")

	parameterPrompt["apply_method"] = types.TfPrompt{
		Label: "Enter apply_method:\n(Optional) \"immediate\" (default), or \"pending-reboot\". Some engines can't apply some parameters without a reboot, and you will need to specify \"pending-reboot\" here",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "apply_method")
	selectOrder = append(selectOrder, "parameter")

	resourceBlock["parameter"] = builder.NestedPSOrder(nestedPromptOrder, nil, parameterPrompt, nil)

	builder.ResourceBuilder("aws_db_parameter_group", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSDbProxyPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder, selectOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) The identifier for the proxy. This name must be unique for all proxies owned by your AWS " +
			"\naccount in the specified AWS Region. An identifier must begin with a letter and must " +
			"\ncontain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["debug_logging"] = types.TfPrompt{
		Label: "Enter debug_logging:\n(Optional) Whether the proxy includes detailed information about SQL statements " +
			"\nin its logs. This information helps you to debug issues involving SQL " +
			"\nbehavior or the performance and scalability of the proxy connections. " +
			"\nThe debug information includes the text of SQL statements that you submit " +
			"\nthrough the proxy. Thus, only enable this setting when needed for debugging, " +
			"\nand only when you have security measures in place to safeguard any sensitive information that appears in the logs.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "debug_logging")

	prompts["idle_client_timeout"] = types.TfPrompt{
		Label: "Enter idle_client_timeout:\n(Optional) The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "idle_client_timeout")

	prompts["require_tls"] = types.TfPrompt{
		Label: "Enter require_tls(true/false):\n(Optional) The number of seconds that a connection to the proxy " +
			"\ncan be inactive before the proxy disconnects it. You can set this " +
			"\nvalue higher or lower than the connection timeout limit for the " +
			"\nassociated database.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "require_tls")

	prompts["role_arn"] = types.TfPrompt{
		Label: "Enter role_arn(true/false):\n(Required) The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "role_arn")

	prompts["vpc_security_group_ids"] = types.TfPrompt{
		Label: "Enter vpc_security_group_ids:\n(Optional) One or more VPC security group IDs to associate with the new proxy.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_security_group_ids")

	prompts["vpc_subnet_ids"] = types.TfPrompt{
		Label: "Enter vpc_subnet_ids:\n(Required) One or more VPC subnet IDs to associate with the new proxy.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_subnet_ids")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g. k1=v1,k2=v2\n(Optional) A mapping of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}

	selects["engine_family"] = types.TfSelect{
		Label: "Enter engine_family:\n(Required, Forces new resource) The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are MYSQL and POSTGRESQL",
		Select: promptui.Select{
			Label: "",
			Items: []string{"MYSQL", "POSTGRESQL"},
		},
	}
	selectOrder = append(selectOrder, "engine_family")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Yellow("\nConfigure nested settings like auth etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_db_proxy", blockName, promptOrder, selectOrder, resourceBlock)
		return
	}

	color.Green("\nEnter auth:\n(Required) Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters" +
		"\nThe auth block supports the following arguments:" +
		"\n1.auth_scheme\n2.description\n3.iam_auth\n4.secret_arn\n5.username\n")

	authPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder, nestedSelectOrder []string

	authPrompt["auth_scheme"] = types.TfPrompt{
		Label: "Enter auth_scheme:\n(Optional) The type of authentication that the proxy uses for connections " +
			"\nfrom the proxy to the underlying database. One of SECRETS",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "auth_scheme")

	authPrompt["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A user-specified description about the authentication used by a proxy to log in as a specific database user.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "description")

	authPrompt["secret_arn"] = types.TfPrompt{
		Label: "Enter secret_arn:\n(Optional) The Amazon Resource Name (ARN) representing the secret that the proxy uses " +
			"\nto authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "secret_arn")

	authPrompt["username"] = types.TfPrompt{
		Label: "Enter username:\n(Optional) The name of the database user to which the proxy connects.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "username")

	authSelect := map[string]types.TfSelect{}

	authSelect["iam_auth"] = types.TfSelect{
		Label: "Enter iam_auth:\n(Optional) Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. One of DISABLED, REQUIRED",
		Select: promptui.Select{
			Label: "",
			Items: []string{"DISABLED", "REQUIRED"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "iam_auth")
	selectOrder = append(selectOrder, "auth")

	resourceBlock["auth"] = builder.NestedPSOrder(nestedPromptOrder, nestedSelectOrder, authPrompt, authSelect)

	builder.ResourceBuilder("aws_db_proxy", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSDBProxyDefaultTargetGroupPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder, selectOrder []string

	prompts["db_proxy_name"] = types.TfPrompt{
		Label: "Enter db_proxy_name:\n(Required) Name of the RDS DB Proxy.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_proxy_name")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	connectionPoolConfigPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	connectionPoolConfigPrompt["connection_borrow_timeout"] = types.TfPrompt{
		Label: "Enter connection_borrow_timeout:\n(Optional) The number of seconds for a proxy to wait for a connection to become " +
			"\navailable in the connection pool. Only applies when the proxy has " +
			"\nopened its maximum number of connections and all connections are busy with client sessions.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "connection_borrow_timeout")

	connectionPoolConfigPrompt["init_query"] = types.TfPrompt{
		Label: "Enter init_query:\n(Optional) One or more SQL statements for the proxy to run when opening each " +
			"\nnew database connection. Typically used with SET statements to make sure that " +
			"\neach connection has identical settings such as time zone and character set. " +
			"\nThis setting is empty by default. For multiple statements, use semicolons as the separator. " +
			"\nYou can also include multiple variables in a single SET statement, such as SET x=1, y=2.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "init_query")

	connectionPoolConfigPrompt["max_connections_percent"] = types.TfPrompt{
		Label: "Enter max_connections_percent:\n(Optional) The maximum size of the connection pool for each target in a " +
			"\ntarget group. For Aurora MySQL, it is expressed as a " +
			"\npercentage of the max_connections setting for the RDS DB " +
			"\ninstance or Aurora DB cluster used by the target group.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "max_connections_percent")

	connectionPoolConfigPrompt["max_idle_connections_percent"] = types.TfPrompt{
		Label: "Enter max_idle_connections_percent:\n(Optional) The maximum size of the connection pool for each target in a " +
			"\ntarget group. For Aurora MySQL, it is expressed as a " +
			"\npercentage of the max_connections setting for the RDS DB " +
			"\ninstance or Aurora DB cluster used by the target group.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "max_idle_connections_percent")

	connectionPoolConfigPrompt["session_pinning_filters"] = types.TfPrompt{
		Label: "Enter session_pinning_filters:\n(Optional) Each item in the list represents a class of SQL operations " +
			"\nthat normally cause all later statements in a session using " +
			"\na proxy to be pinned to the same underlying database connection. Including an item in the list exempts that class of SQL operations from the pinning behavior. Currently, the only allowed value is EXCLUDE_VARIABLE_SETS",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "session_pinning_filters")
	selectOrder = append(selectOrder, "connection_pool_config")

	resourceBlock["connection_pool_config"] = builder.NestedPSOrder(nestedPromptOrder, nil, connectionPoolConfigPrompt, nil)

	builder.ResourceBuilder("aws_db_proxy_default_target_group", blockName, promptOrder, selectOrder, resourceBlock)

}

func AWSDBProxyTargetPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder []string

	prompts["db_proxy_name"] = types.TfPrompt{
		Label: "Enter db_proxy_name:\n(Required, Forces new resource) The name of the DB proxy.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_proxy_name")

	prompts["target_group_name"] = types.TfPrompt{
		Label: "Enter target_group_name:\n(Required, Forces new resource) The name of the target group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "target_group_name")

	color.Yellow("\nEither db_instance_identifier or db_cluster_identifier should be specified and both should not be specified together\n")

	prompts["db_instance_identifier"] = types.TfPrompt{
		Label: "Enter db_instance_identifier:\n(Optional, Forces new resource) DB instance identifier.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_instance_identifier")

	prompts["db_cluster_identifier"] = types.TfPrompt{
		Label: "Enter db_cluster_identifier:\n(Optional, Forces new resource) DB cluster identifier.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_cluster_identifier")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_db_proxy_target", blockName, promptOrder, nil, resourceBlock)

}

func AWSDBSecurityGroupPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder, selectOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) The name of the DB security group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) The description of the DB security group. Defaults to \"Managed by Terraform\".",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g.k1=v1,k2=v2\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter ingress:\n(Required) A list of ingress rules." +
		"\nThe auth block supports the following arguments:" +
		"\n1.cidr\n2.security_group_name\n3.security_group_id\n4.security_group_owner_id\n")

	ingressPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	ingressPrompt["cidr"] = types.TfPrompt{
		Label: "Enter cidr:\n  The CIDR block to accept",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cidr")

	ingressPrompt["security_group_name"] = types.TfPrompt{
		Label: "Enter security_group_name:\nThe name of the security group to authorize",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "security_group_name")

	ingressPrompt["security_group_id"] = types.TfPrompt{
		Label: "Enter security_group_id:\nThe ID of the security group to authorize",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "security_group_id")

	ingressPrompt["security_group_owner_id"] = types.TfPrompt{
		Label: "Enter security_group_owner_id:\nThe owner Id of the security group provided by security_group_name.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "security_group_owner_id")
	selectOrder = append(selectOrder, "ingress")

	resourceBlock["ingress"] = builder.NestedPSOrder(nestedPromptOrder, nil, ingressPrompt, nil)

	builder.ResourceBuilder("aws_db_security_group", blockName, promptOrder, selectOrder, resourceBlock)

}

func AWSDBSnapshotPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder []string

	prompts["db_instance_identifier"] = types.TfPrompt{
		Label: "Enter db_instance_identifier:\n(Required) The DB Instance Identifier from which to take the snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_instance_identifier")

	prompts["db_snapshot_identifier"] = types.TfPrompt{
		Label: "Enter db_snapshot_identifier:\n(Required) The Identifier for the snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "db_snapshot_identifier")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) Key-value map of resource tags",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_db_snapshot", blockName, promptOrder, nil, resourceBlock)
}

func AWSDBSubnetGroupPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}

	var promptOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Optional, Forces new resource) The name of the DB subnet group. " +
			"\nIf omitted, Terraform will assign a random, unique name.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["name_prefix"] = types.TfPrompt{
		Label: "Enter name_prefix:\n(Optional, Forces new resource) Creates a unique name beginning with the specified prefix. " +
			"\nConflicts with name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name_prefix")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) The description of the DB subnet group. Defaults to \"Managed by Terraform\".",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["subnet_ids"] = types.TfPrompt{
		Label: "Enter subnet_ids: e.g.[\"id1\",\"id2\"]\n(Required) A list of VPC subnet IDs.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_ids")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g. k1=v1,k2=v2\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_db_subnet_group", blockName, promptOrder, nil, resourceBlock)
}
