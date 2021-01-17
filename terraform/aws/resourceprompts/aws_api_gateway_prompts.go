package resourceprompts

import (
	"fmt"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/utils"

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

	schema := []types.Schema{
		{
			Field: "cloudwatch_role_arn",
			Ex:    "",
			Doc: "(Optional) The ARN of an IAM role for CloudWatch (to allow logging & monitoring). " +
				"\nCheckout https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console " +
				"\nLogging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_api_gateway_account", blockName, resourceBlock)
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

	schema := []types.Schema{
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) The name of the API key",
		},
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) The API key description. Defaults to \"Managed by Terraform\".",
		},
		{
			Field:     "enabled",
			Ex:        "(true/false)",
			Doc:       "(Optional) Specifies whether the API key can be used by callers. Defaults to true.",
			Validator: utils.BoolValidator,
		},
		{
			Field: "value",
			Ex:    "api-key-123",
			Doc:   "(Optional) The value of the API key. If not specified, it will be automatically generated by AWS on creation.",
		},
		{
			Field: "tags",
			Ex:    "k1=v1,k2=v2",
			Doc:   "(Optional) Key-value map of resource tags",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_api_gateway_api_key", blockName, resourceBlock)
}

func AWSAPIGatewayAuthorizerPrompt() {
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
			Field: "name",
			Ex:    "auth-name-123",
			Doc:   "(Required) The name of the authorizer",
		},
		{
			Field: "rest_api_id",
			Ex:    "rai-123",
			Doc:   "(Required) The ID of the associated REST API",
		},
		{
			Field: "identity_source",
			Ex:    "",
			Doc: "(Optional) The source of the identity in an incoming request. Defaults to method.request.header.Authorization. " +
				"\nFor REQUEST type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. " +
				"\n\"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName\"",
		},
		{
			Type:  "select",
			Field: "type",
			Doc: "(Optional) The type of the authorizer. Possible values are TOKEN for a " +
				"\nLambda function using a single authorization token submitted in a " +
				"\ncustom header, REQUEST for a Lambda function using incoming request " +
				"\nparameters, or COGNITO_USER_POOLS for using an Amazon Cognito user pool. " +
				"\nDefaults to TOKEN",
			Items: []string{"TOKEN", "REQUEST", "COGNITO_USER_POOLS"},
		},
		{
			Field: "authorizer_credentials",
			Ex:    "",
			Doc: "(Optional) The credentials required for the authorizer. To specify an IAM Role for API Gateway " +
				"\nto assume, use the IAM Role ARN.",
		},
		{
			Field:     "authorizer_result_ttl_in_seconds",
			Ex:        "300",
			Doc:       "(Optional) The TTL of cached authorizer results in seconds. Defaults to 300",
			Validator: utils.IntValidator,
		},
		{
			Field: "identity_validation_expression",
			Ex:    "",
			Doc: "(Optional) A validation expression for the incoming identity. For TOKEN type, " +
				"\nthis value should be a regular expression. The incoming token from the client " +
				"\nis matched against this expression, and will proceed if the token matches. " +
				"\nIf the token doesn't match, the client receives a 401 Unauthorized response.",
		},
		{
			Field: "provider_arns",
			Ex:    "",
			Doc: "(Optional, required for type COGNITO_USER_POOLS) A list of the Amazon Cognito user pool ARNs. " +
				"\nEach element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_api_gateway_authorizer", blockName, resourceBlock)
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

	builder.ResourceBuilder("aws_api_gateway_base_path_mapping", blockName, resourceBlock)
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
	var promptOrder, nestedOrder []string

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

	resourceBlock["tags"] = builder.PSOrder(nestedOrder, nil, tagsPrompt, nil)

	builder.ResourceBuilder("aws_api_gateway_client_certificate", blockName, resourceBlock)
}

func AWSAPIGatewayDeploymentPrompt() {
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
		Label: "Enter tags:\n For e.g. k1=v1,k2=v2",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["variables"] = types.TfPrompt{
		Label: "Enter variables: For e.g. k1=v1,k2=v2\n(Optional) A map that defines variables for the stage",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "variables")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	lifecyclePrompt := map[string]types.TfPrompt{}
	var nestedOrder []string

	color.Green("\nEnter lifecycle block(Recommended):\n")

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

	resourceBlock["lifecycle"] = builder.PSOrder(nestedOrder, selectOrder, lifecyclePrompt, nil)

	builder.ResourceBuilder("aws_api_gateway_deployment", blockName, resourceBlock)
}

func AWSAPIGatewayDocumentationPartPrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, nestedPromptOrder, nestedSelectOrder []string

	prompts["properties"] = types.TfPrompt{
		Label: "Enter properties:\n(Required) A content map of API-specific key-value pairs describing \n" +
			"the targeted API entity. The map must be encoded as a JSON string, \n" +
			"e.g., \"{ \\\"description\\\": \\\"The API does …\\\" }\". \n" +
			"Only Swagger-compliant key-value pairs can be exported and, hence, published.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "properties")

	prompts["rest_api_id"] = types.TfPrompt{
		Label: "Enter rest_api_id:\n(Required) The ID of the associated REST API",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "rest_api_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter location:\n(Required) The location of the targeted API entity of the to-be-created documentation part.\n" +
		"The location block supports:" +
		"\n1.method\n2.name\n3.path\n4.status_code\n5.type")

	locationPrompt := map[string]types.TfPrompt{}

	locationPrompt["method"] = types.TfPrompt{
		Label: "Enter method:\n(Optional) The HTTP verb of a method. The default value is * for any method.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "method")

	locationPrompt["name"] = types.TfPrompt{
		Label: "Enter name:\n(Optional) The name of the targeted API entity.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "name")

	locationPrompt["path"] = types.TfPrompt{
		Label: "Enter path:\n(Optional) The URL path of the target. The default value is / for the root resource.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "path")

	locationPrompt["status_code"] = types.TfPrompt{
		Label: "Enter status_code:\n(Optional) The HTTP status code of a response. The default value is * for any status code.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "status_code")

	locationSelect := map[string]types.TfSelect{}

	locationSelect["type"] = types.TfSelect{
		Label: "Enter type:\n(Required) The type of API entity to which the documentation content applies",
		Select: promptui.Select{
			Label: "",
			Items: []string{"API", "METHOD", "REQUEST_BODY"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "type")

	resourceBlock["location"] = builder.PSOrder(nestedPromptOrder, nestedSelectOrder, locationPrompt, locationSelect)

	builder.ResourceBuilder("aws_api_gateway_documentation_part", blockName, resourceBlock)
}

func AWSAPIGatewayDocumentationVersionPrompt() {
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

	prompts["version"] = types.TfPrompt{
		Label: "Enter version:\n(Required) The version identifier of the API documentation snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "version")

	prompts["rest_api_id"] = types.TfPrompt{
		Label: "Enter rest_api_id:\n(Required) The ID of the associated Rest API" +
			"",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "rest_api_id")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) The description of the API documentation version." +
			"",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	builder.ResourceBuilder("aws_api_gateway_documentation_version", blockName, builder.PSOrder(promptOrder, nil, prompts, nil))
}

// aws_api_gateway_domain_name
func AWSAPIGatewayDomainNamePrompt() {
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

	var promptOrder, selectOrder, nestedSelectOrder []string

	prompts["domain_name"] = types.TfPrompt{
		Label: "Enter domain_name:\n(Required) The fully-qualified domain name to register",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "domain_name")

	prompts["certificate_arn"] = types.TfPrompt{
		Label: "Enter certificate_arn:\n(Optional) The ARN for an AWS-managed certificate. \n" +
			"AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. \n" +
			"Conflicts with certificate_name, certificate_body, certificate_chain, certificate_private_key, \nregional_certificate_arn, and regional_certificate_name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_arn")

	prompts["regional_certificate_arn"] = types.TfPrompt{
		Label: "Enter regional_certificate_arn:\n(Optional) The unique name to use when registering this \n" +
			"certificate as an IAM server certificate. \n" +
			"Conflicts with certificate_arn, regional_certificate_arn, and \n" +
			"regional_certificate_name. Required if certificate_arn is not set.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "regional_certificate_arn")

	prompts["certificate_name"] = types.TfPrompt{
		Label: "Enter certificate_name:\n(Optional) The unique name to use when registering this certificate as an IAM server certificate. \n" +
			"Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name. \n" +
			"Required if certificate_arn is not set.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_name")

	prompts["certificate_body"] = types.TfPrompt{
		Label: "Enter certificate_body:\n(Optional) The certificate issued for the domain name \n" +
			"being registered, in PEM format. Only valid for EDGE endpoint configuration type. \n" +
			"Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_body")

	prompts["certificate_chain"] = types.TfPrompt{
		Label: "Enter certificate_chain:\n(Optional) The certificate for the CA that issued the \n" +
			"certificate, along with any intermediate CA certificates required to create an \n" +
			"unbroken chain to a certificate trusted by the intended API clients. \n" +
			"Only valid for EDGE endpoint configuration type. \n" +
			"Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_chain")

	prompts["certificate_private_key"] = types.TfPrompt{
		Label: "Enter certificate_private_key:\n(Optional) The private key associated with the domain \n" +
			"certificate given in certificate_body. Only valid for EDGE endpoint configuration type. \n" +
			"Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "certificate_private_key")

	prompts["regional_certificate_name"] = types.TfPrompt{
		Label: "Enter regional_certificate_name:\n(Optional) The user-friendly name of the certificate that will be used by \n" +
			"regional endpoint for this domain name. \n" +
			"Conflicts with certificate_arn, certificate_name, certificate_body, certificate_chain, and certificate_private_key",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "regional_certificate_name")

	selects["security_policy"] = types.TfSelect{
		Label: "Enter domain_name:\n(Optional) The Transport Layer Security (TLS) version + cipher suite for this DomainName.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"TLS_1_0", "TLS_1_2"},
		},
	}
	selectOrder = append(selectOrder, "security_policy")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Yellow("\nConfigure nested settings like endpoint_configuration/tags [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ProviderBuilder("aws_api_gateway_domain_name", resourceBlock)
		return
	}

	endpointConfigSelect := map[string]types.TfSelect{}
	endpointConfigSelect["types"] = types.TfSelect{
		Label: "Enter types: [\"t1\",\"t2\"]\n(Required) A list of endpoint types. This resource currently only supports managing a single value. \n" +
			"Valid values: EDGE or REGIONAL. If unspecified, defaults to EDGE. \n" +
			"Must be declared as REGIONAL in non-Commercial partitions. \n" +
			"Refer to https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html \n" +
			"for more information on the difference between edge-optimized and regional APIs.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"EDGE", "REGIONAL"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "types")

	resourceBlock["endpoint_configuration"] = builder.PSOrder(nil, nestedSelectOrder, nil, endpointConfigSelect)
	builder.ResourceBuilder("aws_api_gateway_domain_name", blockName, resourceBlock)
}

func AWSAPIGatewayGatewayResponsePrompt() {
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

	prompts["rest_api_id"] = types.TfPrompt{
		Label: "Enter rest_api_id:\n(Required) The string identifier of the associated REST API.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "rest_api_id")

	prompts["response_type"] = types.TfPrompt{
		Label: "Enter response_type:\n(Required) The response type of the associated GatewayResponse.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "response_type")

	prompts["status_code"] = types.TfPrompt{
		Label: "Enter status_code:\n(Optional) The HTTP status code of the Gateway Response.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "status_code")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter response_templates:\n(Optional) A map specifying the templates used to transform the response body.")

	responseTemplatesPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	responseTemplatesPrompt["application/json"] = types.TfPrompt{
		Label: "Enter application/json:\nCheck https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_gateway_response",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "application/json")

	resourceBlock["response_templates"] = builder.PSOrder(nestedPromptOrder, nil, responseTemplatesPrompt, nil)

	builder.ResourceBuilder("aws_api_gateway_account", blockName, resourceBlock)
}

func AWSAPIGatewayIntegration() {

}
