package builder

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/manifoldco/promptui"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"tf/file"
)

func ProviderBuilder(provider string, providerBlock map[string]interface{}) {
	var providerInfo strings.Builder

	providerInfo.WriteString("\nprovider \"" + provider + "\" {\n")
	providerInfo = infoBuilder(&providerInfo, providerBlock)
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

	providerInfo.WriteString("\nresource \"" + resource + "\" \"" + blockName + "\" {\n")

	providerInfo = infoBuilder(&providerInfo, resourceBlock)
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

func PSOrder(promptOrder, selectOrder []string,
	prompts map[string]promptui.Prompt,
	selects map[string]promptui.Select) map[string]interface{} {
	resourceBlock := map[string]interface{}{}

	for _, v := range promptOrder {
		p := prompts[v]
		value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		resourceBlock[v] = value
	}

	for _, v := range selectOrder {
		p := selects[v]
		_, value, err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
		resourceBlock[v] = value
	}

	return resourceBlock
}

func infoBuilder(strBuilder *strings.Builder, infoBlock map[string]interface{}) strings.Builder {
	for k, v := range infoBlock {
		switch v.(type) {
		case string:
			if v.(string) != "" {
				if govalidator.IsInt(v.(string)) {
					temp, err := strconv.Atoi(v.(string))
					if err != nil {
						log.Fatal(err)
					}
					s := fmt.Sprintf("  "+k+"= %d \n", temp)
					strBuilder.WriteString(s)
				} else if v.(string) == "true" || v.(string) == "false" {
					b, err := strconv.ParseBool(v.(string))
					if err != nil {
						fmt.Println(err)
					}
					s := fmt.Sprintf("  "+k+"= %t \n", b)
					strBuilder.WriteString(s)
				} else {
					s := fmt.Sprintf("  " + k + "= \"" + v.(string) + "\"\n")
					strBuilder.WriteString(s)
				}
			}
		case map[string]interface{}:
			strBuilder.WriteString("  " + k + " {\n")
			for nestedK, i := range v.(map[string]interface{}) {
				if i.(string) != "" {
					if govalidator.IsInt(i.(string)) {
						temp, err := strconv.Atoi(i.(string))
						if err != nil {
							log.Fatal(err)
						}
						s := fmt.Sprintf("  "+nestedK+" = %d \n", temp)
						strBuilder.WriteString(s)
					} else if i.(string) == "true" || i.(string) == "false" {
						b, err := strconv.ParseBool(i.(string))
						if err != nil {
							fmt.Println(err)
						}
						s := fmt.Sprintf("  "+nestedK+" = %t \n", b)
						strBuilder.WriteString(s)
					} else {
						s := fmt.Sprintf("  " + nestedK + " = \"" + i.(string) + "\"\n")
						strBuilder.WriteString(s)
					}
				}
			}
			strBuilder.WriteString("}\n")
		}
	}

	return *strBuilder
}

func terraformExists() bool {
	_, err := exec.LookPath("terraform")
	return err == nil
}
