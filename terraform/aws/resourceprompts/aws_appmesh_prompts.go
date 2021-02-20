package resourceprompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
)

func AWSAppmeshGatewayRoutePrompt() {
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
			Doc:   "(Required) The name to use for the gateway route. Must be between 1 and 255 characters in length.",
		},
		{
			Field: "mesh_name",
			Doc:   "(Required) The name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.",
		},
		{
			Field: "virtual_gateway_name",
			Doc:   "(Required) The name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.",
		},
		{
			Field: "mesh_owner",
			Doc:   "(Optional) The AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Green("\nEnter spec:\n(Required) The gateway route specification to apply." +
		"\nThe spec object supports the following:" +
		"\n1.grpc_route\n2.http_route\n3.http2_route")

	color.Green("\nEnter grpc_route:\n(Optional) The specification of a gRPC gateway route." +
		"\nThe The grpc_route, http_route and http2_route objects supports the following:" +
		"\n1.action\n2.match")

	color.Green("\nEnter action(of grpc_route):\n(Required) The action to take if a match is determined." +
		"\nThe action object supports the following:" +
		"\n1.target\n")

	color.Green("\nEnter target:\n(Required) The target that traffic is routed to when a request matches the gateway route." +
		"\nThe target object supports the following:" +
		"\n1.virtual_service")

	color.Green("\nEnter virtual_service:\n(Required) The virtual service gateway route target." +
		"\nThe virtual_service object supports the following:" +
		"\n1.virtual_service_name")

	virtualServiceNameSchema := []types.Schema{
		{
			Field: "virtual_service_name",
			Doc:   "(Required) The name of the virtual service that traffic is routed to. Must be between 1 and 255 characters in length.",
		},
	}

	serviceNameSchema := []types.Schema{
		{
			Field: "service_name",
			Doc:   "(Required) The fully qualified domain name for the service to match from the request.",
		},
	}

	virtualServiceNameBlockGRPC := builder.PSOrder(types.ProvidePS(virtualServiceNameSchema))
	serviceNameBlock := builder.PSOrder(types.ProvidePS(serviceNameSchema))

	specBlock := map[string]interface{}{
		"grpc_route": map[string]interface{}{
			"action": map[string]interface{}{
				"target": map[string]interface{}{
					"virtual_service": virtualServiceNameBlockGRPC,
				},
			},
			"match": map[string]interface{}{
				"service_name": serviceNameBlock,
			},
		},
	}

	color.Green("\nEnter virtual_service_name for http_route:\n")

	virtualServiceNameSchema = []types.Schema{
		{
			Field: "virtual_service_name",
			Doc:   "(Required) The name of the virtual service that traffic is routed to. Must be between 1 and 255 characters in length.",
		},
	}

	prefixSchema := []types.Schema{
		{
			Field: "prefix",
			Doc:   "(Required) Specifies the path to match requests with. This parameter must always start with /, which by itself matches all requests to the virtual service name.",
		},
	}
	virtualServiceNameBlockHTTP := builder.PSOrder(types.ProvidePS(virtualServiceNameSchema))
	prefixBlock := builder.PSOrder(types.ProvidePS(prefixSchema))

	specBlock["http_route"] = map[string]interface{}{
		"action": map[string]interface{}{
			"target": map[string]interface{}{
				"virtual_service": virtualServiceNameBlockHTTP,
			},
		},
		"match": map[string]interface{}{
			"prefix": prefixBlock,
		},
	}

	color.Green("\nEnter virtual_service_name for http2_route:\n")

	virtualServiceNameBlockHTTP2 := builder.PSOrder(types.ProvidePS(virtualServiceNameSchema))

	color.Green("\nEnter match for http2_route:\n")

	prefixBlock = builder.PSOrder(types.ProvidePS(prefixSchema))

	specBlock["http2_route"] = map[string]interface{}{
		"action": map[string]interface{}{
			"target": map[string]interface{}{
				"virtual_service": virtualServiceNameBlockHTTP2,
			},
		},
		"match": map[string]interface{}{
			"prefix": prefixBlock,
		},
	}

	resourceBlock["spec"] = specBlock

	builder.ResourceBuilder("aws_appmesh_gateway_route", blockName, specBlock)
}

func AWSAppMeshMeshPrompt() {
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
			Doc:   "(Required) The name to use for the service mesh. Must be between 1 and 255 characters in length.",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	color.Yellow("\nConfigure nested settings like spec [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_appmesh_mesh", blockName, resourceBlock)
		return
	}

	color.Green("Enter spec:\n(Optional) The service mesh specification to apply." +
		"\nThe spec object supports the following:" +
		"\n1.egress_filter")

	egressFilterSchema := []types.Schema{
		{
			Type:  "select",
			Field: "type",
			Doc:   "(Optional) The egress filter rules for the service mesh.",
			Items: []string{"DROP_ALL", "ALLOW_ALL"},
		},
	}

	egressFilterBlock := builder.PSOrder(types.ProvidePS(egressFilterSchema))

	resourceBlock["spec"] = egressFilterBlock

	builder.ResourceBuilder("aws_appmesh_mesh", blockName, resourceBlock)
}
