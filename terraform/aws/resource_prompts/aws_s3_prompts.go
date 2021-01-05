package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSS3AccessPointPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["bucket"] = types.TfPrompt{
		Label: "Enter bucket:\n(Required) The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) " +
			"\nof S3 on Outposts Bucket that you want to associate this access point with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) The name you want to assign to this access point.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["account_id"] = types.TfPrompt{
		Label: "Enter account_id:\n(Optional) The AWS account ID for the owner of the bucket for which " +
			"\nyou want to create an access point. Defaults to automatically determined account ID of the Terraform AWS provider.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "account_id")

	prompts["policy"] = types.TfPrompt{
		Label: "Enter policy:\n(Optional) A valid JSON document that specifies the policy that you want to apply to this access point.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "policy")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like public_access_block_configuration/vpc_configuration etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_access_point", blockName, promptOrder, nil, resourceBlock)
		return
	}

	publicAccessBlockConfigurationPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	publicAccessBlockConfigurationPrompt["block_public_acls"] = types.TfPrompt{
		Label: "Enter block_public_acls(true/false):\nOptional) Whether Amazon S3 should block public ACLs for buckets in this account. " +
			"\nDefaults to true. Enabling this setting does not affect existing policies or ACLs. " +
			"\nWhen set to true causes the following behavior:\n\n    " +
			"PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n    " +
			"PUT Object calls fail if the request includes a public ACL.\n    " +
			"PUT Bucket calls fail if the request includes a public ACL.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "block_public_acls")

	publicAccessBlockConfigurationPrompt["block_public_policy"] = types.TfPrompt{
		Label: "Enter block_public_policy(true/false):\n(Optional) Whether Amazon S3 should block public bucket policies for buckets in this account. " +
			"\nDefaults to true. Enabling this setting does not affect existing bucket policies. " +
			"\nWhen set to true causes Amazon S3 to:\n\n    Reject calls to PUT Bucket policy if the specified bucket policy allows public access.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "block_public_policy")

	publicAccessBlockConfigurationPrompt["ignore_public_acls"] = types.TfPrompt{
		Label: "Enter ignore_public_acls(true/false):\n(Optional) Whether Amazon S3 should ignore public ACLs for buckets in this account. " +
			"\nDefaults to true. Enabling this setting does not affect the persistence of any " +
			"\nexisting ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:\n\n    Ignore all public ACLs on buckets in this account and any objects that they contain.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ignore_public_acls")

	publicAccessBlockConfigurationPrompt["restrict_public_buckets"] = types.TfPrompt{
		Label: "Enter restrict_public_buckets(true/false):\n(Optional) Whether Amazon S3 should restrict public bucket policies for buckets " +
			"\nin this account. Defaults to true. Enabling this setting does not affect previously stored " +
			"\nbucket policies, except that public and cross-account access within any public bucket policy, " +
			"\nincluding non-public delegation to specific accounts, is blocked. When set to true:\n\n    Only the bucket owner and AWS Services can access buckets with public policies.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "restrict_public_buckets")
	selectOrder = append(selectOrder, "public_access_block_configuration")

	resourceBlock["public_access_block_configuration"] = builder.NestedPSOrder(nestedPromptOrder, nil, publicAccessBlockConfigurationPrompt, nil)

	vpcConfigPrompt := map[string]types.TfPrompt{}

	vpcConfigPrompt["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) This access point will only allow connections from the specified VPC ID.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "vpc_id")
	selectOrder = append(selectOrder, "vpc_configuration")

	resourceBlock["vpc_configuration"] = builder.NestedPSOrder(nestedPromptOrder[len(nestedPromptOrder)-1:], nil, vpcConfigPrompt, nil)

	builder.ResourceBuilder("aws_s3_access_point", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSS3AccountPublicAccessBlockPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder []string

	prompts["account_id"] = types.TfPrompt{
		Label: "Enter account_id:\n(Optional) AWS account ID to configure. Defaults to " +
			"\nautomatically determined account ID of the Terraform AWS provider.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "account_id")

	prompts["block_public_acls"] = types.TfPrompt{
		Label: "Enter block_public_acls:\n(Optional) Whether Amazon S3 should block public ACLs for buckets in this account. " +
			"\nDefaults to false. Enabling this setting does not affect existing policies or ACLs. " +
			"\nWhen set to true causes the following behavior:\n\n    PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.\n    PUT Object calls will fail if the request includes an object ACL.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "block_public_acls")

	prompts["block_public_policy"] = types.TfPrompt{
		Label: "Enter block_public_policy:\n(Optional) Whether Amazon S3 should block public bucket policies for buckets in this account. " +
			"\nDefaults to false. Enabling this setting does not affect existing bucket policies. " +
			"\nWhen set to true causes Amazon S3 to:\n\n    Reject calls to PUT Bucket policy if the specified bucket policy allows public access.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "block_public_policy")

	prompts["ignore_public_acls"] = types.TfPrompt{
		Label: "Enter ignore_public_acls:\n(Optional) Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to false. " +
			"\nEnabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new " +
			"\npublic ACLs from being set. When set to true causes Amazon S3 to:\n\n    Ignore all public ACLs on buckets in this account and any objects that they contain.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ignore_public_acls")

	prompts["restrict_public_buckets"] = types.TfPrompt{
		Label: "Enter restrict_public_buckets:\n(Optional) Whether Amazon S3 should restrict public bucket policies for buckets in this account. " +
			"\nDefaults to false. Enabling this setting does not affect previously stored bucket policies, " +
			"\nexcept that public and cross-account access within any public bucket policy, " +
			"\nincluding non-public delegation to specific accounts, is blocked. When set to true:\n\n    Only the bucket owner and AWS Services can access buckets with public policies.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "restrict_public_buckets")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_s3_account_public_access_block", blockName, promptOrder, nil, resourceBlock)
}

func AWSS3BucketPrompt() {

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder []string

	prompts["bucket"] = types.TfPrompt{
		Label: "The name of the bucket. If omitted, Terraform will assign a random, unique name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	selects := map[string]types.TfSelect{}
	var selectOrder []string

	selects["acl"] = types.TfSelect{
		Label: "Enter acl:\nThe canned ACL to apply",
		Select: promptui.Select{
			Label: "",
			Items: []string{"private", "public-read", "public-read-write", "aws-exec-read", "authenticated-read", "log-delivery-write"},
		},
	}
	selectOrder = append(selectOrder, "acl")

	selects["force_destroy"] = types.TfSelect{
		Label: "Enter force_destroy:\n(Optional, Default:false) A boolean that indicates all objects \n" +
			"(including any locked objects) should be deleted from the bucket \n" +
			"so that the bucket can be destroyed without error. These objects are not recoverable.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"true", "false"},
		},
	}
	selectOrder = append(selectOrder, "force_destroy")

	selects["acceleration_status"] = types.TfSelect{
		Label: "Enter acceleration_status:\n(Optional) Sets the accelerate " +
			"configuration of an existing bucket. Can be Enabled or Suspended",
		Select: promptui.Select{
			Label: "",
			Items: []string{"Enabled", "Suspended"},
		},
	}
	selectOrder = append(selectOrder, "acceleration_status")
	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Yellow("\nConfigure nested settings like cors_rule/versioning/website etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket", blockName, promptOrder, selectOrder, resourceBlock)
		return
	}

	corsRulePrompt := map[string]types.TfPrompt{}
	var nestedOrder []string

	color.Green("\nEnter cors_rule (Optional) A rule of Cross-Origin Resource Sharing :\n\n")

	corsRulePrompt["allowed_headers"] = types.TfPrompt{
		Label: "Enter allowed_headers:\n(Optional) Specifies which headers are allowed",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "allowed_headers")

	corsRulePrompt["allowed_methods"] = types.TfPrompt{
		Label: "Enter allowed_methods:\nRequired) Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "allowed_methods")

	corsRulePrompt["allowed_origins"] = types.TfPrompt{
		Label: "Enter allowed_origins:\n(Required) Specifies which origins are allowed.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "allowed_origins")

	corsRulePrompt["exposed_headers"] = types.TfPrompt{
		Label: "Enter exposed_headers:\n(Optional) Specifies expose header in the response.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "exposed_headers")

	corsRulePrompt["max_age_seconds"] = types.TfPrompt{
		Label: "Enter max_age_seconds:\n(Optional) Specifies time in seconds " +
			"that browser can cache the response for a preflight request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "max_age_seconds")
	selectOrder = append(selectOrder, "cors_rule")

	resourceBlock["cors_rule"] = builder.NestedPSOrder(nestedOrder, nil, corsRulePrompt, nil)

	color.Green("\nEnter website:\nThe website object supports the following:" +
		"\n1.index_document\n2.error_document\n3.redirect_all_requests_to\n4.routing_rules\n\n")

	websitePrompt := map[string]types.TfPrompt{}

	websitePrompt["index_document"] = types.TfPrompt{
		Label: "Enter index_document:\n(Required), unless using redirect_all_requests_to) Amazon S3\n" +
			" returns this index document when requests are made to the" +
			" root domain or any of the subfolders.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "index_document")

	websitePrompt["error_document"] = types.TfPrompt{
		Label: "Enter error_document:\n(Optional) An absolute path to the document to return in case of a 4XX error.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "error_document")

	websitePrompt["redirect_all_requests_to"] = types.TfPrompt{
		Label: "Enter redirect_all_requests_to:\nOptional) A hostname to redirect all website requests for \n" +
			"this bucket to. Hostname can optionally be prefixed with a \n" +
			"protocol (http:// or https://) to use when redirecting requests. \n" +
			"The default is the protocol that is used in the original request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "redirect_all_requests_to")

	websitePrompt["routing_rules"] = types.TfPrompt{
		Label: "Enter routing_rules:\n(Optional) A json array containing routing rules describing redirect behavior and when redirects are applied.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "routing_rules")
	selectOrder = append(selectOrder, "website")

	resourceBlock["website"] = builder.NestedPSOrder(nestedOrder[len(nestedOrder)-4:], nil, websitePrompt, nil)

	versioningPrompt := map[string]types.TfPrompt{}
	versioningPrompt["enabled"] = types.TfPrompt{
		Label: "Enter enabled:(true/false)(Optional) A state of versioning",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "enabled")

	versioningPrompt["mfa_delete"] = types.TfPrompt{
		Label: "Enter mfa_delete:\n(Optional) Enable MFA delete for either Change the versioning\n" +
			"state of your bucket or Permanently delete an object version.\n" +
			"Default is false. This cannot be used to toggle this setting but is \n" +
			"available to allow managed buckets to reflect the state in AWS.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "mfa_delete")
	selectOrder = append(selectOrder, "versioning")

	resourceBlock["versioning"] = builder.NestedPSOrder(nestedOrder[len(nestedOrder)-2:], nil, versioningPrompt, nil)

	loggingPrompt := map[string]types.TfPrompt{}

	loggingPrompt["target_bucket"] = types.TfPrompt{
		Label: "Enter target_bucket:\n(Required) The name of the bucket that will receive the log objects.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "target_bucket")
	loggingPrompt["target_prefix"] = types.TfPrompt{
		Label: "Enter target_prefix:\n(Optional) To specify a key prefix for log objects.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "target_prefix")
	selectOrder = append(selectOrder, "logging")

	resourceBlock["logging"] = builder.NestedPSOrder(nestedOrder[len(nestedOrder)-2:], nil, loggingPrompt, nil)

	builder.ResourceBuilder("aws_s3_bucket", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSS3BucketAnalyticsConfigurationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder []string

	prompts["bucket"] = types.TfPrompt{
		Label: "Enter bucket:\n(Required) The name of the bucket this analytics configuration is associated with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) Unique identifier of the analytics configuration for the bucket.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_s3_bucket_analytics_configuration", blockName, promptOrder, nil, resourceBlock)
}

func AWSS3BucketMetricPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["bucket"] = types.TfPrompt{
		Label: "Enter bucket:\n(Required) The name of the bucket to put metric configuration.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) Unique identifier of the metrics configuration for the bucket.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like filter [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket_metric", blockName, promptOrder, nil, resourceBlock)
		return
	}

	color.Green("\nEnter filter:\n(Optional) Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags" +
		"\nThe filter block supports:\n1.prefix\n2.tags(not supported by this cli yet)")

	filterPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	filterPrompt["prefix"] = types.TfPrompt{
		Label: "Enter prefix:\n(Optional) Object prefix for filtering (singular).",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "prefix")
	selectOrder = append(selectOrder, "filter")

	resourceBlock["filter"] = builder.NestedPSOrder(nestedPromptOrder, nil, filterPrompt, nil)

	builder.ResourceBuilder("aws_s3_bucket_metric", blockName, promptOrder, selectOrder, resourceBlock)
}
