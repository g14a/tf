package resourcebps

import "github.com/fatih/color"

func AWSSNSPlatformApplicationBP() {
	color.Green("\nresource \"aws_sns_platform_application\" \"apns_application\" {\n  name                = \"apns_application\"\n  platform            = \"APNS\"\n  platform_credential = \"<APNS PRIVATE KEY>\"\n  platform_principal  = \"<APNS CERTIFICATE>\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_platform_application\n\n")
}
