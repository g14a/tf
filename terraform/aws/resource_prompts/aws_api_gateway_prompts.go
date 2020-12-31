package resource_prompts

import (
	"fmt"
	"tf/builder"
	"tf/types"
	"tf/utils"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func AWSAPIGatewayAccountPrompt() {
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

	prompts["cloudwatch_role_arn"] = types.TfPrompt{
		Label: "Enter cloudwatch_role_arn:\n(Optional) The ARN of an IAM role for CloudWatch (to allow logging & monitoring).\n" +
			"See more in AWS Docs. Logging & monitoring can be enabled/disabled \n" +
			"and otherwise tuned on the API Gateway Stage level.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "cloudwatch_role_arn")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_api_gateway_account", blockName, promptOrder, nil, resourceBlock)
}

func AWSAPIGatewayApiKeyPrompt() {
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
		Label: "Enter name:\n(Required) The name of the API key",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) The API key description. Defaults to \"Managed by Terraform\".",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["enabled"] = types.TfPrompt{
		Label: "Enter enabled:\n(Optional) Specifies whether the API key can be used by callers. Defaults to true",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "enabled")

	prompts["value"] = types.TfPrompt{
		Label: "Enter value:\n(Optional) The value of the API key. If not specified, \n" +
			"it will be automatically generated by AWS on creation.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "value")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_api_gateway_api_key", blockName, promptOrder, nil, resourceBlock)
}

func AWSAPiGatewayAuthorizer() {
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
		Label: "Enter name:\n(Required) The name of the authorizer",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["rest_api_id"] = types.TfPrompt{
		Label: "Enter rest_api_id:\n(Required) The ID of the associated REST API",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "rest_api_id")

	prompts["identity_source"] = types.TfPrompt{
		Label: "Enter identity_source:\n(Optional) The source of the identity in an incoming request. Defaults to method.request.header.Authorization." +
			"\n Check https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_authorizer#identity_source",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "identity_source")

	prompts["authorizer_credentials"] = types.TfPrompt{
		Label: "Enter authorizer_credentials:\n(Optional) The credentials required for the authorizer. \n" +
			"To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "authorizer_credentials")

	prompts["authorizer_result_ttl_in_seconds"] = types.TfPrompt{
		Label: "Enter authorizer_result_ttl_in_seconds:\n (Optional) The TTL of cached authorizer results in seconds. Defaults to 300.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: utils.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "authorizer_result_ttl_in_seconds")

	prompts["identity_validation_expression"] = types.TfPrompt{
		Label: "Enter identity_validation_expression:\n(Optional) A validation expression for the incoming identity. " +
			"\nFor TOKEN type, this value should be a regular expression. " +
			"\nThe incoming token from the client is matched against this expression, " +
			"\nand will proceed if the token matches. If the token doesn't match, the client \n" +
			"receives a 401 Unauthorized response.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "identity_validation_expression")

	prompts["provider_arns"] = types.TfPrompt{
		Label: "Enter provider_arns:\n(Optional, required for type COGNITO_USER_POOLS) A list of the Amazon Cognito user pool ARNs. \n" +
			"Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "provider_arns")

	selects := map[string]types.TfSelect{}

	selects["type"] = types.TfSelect{
		Label: "Enter type:\n(Optional) The type of the authorizer. Possible values are " +
			"\nTOKEN for a Lambda function using a single authorization token submitted" +
			"\nin a custom header, REQUEST for a Lambda function using incoming request parameters, " +
			"\nor COGNITO_USER_POOLS for using an Amazon Cognito user pool. Defaults to TOKEN",
		Select: promptui.Select{
			Label: "",
			Items: []string{"TOKEN", "REQUEST", "COGNITO_USER_POOLS"},
		},
	}
	selectOrder = append(selectOrder, "type")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)
	builder.ResourceBuilder("aws_api_gateway_authorizer", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSAPIGatewayBasePathMappingPrompt() {
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

	prompts["api_id"] = types.TfPrompt{
		Label: "Enter api_id:\n(Required) The id of the API to connect.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "api_id")

	prompts["domain_name"] = types.TfPrompt{
		Label: "Enter domain_name:\n(Required) The already-registered domain name to connect the API to.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "domain_name")

	prompts["stage_name"] = types.TfPrompt{
		Label: "Enter stage_name:\n(Optional) The name of a specific deployment stage to expose at the given path. \n" +
			"If omitted, callers may select any stage by including its name as a path element after the base path.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "stage_name")

	prompts["base_path"] = types.TfPrompt{
		Label: "Enter base_path:\n(Optional) Path segment that must be prepended to the path when accessing \n" +
			"the API via this mapping. If omitted, the API is exposed at the root of the given domain.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "base_path")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_api_gateway_base_path_mapping", blockName, promptOrder, nil, resourceBlock)
}

func AWSAPIGatewayClientCertificatePrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder, nestedOrder []string

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) The description of the client certificate.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter tags:\n")

	tagsPrompt := map[string]types.TfPrompt{}

	tagsPrompt["Name"] = types.TfPrompt{
		Label: "Enter Name:\n",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedOrder = append(nestedOrder, "Name")
	selectOrder = append(selectOrder, "tags")

	resourceBlock["tags"] = builder.NestedPSOrder(nestedOrder, nil, tagsPrompt, nil)

	builder.ResourceBuilder("aws_api_gateway_client_certificate", blockName, promptOrder, selectOrder, resourceBlock)
}

func AWSAPIGatewayDeploymentPrompt()  {
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

	prompts["rest_api_id"] = types.TfPrompt{
		Label: "Enter rest_api_id:\n(Required) The ID of the associated REST API",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "rest_api_id")

	prompts["stage_name"] = types.TfPrompt{
		Label: "Enter stage_name:\n(Optional) The name of the stage. If the specified stage already exists, \n" +
			"it will be updated to point to the new deployment. If the stage does not exist, \n" +
			"a new one will be created and point to this deployment.\n",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "stage_name")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) The description of the deployment",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["stage_description"] = types.TfPrompt{
		Label: "Enter stage_description:\n(Optional) The description of the stage",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "stage_description")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n Follow the format k1=v1,k2=v2",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like variables etc [y/n]?\n\n")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_api_gateway_deployment", blockName, promptOrder, selectOrder, resourceBlock)
		return
	}

	builder.ResourceBuilder("aws_api_gateway_deployment", blockName, promptOrder, selectOrder, resourceBlock)
}