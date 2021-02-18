package resourceprompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/manifoldco/promptui"
)

func AWSPrometheusWorkspacePrompt() {
	color.Green("\nEnter block name(Required) e.g. foo/bar\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	color.Yellow("\nThis AWS functionality is in Preview and may change " +
		"\nbefore General Availability release. Backwards compatibility is not " +
		"\bguaranteed between Terraform AWS Provider releases.\n")

	schema := []types.Schema{
		{
			Field: "alias",
			Doc: "(Optional) The alias of the prometheus workspace." +
				"\nCheckout https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-onboard-create-workspace.html",
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_prometheus_workspace", blockName, resourceBlock)
}
