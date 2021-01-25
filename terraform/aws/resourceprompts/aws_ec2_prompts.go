package resourceprompts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/g14a/tf/builder"
	"github.com/g14a/tf/types"
	"github.com/g14a/tf/validators"
	"github.com/manifoldco/promptui"
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

	schema := []types.Schema{
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) A region-unique name for the AMI.",
		},
		{
			Field: "description",
			Ex:    "",
			Doc:   "(Optional) A longer, human-readable description for the AMI.",
		},
		{
			Field:     "ena_support",
			Ex:        "(true/false)",
			Doc:       "(Optional) Specifies whether enhanced networking with ENA is enabled. Defaults to false",
			Validator: validators.BoolValidator,
		},
		{
			Field: "root_device_name",
			Ex:    "",
			Doc:   "(Optional) The name of the root device (for example, /dev/sda1, or /dev/xvda)",
		},
		{
			Type:  "select",
			Field: "architecture",
			Doc:   "(Optional) Machine architecture for created instances. Defaults to \"x86_64\".",
			Items: []string{"i386", "x86_64", "arm64"},
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: validators.RCValidator,
		},
	}

	color.Green("\nEnter virtualization_type:\n(Optional) Keyword to choose what virtualization mode created " +
		"\ninstances will use. Can be either \"paravirtual\" (the default) or \"hvm\". The choice of virtualization " +
		"\ntype changes the set of further arguments that are required, as described below.")

	vTypePrompt := promptui.Select{
		Label: "",
		Items: []string{"paravirtual", "hvm"},
	}

	_, vType, err := vTypePrompt.Run()

	if vType == "paravirtual" {
		schema = append(schema, []types.Schema{
			{
				Field: "image_location",
				Ex:    "",
				Doc: "(Required) Path to an S3 object containing an image manifest, e.g. " +
					"\ncreated by the ec2-upload-bundle command in the EC2 command line tools.",
			},
			{
				Field: "kernel_id",
				Ex:    "",
				Doc: "(Required) The id of the kernel image (AKI) that will be used as the " +
					"\nparavirtual kernel in created instances.",
			},
			{
				Field: "ramdisk_id",
				Ex:    "",
				Doc:   "(Optional) The id of an initrd image (ARI) that will be used when booting the created instances.",
			},
		}...)

	} else if vType == "hvm" {
		schema = append(schema, []types.Schema{
			{
				Field: "sriov_net_support",
				Ex:    "",
				Doc: "(Optional) When set to \"simple\" (the default), enables enhanced networking " +
					"\nfor created instances. No other value is supported at this time.",
			},
		}...)
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

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

	color.Green("\nEnter ebs_block_device:\n(Optional) Nested block describing an EBS block device " +
		"\nthat should be attached to created instances." +
		"\nThe ebs_block_device supports the following arguments:" +
		"\n1.device_name\n2.delete_on_termination\n3.encrypted\n4.iops\n5.snapshot_id\n6.volume_size\n7.volume_type\n8.kms_key_id")

	color.Yellow("\nYou can specify encrypted or snapshot_id but not both.\n")

	ebsBlockDeviceSchema := []types.Schema{
		{
			Field: "device_name",
			Ex:    "",
			Doc:   "(Required) The path at which the device is exposed to created instances.",
		},
		{
			Field: "delete_on_termination",
			Ex:    "(true/false)",
			Doc: "(Optional) Boolean controlling whether the EBS volumes created to support each created " +
				"\ninstance will be deleted once that instance is terminated.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "encrypted",
			Ex:    "(true/false)",
			Doc: "(Optional) Boolean controlling whether the created EBS volumes will be encrypted. " +
				"\nCan't be used with snapshot_id.",
			Validator: validators.BoolValidator,
		},
		{
			Field: "iops",
			Ex:    "",
			Doc: "(Required only when volume_type is io1 or io2) Number of I/O operations per second " +
				"\nthe created volumes will support.",
		},
		{
			Field: "snapshot_id",
			Ex:    "",
			Doc: "(Optional) The id of an EBS snapshot that will be used to initialize the created " +
				"\nEBS volumes. If set, the volume_size attribute must be at least as large as the referenced snapshot.",
		},
		{
			Field: "volume_size",
			Ex:    "",
			Doc: "(Required unless snapshot_id is set) The size of created volumes in GiB. " +
				"\nIf snapshot_id is set and volume_size is omitted then the volume will have " +
				"\nthe same size as the selected snapshot.",
		},
		{
			Field: "volume_type",
			Doc:   "(Optional) The type of EBS volume to create. Defaults to standard",
			Items: []string{"standard", "io1", "io2", "gp2"},
		},
		{
			Field: "kms_key_id",
			Ex:    "",
			Doc: "(Optional) The full ARN of the AWS Key Management Service (AWS KMS) " +
				"\nCMK to use when encrypting the snapshots of an image during a copy operation. " +
				"\nThis parameter is only required if you want to use a non-default CMK; " +
				"\nif this parameter is not specified, the default CMK for EBS is used",
		},
	}

	resourceBlock["ebs_block_device"] = builder.PSOrder(types.ProvidePS(ebsBlockDeviceSchema))

	color.Green("\nEnter ephemeral_block_device:\n(Optional) Nested block describing an ephemeral block device that should be attached to created instances." +
		"\nThe ephemeral_block_device supports the following structure:" +
		"\n1.device_name\n2.virtual_name\n")

	ephemeralBlockDeviceSchema := []types.Schema{
		{
			Field: "device_name",
			Ex:    "",
			Doc:   "(Required) The path at which the device is exposed to created instances.",
		},
		{
			Field: "virtual_name",
			Ex:    "",
			Doc:   "(Required) A name for the ephemeral device, of the form \"ephemeralN\" where N is a volume number starting from zero.",
		},
	}

	resourceBlock["ephemeral_block_device"] = builder.PSOrder(types.ProvidePS(ephemeralBlockDeviceSchema))

	color.Green("\nEnter timeouts block:\n" +
		"The timeout block supports the following arguments:" +
		"\n1.create\n2.delete\n3.update")

	timeoutSchema := []types.Schema{
		{
			Field: "create",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for creating the AMI",
		},
		{
			Field: "update",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for updating the AMI",
		},
		{
			Field: "delete",
			Ex:    "60s | 10m | 2h",
			Doc:   "Used for deregistering the AMI",
		},
	}

	resourceBlock["timeouts"] = builder.PSOrder(types.ProvidePS(timeoutSchema))

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

	schema := []types.Schema{
		{
			Field: "name",
			Ex:    "",
			Doc:   "(Required) A region-unique name for the AMI.",
		},
		{
			Field: "source_ami_id",
			Ex:    "",
			Doc:   "(Required) The id of the AMI to copy. This id must be valid in the region given by source_ami_region",
		},
		{
			Field: "source_ami_region",
			Ex:    "",
			Doc:   "(Required) The region from which the AMI will be copied. This may be the same as the AWS provider region in order to create a copy within the same region.",
		},
		{
			Field:     "encrypted",
			Ex:        "(true/false)",
			Doc:       "(Optional) Specifies whether the destination snapshots of the copied image should be encrypted. Defaults to false",
			Validator: validators.BoolValidator,
		},
		{
			Field: "kms_key_id",
			Ex:    "",
			Doc:   "(Optional) The full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used",
		},
		{
			Field:     "tags",
			Ex:        "k1=v1,k2=v2",
			Doc:       "(Optional) A map of tags to assign to the resource.",
			Validator: validators.RCValidator,
		},
	}

	resourceBlock := builder.PSOrder(types.ProvidePS(schema))

	builder.ResourceBuilder("aws_ami_copy", blockName, resourceBlock)

}

func AWSAMIFromInstancePrompt() {
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
		Label: "Enter name:\n(Required) A region-unique name for the AMI",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["source_instance_id"] = types.TfPrompt{
		Label: "Enter source_instance_id:\n(Required) The id of the instance to use as the basis of the AMI.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_instance_id")

	prompts["snapshot_without_reboot"] = types.TfPrompt{
		Label: "Enter snapshot_without_reboot(true/false):\n(Optional) Boolean that overrides the behavior of stopping the instance before snapshotting. " +
			"\nThis is risky since it may cause a snapshot of an inconsistent filesystem state, but can be used to avoid downtime if the user otherwise " +
			"\nguarantees that no filesystem writes will be underway at the time of snapshot.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "snapshot_without_reboot")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like timeouts etc [y/n]?\n\n", "text")

	ynPrompt := promptui.Prompt{
		Label: "",
	}

	yn, err := ynPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	if yn == "n" || yn == "" {
		builder.ResourceBuilder("aws_ami_from_instance", blockName, resourceBlock)
		return
	}

	color.Green("\nEnter timeouts block:\n" +
		"The timeout block supports the following arguments:" +
		"\n1.create\n2.delete\n3.update")

	timeoutsPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	timeoutsPrompt["create"] = types.TfPrompt{
		Label: "Enter create: e.g. 40m\n(Defaults to 40 mins) Used when creating the AMI",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "create")

	timeoutsPrompt["update"] = types.TfPrompt{
		Label: "Enter update: e.g. 40m\n(Defaults to 40 mins) Used when updating the AMI",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "update")

	timeoutsPrompt["delete"] = types.TfPrompt{
		Label: "Enter delete: e.g. 40m\n(Defaults to 90 mins) Used when deregistering the AMI",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "delete")

	resourceBlock["timeouts"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-3:], nil, timeoutsPrompt, nil)

	builder.ResourceBuilder("aws_ami_from_instance", blockName, resourceBlock)

}

func AWSAMILaunchPermissionPrompt() {
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

	prompts["image_id"] = types.TfPrompt{
		Label: "Enter image_id:\n(Required) A region-unique name for the AMI.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "image_id")

	prompts["account_id"] = types.TfPrompt{
		Label: "Enter account_id:\n(required) An AWS Account ID to add launch permissions.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "account_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ami_launch_permission", blockName, resourceBlock)
}

func AWSEBSDefaultKMSKeyPrompt() {
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

	prompts["key_arn"] = types.TfPrompt{
		Label: "Enter key_arn:\n(Required, ForceNew) The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "key_arn")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ebs_default_kms_key", blockName, resourceBlock)
}

func AWSEBSEncryptionByDefaultPrompt() {
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

	prompts["enabled"] = types.TfPrompt{
		Label: "Enter enabled(true/false):\n(Required, ForceNew) The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "enabled")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ebs_encryption_by_default", blockName, resourceBlock)
}

func AWSEBSSnapshotPrompt() {
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

	prompts["volume_id"] = types.TfPrompt{
		Label: "Enter volume_id(true/false):\n(Required) The Volume ID of which to make a snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "volume_id")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A description of what the snapshot is.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the snapshot",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter timeouts block:\n" +
		"The timeout block supports the following arguments:" +
		"\n1.create\n2.delete\n")

	timeoutsPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	timeoutsPrompt["create"] = types.TfPrompt{
		Label: "Enter create: e.g. 40m\n(Defaults to 40 mins) Used when creating the AMI",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "create")

	timeoutsPrompt["delete"] = types.TfPrompt{
		Label: "Enter delete: e.g. 40m\n(Defaults to 90 mins) Used when deregistering the AMI",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "delete")

	resourceBlock["timeouts"] = builder.PSOrder(nestedPromptOrder, nil, timeoutsPrompt, nil)

	builder.ResourceBuilder("aws_ebs_snapshot", blockName, resourceBlock)
}

func AWSEBSSnapshotCopyPrompt() {
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

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A description of what the snapshot is.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["encrypted"] = types.TfPrompt{
		Label: "Enter encrypted(true/false):\nWhether the snapshot is encrypted.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "encrypted")

	prompts["kms_key_id"] = types.TfPrompt{
		Label: "Enter kms_key_id:\nThe ARN for the KMS encryption key.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "kms_key_id")

	prompts["source_snapshot_id"] = types.TfPrompt{
		Label: "Enter source_snapshot_id:\nThe ARN for the snapshot to be copied.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_snapshot_id")

	prompts["source_region"] = types.TfPrompt{
		Label: "Enter source_region:\nThe region of the source snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_region")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\nA map of tags for the snapshot.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ebs_snapshot_copy", blockName, resourceBlock)
}

func AWSEBSVolumePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["availability_zone"] = types.TfPrompt{
		Label: "Enter availability_zone:\n(Required) The AZ where the EBS volume will exist.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "availability_zone")

	prompts["iops"] = types.TfPrompt{
		Label: "Enter iops:\n(Optional) The amount of IOPS to provision for the disk. Only valid for type of io1 or io2",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "iops")

	prompts["multi_attach_enabled"] = types.TfPrompt{
		Label: "Enter multi_attach_enabled(true/false):\n(Optional) Specifies whether to enable Amazon EBS Multi-Attach. " +
			"\nMulti-Attach is supported exclusively on io1 volumes.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "multi_attach_enabled")

	prompts["snapshot_id"] = types.TfPrompt{
		Label: "Enter snapshot_id:\n(Optional) A snapshot to base the EBS volume off of.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "snapshot_id")

	prompts["outpost_arn"] = types.TfPrompt{
		Label: "Enter outpost_arn:\n(Optional) The Amazon Resource Name (ARN) of the Outpost.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "outpost_arn")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["kms_key_id"] = types.TfPrompt{
		Label: "Enter kms_key_id:\n(Optional) The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "kms_key_id")

	color.Yellow("\nWhen changing the size, iops or type of an instance, there are considerations to be aware of that Amazon have written about this." +
		"\nCheckout http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/considerations.html\n")

	prompts["encrypted"] = types.TfPrompt{
		Label: "Enter encrypted(true/false):\n(Optional) If true, the disk will be encrypted.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "encrypted")

	prompts["size"] = types.TfPrompt{
		Label: "Enter size(true/false):\n(Optional) The size of the drive in GiBs",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "size")

	selects := map[string]types.TfSelect{}

	selects["type"] = types.TfSelect{
		Label: "Enter type:\n(Optional) The type of EBS volume. Defaults to gp2",
		Select: promptui.Select{
			Label: "",
			Items: []string{"standard", "gp2", "io1", "io2", "sc1", "st1"},
		},
	}
	selectOrder = append(selectOrder, "type")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_ebs_volume", blockName, resourceBlock)
}

func AWSEC2AvailabilityZoneGroupPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["group_name"] = types.TfPrompt{
		Label: "Enter group_name:\n:(Required) Name of the Availability Zone Group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "group_name")

	selects := map[string]types.TfSelect{}

	selects["opt_in_status"] = types.TfSelect{
		Label: "Enter opt_in_status:\n(Required) Indicates whether to enable or disable Availability Zone Group. Valid values: opted-in or not-opted-in",
		Select: promptui.Select{
			Label: "",
			Items: []string{"opted-in", "not-opted-in"},
		},
	}
	selectOrder = append(selectOrder, "opt_in_status")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_ec2_availability_zone_group", blockName, resourceBlock)
}

func AWSEC2CapacityReservationPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["availability_zone"] = types.TfPrompt{
		Label: "Enter availability_zone:\n(Required) The Availability Zone in which to create the Capacity Reservation.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "availability_zone")

	prompts["ebs_optimized"] = types.TfPrompt{
		Label: "Enter ebs_optimized(true/false):\n(Optional) Indicates whether the Capacity Reservation supports EBS-optimized instances.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ebs_optimized")

	prompts["end_date"] = types.TfPrompt{
		Label: "Enter end_date:\n(Optional) The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, " +
			"\nthe reserved capacity is released and you can no longer launch instances into it.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "end_date")

	prompts["ephemeral_storage"] = types.TfPrompt{
		Label: "Enter ephemeral_storage(true/false):\n(Optional) Indicates whether the Capacity Reservation supports instances with temporary, block-level storage",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ephemeral_storage")

	prompts["instance_count"] = types.TfPrompt{
		Label: "Enter instance_count:\n(Required) The number of instances for which to reserve capacity.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "instance_count")

	prompts["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type:\n(Required) The instance type for which to reserve capacity.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_type")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}

	selects["end_date_type"] = types.TfSelect{
		Label: "Enter end_date_type:\n(Optional) Indicates the way in which the Capacity Reservation ends.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"unlimited", "limited"},
		},
	}
	selectOrder = append(selectOrder, "end_date_type")

	selects["instance_match_criteria"] = types.TfSelect{
		Label: "Enter instance_match_criteria:\n(Optional) Indicates the type of instance launches that the Capacity Reservation accepts.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"open", "targeted"},
		},
	}
	selectOrder = append(selectOrder, "instance_match_criteria")

	selects["instance_platform"] = types.TfSelect{
		Label: "Enter instance_platform:\n(Required) The type of operating system for which to reserve capacity.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"Linux/Unix", "Red Hat Enterprise Linux", "SUSE Linux", "Windows", "Windows with SQL Server", "Windows with SQL Server Enterprise", "Windows with SQL Server Standard", "Windows with SQL Server Web"},
		},
	}
	selectOrder = append(selectOrder, "instance_platform")

	selects["tenancy"] = types.TfSelect{
		Label: "Enter tenancy:\n(Required) The type of operating system for which to reserve capacity.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"default", "dedicated"},
		},
	}
	selectOrder = append(selectOrder, "tenancy")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_ec2_capacity_reservation", blockName, resourceBlock)
}

func AWSEC2CarrierGatewayPrompt() {
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

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) The ID of the VPC to associate with the carrier gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_carrier_gateway", blockName, resourceBlock)

}

func AWSEC2ClientVPNAuthorizationRulePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["client_vpn_endpoint_id"] = types.TfPrompt{
		Label: "Enter client_vpn_endpoint_id:\n(Required) The ID of the Client VPN endpoint.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "client_vpn_endpoint_id")

	prompts["target_network_cidr"] = types.TfPrompt{
		Label: "Enter target_network_cidr:\n(Required) The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "target_network_cidr")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A brief description of the authorization rule.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	selects := map[string]types.TfSelect{}

	selects["access_group_id"] = types.TfSelect{
		Label: "Enter access_group_id:\n(Optional) The ID of the group to which the authorization rule grants access.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"access_group_id", "authorize_all_groups"},
		},
	}
	selectOrder = append(selectOrder, "access_group_id")

	selects["authorize_all_groups"] = types.TfSelect{
		Label: "Enter authorize_all_groups:\n(Optional) Indicates whether the authorization rule grants access to all clients.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"access_group_id", "authorize_all_groups"},
		},
	}
	selectOrder = append(selectOrder, "authorize_all_groups")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_ec2_client_vpn_authorization_rule", blockName, resourceBlock)
}

func AWSEC2ClientVPNEndpointPrompt() {
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

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Name of the repository.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["server_certificate_arn"] = types.TfPrompt{
		Label: "Enter server_certificate_arn:\n(Required) The ARN of the ACM server certificate.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "server_certificate_arn")

	prompts["client_cidr_block"] = types.TfPrompt{
		Label: "Enter client_cidr_block:\n(Required) The IPv4 address range, in CIDR notation, from which to assign " +
			"\nclient IP addresses. The address range cannot overlap with the local CIDR of the " +
			"\nVPC in which the associated subnet is located, or the routes that you add manually. " +
			"\nThe address range cannot be changed after the Client VPN endpoint has been created. " +
			"\nThe CIDR block should be /22 or greater.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "client_cidr_block")

	prompts["dns_servers"] = types.TfPrompt{
		Label: "Enter dns_servers e.g.[\"s1\",\"s2\"]:\n(Optional) Information about the DNS servers to be used for DNS resolution. " +
			"\nA Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address " +
			"\nof the VPC that is to be associated with Client VPN endpoint is used as the DNS server.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "dns_servers")

	prompts["split_tunnel"] = types.TfPrompt{
		Label: "Enter split_tunnel(true/false):\n(Optional) Indicates whether split-tunnel is enabled on VPN endpoint. Default value is false",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "split_tunnel")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A mapping of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["transport_protocol"] = types.TfPrompt{
		Label: "Enter transport_protocol:\n(Optional) The transport protocol to be used by the VPN session. Default value is udp",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transport_protocol")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Green("\nEnter authentication_options:\n(Required) Information about the authentication method to be used to authenticate clients." +
		"\nauthentication_options supports the following arguments:" +
		"\n1.type\n2.active_directory_id\n3.root_certificate_chain_arn\n4.saml_provider_arn\n")

	authenticationOptionsSelect := map[string]types.TfSelect{}
	var nestedPromptOrder, nestedSelectOrder []string

	authenticationOptionsPrompt := map[string]types.TfPrompt{}

	authenticationOptionsPrompt["active_directory_id"] = types.TfPrompt{
		Label: "Enter active_directory_id:\n(Optional) The ID of the Active Directory to be used for authentication if type is directory-service-authentication",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "active_directory_id")

	authenticationOptionsPrompt["root_certificate_chain_arn"] = types.TfPrompt{
		Label: "Enter root_certificate_chain_arn:\n(Optional) The ARN of the client certificate. The certificate must be signed by a certificate authority (CA) and " +
			"\nit must be provisioned in AWS Certificate Manager (ACM). Only necessary when type is set to certificate-authentication.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "root_certificate_chain_arn")

	authenticationOptionsPrompt["saml_provider_arn"] = types.TfPrompt{
		Label: "Enter saml_provider_arn:\n(Optional) The ARN of the IAM SAML identity provider if type is federated-authentication.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "saml_provider_arn")

	authenticationOptionsSelect["type"] = types.TfSelect{
		Label: "Enter type:\n(Required) The type of client authentication to be used. Specify certificate-authentication to use certificate-based authentication, " +
			"\ndirectory-service-authentication to use Active Directory authentication, or federated-authentication to use Federated Authentication via SAML 2.0.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"certificate-authentication", "directory-service-authentication", "federated-authentication"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "type")

	resourceBlock["authentication_options"] = builder.PSOrder(nestedPromptOrder, nestedSelectOrder, authenticationOptionsPrompt, authenticationOptionsSelect)

	connectionLogOptionsPrompt := map[string]types.TfPrompt{}

	color.Green("\nEnter connection_log_options:\n(Required) Information about the authentication method to be used to authenticate clients." +
		"\nauthentication_options supports the following arguments:" +
		"\n1.type\n2.active_directory_id\n3.root_certificate_chain_arn\n4.saml_provider_arn\n")

	connectionLogOptionsPrompt["enabled"] = types.TfPrompt{
		Label: "Enter enabled:(true/false):\n(Required) Indicates whether connection logging is enabled.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "enabled")

	connectionLogOptionsPrompt["enabled"] = types.TfPrompt{
		Label: "Enter enabled:(true/false):\n(Required) Indicates whether connection logging is enabled.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "enabled")

	connectionLogOptionsPrompt["cloudwatch_log_group"] = types.TfPrompt{
		Label: "Enter cloudwatch_log_group:\n(Optional) The name of the CloudWatch Logs log group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cloudwatch_log_group")

	connectionLogOptionsPrompt["cloudwatch_log_stream"] = types.TfPrompt{
		Label: "Enter cloudwatch_log_stream:\n(Optional) The name of the CloudWatch Logs log stream to which the connection data is published.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "cloudwatch_log_stream")

	resourceBlock["connection_log_options"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-3:], nil, connectionLogOptionsPrompt, nil)

	builder.ResourceBuilder("aws_ec2_client_vpn_endpoint", blockName, resourceBlock)
}

func AWSEC2ClientVPNNetworkAssociationPrompt() {
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

	prompts["client_vpn_endpoint_id"] = types.TfPrompt{
		Label: "Enter client_vpn_endpoint_id:\n(Required) The ID of the Client VPN endpoint.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "client_vpn_endpoint_id")

	prompts["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Required) The ID of the subnet to associate with the Client VPN endpoint.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["security_groups"] = types.TfPrompt{
		Label: "Enter security_groups e.g.[\"g1\",\"g2\"]:\n(Required) The ID of the subnet to associate with the Client VPN endpoint.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "security_groups")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_client_vpn_network_association", blockName, resourceBlock)
}

func AWSEC2ClientVPNRoutePrompt() {
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

	prompts["client_vpn_endpoint_id"] = types.TfPrompt{
		Label: "Enter client_vpn_endpoint_id:\n(Required) The ID of the Client VPN endpoint.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "client_vpn_endpoint_id")

	prompts["destination_cidr_block"] = types.TfPrompt{
		Label: "Enter destination_cidr_block:\n(Required) The IPv4 address range, in CIDR notation, of the route destination.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "destination_cidr_block")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A brief description of the authorization rule.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["target_vpc_subnet_id"] = types.TfPrompt{
		Label: "Enter target_vpc_subnet_id:\n(Required) The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "target_vpc_subnet_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_client_vpn_route", blockName, resourceBlock)
}

func AWSEC2FleetPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["terminate_instances"] = types.TfPrompt{
		Label: "Enter terminate_instances(true/false):\n(Optional) Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to false",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "terminate_instances")

	prompts["terminate_instances_with_expiration"] = types.TfPrompt{
		Label: "Enter terminate_instances_with_expiration(true/false):\n(Optional) Whether running instances should be terminated when the EC2 Fleet expires. Defaults to false",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "terminate_instances_with_expiration")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["replace_unhealthy_instances"] = types.TfPrompt{
		Label: "Enter replace_unhealthy_instances(true/false): (Optional) Whether EC2 Fleet should replace unhealthy instances. Defaults to false",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "replace_unhealthy_instances")

	selects := map[string]types.TfSelect{}

	selects["type"] = types.TfSelect{
		Label: "Enter type:\n(Optional) The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Defaults to \"maintain\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"maintain", "request"},
		},
	}
	selectOrder = append(selectOrder, "type")

	selects["excess_capacity_termination_policy"] = types.TfSelect{
		Label: "Enter excess_capacity_termination_policy:\n(Optional) Whether running instances should be terminated if the total " +
			"\ntarget capacity of the EC2 Fleet is decreased below the current size of the EC2. Defaults to \"termination\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"no-termination", "termination"},
		},
	}
	selectOrder = append(selectOrder, "excess_capacity_termination_policy")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	spotOptionsPrompt := map[string]types.TfPrompt{}
	spotOptionsSelect := map[string]types.TfSelect{}

	var nestedPromptOrder, nestedSelectOrder []string

	color.Green("\nEnter spot_options:\n(Optional) Nested argument containing Spot configurations." +
		"\nspot_options support the following arguments:" +
		"\n1.allocation_strategy\n2.instance_interruption_behavior\n3.instance_pools_to_use_count\n4.maintenance_strategies\n")

	spotOptionsSelect["allocation_strategy"] = types.TfSelect{
		Label: "Enter allocation_strategy:\n(Optional) How to allocate the target capacity across the Spot pools.Defaults to \"lowestPrice\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"diversified", "lowestPrice"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "allocation_strategy")

	spotOptionsSelect["instance_interruption_behavior"] = types.TfSelect{
		Label: "Enter instance_interruption_behavior:\n(Optional) Behavior when a Spot Instance is interrupted. Defaults to \"terminate\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"hibernate", "stop", "terminate"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "instance_interruption_behavior")

	spotOptionsPrompt["instance_pools_to_use_count"] = types.TfPrompt{
		Label: "Enter instance_pools_to_use_count:\n(Optional) Number of Spot pools across which to allocate your target Spot capacity. Valid only when Spot allocation_strategy is set to lowestPrice. Default: 1",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "instance_pools_to_use_count")

	spotOptionsBlock := builder.PSOrder(nestedPromptOrder, nestedSelectOrder, spotOptionsPrompt, spotOptionsSelect)

	replacementStrategySelect := map[string]types.TfSelect{}

	replacementStrategySelect["replacement_strategy"] = types.TfSelect{
		Label: "Enter replacement_strategy:\n(Optional) The replacement strategy to use. Only available for fleets of type set to maintain. Valid values: launch",
		Select: promptui.Select{
			Label: "",
			Items: []string{"launch"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "replacement_strategy")

	color.Yellow("\nCheckout about maintenance_strategies at https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_fleet#maintenance_strategies\n")

	replacementStategyBlock := builder.PSOrder(nil, nestedSelectOrder[len(nestedSelectOrder)-1:], nil, replacementStrategySelect)

	capacityRebalanceBlock := map[string]interface{}{
		"capacity_rebalance": replacementStategyBlock,
	}

	spotOptionsBlock["maintenance_strategies"] = capacityRebalanceBlock
	resourceBlock["spot_options"] = spotOptionsBlock

	color.Green("\nEnter on_demand_options:\n(Optional) Nested argument containing On-Demand configurations." +
		"\non_demand_options currently support allocation_strategy:\n")

	onDemandOptionSelect := map[string]types.TfSelect{}

	onDemandOptionSelect["allocation_strategy"] = types.TfSelect{
		Label: "Enter allocation_strategy:\n(Optional) The order of the launch template overrides to use in fulfilling On-Demand capacity. Defaults to \"lowestPrice\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"lowestPrice", "prioritized"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "allocation_strategy")

	resourceBlock["on_demand_options"] = builder.PSOrder(nil, nestedSelectOrder[len(nestedSelectOrder)-1:], nil, onDemandOptionSelect)

	targetCapacitySpecificationPrompt := map[string]types.TfPrompt{}
	targetCapacitySpecificationSelect := map[string]types.TfSelect{}

	targetCapacitySpecificationSelect["default_target_capacity_type"] = types.TfSelect{
		Label: "Enter default_target_capacity_type:\n(Required) Default target capacity type.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"on-demand", "spot"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "default_target_capacity_type")

	targetCapacitySpecificationPrompt["total_target_capacity"] = types.TfPrompt{
		Label: "Enter total_target_capacity:\n(Required) The number of units to request, filled using default_target_capacity_type.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "total_target_capacity")

	targetCapacitySpecificationPrompt["on_demand_target_capacity"] = types.TfPrompt{
		Label: "Enter on_demand_target_capacity:\n(Optional) The number of On-Demand units to request.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "on_demand_target_capacity")

	targetCapacitySpecificationPrompt["spot_target_capacity"] = types.TfPrompt{
		Label: "Enter spot_target_capacity:\n(Optional) The number of Spot units to request.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "spot_target_capacity")

	resourceBlock["target_capacity_specification"] = builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-3:], nestedSelectOrder[len(nestedSelectOrder)-1:], targetCapacitySpecificationPrompt, targetCapacitySpecificationSelect)

	launchTemplateSpecificationPrompt := map[string]types.TfPrompt{}

	launchTemplateSpecificationPrompt["version"] = types.TfPrompt{
		Label: "Enter version:\n(Required) Version number of the launch template.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "version")

	launchTemplateSpecificationPrompt["launch_template_id"] = types.TfPrompt{
		Label: "Enter launch_template_id:\n(Optional) ID of the launch template.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "launch_template_id")

	launchTemplateSpecificationPrompt["launch_template_name"] = types.TfPrompt{
		Label: "Enter launch_template_name:\n(Optional) Name of the launch template.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "launch_template_name")

	launchTemplateSpecificationBlock := builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-3:], nil, launchTemplateSpecificationPrompt, nil)

	overridePrompt := map[string]types.TfPrompt{}

	overridePrompt["availability_zone"] = types.TfPrompt{
		Label: "Enter availability_zone:\n(Optional) Availability Zone in which to launch the instances.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "availability_zone")

	overridePrompt["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type:\n(Optional) Instance type.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "instance_type")

	overridePrompt["max_price"] = types.TfPrompt{
		Label: "Enter max_price:\n(Optional) Maximum price per unit hour that you are willing to pay for a Spot Instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "max_price")

	overridePrompt["priority"] = types.TfPrompt{
		Label: "Enter priority:\n(Optional) Priority for the launch template override. If on_demand_options allocation_strategy " +
			"\nis set to prioritized, EC2 Fleet uses priority to determine which launch template override to use " +
			"\nfirst in fulfilling On-Demand capacity. The highest priority is launched first. The lower the number, " +
			"\nthe higher the priority. If no number is set, the launch template override has the lowest priority. " +
			"\nValid values are whole numbers starting at 0.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "priority")

	overridePrompt["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Optional) ID of the subnet in which to launch the instances.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "subnet_id")

	overridePrompt["weighted_capacity"] = types.TfPrompt{
		Label: "Enter weighted_capacity:\n(Optional) Number of units provided by the specified instance type.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "weighted_capacity")

	overrideBlock := builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-6:], nil, overridePrompt, nil)

	launchTemplateConfig := map[string]interface{}{
		"launch_template_specification": launchTemplateSpecificationBlock,
		"override":                      overrideBlock,
	}

	resourceBlock["launch_template_config"] = launchTemplateConfig

	builder.ResourceBuilder("aws_ec2_fleet", blockName, resourceBlock)
}

func AWSEC2LocalGatewayRoutePrompt() {
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

	prompts["destination_cidr_block"] = types.TfPrompt{
		Label: "Enter destination_cidr_block:\n(Required) IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "destination_cidr_block")

	prompts["local_gateway_route_table_id"] = types.TfPrompt{
		Label: "Enter local_gateway_route_table_id:\n(Required) Identifier of EC2 Local Gateway Route Table.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "local_gateway_route_table_id")

	prompts["local_gateway_virtual_interface_group_id"] = types.TfPrompt{
		Label: "Enter local_gateway_virtual_interface_group_id:\n(Required) Identifier of EC2 Local Gateway Virtual Interface Group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "local_gateway_virtual_interface_group_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_local_gateway_route", blockName, resourceBlock)
}

func AWSEC2LocalGatewayRouteTableVPCAssociationPrompt() {
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

	prompts["local_gateway_route_table_id"] = types.TfPrompt{
		Label: "Enter local_gateway_route_table_id:\n(Required) Identifier of EC2 Local Gateway Route Table",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "local_gateway_route_table_id")

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) Identifier of EC2 VPC.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) Key-value map of resource tags.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_local_gateway_route_table_vpc_association", blockName, resourceBlock)

}

func AWSEC2TagPrompt() {
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

	color.Yellow("This tagging resource should not be combined with the Terraform resource " +
		"\nfor managing the parent resource. For example, using aws_vpc and aws_ec2_tag to manage " +
		"\ntags of the same VPC will cause a perpetual difference where the aws_vpc resource will try " +
		"\nto remove the tag being added by the aws_ec2_tag resource.")

	prompts["resource_id"] = types.TfPrompt{
		Label: "Enter resource_id:\n(Required) The ID of the EC2 resource to manage the tag for.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "resource_id")

	prompts["key"] = types.TfPrompt{
		Label: "Enter key:\n(Required) The tag name.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "key")

	prompts["value"] = types.TfPrompt{
		Label: "Enter value:\n(Required) The value of the tag",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "value")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_tag", blockName, resourceBlock)
}

func AWSEC2TrafficMirrorFilterPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	color.Yellow("Checkout https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html" +
		"\nfor traffic mirroring")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional, Forces new resource) A description of the filter.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags:\n(Optional) Key-value map of resource tags.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}

	selects["network_services"] = types.TfSelect{
		Label: "Enter network_services:\n(Optional) List of amazon network services that should be mirrored. Valid values: amazon-dns",
		Select: promptui.Select{
			Label: "",
			Items: []string{"amazon-dns"},
		},
	}
	selectOrder = append(selectOrder, "network_services")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_ec2_traffic_mirror_filter", blockName, resourceBlock)

}

func AWSEC2TrafficMirrorFilterRulePrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A description of the traffic mirror filter rule.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["traffic_mirror_filter_id"] = types.TfPrompt{
		Label: "Enter traffic_mirror_filter_id:\n(Required) ID of the traffic mirror filter to which this rule should be added",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "traffic_mirror_filter_id")

	prompts["destination_cidr_block"] = types.TfPrompt{
		Label: "Enter destination_cidr_block:\n(Required) The destination CIDR block to assign to the Traffic Mirror rule.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "destination_cidr_block")

	prompts["protocol"] = types.TfPrompt{
		Label: "Enter protocol:\nOptional) The protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule." +
			"\nCheckout https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "protocol")

	prompts["rule_number"] = types.TfPrompt{
		Label: "Enter rule_number:\n(Required) The number of the Traffic Mirror rule. This number must be unique for each Traffic " +
			"\nMirror rule in a given direction. The rules are processed in ascending order by rule number.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "rule_number")

	prompts["source_cidr_block"] = types.TfPrompt{
		Label: "Enter source_cidr_block:\n(Required) The source CIDR block to assign to the Traffic Mirror rule.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "source_cidr_block")

	selects := map[string]types.TfSelect{}

	selects["rule_action"] = types.TfSelect{
		Label: "Enter rule_action:\n(Required) The action to take (accept | reject) on the filtered traffic.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"accept", "reject"},
		},
	}
	selectOrder = append(selectOrder, "rule_action")

	selects["traffic_direction"] = types.TfSelect{
		Label: "Enter traffic_direction:\n(Required) The direction of traffic to be captured.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"ingress", "egress"},
		},
	}
	selectOrder = append(selectOrder, "traffic_direction")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	color.Green("\nEnter destination_port_range:\n(Optional) The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17)" +
		"\ndestination_port_range supports the following arguments:" +
		"\n1.from_port\n2.to_port\n")

	destinationPortRangePrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder []string

	destinationPortRangePrompt["from_port"] = types.TfPrompt{
		Label: "Enter from_port:\n(Optional) Starting port of the range",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "from_port")

	destinationPortRangePrompt["to_port"] = types.TfPrompt{
		Label: "Enter to_port:\n(Optional) Ending port of the range",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "to_port")

	resourceBlock["destination_port_range"] = builder.PSOrder(nestedPromptOrder, nil, destinationPortRangePrompt, nil)

	color.Green("\nEnter source_port_range:\n(Optional) The destination port range. Supported only when the protocol is set to TCP(6) or UDP(17)" +
		"\nsource_port_range supports the following arguments:" +
		"\n1.from_port\n2.to_port\n")

	sourcePortRangePrompt := map[string]types.TfPrompt{}

	sourcePortRangePrompt["from_port"] = types.TfPrompt{
		Label: "Enter from_port:\n(Optional) Starting port of the range",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}

	sourcePortRangePrompt["to_port"] = types.TfPrompt{
		Label: "Enter to_port:\n(Optional) Ending port of the range",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}

	resourceBlock["source_port_range"] = builder.PSOrder(nestedPromptOrder, nil, sourcePortRangePrompt, nil)

	builder.ResourceBuilder("aws_ec2_traffic_mirror_filter_rule", blockName, resourceBlock)
}

func AWSEC2TrafficMirrorSessionPrompt() {
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

	color.Yellow("\nCheckout https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html" +
		"\nto know more about limits and consideration for traffic mirroring\n")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) A description of the traffic mirror session.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["network_interface_id"] = types.TfPrompt{
		Label: "Enter network_interface_id:\n(Required, Forces new) ID of the source network interface. Not all network " +
			"\ninterfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_interface_id")

	prompts["traffic_mirror_filter_id"] = types.TfPrompt{
		Label: "Enter traffic_mirror_filter_id:\n(Required) ID of the traffic mirror filter to be used",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "traffic_mirror_filter_id")

	prompts["traffic_mirror_target_id"] = types.TfPrompt{
		Label: "Enter traffic_mirror_target_id:\n(Required) ID of the traffic mirror target to be used",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "traffic_mirror_target_id")

	prompts["packet_length"] = types.TfPrompt{
		Label: "Enter packet_length:\n(Required) ID of the traffic mirror target to be used",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "packet_length")

	prompts["session_number"] = types.TfPrompt{
		Label: "Enter session_number:\n(Required) - The session number determines the order in which sessions are evaluated when an " +
			"\ninterface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "session_number")

	prompts["virtual_network_id"] = types.TfPrompt{
		Label: "Enter virtual_network_id:\n(Optional) - The VXLAN ID for the Traffic Mirror session. For more information about " +
			"\nthe VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "virtual_network_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) Key-value map of resource tags.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_traffic_mirror_session", blockName, resourceBlock)

}

func AWSEC2TrafficMirrorTargetPrompt() {
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

	color.Yellow("\nEither network_interface_id or network_load_balancer_arn should be " +
		"\nspecified and both should not be specified together\n")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional, Forces new) A description of the traffic mirror session.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["network_interface_id"] = types.TfPrompt{
		Label: "Enter network_interface_id:\n(Optional, Forces new) The network interface ID that is associated with the target.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_interface_id")

	prompts["network_load_balancer_arn"] = types.TfPrompt{
		Label: "Enter network_load_balancer_arn:\n(Optional, Forces new) The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_load_balancer_arn")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) Key-value map of resource tags.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_traffic_mirror_target", blockName, resourceBlock)
}

func AWSEC2TransitGatewayPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["amazon_side_asn"] = types.TfPrompt{
		Label: "Enter amazon_side_asn:\n(Optional) Private Autonomous System Number (ASN) for the Amazon side of a BGP session. " +
			"\nThe range is 64512 to 65534 for 16-bit ASNs and 4200000000 to 4294967294 for 32-bit ASNs. Default value: 64512",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "amazon_side_asn")

	prompts["description"] = types.TfPrompt{
		Label: "Enter description:\n(Optional) Description of the EC2 Transit Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "description")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) Key-value tags for the EC2 Transit Gateway.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}

	selects["auto_accept_shared_attachments"] = types.TfSelect{
		Label: "Enter auto_accept_shared_attachments:\n(Optional) Whether resource attachment requests are automatically accepted. Defaults to \"disable\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"disable", "enable"},
		},
	}
	selectOrder = append(selectOrder, "auto_accept_shared_attachments")

	selects["default_route_table_association"] = types.TfSelect{
		Label: "Enter default_route_table_association:\n(Optional) Whether resource attachments automatically propagate routes to the default propagation route table. Defaults to \"enable\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"disable", "enable"},
		},
	}
	selectOrder = append(selectOrder, "default_route_table_association")

	selects["dns_support"] = types.TfSelect{
		Label: "Enter dns_support:\n(Optional) Whether DNS support is enabled. Defaults to \"enable\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"disable", "enable"},
		},
	}
	selectOrder = append(selectOrder, "dns_support")

	selects["vpn_ecmp_support"] = types.TfSelect{
		Label: "Enter vpn_ecmp_support:\n(Optional) Whether VPN Equal Cost Multipath Protocol support is enabled. Defaults to \"enable\"",
		Select: promptui.Select{
			Label: "",
			Items: []string{"disable", "enable"},
		},
	}
	selectOrder = append(selectOrder, "vpn_ecmp_support")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_ec2_transit_gateway", blockName, resourceBlock)
}

func AWSEC2TransitGatewayPeeringAttachmentPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["peer_account_id"] = types.TfPrompt{
		Label: "Enter peer_account_id:\n(Optional) Account ID of EC2 Transit Gateway to peer with. " +
			"\nDefaults to the account ID the AWS provider is currently connected to.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "peer_account_id")

	prompts["peer_region"] = types.TfPrompt{
		Label: "Enter peer_region:\n(Required) Region of EC2 Transit Gateway to peer with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "peer_region")

	prompts["peer_transit_gateway_id"] = types.TfPrompt{
		Label: "Enter peer_transit_gateway_id:\n(Required) Identifier of EC2 Transit Gateway to peer with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "peer_transit_gateway_id")

	prompts["transit_gateway_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_id:\n(Required) Identifier of EC2 Transit Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) Key-value tags for the EC2 Transit Gateway Peering Attachment.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_peering_attachment", blockName, resourceBlock)
}

func AWSEC2TransitGatewayPeeringAttachmentAccepterPrompt() {
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

	prompts["transit_gateway_attachment_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_attachment_id:\n(Required) The ID of the EC2 Transit Gateway Peering Attachment to manage.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_attachment_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) Key-value tags for the EC2 Transit Gateway Peering Attachment.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_peering_attachment_accepter", blockName, resourceBlock)
}

func AWSEC2TransitGatewayRoutePrompt() {
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

	prompts["destination_cidr_block"] = types.TfPrompt{
		Label: "Enter destination_cidr_block:\n(Required) IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "destination_cidr_block")

	prompts["transit_gateway_attachment_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_attachment_id:\n(Optional) Identifier of EC2 Transit Gateway Attachment (required if blackhole is set to false).",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_attachment_id")

	prompts["blackhole"] = types.TfPrompt{
		Label: "Enter blackhole(true/false):\n(Optional) Indicates whether to drop traffic that matches this route (default to false).",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "blackhole")

	prompts["transit_gateway_route_table_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_route_table_id:\n(Required) Identifier of EC2 Transit Gateway Route Table.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_route_table_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_route", blockName, resourceBlock)
}

func AWSEC2TransitGatewayRouteTablePrompt() {
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

	prompts["transit_gateway_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_id:\n(Required) Identifier of EC2 Transit Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_id")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) Key-value tags for the EC2 Transit Gateway Route Table.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_route_table", blockName, resourceBlock)

}

func AWSEC2TransitGatewayRouteTableAssociationPrompt() {
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

	prompts["transit_gateway_attachment_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_attachment_id:\n(Required) Identifier of EC2 Transit Gateway Attachment.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_attachment_id")

	prompts["transit_gateway_route_table_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_route_table_id:\n(Required) Identifier of EC2 Transit Gateway Route Table.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_route_table_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_route_table_association", blockName, resourceBlock)

}

func AWSEC2TransitGatewayRouteTablePropagationPrompt() {
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

	prompts["transit_gateway_attachment_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_attachment_id:\n(Required) Identifier of EC2 Transit Gateway Attachment.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_attachment_id")

	prompts["transit_gateway_route_table_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_route_table_id:\n(Required) Identifier of EC2 Transit Gateway Route Table.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_route_table_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_route_table_propagation", blockName, resourceBlock)
}

func AWSEC2TransitGatewayVPCAttachmentPrompt() {
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

	prompts["subnet_ids"] = types.TfPrompt{
		Label: "Enter subnet_ids:\n(Required) Identifiers of EC2 Subnets.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "subnet_ids")

	prompts["transit_gateway_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_id:\n(Required) Identifier of EC2 Transit Gateway.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_id")

	prompts["vpc_id"] = types.TfPrompt{
		Label: "Enter vpc_id:\n(Required) Identifier of EC2 VPC.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_id")

	prompts["appliance_mode_support"] = types.TfPrompt{
		Label: "Enter appliance_mode_support(true/false):\n(Optional) Whether Appliance Mode support is enabled. " +
			"\nIf enabled, a traffic flow between a source and destination uses the same " +
			"\nAvailability Zone for the VPC attachment for the lifetime of that flow.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "appliance_mode_support")

	prompts["dns_support"] = types.TfPrompt{
		Label: "Enter dns_support(true/false):\n(Optional) Whether DNS support is enabled. Defaults to \"enable\"",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "dns_support")

	prompts["ipv6_support"] = types.TfPrompt{
		Label: "Enter ipv6_support(true/false):\n(Optional) Whether IPv6 support is enabled. Defaults to \"disable\"",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ipv6_support")

	prompts["transit_gateway_default_route_table_association"] = types.TfPrompt{
		Label: "Enter transit_gateway_default_route_table_association(true/false):\n(Optional) Boolean whether the VPC Attachment should be associated with the " +
			"\nEC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with " +
			"\nResource Access Manager shared EC2 Transit Gateways. Defaults to \"true\"",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_default_route_table_association")

	prompts["transit_gateway_default_route_table_propagation"] = types.TfPrompt{
		Label: "Enter transit_gateway_default_route_table_propagation(true/false):\n(Optional) Boolean whether the VPC Attachment should propagate routes with the " +
			"\nEC2 Transit Gateway propagation default route table. This cannot be configured " +
			"\nor perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Defaults to \"true\"",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_default_route_table_propagation")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) Key-value tags for the EC2 Transit Gateway VPC Attachment.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_vpc_attachment", blockName, resourceBlock)
}

func AWSEC2TransitGatewayVPCAttachmentAccepterPrompt() {
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

	prompts["transit_gateway_attachment_id"] = types.TfPrompt{
		Label: "Enter transit_gateway_attachment_id:\n(Required) The ID of the EC2 Transit Gateway Attachment to manage.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_attachment_id")

	prompts["transit_gateway_default_route_table_association"] = types.TfPrompt{
		Label: "Enter transit_gateway_default_route_table_association:\n(Optional) Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_default_route_table_association")

	prompts["transit_gateway_default_route_table_propagation"] = types.TfPrompt{
		Label: "Enter transit_gateway_default_route_table_propagation:\n(Optional) Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "transit_gateway_default_route_table_propagation")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) Key-value tags for the EC2 Transit Gateway VPC Attachment.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_ec2_transit_gateway_vpc_attachment_accepter", blockName, resourceBlock)
}

func AWSEIPPrompt() {
	color.Yellow("\n EIP may require IGW to exist prior to association. Use depends_on to set an explicit dependency on the IGW." +
		"\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip")

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

	color.Yellow("\nYou can specify either the instance ID or the network_interface ID, but not both. " +
		"\nIncluding both will not return an error from the AWS API, but will have undefined behavior." +
		"\nCheckout https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociateAddress.html\n")

	prompts["vpc"] = types.TfPrompt{
		Label: "Enter vpc(true/false):\n(Optional) Boolean if the EIP is in a VPC or not.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "vpc")

	prompts["instance"] = types.TfPrompt{
		Label: "Enter instance:\n(Optional) EC2 instance ID.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance")

	prompts["network_interface"] = types.TfPrompt{
		Label: "Enter network_interface:\n(Optional) Network interface ID to associate with.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_interface")

	prompts["associate_with_private_ip"] = types.TfPrompt{
		Label: "Enter associate_with_private_ip:\n(Optional) A user specified primary or secondary private IP address to associate with the " +
			"\nElastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "associate_with_private_ip")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["public_ipv4_pool"] = types.TfPrompt{
		Label: "Enter public_ipv4_pool:\n(Optional) EC2 IPv4 address pool identifier or amazon. This option is only available for VPC EIPs.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "public_ipv4_pool")

	prompts["customer_owned_ipv4_pool"] = types.TfPrompt{
		Label: "Enter customer_owned_ipv4_pool:\nThe ID of a customer-owned address pool." +
			"\nCheckout https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "customer_owned_ipv4_pool")

	prompts["network_border_group"] = types.TfPrompt{
		Label: "Enter network_border_group:\nThe location from which the IP address is advertised. Use this parameter to limit the address to this location.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_border_group")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_eip", blockName, resourceBlock)
}

func AWSEIPAssociationPrompt() {

	color.Yellow("\nDo not use this resource to associate an EIP to aws_lb or aws_nat_gateway " +
		"\nresources. Instead use the allocation_id available in those resources to allow AWS " +
		"\nto manage the association, otherwise you will see \"AuthFailure\" errors.\n")

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

	prompts["allocation_id"] = types.TfPrompt{
		Label: "Enter allocation_id:\n(Optional) The allocation ID. This is required for EC2-VPC.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "allocation_id")

	prompts["allow_reassociation"] = types.TfPrompt{
		Label: "Enter allow_reassociation(true/false):\n(Optional, Boolean) Whether to allow an Elastic IP to be re-associated. Defaults to true in VPC.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "allow_reassociation")

	prompts["instance_id"] = types.TfPrompt{
		Label: "Enter instance_id:\n(Optional) The ID of the instance. This is required for EC2-Classic. For EC2-VPC, you can specify " +
			"\neither the instance ID or the network interface ID, but not both. The operation fails if you specify an " +
			"\ninstance ID unless exactly one network interface is attached.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_id")

	prompts["network_interface_id"] = types.TfPrompt{
		Label: "Enter network_interface_id:\n(Optional) The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "network_interface_id")

	prompts["private_ip_address"] = types.TfPrompt{
		Label: "Enter private_ip_address:\n(Optional) The primary or secondary private IP address to associate with " +
			"\nthe Elastic IP address. If no private IP address is specified, the Elastic IP " +
			"\naddress is associated with the primary private IP address.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "private_ip_address")

	prompts["public_ip"] = types.TfPrompt{
		Label: "Enter public_ip:\n(Optional) The Elastic IP address. This is required for EC2-Classic.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "public_ip")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_eip_association", blockName, resourceBlock)
}

func AWSInstancePrompt() {
	prompts := map[string]types.TfPrompt{}

	color.Green("\nEnter block name(Required) e.g. web\n\n")
	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var promptOrder []string
	prompts["ami"] = types.TfPrompt{
		Label: "Enter ami(Required):\nThe AMI to use for the instance",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ami")

	prompts["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type(Required) e.g. t2.micro\nThe type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_type")

	prompts["cpu_core_count"] = types.TfPrompt{
		Label: "Enter cpu_core_count(number):\n(Optional)Sets the number of CPU cores for an instance. " +
			"This option is only supported on creation of instance type that support CPU Options - " +
			"specifying this option for unsupported instance types will return an error from the EC2 API. Checkout " +
			"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values for more info.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "cpu_core_count")

	prompts["cpu_threads_per_core"] = types.TfPrompt{
		Label: "Enter cpu_threads_per_core(number):\n(Optional) - has no effect unless cpu_core_count is also set) " +
			"If set to to 1, hyperthreading is disabled on the launched instance. " +
			"Defaults to 2 if not set. See Optimizing CPU Options for more information.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "cpu_threads_per_core")

	prompts["ebs_optimized"] = types.TfPrompt{
		Label: "Enter EBS-optimized(true/false):\n(Optional) If true, the launched EC2 instance will be EBS-optimized. " +
			"\nNote that if this is not set on an instance type that is optimized by default then this will show " +
			"\nas disabled but if the instance type is optimized by default then there is no " +
			"\nneed to set this and there is no effect to disabling it. " +
			"\nCheckout https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html of AWS User Guide for more information.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "ebs_optimized")

	prompts["monitoring"] = types.TfPrompt{
		Label: "Select true/false for monitoring:\n(Optional) " +
			"If true, the launched EC2 instance will have detailed monitoring enabled",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "monitoring")

	prompts["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Optional) The VPC Subnet ID to launch in.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "subnet_id")

	prompts["private_ip"] = types.TfPrompt{
		Label: "Enter private_ip:\n(Optional) Private IP address to associate with the instance in a VPC.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "private_ip")

	prompts["iam_instance_profile"] = types.TfPrompt{
		Label: "Enter iam_instance_profile:\n(Optional) The IAM Instance Profile to launch the " +
			"instance with. Specified as the name of the Instance Profile. " +
			"Ensure your credentials have the correct permission to assign " +
			"the instance profile according to the EC2 documentation, notably iam:PassRole",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.StringValidator,
		},
	}
	promptOrder = append(promptOrder, "iam_instance_profile")

	prompts["security_groups"] = types.TfPrompt{
		Label: "A list of security group names (EC2-Classic) or IDs (default VPC) to associate with\ne.g.[\"a\",\"b\",\"c\"]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "security_groups")

	prompts["vpc_security_group_ids"] = types.TfPrompt{
		Label: "A list of security group IDs to associate with(Only VPC) e.g. [\"a\",\"b\",\"c\"]",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "vpc_security_group_ids")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags: e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	prompts["associate_public_ip_address"] = types.TfPrompt{
		Label: "Enter associate_public_ip_address(true/false):\n(Optional)Associate a public ip address with an instance in a VPC.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "associate_public_ip_address")

	prompts["hibernation"] = types.TfPrompt{
		Label: "Enter hibernation(true/false).\n(Optional)If true, the launched EC2 instance will support hibernation.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "hibernation")

	selects := map[string]types.TfSelect{}
	var selectOrder []string

	selects["placement_group"] = types.TfSelect{
		Label: "Enter placement_group:\nThe Placement Group to start the instance in",
		Select: promptui.Select{
			Label: "",
			Items: []string{"cluster", "partition", "spread"},
		},
	}
	selectOrder = append(selectOrder, "placement_group")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_instance", blockName, resourceBlock)
}

func AWSKeyPairPrompt() {
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

	prompts["key_name"] = types.TfPrompt{
		Label: "Enter key_name:\n(Optional) The name for the key pair.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "key_name")

	prompts["key_name_prefix"] = types.TfPrompt{
		Label: "Enter key_name_prefix:\n(Optional) Creates a unique name beginning with the specified prefix. Conflicts with key_name",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "key_name_prefix")

	prompts["public_key"] = types.TfPrompt{
		Label: "Enter public_key:\n(Required) The public key material.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "public_key")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) Key-value map of resource tags",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.RCValidator,
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_key_pair", blockName, resourceBlock)

}

func AWSPlacementGroupPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["name"] = types.TfPrompt{
		Label: "Enter name:\n(Required) The name of the placement group.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "name")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) Key-value map of resource tags.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}

	selects["strategy"] = types.TfSelect{
		Label: "Enter strategy:\n(Required) The placement strategy.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"cluster", "partition", "spread"},
		},
	}
	selectOrder = append(selectOrder, "strategy")

	resourceBlock := builder.PSOrder(promptOrder, selectOrder, prompts, selects)

	builder.ResourceBuilder("aws_placement_group", blockName, resourceBlock)
}

func AWSSnapshotCreateVolumePermissionPrompt() {
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

	prompts["snapshot_id"] = types.TfPrompt{
		Label: "Enter snapshot_id:\n(required) A snapshot ID",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "snapshot_id")

	prompts["account_id"] = types.TfPrompt{
		Label: "Enter account_id:\n(required) An AWS Account ID to add create volume permissions",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "account_id")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_snapshot_create_volume_permission", blockName, resourceBlock)

}

func AWSSpotDatafeedSubscriptionPrompt() {
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

	prompts["bucket"] = types.TfPrompt{
		Label: "Enter bucket:\n(Required) The Amazon S3 bucket in which to store the Spot instance data feed.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "bucket")

	prompts["prefix"] = types.TfPrompt{
		Label: "Enter prefix:\n(Optional) Path of folder inside bucket to place spot pricing data.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "prefix")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_spot_datafeed_subscription", blockName, resourceBlock)

}

func AWSSpotFleetRequestPrompt() {
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

	prompts["iam_fleet_role"] = types.TfPrompt{
		Label: "Enter iam_fleet_role:\n(Required) Grants the Spot fleet permission to terminate Spot instances " +
			"\non your behalf when you cancel its Spot fleet request using CancelSpotFleetRequests " +
			"\nor when the Spot fleet request expires, if you set terminateInstancesWithExpiration.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "iam_fleet_role")

	prompts["replace_unhealthy_instances"] = types.TfPrompt{
		Label: "Enter replace_unhealthy_instances(true/false):\n(Optional) Indicates whether Spot fleet should replace unhealthy instances. Defaults to \"false\"",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "replace_unhealthy_instances")

	prompts["spot_price"] = types.TfPrompt{
		Label: "Enter spot_price:\n(Optional; Default: On-demand price) The maximum bid price per unit hour.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "spot_price")

	prompts["target_capacity"] = types.TfPrompt{
		Label: "Enter target_capacity:\nThe number of units to request. You can choose to set the target capacity " +
			"\nin terms of instances or a performance characteristic that is important to your " +
			"\napplication workload, such as vCPUs, memory, or I/O.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "target_capacity")

	prompts["instance_pools_to_use_count"] = types.TfPrompt{
		Label: "Enter instance_pools_to_use_count:\n(Optional; Default: 1) The number of Spot pools across which to allocate your target " +
			"\nSpot capacity. Valid only when allocation_strategy is set to lowestPrice. " +
			"\nSpot Fleet selects the cheapest Spot pools and evenly allocates your target " +
			"\nSpot capacity across the number of Spot pools that you specify.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	promptOrder = append(promptOrder, "instance_pools_to_use_count")

	prompts["excess_capacity_termination_policy"] = types.TfPrompt{
		Label: "Enter excess_capacity_termination_policy:\nIndicates whether running Spot instances should be terminated if the target capacity of the Spot fleet request is decreased below the current size of the Spot fleet.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "excess_capacity_termination_policy")

	prompts["instance_interruption_behaviour"] = types.TfPrompt{
		Label: "Enter instance_interruption_behaviour:\n(Optional) Indicates whether a Spot instance stops or terminates when it is interrupted. Default is terminate",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_interruption_behaviour")

	prompts["fleet_type"] = types.TfPrompt{
		Label: "Enter fleet_type:\n(Optional) The type of fleet request. Indicates whether the Spot Fleet only requests the " +
			"\ntarget capacity or also attempts to maintain it. Default is maintain",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "fleet_type")

	prompts["valid_from"] = types.TfPrompt{
		Label: "Enter valid_from:\n(Optional) The start date and time of the request, in UTC RFC3339 format" +
			"\n(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "valid_from")

	prompts["valid_until"] = types.TfPrompt{
		Label: "Enter valid_until:\n(Optional) The end date and time of the request, in UTC RFC3339 format" +
			"\n(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "valid_until")

	prompts["load_balancers"] = types.TfPrompt{
		Label: "Enter load_balancers e.g. [\"b1\",\"b2\"]:\n(Optional) A list of elastic load balancer names to add to the Spot fleet.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "load_balancers")

	prompts["target_group_arns"] = types.TfPrompt{
		Label: "Enter target_group_arns e.g. [\"a1\",\"a2\"]:\n(Optional) A list of aws_alb_target_group ARNs, for use with Application Load Balancing.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "target_group_arns")

	prompts["wait_for_fulfillment"] = types.TfPrompt{
		Label: "Enter wait_for_fulfillment:\n(Optional; Default: false) If set, Terraform will wait for the Spot Request to be fulfilled, and will throw an error if the timeout of 10m is reached",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "wait_for_fulfillment")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g. k1=v1,k2=v2:\n(Optional) A map of tags to assign to the resource.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	color.Yellow("\nConfigure nested settings like launch_specification/root_block_device etc [y/n]?\n\n", "text")

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

	color.Green("\nEnter launch_template_config:\n(Optional) Launch template configuration block" +
		"\nThe launch_template_config block supports the following:" +
		"\n1.launch_template_specification\n2.overrides")
	color.Yellow("\nConflicts with launch_specification. At least one of launch_specification or launch_template_config is required.")

	color.Green("\nEnter launch_template_specification:\n(Required) Launch template specification." +
		"\nThe launch_template_specification supports the following arguments:" +
		"\n1.id\n2.name\n3.version\n")

	launchTemplateSpecificationPrompt := map[string]types.TfPrompt{}
	var nestedPromptOrder, nestedSelectOrder []string

	launchTemplateSpecificationPrompt["id"] = types.TfPrompt{
		Label: "Enter id:\nThe ID of the launch template. Conflicts with name.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "id")

	launchTemplateSpecificationPrompt["name"] = types.TfPrompt{
		Label: "Enter name:\nThe name of the launch template. Conflicts with id",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "name")

	launchTemplateSpecificationPrompt["version"] = types.TfPrompt{
		Label: "Enter version:\n(Optional) Template version. Unlike the autoscaling equivalent, does not support $Latest or $Default, so use the launch_template resource's attribute, e.g. \"${aws_launch_template.foo.latest_version}\". It will use the default version if omitted.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "version")

	launchTemplateSpecificationBlock := builder.PSOrder(nestedPromptOrder, nil, launchTemplateSpecificationPrompt, nil)

	overridesPrompt := map[string]types.TfPrompt{}

	color.Green("\nEnter overrides:\n(Optional) One or more override configurations" +
		"\nThe overrides block supports the following:\n" +
		"\n1.availability_zone\n2.instance_type\n3.priority\n4.spot_price\n5.subnet_id\n6.weighted_capacity\n")

	overridesPrompt["availability_zone"] = types.TfPrompt{
		Label: "Enter availability_zone:\n(Optional) The availability zone in which to place the request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "availability_zone")

	overridesPrompt["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type:\n(Optional) The type of instance to request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "instance_type")

	overridesPrompt["spot_price"] = types.TfPrompt{
		Label: "Enter spot_price:\n(Optional) The maximum spot bid for this override request.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "spot_price")

	overridesPrompt["priority"] = types.TfPrompt{
		Label: "Enter priority:\n(Optional) The priority for the launch template override. The lower the number, " +
			"\nthe higher the priority. If no number is set, the launch template override has the lowest priority.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "priority")

	overridesPrompt["subnet_id"] = types.TfPrompt{
		Label: "Enter subnet_id:\n(Optional) The subnet in which to launch the requested instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "subnet_id")

	overridesPrompt["weighted_capacity"] = types.TfPrompt{
		Label: "Enter weighted_capacity:\n(Optional) The capacity added to the fleet by a fulfilled request.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.IntValidator,
		},
	}
	nestedPromptOrder = append(nestedPromptOrder, "weighted_capacity")

	overrideBlock := builder.PSOrder(nestedPromptOrder[len(nestedPromptOrder)-6:], nil, overridesPrompt, nil)

	launchTemplateConfigBlock := map[string]interface{}{
		"launch_template_specification": launchTemplateSpecificationBlock,
		"overrides":                     overrideBlock,
	}

	resourceBlock["launch_template_config"] = launchTemplateConfigBlock

	replacementStrategySelect := map[string]types.TfSelect{}

	replacementStrategySelect["replacement_strategy"] = types.TfSelect{
		Label: "Enter replacement_strategy:\n(Optional) The replacement strategy to use. Only available for fleets of type set to maintain. Valid values: launch",
		Select: promptui.Select{
			Label: "",
			Items: []string{"launch"},
		},
	}
	nestedSelectOrder = append(nestedSelectOrder, "replacement_strategy")

	replacementStategyBlock := builder.PSOrder(nil, nestedSelectOrder, nil, replacementStrategySelect)

	spotOptionsBlock := map[string]interface{}{
		"capacity_rebalance": replacementStategyBlock,
	}

	resourceBlock["spot_maintenance_strategies"] = spotOptionsBlock

	builder.ResourceBuilder("aws_spot_fleet_request", blockName, resourceBlock)
}

func AWSSpotInstanceRequestPrompt() {
	color.Green("\nEnter block name(Required) e.g. web\n\n")

	blockPrompt := promptui.Prompt{
		Label: "",
	}

	blockName, err := blockPrompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	prompts := map[string]types.TfPrompt{}
	var promptOrder, selectOrder []string

	prompts["ami"] = types.TfPrompt{
		Label: "Enter ami(Required):\nThe AMI to use for the instance",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "ami")

	prompts["instance_type"] = types.TfPrompt{
		Label: "Enter instance_type(Required) e.g. t2.micro\nThe type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_type")

	prompts["spot_price"] = types.TfPrompt{
		Label: "Enter spot_price:\n(Optional; Default: On-demand price) The maximum price to request on the spot market.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "spot_price")

	prompts["wait_for_fulfillment"] = types.TfPrompt{
		Label: "Enter wait_for_fulfillment(true/false):\n(Optional; Default: false) If set, Terraform will wait for the Spot Request to be fulfilled, and will throw an error if the timeout of 10m is reached.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BoolValidator,
		},
	}
	promptOrder = append(promptOrder, "wait_for_fulfillment")

	prompts["launch_group"] = types.TfPrompt{
		Label: "Enter launch_group:\n(Optional) A launch group is a group of spot instances that launch together and terminate together. If left empty instances are launched and terminated individually.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "launch_group")

	prompts["block_duration_minutes"] = types.TfPrompt{
		Label: "Enter block_duration_minutes:\n(Optional) The required duration for the Spot instances, in minutes. This value must be a " +
			"\nmultiple of 60 (60, 120, 180, 240, 300, or 360). The duration period starts as soon " +
			"\nas your Spot instance receives its instance ID. At the end of the duration period, " +
			"\nAmazon EC2 marks the Spot instance for termination and provides a Spot instance " +
			"\ntermination notice, which gives the instance a two-minute warning before it terminates. " +
			"\nNote that you can't specify an Availability Zone group or a launch group if you specify a duration.",
		Prompt: promptui.Prompt{
			Label:    "",
			Validate: validators.BlockDurationValidator,
		},
	}
	promptOrder = append(promptOrder, "block_duration_minutes")

	prompts["instance_interruption_behaviour"] = types.TfPrompt{
		Label: "Enter instance_interruption_behaviour:\n(Optional) Indicates whether a Spot instance stops or terminates " +
			"\nwhen it is interrupted. Default is terminate as this is the current AWS behaviour.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "instance_interruption_behaviour")

	prompts["valid_from"] = types.TfPrompt{
		Label: "Enter valid_from:\n(Optional) The start date and time of the request, in UTC RFC3339 format" +
			"\n(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "valid_from")

	prompts["valid_until"] = types.TfPrompt{
		Label: "Enter valid_until:\n(Optional) The end date and time of the request, in UTC RFC3339 format" +
			"\n(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed " +
			"\nor enabled to fulfill the request. The default end date is 7 days from the current date.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "valid_until")

	prompts["tags"] = types.TfPrompt{
		Label: "Enter tags e.g.k1=v1,k2=v2:\n(Optional) A map of tags to assign to the Spot Instance Request. These tags are not automatically applied to the launched Instance.",
		Prompt: promptui.Prompt{
			Label: "",
		},
	}
	promptOrder = append(promptOrder, "tags")

	selects := map[string]types.TfSelect{}

	selects["spot_type"] = types.TfSelect{
		Label: "Enter spot_type:\n(Optional; Default: persistent) If set to one-time, after the instance is terminated, the spot request will be closed.",
		Select: promptui.Select{
			Label: "",
			Items: []string{"persistent", "one-time"},
		},
	}
	selectOrder = append(selectOrder, "spot_type")

	resourceBlock := builder.PSOrder(promptOrder, nil, prompts, nil)

	builder.ResourceBuilder("aws_spot_instance_request", blockName, resourceBlock)
}
