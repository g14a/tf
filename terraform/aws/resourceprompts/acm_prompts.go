package resourceprompts

import (
	"fmt"

	"github.com/g14a/tf/utils"

	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/manifoldco/promptui"
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

	schema := []types.Schema{
		{
			Field: "domain_name",
			Ex:    "",
			Doc:   "(Required) A domain name for which the certificate should be issued",
		},
		{
			Field: "subject_alternative_names",
			Ex:    "",
			Doc:   "(Optional) Set of domains that should be SANs in the issued certificate. To remove all elements of a previously configured list, set this value equal to an empty list ([]) or use the terraform taint command to trigger recreation.",
		},
		{
			Type:  "select",
			Field: "validation_method",
			Doc:   "(Required) Which method to use for validation. DNS or EMAIL are valid, NONE can be used for certificates that were imported into ACM and then into Terraform.",
			Items: []string{"DNS", "EMAIL", "NONE"},
		},
		{
			Field: "private_key",
			Ex:    "",
			Doc:   "(Required) The certificate's PEM-formatted private key",
		},
		{
			Field: "certificate_body",
			Ex:    "",
			Doc:   "(Required) The certificate's PEM-formatted public key",
		},
		{
			Field: "certificate_chain",
			Ex:    "",
			Doc:   "(Optional) The certificate's PEM-formatted chain",
		},
		{
			Field: "certification_authority_arn",
			Ex:    "",
			Doc:   "(Required) ARN of an ACMPCA",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: utils.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like options [y/n]?\n\n", "text")

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

	optionsSchema := []types.Schema{
		{
			Type:  "select",
			Field: "certificate_transparency_logging_preference",
			Doc: "(Optional) Specifies whether certificate details should be added to a certificate transparency log. " +
				"\nCheckout https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency",
			Items: []string{"ENABLED", "DISABLED"},
		},
	}
	resourceBlock["options"] = builder.PSOrder(types.ProvidePS(optionsSchema))

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

func AWSACMCertificationValidationPrompt() {
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

	prompts["certificate_arn"] = types.TfPrompt{
		Label: "Enter certificate_arn:\n(Required) The ARN of the certificate that is being validated.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_arn")

	prompts["validation_record_fqdns"] = types.TfPrompt{
		Label: "Enter validation_record_fqdns:\n(Optional) List of FQDNs that implement the validation. Only valid for DNS validation method " +
			"\nACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the " +
			"\nresource that is implementing the validation",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "validation_record_fqdns")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_acm_certificate_validation", blockName, resourceBlock)
}
