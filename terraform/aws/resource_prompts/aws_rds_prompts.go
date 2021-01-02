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
			Label:    "",
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
