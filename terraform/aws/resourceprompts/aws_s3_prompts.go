package resourceprompts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
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

	schema := []types.Schema{
		{
			Field: "bucket",
			Ex:    "",
			Doc:   "(Required) The name of an AWS Partition S3 Bucket or the Amazon Resource Name (ARN) of S3 on Outposts Bucket that you want to associate this access point with.",
		},
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) The name you want to assign to this access point.",
		},
		{
			Field: "account_id",
			Ex:    "",
			Doc:   "(Optional) The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the Terraform AWS provider.",
		},
		{
			Field: "policy",
			Ex:    "",
			Doc:   "(Optional) A valid JSON document that specifies the policy that you want to apply to this access point.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like public_access_block_configuration/vpc_configuration etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_access_point", blockName, resourceBlock)
		return
	}

	publicAccessBlockConfigurationSchema := []types.Schema{
		{
			Field: "block_public_acls",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether Amazon S3 should block public ACLs for buckets in this account. " +
				"\nDefaults to true. Enabling this setting does not affect existing policies or ACLs. " +
				"\nWhen set to true causes the following behavior:\n\n    " +
				"PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n    " +
				"PUT Object calls fail if the request includes a public ACL.\n    " +
				"PUT Bucket calls fail if the request includes a public ACL.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "block_public_policy",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether Amazon S3 should block public bucket policies for buckets in this account. " +
				"\nDefaults to true. Enabling this setting does not affect existing bucket policies. " +
				"\nWhen set to true causes Amazon S3 to:\n\n    " +
				"Reject calls to PUT Bucket policy if the specified bucket policy allows public access.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "ignore_public_acls",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether Amazon S3 should ignore public ACLs for buckets in this account. " +
				"\nDefaults to true. Enabling this setting does not affect the persistence of any existing " +
				"\nACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:\n\n    " +
				"Ignore all public ACLs on buckets in this account and any objects that they contain.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "restrict_public_buckets",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether Amazon S3 should restrict public bucket policies for buckets in " +
				"\nthis account. Defaults to true. Enabling this setting does not affect previously stored " +
				"\nbucket policies, except that public and cross-account access within any public bucket policy, " +
				"\nincluding non-public delegation to specific accounts, is blocked. When set to true:\n\n    " +
				"Only the bucket owner and AWS Services can access buckets with public policies.",
			Validator: validators.BoolValidator,
		},
	}

	resourceBlock["public_access_block_configuration"] = builder.PSOrder(types.ProvidePS(publicAccessBlockConfigurationSchema))

	vpcConfigSchema := []types.Schema{
		{
			Field: "vpc_id",
			Ex:    "vpc-123",
			Doc:   "(Required) This access point will only allow connections from the specified VPC ID.",
		},
	}

	resourceBlock["vpc_configuration"] = builder.PSOrder(types.ProvidePS(vpcConfigSchema))

	builder.ResourceBuilder("aws_s3_access_point", blockName, resourceBlock)
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

	schema := []types.Schema{
		{
			Field: "account_id",
			Ex:    "",
			Doc:   "(Optional) AWS account ID to configure. Defaults to automatically determined account ID of the Terraform AWS provider.",
		},
		{
			Field: "block_public_acls",
			Ex:    "",
			Doc: "(Optional) Whether Amazon S3 should block public ACLs for buckets in this account. " +
				"\nDefaults to false. Enabling this setting does not affect existing policies or ACLs. " +
				"\nWhen set to true causes the following behavior:\n\n    " +
				"PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.\n    " +
				"PUT Object calls will fail if the request includes an object ACL.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "block_public_policy",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether Amazon S3 should block public bucket policies for buckets in this account. " +
				"\nDefaults to true. Enabling this setting does not affect existing bucket policies. " +
				"\nWhen set to true causes Amazon S3 to:\n\n    " +
				"Reject calls to PUT Bucket policy if the specified bucket policy allows public access.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "ignore_public_acls",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether Amazon S3 should ignore public ACLs for buckets in this account. " +
				"\nDefaults to true. Enabling this setting does not affect the persistence of any existing " +
				"\nACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:\n\n    " +
				"Ignore all public ACLs on buckets in this account and any objects that they contain.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "restrict_public_buckets",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether Amazon S3 should restrict public bucket policies for buckets in " +
				"\nthis account. Defaults to true. Enabling this setting does not affect previously stored " +
				"\nbucket policies, except that public and cross-account access within any public bucket policy, " +
				"\nincluding non-public delegation to specific accounts, is blocked. When set to true:\n\n    " +
				"Only the bucket owner and AWS Services can access buckets with public policies.",
			Validator: validators.BoolValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_s3_account_public_access_block", blockName, resourceBlock)
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

	schema := []types.Schema{
		{
			Field: "bucket",
			Ex:    "",
			Doc:   "(Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name. Must be less than or equal to 63 characters in length.",
		},
		{
			Type:  "select",
			Field: "acl",
			Doc:   "(Optional) The canned ACL to apply.",
			Items: []string{"private", "public-read", "public-read-write", "aws-exec-read", "authenticated-read", "log-delivery-write"},
		},
		{
			Field: "force_destroy",
			Doc:   "(Optional) The canned ACL to apply.",
			Items: []string{"private", "public-read", "public-read-write", "aws-exec-read", "authenticated-read", "log-delivery-write"},
		},
		{
			Field:     "force_destroy",
			Ex:        "(true/false)",
			Doc:       "(Optional, Default:false) A boolean that indicates all objects (including any locked objects) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable.",
			Validator: validators.BoolValidator,
		},
		{
			Type:  "select",
			Field: "acceleration_status",
			Doc:   "(Optional) Sets the accelerate configuration of an existing bucket.",
			Items: []string{"Enabled", "Suspended"},
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like cors_rule/versioning/website etc [y/n]?\n\n", "text")

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

	color.Green("\nEnter cors_rule (Optional) A rule of Cross-Origin Resource Sharing :\n\n")

	corsRuleSchema := []types.Schema{
		{
			Field: "allowed_headers",
			Ex:    "",
			Doc:   "(Optional) Specifies which headers are allowed.",
		},
		{
			Type:  "select",
			Field: "allowed_methods",
			Doc:   "(Required) Specifies which methods are allowed.",
			Items: []string{"GET", "PUT", "POST", "DELETE", "HEAD"},
		},
		{
			Field: "allowed_origins",
			Ex:    "",
			Doc:   "(Required) Specifies which origins are allowed.",
		},
		{
			Field: "exposed_headers",
			Ex:    "",
			Doc:   "(Optional) Specifies expose header in the response.",
		},
		{
			Field:     "max_age_seconds",
			Ex:        "30",
			Doc:       "(Optional) Specifies time in seconds that browser can cache the response for a preflight request.",
			Validator: validators.IntValidator,
		},
	}

	resourceBlock["cors_rule"] = builder.PSOrder(types.ProvidePS(corsRuleSchema))

	color.Green("\nEnter website:\nThe website object supports the following:" +
		"\n1.index_document\n2.error_document\n3.redirect_all_requests_to\n4.routing_rules\n\n")

	websiteSchema := []types.Schema{
		{
			Field: "index_document",
			Ex:    "index.html",
			Doc:   "(Required, unless using redirect_all_requests_to) Amazon S3 returns this index document when requests are made to the root domain or any of the subfolders.",
		},
		{
			Field: "error_document",
			Ex:    "",
			Doc:   "(Optional) An absolute path to the document to return in case of a 4XX error.",
		},
		{
			Field: "redirect_all_requests_to",
			Ex:    "",
			Doc:   "(Optional) A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting requests. The default is the protocol that is used in the original request.",
		},
		{
			Field: "routing_rules",
			Ex:    "",
			Doc:   "(Optional) A json array containing routing rules describing redirect behavior and when redirects are applied.",
		},
	}

	resourceBlock["website"] = builder.PSOrder(types.ProvidePS(websiteSchema))

	versioningSchema := []types.Schema{
		{
			Field:     "enabled",
			Ex:        "(true/false)",
			Doc:       "(Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.",
			Validator: validators.BoolValidator,
		},
		{
			Field:     "mfa_delete",
			Ex:        "(true/false)",
			Doc:       "(Optional) Enable MFA delete for either Change the versioning state of your bucket or Permanently delete an object version. Default is false. This cannot be used to toggle this setting but is available to allow managed buckets to reflect the state in AWS",
			Validator: validators.BoolValidator,
		},
	}

	resourceBlock["versioning"] = builder.PSOrder(types.ProvidePS(versioningSchema))

	loggingSchema := []types.Schema{
		{
			Field: "target_bucket",
			Ex:    "",
			Doc:   "(Required) The name of the bucket that will receive the log objects.",
		},
		{
			Field: "target_prefix",
			Ex:    "",
			Doc:   "(Optional) To specify a key prefix for log objects.",
		},
	}

	resourceBlock["logging"] = builder.PSOrder(types.ProvidePS(loggingSchema))

	builder.ResourceBuilder("aws_s3_bucket", blockName, resourceBlock)
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

	builder.ResourceBuilder("aws_s3_bucket_analytics_configuration", blockName, resourceBlock)
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
	var promptOrder []string

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
		builder.ResourceBuilder("aws_s3_bucket_metric", blockName, resourceBlock)
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

	resourceBlock["filter"] = builder.PSOrder(nestedPromptOrder, nil, filterPrompt, nil)

	builder.ResourceBuilder("aws_s3_bucket_metric", blockName, resourceBlock)
}

func AWSS3BucketNotificationPrompt() {
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
		Label: "Enter bucket:\n(Required) The name of the bucket to put notification configuration.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like topic/queue/lambda_function etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket_notification", blockName, resourceBlock)
		return
	}

	topicPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	color.Green("\nEnter topic:\n(Optional) The notification configuration to SNS Topic (documented below)." +
		"\nThe topic notification configuration supports the following:" +
		"\n1.id\n2.topic_arn\n3.events\n4.filter_prefix\n5.filter_suffix\n")

	topicPrompt["id"] = types.TfPrompt{
		Label: "Enter id:\n(Optional) Specifies unique identifier for each of the notification configurations.\n",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "id")

	topicPrompt["topic_arn"] = types.TfPrompt{
		Label: "Enter topic_arn:\n(Required) Specifies Amazon SNS topic ARN.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "topic_arn")

	topicPrompt["events"] = types.TfPrompt{
		Label: "Enter events:\n(Required) Specifies event for which to send notifications.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "events")

	topicPrompt["filter_prefix"] = types.TfPrompt{
		Label: "Enter filter_prefix:\n(Optional) Specifies object key name prefix.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "filter_prefix")

	topicPrompt["filter_suffix"] = types.TfPrompt{
		Label: "Enter filter_suffix:\n(Optional) Specifies object key name suffix.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "filter_suffix")

	resourceBlock["topic"] = builder.PSOrder(nestedPromptOrder, nil, topicPrompt, nil)

	color.Green("\nEnter queue:\n(Optional) The notification configuration to SQS Queue (documented below)." +
		"\nThe queue notification configuration supports the following:" +
		"\n1.id\n2.queue_arn\n3.events\n4.filter_prefix\n5.filter_suffix\n")

	queuePrompt := map[string]types.TfPrompt{}

	queuePrompt["id"] = types.TfPrompt{
		Label: "Enter id:\n(Optional) Specifies unique identifier for each of the notification configurations.\n",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "id")

	queuePrompt["queue_arn"] = types.TfPrompt{
		Label: "Enter queue_arn:\n(Required) Specifies Amazon SQS queue ARN.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "queue_arn")

	queuePrompt["events"] = types.TfPrompt{
		Label: "Enter events:\n(Required) Specifies event for which to send notifications.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "events")

	queuePrompt["filter_prefix"] = types.TfPrompt{
		Label: "Enter filter_prefix:\n(Optional) Specifies object key name prefix.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "filter_prefix")

	queuePrompt["filter_suffix"] = types.TfPrompt{
		Label: "Enter filter_suffix:\n(Optional) Specifies object key name suffix.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "filter_suffix")

	resourceBlock["queue"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-5:], nil, queuePrompt, nil)

	color.Green("\nEnter lambda_function:\n(Optional, Multiple) Used to configure notifications to a Lambda Function" +
		"\nThe queue notification configuration supports the following:" +
		"\n1.id\n2.lambda_function_arn\n3.events\n4.filter_prefix\n5.filter_suffix\n")

	lambdaFunctionPrompt := map[string]types.TfPrompt{}

	lambdaFunctionPrompt["id"] = types.TfPrompt{
		Label: "Enter id:\n(Optional) Specifies unique identifier for each of the notification configurations.\n",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "id")

	lambdaFunctionPrompt["lambda_function_arn"] = types.TfPrompt{
		Label: "Enter lambda_function_arn:\n(Required) Specifies Amazon Lambda function ARN.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "lambda_function_arn")

	lambdaFunctionPrompt["events"] = types.TfPrompt{
		Label: "Enter events:\n(Required) Specifies event for which to send notifications.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "events")

	lambdaFunctionPrompt["filter_prefix"] = types.TfPrompt{
		Label: "Enter filter_prefix:\n(Optional) Specifies object key name prefix.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "filter_prefix")

	lambdaFunctionPrompt["filter_suffix"] = types.TfPrompt{
		Label: "Enter filter_suffix:\n(Optional) Specifies object key name suffix.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "filter_suffix")

	resourceBlock["lambda_function"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-5:], nil, lambdaFunctionPrompt, nil)

	builder.ResourceBuilder("aws_s3_bucket_notification", blockName, resourceBlock)
}

func AWSS3BucketObjectPrompt() {
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
		Label: "Enter bucket:\n(Required) The name of the bucket to put the file in. Alternatively, an S3 access point ARN can be specified.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	prompts["key"] = types.TfPrompt{
		Label: "Enter key:\n(Required) The name of the object once it is in the bucket.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "key")

	prompts["source"] = types.TfPrompt{
		Label: "Enter source:\n(Optional, conflicts with content and content_base64) " +
			"\nThe path to a file that will be read and uploaded as raw bytes for the object content.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source")

	prompts["content"] = types.TfPrompt{
		Label: "Enter content:\n(Optional, conflicts with source and content_base64) Literal string " +
			"\nvalue to use as the object content, which will be uploaded as UTF-8-encoded text.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "content")

	prompts["content_base64"] = types.TfPrompt{
		Label: "Enter content_base64:\n(Optional, conflicts with source and content) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "content_base64")

	prompts["cache_control"] = types.TfPrompt{
		Label: "Enter cache_control:\n(Optional) Specifies caching behavior along the request/reply chain. " +
			"\nCheckout http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9 for further details.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "cache_control")

	prompts["content_disposition"] = types.TfPrompt{
		Label: "Enter content_disposition:\n(Optional) Specifies presentational information for the object. " +
			"\nCheckout http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1 for further information.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "content_disposition")

	prompts["content_encoding"] = types.TfPrompt{
		Label: "Enter content_encoding:\n(Optional) Specifies what content encodings have been applied " +
			"\nto the object and thus what decoding mechanisms must be applied to obtain the " +
			"\nmedia-type referenced by the Content-Type header field. " +
			"\nCheckout http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11 for further information.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "content_encoding")

	prompts["content_language"] = types.TfPrompt{
		Label: "Enter content_language:\n(Optional) The language the content is in e.g. en-US or en-GB.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "content_language")

	prompts["content_type"] = types.TfPrompt{
		Label: "Enter content_type:\n(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. " +
			"\nAll Valid MIME Types are valid for this input.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "content_type")

	prompts["website_redirect"] = types.TfPrompt{
		Label: "Enter website_redirect:\n(Optional) Specifies a target URL for website redirect." +
			"\nCheckout http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "website_redirect")

	prompts["etag"] = types.TfPrompt{
		Label: "Enter etag:\n(Optional) Used to trigger updates. The only meaningful value is " +
			"\n${filemd5(\"path/to/file\")} (Terraform 0.11.12 or later) or ${md5(file(\"path/to/file\"))} (Terraform 0.11.11 or earlier). " +
			"\nThis attribute is not compatible with KMS encryption, kms_key_id or server_side_encryption = \"aws:kms\"",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "etag")

	prompts["kms_key_id"] = types.TfPrompt{
		Label: "Enter kms_key_id:\n(Optional) Amazon Resource Name (ARN) of the KMS Key to use for object encryption. " +
			"\nIf the S3 Bucket has server-side encryption enabled, that value will automatically be used. " +
			"\nIf referencing the aws_kms_key resource, use the arn attribute. If referencing the " +
			"\naws_kms_alias data source or resource, use the target_key_arn attribute. " +
			"\nTerraform will only perform drift detection if a configuration value is provided.\n",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "kms_key_id")

	prompts["metadata"] = types.TfPrompt{
		Label: "Enter metadata: e.g.k1=v1,k2=v2\n(Optional) A map of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-).",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "metadata")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g.k1=v1,k2=v2\n(Optional) A map of tags to assign to the object.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["force_destroy"] = types.TfPrompt{
		Label: "Enter force_destroy(true/false):\n(Optional) Allow the object to be deleted by removing any legal hold on any object version. " +
			"\nDefault is false. This value should be set to true only if the bucket has S3 object lock enabled.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "force_destroy")

	prompts["object_lock_retain_until_date"] = types.TfPrompt{
		Label: "Enter object_lock_retain_until_date:\n(Optional) The date and time, in RFC3339 format, when this object's object lock will expire." +
			"\nCheckout https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "object_lock_retain_until_date")

	selects := map[string]types.TfSelect{}

	selects["acl"] = types.TfSelect{
		Label: "Enter acl:\n(Optional) The canned ACL to apply. Defaults to private.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"private", "public-read", "public-read-write", "aws-exec-read", "authenticated-read", "bucket-owner-read", "bucket-owner-full-control"},
		},
	}
	selectOrder = append(selectOrder, "acl")

	selects["storage_class"] = types.TfSelect{
		Label: "Enter storage_class:\n(Optional) Specifies the desired Storage Class for the object. Can be either \"STANDARD\", \"REDUCED_REDUNDANCY\", \"ONEZONE_IA\", \"INTELLIGENT_TIERING\", \"GLACIER\", \"DEEP_ARCHIVE\", or \"STANDARD_IA\". Defaults to \"STANDARD",
		Select: promptui.Select{
			Label: "",
			Items: []string{"STANDARD", "REDUCED_REDUNDANCY", "ONEZONE_IA", "INTELLIGENT_TIERING", "GLACIER", "DEEP_ARCHIVE", "STANDARD_IA"},
		},
	}
	selectOrder = append(selectOrder, "storage_class")

	selects["object_lock_legal_hold_status"] = types.TfSelect{
		Label: "Enter object_lock_legal_hold_status:\n(Optional) The legal hold status that you want to apply to the specified object. Valid values are ON and OFF." +
			"\nCheckout https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds",
		Select: promptui.Select{
			Label: "",
			Items: []string{"ON", "OFF"},
		},
	}
	selectOrder = append(selectOrder, "object_lock_legal_hold_status")

	selects["object_lock_mode"] = types.TfSelect{
		Label: "Enter object_lock_mode:\n(Optional) The object lock retention mode that you want to apply to this object.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"ON", "OFF"},
		},
	}
	selectOrder = append(selectOrder, "object_lock_mode")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_s3_bucket_object", blockName, resourceBlock)
}

func AWSS3BucketOwnershipControlsPrompt() {
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
		Label: "Enter bucket:\n(Required) The name of the bucket that you want to associate this access point with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like rule [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket_ownership_controls", blockName, resourceBlock)
		return
	}

	roleSelect := map[string]types.TfSelect{}
	var nestedSelectOrder []string
	roleSelect["object_ownership"] = types.TfSelect{
		Label: "Enter object_ownership:\n(Optional) Object ownership.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"BucketOwnerPreferred", "ObjectWriter"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "object_ownership")

	resourceBlock["rule"] = builder.PSOrder(nil, nestedSelectOrder, nil, roleSelect)

	builder.ResourceBuilder("aws_s3_bucket_ownership_controls", blockName, resourceBlock)
}

func AWSS3BucketPublicAccessBlockPrompt() {
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

	prompts["block_public_acls"] = types.TfPrompt{
		Label: "Enter block_public_acls(true/false):\nOptional) Whether Amazon S3 should block public ACLs for buckets in this account. " +
			"\nDefaults to true. Enabling this setting does not affect existing policies or ACLs. " +
			"\nWhen set to true causes the following behavior:\n\n    " +
			"PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n    " +
			"PUT Object calls fail if the request includes a public ACL.\n    " +
			"PUT Bucket calls fail if the request includes a public ACL.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "block_public_acls")

	prompts["block_public_policy"] = types.TfPrompt{
		Label: "Enter block_public_policy(true/false):\n(Optional) Whether Amazon S3 should block public bucket policies for buckets in this account. " +
			"\nDefaults to true. Enabling this setting does not affect existing bucket policies. " +
			"\nWhen set to true causes Amazon S3 to:\n\n    Reject calls to PUT Bucket policy if the specified bucket policy allows public access.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "block_public_policy")

	prompts["ignore_public_acls"] = types.TfPrompt{
		Label: "Enter ignore_public_acls(true/false):\n(Optional) Whether Amazon S3 should ignore public ACLs for buckets in this account. " +
			"\nDefaults to true. Enabling this setting does not affect the persistence of any " +
			"\nexisting ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:\n\n    Ignore all public ACLs on buckets in this account and any objects that they contain.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ignore_public_acls")

	prompts["restrict_public_buckets"] = types.TfPrompt{
		Label: "Enter restrict_public_buckets(true/false):\n(Optional) Whether Amazon S3 should restrict public bucket policies for buckets " +
			"\nin this account. Defaults to true. Enabling this setting does not affect previously stored " +
			"\nbucket policies, except that public and cross-account access within any public bucket policy, " +
			"\nincluding non-public delegation to specific accounts, is blocked. When set to true:\n\n    Only the bucket owner and AWS Services can access buckets with public policies.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "restrict_public_buckets")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_s3_bucket_public_access_block", blockName, resourceBlock)
}
