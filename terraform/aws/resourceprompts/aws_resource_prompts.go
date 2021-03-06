package resourceprompts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
)

func AWSELBPrompt() {
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
		Label: "Enter name:\n(Optional) The name of the ELB. By default generated by Terraform.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["instances"] = types.TfPrompt{
		Label: "Enter instances:\n(Optional) A list of instance ids to place in the ELB pool. e.g. [\"id1\",\"id2\"]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instances")

	prompts["cross_zone_load_balancing"] = types.TfPrompt{
		Label: "Enter cross_zone_load_balancing(true/false):\n(Optional) Enable cross-zone load balancing. Default: true",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "cross_zone_load_balancing")

	prompts["idle_timeout"] = types.TfPrompt{
		Label: "Enter idle_timeout:\n(Optional) The time in seconds that the connection is allowed to be idle. Default: 60",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "idle_timeout")

	prompts["connection_draining"] = types.TfPrompt{
		Label: "Enter connection_draining(true/false):\n(Optional) Boolean to enable connection draining. Default: false",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "connection_draining")

	prompts["connection_draining_timeout"] = types.TfPrompt{
		Label: "Enter connection_draining_timeout(true/false):\n(Optional) The time in seconds to allow for connections to drain. Default: 300",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "connection_draining_timeout")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like listener/access_logs/website etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_s3_bucket", blockName, resourceBlock)
		return
	}

	listenerPrompt := map[string]types.TfPrompt{}
	listenerSelect := map[string]types.TfSelect{}

	var nestedPromptOrder, nestedSelectOrder []string

	listenerPrompt["instance_port"] = types.TfPrompt{
		Label: "Enter instance_port:\n(Required) The port on the instance to route to",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "instance_port")

	listenerPrompt["lb_port"] = types.TfPrompt{
		Label: "Enter lb_port:\n(Required) The port to listen on for the load balancer",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "lb_port")

	listenerPrompt["ssl_certificate_id"] = types.TfPrompt{
		Label: "Enter ssl_certificate_id:\n(Optional) The ARN of an SSL certificate you have uploaded to AWS IAM. Note ECDSA-specific restrictions below. Only valid when lb_protocol is either HTTPS or SSL",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "ssl_certificate_id")

	listenerSelect["instance_protocol"] = types.TfSelect{
		Label: "Enter instance_protocol:\n(Required) The protocol to use to the instance.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"HTTP", "HTTPS", "TCP", "SSL"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "instance_protocol")

	listenerSelect["lb_protocol"] = types.TfSelect{
		Label: "Enter lb_protocol:\n(Required) The protocol to listen on.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"HTTP", "HTTPS", "TCP", "SSL"},
		},
	}

	nestedSelectOrder = append(nestedSelectOrder, "lb_protocol")

	resourceBlock["listener"] = builder.PSOrder(nestedPromptOrder, nestedSelectOrder, listenerPrompt, listenerSelect)

	accessLogsPrompt := map[string]types.TfPrompt{}
	accessLogsPrompt["bucket"] = types.TfPrompt{
		Label: "Enter bucket:\n(Required) The S3 bucket name to store the logs in.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "bucket")

	accessLogsPrompt["bucket_prefix"] = types.TfPrompt{
		Label: "Enter bucket_prefix:\n(Optional) The S3 bucket prefix. Logs are stored in the root if not configured.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "bucket_prefix")

	accessLogsPrompt["interval"] = types.TfPrompt{
		Label: "Enter interval:\n(Optional) The publishing interval in minutes. Default: 60 minutes.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "interval")

	accessLogsPrompt["enabled"] = types.TfPrompt{
		Label: "Enter enabled:(true/false)\n(Optional) Boolean to enable / disable access_logs. Default is true",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "enabled")

	resourceBlock["access_logs"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-4:], nil, accessLogsPrompt, nil)

	healthCheckPrompt := map[string]types.TfPrompt{}
	healthCheckPrompt["healthy_threshold"] = types.TfPrompt{
		Label: "Enter healthy_threshold:\n(Required) The number of checks before the instance is declared healthy.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "healthy_threshold")

	healthCheckPrompt["unhealthy_threshold"] = types.TfPrompt{
		Label: "Enter unhealthy_threshold:\n(Required) The number of checks before the instance is declared unhealthy.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "unhealthy_threshold")

	healthCheckPrompt["target"] = types.TfPrompt{
		Label: "Enter target:\n(Required) The target of the check. Valid pattern is \"${PROTOCOL}:${PORT}${PATH}\" e.g. HTTP:8000/",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "target")

	healthCheckPrompt["timeout"] = types.TfPrompt{
		Label: "Enter timeout:\n(Required) The length of time before the check times out.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "timeout")

	healthCheckPrompt["interval"] = types.TfPrompt{
		Label: "Enter interval:\n(Required) The interval between checks.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "interval")

	resourceBlock["health_check"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-5:], nil, healthCheckPrompt, nil)

	builder.ResourceBuilder("aws_elb", blockName, resourceBlock)
}
