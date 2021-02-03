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
