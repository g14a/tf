package builder

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"log"
	"strconv"
	"strings"
	"tf/file"
)

func ProviderBuilder(provider string, providerBlock map[string]interface{}) {
	var providerInfo strings.Builder

	providerInfo.WriteString("# ------------THIS FILE IS GENERATED AUTOMATICALLY------------\n\n")

	providerInfo.WriteString("provider \"" + provider + "\" {\n")
	for k, v := range providerBlock {
		if govalidator.IsInt(v.(string)) {
			temp, err := strconv.Atoi(v.(string))
			if err != nil {
				log.Fatal(err)
			}
			s := fmt.Sprintf("  "+k+"= %d \n", temp)
			providerInfo.WriteString(s)
		} else {
			providerInfo.WriteString("  " + k + "= \"" + v.(string) + "\"\n")
		}
	}
	providerInfo.WriteString("}")

	fmt.Println(providerInfo.String())

	_, err := file.TerraformFile.WriteString(providerInfo.String())
	if err != nil {
		fmt.Println(err)
	}
}

func ResourceBuilder(resource, blockName string, resourceBlock map[string]interface{}) {
	var providerInfo strings.Builder

	providerInfo.WriteString("\n\nresource \"" + resource + "\" \"" + blockName + "\" {\n")
	for k, v := range resourceBlock {
		if govalidator.IsInt(v.(string)) {
			temp, err := strconv.Atoi(v.(string))
			if err != nil {
				log.Fatal(err)
			}
			s := fmt.Sprintf("  "+k+"= %d \n", temp)
			providerInfo.WriteString(s)
		} else {
			providerInfo.WriteString("  " + k + "= \"" + v.(string) + "\"\n")
		}
	}
	providerInfo.WriteString("}")

	fmt.Println(providerInfo.String())

	_, err := file.TerraformFile.WriteString(providerInfo.String())
	if err != nil {
		fmt.Println(err)
	}
}
