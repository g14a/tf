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
			Doc: "(Required) The authorizer type. Valid values: JWT, REQUEST. " +
				"\nSpecify REQUEST for a Lambda function using incoming request parameters. " +
				"\nFor HTTP APIs, specify JWT to use JSON Web Tokens.",
			Items: []string{"JWT", "REQUEST"},
		},
		{
			Field: "name",
			Doc:   "(Required) The name of the authorizer. Must be between 1 and 128 characters in length.",
		},
		{
			Field: "authorizer_credentials_arn",
			Doc: "(Optional) The required credentials as an IAM role for API " +
				"\nGateway to invoke the authorizer. Supported only for REQUEST authorizers.",
		},
		{
			Type:  "select",
			Field: "authorizer_payload_format_version",
			Doc: " (Optional) The format of the payload sent to an HTTP API " +
				"\nLambda authorizer. Required for HTTP API Lambda authorizers.",
			Items: []string{"1.0", "2.0"},
		},
		{
			Field: "authorizer_request_ttl_in_seconds",
			Doc: "(Optional) The time to live (TTL) for cached authorizer results, " +
				"\nin seconds. If it equals 0, authorization caching is disabled. If it " +
				"\nis greater than 0, API Gateway caches authorizer responses. The maximum " +
				"\nvalue is 3600, or 1 hour. Defaults to 300. Supported only for HTTP API Lambda authorizers.",
			Validator: validators.MinMaxIntValidator(0, 3600),
		},
		{
			Field: "authorizer_uri",
			Doc: "(Optional) The authorizer's Uniform Resource Identifier (URI). " +
				"\nFor REQUEST authorizers this must be a well-formed Lambda function " +
				"\nURI, such as the invoke_arn attribute of the aws_lambda_function " +
				"\nresource. Supported only for REQUEST authorizers. Must be between " +
				"\n1 and 2048 characters in length.",
		},
		{
			Field: "enable_simple_responses",
			Ex:    "(true/false)",
			Doc: "(Optional) Whether a Lambda authorizer returns a response " +
				"\nin a simple format. If enabled, the Lambda authorizer can return " +
				"\na boolean value instead of an IAM policy. Supported only for HTTP APIs.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "identity_sources",
			Doc: "(Optional) The identity sources for which authorization is requested. " +
				"\nFor REQUEST authorizers the value is a list of one or more mapping " +
				"\nexpressions of the specified request parameters. For JWT authorizers " +
				"\nthe single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	jwtConfigurationSchema := []types.Schema{
		{
			Field: "audience",
			Doc: "(Optional) A list of the intended recipients of the JWT. " +
				"\nA valid JWT must provide an aud that matches at least one entry in this list.",
		},
		{
			Field: "issuer",
			Doc: "(Optional) The base domain of the identity provider that issues " +
				"\nJSON Web Tokens, such as the endpoint attribute of the aws_cognito_user_pool resource.",
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
			Doc: "(Optional) The description for the deployment resource. " +
				"\nMust be less than or equal to 1024 characters in length.",
		},
		{
			Field: "triggers",
			Doc: "(Optional) A map of arbitrary keys and values that, when changed," +
				"\n will trigger a redeployment. To force a redeployment without changing" +
				"\n these keys/values, use the terraform taint command.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_deployment", blockName, resourceBlock)
}

func AWSAPIGatewayV2DomainNamePrompt() {
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
			Field: "domain_name",
			Doc:   "(Required) The domain name. Must be between 1 and 512 characters in length.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the domain name.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Green("Enter domain_name_configuration:\n(Required) The domain name configuration." +
		"\nThe domain_name_configuration supports the following arguments:" +
		"\n1.certificate_arn\n2.endpoint_type\n3.security_policy\n")

	domainNameConfigurationSchema := []types.Schema{
		{
			Field: "certificate_arn",
			Doc: "(Required) The ARN of an AWS-managed certificate that will be " +
				"\nused by the endpoint for the domain name. AWS Certificate Manager " +
				"\nis the only supported source. Use the aws_acm_certificate resource " +
				"\nto configure an ACM certificate.",
		},
		{
			Type:  "select",
			Field: "endpoint_type",
			Doc:   "(Required) The endpoint type.",
			Items: []string{"REGIONAL"},
		},
		{
			Type:  "select",
			Field: "security_policy",
			Doc: "(Required) The Transport Layer Security (TLS) version of the security " +
				"\npolicy for the domain name.",
			Items: []string{"TLS_1_2"},
		},
	}

	resourceBlock["domain_name_configuration"] = builder.PSOrder(types.ProvidePS(domainNameConfigurationSchema))

	color.Yellow("\nConfigure nested settings like mutual_tls_configuration [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_apigatewayv2_domain_name", blockName, resourceBlock)
		return
	}

	mutualTLSAuthenticationSchema := []types.Schema{
		{
			Field: "truststore_uri",
			Doc: "(Required) An Amazon S3 URL that specifies the truststore for " +
				"\nmutual TLS authentication, for example, s3://bucket-name/key-name. " +
				"\nThe truststore can contain certificates from public or private " +
				"\ncertificate authorities. To update the truststore, upload a new " +
				"\nversion to S3, and then update your custom domain name to use the new version.",
		},
		{
			Field: "truststore_version",
			Doc: "(Optional) The version of the S3 object that contains the truststore. " +
				"\nTo specify a version, you must have versioning enabled for the S3 bucket.",
		},
	}

	resourceBlock["mutual_tls_authentication"] = builder.PSOrder(types.ProvidePS(mutualTLSAuthenticationSchema))

	builder.ResourceBuilder("aws_apigatewayv2_domain_name", blockName, resourceBlock)
}

func AWSAPIGatewayV2IntegrationPrompt() {
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
			Field: "integration_type",
			Doc:   "(Required) The integration type of an integration.",
			Items: []string{"AWS", "AWS_PROXY", "HTTP", "HTTP_PROXY", "MOCK"},
		},
		{
			Field: "connection_id",
			Doc: "(Optional) The ID of the VPC link for a private integration. " +
				"\nSupported only for HTTP APIs. Must be between 1 and 1024 characters in length.",
		},
		{
			Type:  "select",
			Field: "connection_type",
			Doc:   "(Optional) The type of the network connection to the integration endpoint. Default is INTERNET.",
			Items: []string{"INTERNET", "VPC_LINK"},
		},
		{
			Type:  "select",
			Field: "content_handling_strategy",
			Doc: "(Optional) How to handle response payload content type conversions. " +
				"\nSupported only for WebSocket APIs.",
			Items: []string{"CONVERT_TO_BINARY", "CONVERT_TO_TEXT"},
		},
		{
			Field: "credentials_arn",
			Doc:   "(Optional) The credentials required for the integration, if any.",
		},
		{
			Field: "description",
			Doc:   "(Optional) The description of the integration.",
		},
		{
			Field: "integration_method",
			Doc:   "(Optional) The integration's HTTP method. Must be specified if integration_type is not MOCK.",
		},
		{
			Field: "integration_subtype",
			Doc: "(Optional) Specifies the AWS service action to invoke. Supported " +
				"\nonly for HTTP APIs when integration_type is AWS_PROXY. See the " +
				"\nAWS service integration reference documentation for supported values. " +
				"\nMust be between 1 and 128 characters in length.",
		},
		{
			Field: "integration_uri",
			Doc: "(Optional) The URI of the Lambda function for a Lambda proxy integration, " +
				"\nwhen integration_type is AWS_PROXY. For an HTTP integration, specify a " +
				"\nfully-qualified URL. For an HTTP API private integration, specify the " +
				"\nARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.",
		},
		{
			Type:  "select",
			Field: "passthrough_behavior",
			Doc: "(Optional) The pass-through behavior for incoming requests based on " +
				"\nthe Content-Type header in the request, and the available mapping " +
				"\ntemplates specified as the request_templates attribute. " +
				"\nDefault is WHEN_NO_MATCH. Supported only for WebSocket APIs.",
			Items: []string{"WHEN_NO_MATCH", "WHEN_NO_TEMPLATES", "NEVER"},
		},
		{
			Type:  "select",
			Field: "payload_format_version",
			Doc:   "(Optional) The format of the payload sent to an integration.",
			Items: []string{"1.0", "2.0"},
		},
		{
			Field: "request_parameters",
			Doc: "(Optional) For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend. For HTTP APIs with a specified integration_subtype, a key-value map specifying parameters that are passed to AWS_PROXY integrations. For HTTP APIs without a specified integration_subtype, a key-value map specifying how to transform HTTP requests before sending them to the backend." +
				"\nCheckout https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html",
		},
		{
			Field: "request_templates",
			Doc: "(Optional) A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs." +
				"\nCheckout https://velocity.apache.org/",
		},
		{
			Field: "template_selection_expression",
			Doc: "(Optional) The template selection expression for the integration." +
				"\nCheckout https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions",
		},
		{
			Field:     "timeout_milliseconds",
			Doc:       "(Optional) Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs. Terraform will only perform drift detection of its value when present in a configuration.",
			Validator: validators.MinMaxIntValidator(50, 29000),
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	tlsConfigSchema := []types.Schema{
		{
			Field: "server_name_to_verify",
			Doc:   "(Optional) If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.",
		},
	}

	resourceBlock["tls_config"] = builder.PSOrder(types.ProvidePS(tlsConfigSchema))

	builder.ResourceBuilder("aws_apigatewayv2_integration", blockName, resourceBlock)
}

func AWSAPIGatewayV2IntegrationResponsePrompt() {
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
			Field: "integration_id",
			Doc:   "(Required) The identifier of the aws_apigatewayv2_integration.",
		},
		{
			Field: "integration_response_key",
			Doc:   "(Required) The integration response key.",
		},
		{
			Type:  "select",
			Field: "content_handling_strategy",
			Doc:   "(Optional) How to handle response payload content type conversions.",
			Items: []string{"CONVERT_TO_BINARY", "CONVERT_TO_TEXT"},
		},
		{
			Field: "response_templates",
			Doc:   "(Optional) A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.",
		},
		{
			Field: "template_selection_expression",
			Doc:   "(Optional) The template selection expression for the integration response.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_integration_response", blockName, resourceBlock)
}

func AWSAPIGatewayV2ModelPrompt() {
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
			Field: "content_type",
			Ex:    "application/json",
			Doc:   "(Required) The content-type for the model. Must be between 1 and 256 characters in length.",
		},
		{
			Field: "name",
			Doc:   "(Required) The name of the model. Must be alphanumeric. Must be between 1 and 128 characters in length.",
		},
		{
			Field: "schema",
			Doc:   "(Required) The schema for the model. This should be a JSON schema draft 4 model. Must be less than or equal to 32768 characters in length.",
		},
		{
			Field: "description",
			Doc:   "(Optional) The description of the model. Must be between 1 and 128 characters in length.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_model", blockName, resourceBlock)
}

func AWSAPIGatewayV2RoutePrompt() {
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
			Field: "route_key",
			Doc:   "(Required) The route key for the route. For HTTP APIs, the route key can be either $default, or a combination of an HTTP method and resource path, for example, GET /pets",
		},
		{
			Field:     "api_key_required",
			Ex:        "(true/false)",
			Doc:       "(Optional) Boolean whether an API key is required for the route. Defaults to false.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "authorization_scopes",
			Doc:   "(Optional) The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.",
		},
		{
			Type:  "select",
			Field: "authorization_type",
			Doc:   "(Optional) The authorization type for the route. For WebSocket APIs, valid values are NONE for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer. For HTTP APIs, valid values are NONE for open access, or JWT for using JSON Web Tokens. Defaults to NONE.",
			Items: []string{"NONE", "AWS_IAM", "CUSTOM", "JWT"},
		},
		{
			Field: "authorizer_id",
			Doc:   "(Optional) The identifier of the aws_apigatewayv2_authorizer resource to be associated with this route, if the authorizationType is CUSTOM.",
		},
		{
			Field: "model_selection_permission",
			Doc:   "(Optional) The model selection expression for the route.",
		},
		{
			Field: "operation_name",
			Doc:   "(Optional) The operation name for the route. Must be between 1 and 64 characters in length.",
		},
		{
			Field:     "request_models",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) The request models for the route.",
			Validator: validators.RCValidator,
		},
		{
			Field: "route_response_selection_expression",
			Doc:   "(Optional) The route response selection expression for the route.",
		},
		{
			Field: "target",
			Doc:   "(Optional) The target for the route. Must be between 1 and 128 characters in length.",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_route", blockName, resourceBlock)
}

func AWSAPIGatewayV2RouteResponsePrompt() {
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
			Field: "route_id",
			Doc:   "(Required) The identifier of the aws_apigatewayv2_route.",
		},
		{
			Field: "route_response_key",
			Doc:   "(Required) The route response key.",
		},
		{
			Field: "model_selection_permission",
			Doc:   "(Optional) The model selection expression for the route.",
		},
		{
			Field:     "response_models",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) The response models for the route response.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_route_response", blockName, resourceBlock)
}

func AWSAPIGatewayV2StagePrompt() {
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
			Field: "name",
			Doc:   "(Required) The name of the stage. Must be between 1 and 128 characters in length.",
		},
		{
			Field:     "auto_deploy",
			Ex:        "(true/false)",
			Doc:       "(Optional) Whether updates to an API automatically trigger a new deployment. Defaults to false. Applicable for HTTP APIs.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "client_certificate_id",
			Doc:   "(Optional) The identifier of a client certificate for the stage. Use the aws_api_gateway_client_certificate resource to configure a client certificate. Supported only for WebSocket APIs.",
		},
		{
			Field: "deployment_id",
			Doc:   "(Optional) The deployment identifier of the stage. Use the aws_apigatewayv2_deployment resource to configure a deployment.",
		},
		{
			Field: "description",
			Doc:   "(Optional) The description for the stage. Must be less than or equal to 1024 characters in length.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the stage.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like access_log_settings/route_settings [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_apigatewayv2_stage", blockName, resourceBlock)
		return
	}

	accessLogSettingsSchema := []types.Schema{
		{
			Field: "destination_arn",
			Doc:   "(Required) The ARN of the CloudWatch Logs log group to receive access logs. Any trailing :* is trimmed from the ARN.",
		},
		{
			Field: "format",
			Doc: "(Required) A single line format of the access logs of data." +
				"\nCheckout https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-logging.html",
		},
	}

	resourceBlock["access_log_settings"] = builder.PSOrder(types.ProvidePS(accessLogSettingsSchema))

	defaultRouteSettingsSchema := []types.Schema{
		{
			Field:     "data_trace_enabled",
			Ex:        "(true/false)",
			Doc:       "(Optional) Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs. Defaults to false. Supported only for WebSocket APIs.",
			Validator: validators.BoolValidator,
		},
		{
			Field:     "detailed_metrics_enabled",
			Ex:        "(true/false)",
			Doc:       "(Optional) Whether detailed metrics are enabled for the default route. Defaults to false",
			Validator: validators.BoolValidator,
		},
		{
			Type:  "select",
			Field: "logging_level",
			Doc:   "(Optional) The logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs. Defaults to OFF. Supported only for WebSocket APIs. Terraform will only perform drift detection of its value when present in a configuration.",
			Items: []string{"ERROR", "INFO", "OFF"},
		},
		{
			Field:     "throttling_burst_limit",
			Doc:       "(Optional) The throttling burst limit for the default route.",
			Validator: validators.IntValidator,
		},
		{
			Field:     "throttling_rate_limit",
			Doc:       "(Optional) The throttling rate limit for the default route.",
			Validator: validators.IntValidator,
		},
	}

	resourceBlock["default_route_settings"] = builder.PSOrder(types.ProvidePS(defaultRouteSettingsSchema))

	routeSettingsSchema := []types.Schema{
		{
			Field: "route_key",
			Doc:   "(Required) Route key.",
		},
	}

	routeSettingsSchema = append(routeSettingsSchema, defaultRouteSettingsSchema...)

	resourceBlock["route_settings"] = builder.PSOrder(types.ProvidePS(routeSettingsSchema))

	builder.ResourceBuilder("aws_apigatewayv2_stage", blockName, resourceBlock)
}

func AWSAPIGatewayV2VPCLinkPrompt() {
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
			Doc: "(Required) The name of the VPC Link. Must be between 1 and 128 characters in length.",
		},
		{
			Field: "security_group_ids",
			Doc: "(Required) Security group IDs for the VPC Link.",
		},
		{
			Field: "subnet_ids",
			Ex: "[\"id1\",\"id2\"]",
			Doc: "(Required) Subnet IDs for the VPC Link.",
		},
		{
			Field: "tags",
			Ex: "k1=v1,k2=v2",
			Doc: "(Optional) A map of tags to assign to the VPC Link.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_apigatewayv2_vpc_link", blockName, resourceBlock)
}