package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
)

func AWSACMCertificatePrompt() {
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

	prompts["domain_name"] = types.TfPrompt{
		Label: "Enter domain_name:\n(Required) A domain name for which the certificate should be issued",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "domain_name")

	prompts["private_key"] = types.TfPrompt{
		Label: "Enter private_key:\n(Required) The certificate's PEM-formatted private key",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "private_key")

	prompts["certificate_body"] = types.TfPrompt{
		Label: "Enter certificate_body:\n(Required) The certificate's PEM-formatted public key",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_body")

	prompts["certificate_chain"] = types.TfPrompt{
		Label: "Enter certificate_chain:\n(Optional) The certificate's PEM-formatted chain",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_chain")

	prompts["certificate_authority_arn"] = types.TfPrompt{
		Label: "Enter certificate_authority_arn:\n(Required) ARN of an ACMPCA",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_authority_arn")

	prompts["subject_alternative_names"] = types.TfPrompt{
		Label: "Enter subject_alternative_names:\n(Optional) Set of domains that should be SANs \n" +
			"in the issued certificate. To remove all elements of a previously \n" +
			"configured list, set this value equal to an empty list ([]) \n" +
			"or use the terraform taint command to trigger recreation.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subject_alternative_names")

	selects["validation_method"] = types.TfSelect{
		Label: "Enter validation_method:\n(Required) Which method to use for validation. \n" +
			"DNS or EMAIL are valid, NONE can be used for certificates that were \n" +
			"imported into ACM and then into Terraform.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"DNS", "EMAIL", "NONE"},
		},
	}
	selectOrder = append(selectOrder, "validation_method")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Yellow("\nConfigure nested settings like lifecycle/tags etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket", blockName, resourceBlock)
		return
	}

	lifecyclePrompt := map[string]types.TfPrompt{}
	var nestedOrder []string

	color.Green("Enter lifecycle block:\n")

	lifecyclePrompt["create_before_destroy"] = types.TfPrompt{
		Label: "Enter create_before_destroy:(true/false)\nBy default, when Terraform must change a resource argument \n" +
			"that cannot be updated in-place due to remote API limitations, \n" +
			"Terraform will instead destroy the existing object and then \n" +
			"create a new replacement object with the new configured arguments.\n" +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#create_before_destroy",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "create_before_destroy")

	lifecyclePrompt["prevent_destroy"] = types.TfPrompt{
		Label: "Enter prevent_destroy:(true/false)\nThis meta-argument, when set to true, will cause Terraform to \n" +
			"reject with an error any plan that would destroy the infrastructure \n" +
			"object associated with the resource, as long as the argument \n" +
			"remains present in the configuration.\n" +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#prevent_destroy",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "prevent_destroy")

	lifecyclePrompt["ignore_changes"] = types.TfPrompt{
		Label: "Enter ignore_changes: e.g.[\"c1\",\"c2\"]\nBy default, Terraform detects any difference in the " +
			"current settings of a real infrastructure object and plans to " +
			"update the remote object to match configuration." +
			"Check https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#ignore_changes",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "ignore_changes")
	selectOrder = append(selectOrder, "lifecycle")

	resourceBlock["lifecycle"] = builder.PSOrder(nestedOrder, nil, lifecyclePrompt, nil)

	color.Green("\nEnter tags:\n")

	tagsPrompt := map[string]types.TfPrompt{}
	tagsPrompt["Environment"] = types.TfPrompt{
		Label: "Enter Environment:\n",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "Environment")
	selectOrder = append(selectOrder, "tags")

	resourceBlock["tags"] = builder.PSOrder(nestedOrder[len(nestedOrder)-1:], nil, tagsPrompt, nil)

	builder.ResourceBuilder("aws_acm_certificate", blockName, resourceBlock)
}

func AWSACMPCACertificatePrompt() {
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

	prompts["permanent_deletion_time_in_days"] = types.TfPrompt{
		Label: "Enter permanent_deletion_time_in_days\n(Optional) The number of days to make a CA restorable after \n" +
			"it has been deleted, must be between 7 to 30 days, with default to 30 days.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}

	promptOrder = append(promptOrder, "permanent_deletion_time_in_days")

	prompts["enabled"] = types.TfPrompt{
		Label: "Enter enabled\n(Optional) Whether the certificate authority is enabled or disabled. Defaults to true",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}

	promptOrder = append(promptOrder, "enabled")

	selects["type"] = types.TfSelect{
		Label: "Enter type:\n(Optional) The type of the certificate authority. Defaults to SUBORDINATE. Valid values: ROOT and SUBORDINATE",
		Select: promptui.Select{
			Label: "",
			Items: []string{"ROOT", "SUBORDINATE"},
		},
	}
	selectOrder = append(selectOrder, "type")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_acmpca_certificate_authority", blockName, resourceBlock)

}
