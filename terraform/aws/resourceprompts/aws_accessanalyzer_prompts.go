package resourceprompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
)

func AWSAccessAnalyzerAnalyzerPrompt() {
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
			Field: "analyzer_name",
			Doc: "(Required) Name of the Analyzer.",
		},
		{
			Field: "tags",
			Ex: "k1=v1,k2=v2",
			Doc: "(Optional) Key-value map of resource tags.",
			Validator: validators.RCValidator,
		},
		{
			Type: "select",
			Field: "type",
			Doc: "(Optional) Type of Analyzer.",
			Items: []string{"ACCOUNT","ORGANIZATION"},
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_accessanalyzer_analyzer", blockName, resourceBlock)
}