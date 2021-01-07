package resource_prompts

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"tf/builder"
	"tf/types"
	"tf/utils"
)

func AWSAMIPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
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
		Label: "Enter name:\n(Required) A region-unique name for the AMI.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A longer, human-readable description for the AMI.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["ena_support"] = types.TfPrompt{
		Label: "Enter ena_support(true/false):\n(Optional) Specifies whether enhanced networking with ENA is enabled. Defaults to false",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ena_support")

	prompts["root_device_name"] = types.TfPrompt{
		Label: "Enter root_device_name:\n(Optional) The name of the root device (for example, /dev/sda1, or /dev/xvda",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "root_device_name")

	prompts["architecture"] = types.TfPrompt{
		Label: "Enter architecture:\n(Optional) Machine architecture for created instances. Defaults to \"x86_64\".",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "architecture")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	color.Green("\nEnter virtualization_type:\n(Optional) Keyword to choose what virtualization mode created " +
		"\ninstances will use. Can be either \"paravirtual\" (the default) or \"hvm\". The choice of virtualization " +
		"\ntype changes the set of further arguments that are required, as described below.")

	vTypePrompt := promptui.Select{
		Label: "",
		Items: []string{"paravirtual","hvm"},
	}

	_,vType, err := vTypePrompt.Run()

	if vType == "paravirtual" {
		prompts["image_location"] = types.TfPrompt{
			Label: "Enter image_location:\n(Required) Path to an S3 object containing an image manifest, e.g. created by the ec2-upload-bundle command in the EC2 command line tools.",
			Prompt: promptui.Prompt{
				Label: "",
			},
		}
		promptOrder = append(promptOrder, "image_location")

		prompts["kernel_id"] = types.TfPrompt{
			Label: "Enter kernel_id:\n(Required) The id of the kernel image (AKI) that will be used as the paravirtual kernel in created instances.",
			Prompt: promptui.Prompt{
				Label: "",
			},
		}
		promptOrder = append(promptOrder, "kernel_id")

		prompts["ramdisk_id"] = types.TfPrompt{
			Label: "Enter ramdisk_id:\n(Optional) The id of an initrd image (ARI) that will be used when booting the created instances.",
			Prompt: promptui.Prompt{
				Label: "",
			},
		}
		promptOrder = append(promptOrder, "ramdisk_id")
	} else if vType == "hvm" {
		prompts["sriov_net_support"] = types.TfPrompt{
			Label: "Enter sriov_net_support:\n(Optional) When set to \"simple\" (the default), enables enhanced networking for created instances. No other value is supported at this time.",
			Prompt: promptui.Prompt{
				Label: "",
			},
		}
		promptOrder = append(promptOrder, "sriov_net_support")
	}

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like ebs_block_device/ephemeral_block_device etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_ami", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter ebs_block_device:\n(Optional) Nested block describing an EBS block device that should be attached to created instances. " +
		"\nThe structure of this block is described below.")

	ebsBlockDevicePrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	ebsBlockDevicePrompt["device_name"] = types.TfPrompt{
		Label: "Enter device_name:\n(Required) The path at which the device is exposed to created instances.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "device_name")

	ebsBlockDevicePrompt["delete_on_termination"] = types.TfPrompt{
		Label: "Enter delete_on_termination(true/false):\n(Optional) Boolean controlling whether the EBS volumes created " +
			"\nto support each created instance will be deleted once that instance is terminated.",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "delete_on_termination")

	color.Yellow("\nYou can specify encrypted or snapshot_id but not both.\n")

	ebsBlockDevicePrompt["encrypted"] = types.TfPrompt{
		Label: "Enter encrypted(true/false):\n(Optional) Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with snapshot_id",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "encrypted")

	ebsBlockDevicePrompt["snapshot_id"] = types.TfPrompt{
		Label: "Enter snapshot_id:\n(Optional) The id of an EBS snapshot that will be used to initialize the created EBS volumes. " +
			"\nIf set, the volume_size attribute must be at least as large as the referenced snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "snapshot_id")

	ebsBlockDevicePrompt["iops"] = types.TfPrompt{
		Label: "Enter iops:\n(Required only when volume_type is \"io1/io2\") Number of I/O operations per second the created volumes will support.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "iops")

	ebsBlockDevicePrompt["volume_size"] = types.TfPrompt{
		Label: "Enter volume_size:\n(Required unless snapshot_id is set) The size of created volumes in GiB. If snapshot_id is set and " +
			"\nvolume_size is omitted then the volume will have the same size as the selected snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "volume_size")

	ebsBlockDevicePrompt["kms_key_id"] = types.TfPrompt{
		Label: "Enter kms_key_id:\n(Optional) The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the " +
			"\nsnapshots of an image during a copy operation. This parameter is only required if you want to use a non-default CMK; " +
			"\nif this parameter is not specified, the default CMK for EBS is used",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "kms_key_id")

	volumeTypeSelect := map[string]types.TfSelect{}
	var nestedSelectOrder []string

	volumeTypeSelect["volume_type"] = types.TfSelect{
		Label: "(Optional) The type of EBS volume to create. Defaults to \"standard\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"standard","io1","io2","gp2"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "volume_type")

	resourceBlock["ebs_block_device"] = builder.NestedPSOrder(nestedPromptOrder, nestedSelectOrder, ebsBlockDevicePrompt, volumeTypeSelect)

	color.Green("\nEnter ephemeral_block_device:\n(Optional) Nested block describing an ephemeral block device that should be attached to created instances." +
		"\nThe ephemeral_block_device supports the following structure:" +
		"\n1.device_name\n2.virtual_name\n")

	ephemeralBlockDevicePrompt := map[string]types.TfPrompt{}

	ephemeralBlockDevicePrompt["device_name"] = types.TfPrompt{
		Label: "Enter device_name:\n(Required) The path at which the device is exposed to created instances.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "device_name")

	ephemeralBlockDevicePrompt["virtual_name"] = types.TfPrompt{
		Label: "Enter virtual_name:\n(Required) A name for the ephemeral device, of the form \"ephemeralN\" where N is a volume number starting from zero.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "virtual_name")

	resourceBlock["ephemeral_block_device"] = builder.NestedPSOrder(nestedPromptOrder[len(nestedPromptOrder)-2:], nil, ephemeralBlockDevicePrompt, nil)

	color.Green("\nEnter timeouts block:\n" +
		"The timeout block supports the following arguments:" +
		"\n1.create\n2.delete\n3.update")

	timeoutsPrompt := map[string]types.TfPrompt{}

	timeoutsPrompt["create"] = types.TfPrompt{
		Label: "Enter create: e.g. 40m\n(Default 40m) How long to wait for an RDS event notification subscription to be ready.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "create")

	timeoutsPrompt["update"] = types.TfPrompt{
		Label: "Enter update: e.g. 40m\n(Default 40m) How long to wait for an RDS event notification subscription to be updated.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "update")

	timeoutsPrompt["delete"] = types.TfPrompt{
		Label: "Enter delete: e.g. 40m\n(Default 40m) How long to wait for an RDS event notification subscription to be deleted.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "delete")

	resourceBlock["timeouts"] = builder.NestedPSOrder(nestedPromptOrder[len(nestedPromptOrder)-3:], nil, timeoutsPrompt, nil)

	builder.ResourceBuilder("aws_ami", blockName, resourceBlock)

}

func AWSAMICopyPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
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
		Label: "Enter name:\n(Required) A region-unique name for the AMI.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["source_ami_id"] = types.TfPrompt{
		Label: "Enter source_ami_id:\n(Required) The id of the AMI to copy. This id must be valid in the region given by source_ami_region",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_ami_id")

	prompts["source_ami_region"] = types.TfPrompt{
		Label: "Enter source_ami_region:\n(Required) The region from which the AMI will be copied. " +
			"\nThis may be the same as the AWS provider region in order to create a copy within the same region.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_ami_region")

	prompts["encrypted"] = types.TfPrompt{
		Label: "Enter encrypted(true/false):\n(Optional) Specifies whether the destination snapshots of the copied image should be encrypted. Defaults to false",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "encrypted")

	prompts["kms_key_id"] = types.TfPrompt{
		Label: "Enter kms_key_id:\n(Optional) The full ARN of the KMS Key to use when encrypting the snapshots " +
			"\nof an image during a copy operation. If not specified, then the default AWS KMS Key will be used",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "kms_key_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
			Validate: utils.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ami_copy",blockName, resourceBlock)

}