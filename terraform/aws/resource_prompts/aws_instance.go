package resource_prompts

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"tf/builder"
)

func AWSInstanceBuilderPrompt() {
	var prompt promptui.Prompt

	prompt = promptui.Prompt{
		Label:  "Enter block name(required) e.g. web",
	}

	blockName, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompt = promptui.Prompt{
		Label: "Enter ami(required)",
	}

	ami, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompt = promptui.Prompt{
		Label: "Enter instance_type(required) e.g. t2.micro",
	}

	instanceType, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	resourceBlock := map[string]interface{} {
		"ami": ami,
		"instance_type": instanceType,
	}

	builder.ResourceBuilder("aws_instance", blockName, resourceBlock)
}

func AWSVPCBuilderPrompt()  {

}