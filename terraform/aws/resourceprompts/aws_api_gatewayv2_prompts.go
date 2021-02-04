package resourceprompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
)

func AWSAPIGatewayV2APIPrompt() {
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
			Field: "name",
			Doc:   "(Required) The name of the API. Must be less than or equal to 128 characters in length.",
		},
		{
			Type:  "select",
			Field: "protocol_type",
			Doc:   "(Required) The API protocol.",
			Items: []string{"HTTP", "WEBSOCKET"},
		},
		{
			Field: "api_key_selection_expression",
			Doc:   "(Optional) An API key selection expression.",
			Items: []string{"$context.authorizer.usageIdentifierKey", "$request.header.x-api-key"},
		},
		{
			Field: "credentials_arn",
			Doc:   "(Optional) Part of quick create. Specifies any credentials required for the integration. Applicable for HTTP APIs.",
		},
		{
			Field: "description",
			Doc:   "(Optional) The description of the API. Must be less than or equal to 1024 characters in length.",
		},
		{
			Field: "route_key",
			Doc: "(Optional) Part of quick create. Specifies any route key. Applicable for HTTP APIs." +
				"\nCheckout https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the API.",
			Validator: validators.RCValidator,
		},
		{
			Field: "target",
			Doc:   "(Optional) Part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Applicable for HTTP APIs.",
		},
		{
			Field: "version",
			Doc:   "(Optional) A version identifier for the API. Must be between 1 and 64 characters in length.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like cors_configuration [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_apigatewayv2_api", blockName, resourceBlock)
		return
	}

	corsConfigurationSchema := []types.Schema{
		{
			Field: "allow_credentials",
			Doc:   "(Optional) Whether credentials are included in the CORS request.",
		},
		{
			Field: "allow_headers",
			Doc:   "(Optional) The set of allowed HTTP headers.",
		},
		{
			Field: "allow_methods",
			Doc:   "(Optional) The set of allowed HTTP methods.",
		},
		{
			Field: "allow_origins",
			Doc:   "(Optional) The set of allowed origins.",
		},
		{
			Field: "expose_headers",
			Doc:   "(Optional) The set of exposed HTTP headers.",
		},
		{
			Field:     "max_age",
			Ex:        "30",
			Doc:       "(Optional) The number of seconds that the browser should cache preflight request results.",
			Validator: validators.IntValidator,
		},
	}

	resourceBlock["cors_configuration"] = builder.PSOrder(types.ProvidePS(corsConfigurationSchema))

	builder.ResourceBuilder("aws_apigatewayv2_api", blockName, resourceBlock)
}

func AWSAPIGatewayV2APIMappingPrompt() {
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
			Field: "api_id",
			Doc:   "(Required) The API identifier.",
		},
		{
			Field: "domain_name",
			Doc:   "(Required) The domain name. Use the aws_apigatewayv2_domain_name resource to configure a domain name.",
		},
		{
			Field: "stage",
			Doc:   "(Required) The API stage. Use the aws_apigatewayv2_stage resource to configure an API stage.",
		},
		{
			Field: "api_mapping_key",
			Doc: "(Optional) The API mapping key." +
				"\nCheckout https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_api_mapping", blockName, resourceBlock)
}

func AWSAPIGatewayV2AuthorizerPrompt() {
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
			Field: "api_id",
			Doc:   "(Required) The API identifier.",
		},
		{
			Type:  "select",
			Field: "authorizer_type",
			Doc:   "(Required) The authorizer type. Valid values: JWT, REQUEST. Specify REQUEST for a Lambda function using incoming request parameters. For HTTP APIs, specify JWT to use JSON Web Tokens.",
			Items: []string{"JWT", "REQUEST"},
		},
		{
			Field: "name",
			Doc:   "(Required) The name of the authorizer. Must be between 1 and 128 characters in length.",
		},
		{
			Field: "authorizer_credentials_arn",
			Doc:   "(Optional) The required credentials as an IAM role for API Gateway to invoke the authorizer. Supported only for REQUEST authorizers.",
		},
		{
			Type:  "select",
			Field: "authorizer_payload_format_version",
			Doc:   " (Optional) The format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.",
			Items: []string{"1.0", "2.0"},
		},
		{
			Field:     "authorizer_request_ttl_in_seconds",
			Doc:       "(Optional) The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to 300. Supported only for HTTP API Lambda authorizers.",
			Validator: validators.MinMaxIntValidator(0, 3600),
		},
		{
			Field: "authorizer_uri",
			Doc:   "(Optional) The authorizer's Uniform Resource Identifier (URI). For REQUEST authorizers this must be a well-formed Lambda function URI, such as the invoke_arn attribute of the aws_lambda_function resource. Supported only for REQUEST authorizers. Must be between 1 and 2048 characters in length.",
		},
		{
			Field:     "enable_simple_responses",
			Ex:        "(true/false)",
			Doc:       "(Optional) Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "identity_sources",
			Doc:   "(Optional) The identity sources for which authorization is requested. For REQUEST authorizers the value is a list of one or more mapping expressions of the specified request parameters. For JWT authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	jwtConfigurationSchema := []types.Schema{
		{
			Field: "audience",
			Doc:   "(Optional) A list of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.",
		},
		{
			Field: "issuer",
			Doc:   "(Optional) The base domain of the identity provider that issues JSON Web Tokens, such as the endpoint attribute of the aws_cognito_user_pool resource.",
		},
	}

	resourceBlock["jwt_configuration"] = builder.PSOrder(types.ProvidePS(jwtConfigurationSchema))

	builder.ResourceBuilder("aws_apigatewayv2_authorizer", blockName, resourceBlock)
}

func AWSAPIGatewayV2DeploymentPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	color.Yellow("\nCreating a deployment for an API requires at least " +
		"\none aws_apigatewayv2_route resource associated with that API. " +
		"\nTo avoid race conditions when all resources are being created " +
		"\ntogether, you need to add implicit resource references via the " +
		"\ntriggers argument or explicit resource references using the " +
		"\nresource depends_on meta-argument.\n\n" +
		"It is recommended to enable the resource lifecycle configuration " +
		"\nblock create_before_destroy argument in this resource configuration " +
		"\nto properly order redeployments in Terraform.\n")

	schema := []types.Schema{
		{
			Field: "api_id",
			Doc:   "(Required) The API identifier.",
		},
		{
			Field: "description",
			Doc:   "(Optional) The description for the deployment resource. Must be less than or equal to 1024 characters in length.",
		},
		{
			Field:     "triggers",
			Doc:       "(Optional) A map of arbitrary keys and values that, when changed, will trigger a redeployment. To force a redeployment without changing these keys/values, use the terraform taint command.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_deployment", blockName, resourceBlock)
}
