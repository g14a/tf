package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSElasticBeanstalkApplicationPrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) The name of the application, must be unique within your account",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Short description of the application",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g.k1=v1,k2=v2:\n(Optional) Key-value map of tags for the Elastic Beanstalk Application.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}

	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter appversion_lifecycle:\nappversion_lifecycle supports the following arguments:" +
		"\n1.service_role\n2.max_count\n3.max_age_in_days\n4.delete_source_from_s3\n")

	var nestedPromptOrder []string
	appVersionLifecyclePrompt := map[string]types.TfPrompt{}

	appVersionLifecyclePrompt["service_role"] = types.TfPrompt{
		Label: "Enter service_role:\n(Required) The ARN of an IAM service role under which the application " +
			"\nversion is deleted. Elastic Beanstalk must have permission to assume this role.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "service_role")

	appVersionLifecyclePrompt["max_count"] = types.TfPrompt{
		Label: "Enter max_count:\n(Optional) The maximum number of application versions to retain ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "max_count")

	appVersionLifecyclePrompt["max_age_in_days"] = types.TfPrompt{
		Label: "Enter max_age_in_days:\n(Optional) The number of days to retain an application version ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "max_age_in_days")

	appVersionLifecyclePrompt["delete_source_from_s3"] = types.TfPrompt{
		Label: "Enter delete_source_from_s3:(true/false)\n(Optional) Set to true to delete a version's source bundle from S3 when the application version is deleted.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "delete_source_from_s3")

	resourceBlock["appversion_lifecycle"] = builder.PSOrder(nestedPromptOrder, nil, appVersionLifecyclePrompt, nil)

	builder.ResourceBuilder("aws_elastic_beanstalk_application", blockName, resourceBlock)

}

func AWSElasticBeanstalkApplicationVersionPrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) A unique name for the this Application Version.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["application"] = types.TfPrompt{
		Label: "Enter application:\n(Required) Name of the Beanstalk Application the version is associated with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "application")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Short description of the Application Version.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["bucket"] = types.TfPrompt{
		Label: "Enter bucket:\n(Required) S3 bucket that contains the Application Version source bundle.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	prompts["key"] = types.TfPrompt{
		Label: "Enter key:\n(Required) S3 object that is the Application Version source bundle.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "key")

	prompts["force_delete"] = types.TfPrompt{
		Label: "Enter force_delete:\n(Optional) On delete, force an Application Version to be deleted " +
			"\nwhen it may be in use by multiple Elastic Beanstalk Environments.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "force_delete")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g.k1=v1,k2=v2\nKey-value map of tags for the Elastic Beanstalk Application Version.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_elastic_beanstalk_application_version", blockName, resourceBlock)

}

func AWSElasticBeanstalkConfigurationTemplatePrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) A unique name for this Template.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["application"] = types.TfPrompt{
		Label: "Enter application:\n(Required) name of the application to associate with this configuration template",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "application")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Short description of the Template",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["environment_id"] = types.TfPrompt{
		Label: "Enter environment_id:\n(Optional) The ID of the environment used with this configuration template",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "environment_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter setting:\nThe setting block supports the following format:" +
		"\n1.namespace\n2.name\n3.value\n4.resource\n")

	settingPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	settingPrompt["namespace"] = types.TfPrompt{
		Label: "Enter namespace:\nunique namespace identifying the option's associated AWS resource",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "namespace")

	settingPrompt["name"] = types.TfPrompt{
		Label: "Enter name:\nname of the configuration option",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "name")

	settingPrompt["value"] = types.TfPrompt{
		Label: "Enter value:\nvalue of the configuration option",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "value")

	settingPrompt["resource"] = types.TfPrompt{
		Label: "Enter resource:\n(Optional) resource name for scheduled action." +
			"\nCheckout https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "resource")

	resourceBlock["setting"] = builder.PSOrder(nestedPromptOrder, nil, settingPrompt, nil)
	builder.ResourceBuilder("aws_elastic_beanstalk_configuration_template", blockName, resourceBlock)
}
