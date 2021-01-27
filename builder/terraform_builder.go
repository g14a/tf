package builder

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/g14a/tf/file"
	"github.com/g14a/tf/types"
)

// ProviderBuilder builds the provider configuration
func ProviderBuilder(provider string, providerBlock map[string]interface{}) {
	var providerInfo strings.Builder

	if provider != "" {
		providerInfo.WriteString("\nprovider \"" + provider + "\" {\n")
		providerInfo = recursiveBuilder(&providerInfo, reflect.ValueOf(providerBlock))

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
}

// ResourceBuilder builds the resource configuration
func ResourceBuilder(resource, blockName string, resourceBlock map[string]interface{}) {
	var resourceBuilder strings.Builder

	if resource != "" {
		resourceBuilder.WriteString("\nresource \"" + resource + "\" \"" + blockName + "\" {\n")
		resourceBuilder = recursiveBuilder(&resourceBuilder, reflect.ValueOf(resourceBlock))

		_, err := file.TerraformFile.WriteString(resourceBuilder.String())
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
}

// PSOrder collects values from the prompts and selects
// given themselves and the order.
func PSOrder(promptOrder, selectOrder []string,
	prompts map[string]types.TfPrompt,
	selects map[string]types.TfSelect) map[string]interface{} {
	resourceBlock := map[string]interface{}{}

	if len(promptOrder) > 0 {
		for _, v := range promptOrder {
			p := prompts[v]
			value, err := p.Run()
			if err != nil {
				fmt.Println(err)
			}
			if value != "" {
				switch {
				case govalidator.IsInt(value):
					intValue, err := strconv.Atoi(value)
					if err != nil {
						log.Fatal(err)
					}
					resourceBlock[v] = intValue
				case value == "true" || value == "false":
					boolValue, err := strconv.ParseBool(value)
					if err != nil {
						fmt.Println(err)
					}
					resourceBlock[v] = boolValue
				default:
					resourceBlock[v] = value
				}
			}
		}
	}

	if len(selectOrder) > 0 {
		for _, v := range selectOrder {
			p := selects[v]
			value, err := p.Run()
			if err != nil {
				fmt.Println(err)
			}
			if value != "" {
				switch {
				case govalidator.IsInt(value):
					intValue, err := strconv.Atoi(value)
					if err != nil {
						log.Fatal(err)
					}
					resourceBlock[v] = intValue
				case value == "true" || value == "false":
					boolValue, err := strconv.ParseBool(value)
					if err != nil {
						fmt.Println(err)
					}
					resourceBlock[v] = boolValue
				default:
					resourceBlock[v] = value
				}
			}
		}
	}

	return resourceBlock
}

func repeatingConfig(input string) string {
	if input != "" {
		rc := strings.Split(input, ",")
		var rcString strings.Builder
		for _, v := range rc {
			v := strings.Split(v, "=")
			rcString.WriteString("  " + v[0] + " = \"" + v[1] + "\"\n")
		}
		rcString.WriteString("}\n")
		return rcString.String()
	}
	return ""
}

func walk(strBuilder *strings.Builder, v reflect.Value) {

	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			s := fmt.Sprintf("%s", v.MapIndex(k))
			switch {
			case k.String() == "tags" ||
				k.String() == "variables" ||
				k.String() == "metadata" ||
				k.String() == "response_parameters" ||
				k.String() == "response_templates":
				temp := fmt.Sprintf("%s = %s", k, "{\n")
				strBuilder.WriteString(temp)
				strBuilder.WriteString(repeatingConfig(s))
			case strings.HasPrefix(s, "map"):
				strBuilder.WriteString(fmt.Sprintf("%s %s", k, "{\n"))
			case strings.Contains(s, "(int"):
				strBuilder.WriteString(fmt.Sprintf("%s = %d\n", k, v.MapIndex(k)))
			case strings.Contains(s, "(bool"):
				strBuilder.WriteString(fmt.Sprintf("%s = %t\n", k, v.MapIndex(k)))
			case strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]"):
				strBuilder.WriteString(fmt.Sprintf("%s = %s\n", k, v.MapIndex(k)))
			default:
				strBuilder.WriteString(fmt.Sprintf("%s = \"%s\"\n", k, v.MapIndex(k)))
			}
			walk(strBuilder, v.MapIndex(k))
		}
		strBuilder.WriteString("}\n")
	}
}

func recursiveBuilder(str *strings.Builder, v reflect.Value) strings.Builder {
	walk(str, v)
	return *str
}

func terraformExists() bool {
	_, err := exec.LookPath("terraform")
	return err == nil
}
