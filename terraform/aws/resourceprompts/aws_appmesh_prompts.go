package resourceprompts
//
//import (
//	"fmt"
//	"github.com/fatih/color"
//	"github.com/g14a/tf/types"
//	"github.com/g14a/tf/validators"
//	"github.com/manifoldco/promptui"
//)
//
//func AWSAppmeshGatewayRoutePrompt() {
//	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")
//	blockPrompt := promptui.Prompt{
//		Label: "",
//	}
//
//	blockName, err := blockPrompt.Run()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	schema := []types.Schema{
//		{
//			Field: "name",
//			Doc:   "(Required) The name to use for the gateway route. Must be between 1 and 255 characters in length.",
//		},
//		{
//			Field: "mesh_name",
//			Doc:   "(Required) The name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.",
//		},
//		{
//			Field: "virtual_gateway_name",
//			Doc:   "(Required) The name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.",
//		},
//		{
//			Field: "mesh_owner",
//			Doc:   "(Optional) The AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.",
//		},
//		{
//			Field:     "tags",
//			Ex:        "k1=v1,k2=v2",
//			Doc:       "(Optional) A map of tags to assign to the resource.",
//			Validator: validators.RCValidator,
//		},
//	}
//}
