package builder

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"tf/file"
)

func ProviderBuilder(provider string, providerBlock map[string]interface{}) {
	var providerInfo strings.Builder

	providerInfo.WriteString("\n\nprovider \"" + provider + "\" {\n")
	for k, v := range providerBlock {
		if govalidator.IsInt(v.(string)) {
			temp, err := strconv.Atoi(v.(string))
			if err != nil {
				log.Fatal(err)
			}
			s := fmt.Sprintf("  "+k+"= %d \n", temp)
			providerInfo.WriteString(s)
		} else if v.(string) == "true" || v.(string) == "false"{
			b, err := strconv.ParseBool(v.(string))
			if err != nil {
				fmt.Println(err)
			}
			s := fmt.Sprintf("  "+k+"= %t \n", b)
			providerInfo.WriteString(s)
		} else {
			providerInfo.WriteString("  " + k + "= \"" + v.(string) + "\"\n")
		}
	}
	providerInfo.WriteString("}")

	_, err := file.TerraformFile.WriteString(providerInfo.String())
	if err != nil {
		fmt.Println(err)
	}

	if terraformExists() {
		cmd := exec.Command("terraform", "fmt")
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
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
		} else if v.(string) == "true" || v.(string) == "false"{
			b, err := strconv.ParseBool(v.(string))
			if err != nil {
				fmt.Println(err)
			}
			s := fmt.Sprintf("  "+k+"= %t \n", b)
			providerInfo.WriteString(s)
		} else {
			providerInfo.WriteString("  " + k + "= \"" + v.(string) + "\"\n")
		}
	}
	providerInfo.WriteString("}")

	_, err := file.TerraformFile.WriteString(providerInfo.String())
	if err != nil {
		fmt.Println(err)
	}

	if terraformExists() {
		cmd := exec.Command("terraform", "fmt")
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
}

func terraformExists() bool {
	_, err := exec.LookPath("terraform")
	return err == nil
}