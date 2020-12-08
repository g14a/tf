package builder

import (
	"fmt"
	"strings"
	"tf/file"
)

func ProviderBuilder(provider string, providerBlock map[string]interface{}) {
	var providerInfo strings.Builder

	providerInfo.WriteString("# ------------THIS FILE IS GENERATED AUTOMATICALLY------------\n\n")

	providerInfo.WriteString("provider \"" + provider + "\" {\n")
	for k, v := range providerBlock {
		providerInfo.WriteString("\t " + k + ": " + v.(string) + "\n")
	}
	providerInfo.WriteString("}")

	fmt.Println(providerInfo.String())

	_, err := file.TerraformFile.WriteString(providerInfo.String())
	if err != nil {
		fmt.Println(err)
	}
}
