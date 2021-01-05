package builder

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"log"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"tf/file"
	"tf/types"
)

func ProviderBuilder(provider string, promptOrder, selectOrder []string, providerBlock map[string]interface{}) {
	var providerInfo strings.Builder

	if provider != "" {
		providerInfo.WriteString("\nprovider \"" + provider + "\" {\n")
		providerInfo = recursiveBuilder(&providerInfo, reflect.ValueOf(providerBlock))
		providerInfo.WriteString("}\n")

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

func ResourceBuilder(resource, blockName string, promptOrder, selectOrder []string, resourceBlock map[string]interface{}) {
	var providerInfo strings.Builder

	if resource != "" {
		providerInfo.WriteString("\nresource \"" + resource + "\" \"" + blockName + "\" {\n")
		providerInfo = recursiveBuilder(&providerInfo, reflect.ValueOf(resourceBlock))

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

func NestedPSOrder(promptOrder []string, selectOrder []string,
	prompts map[string]types.TfPrompt,
	selects map[string]types.TfSelect) map[string]interface{} {
	nestedBlock := map[string]interface{}{}

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
					nestedBlock[v] = intValue
				case value == "true" || value == "false":
					boolValue, err := strconv.ParseBool(value)
					if err != nil {
						fmt.Println(err)
					}
					nestedBlock[v] = boolValue
				default:
					nestedBlock[v] = value
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
					nestedBlock[v] = intValue
				case value == "true" || value == "false":
					boolValue, err := strconv.ParseBool(value)
					if err != nil {
						fmt.Println(err)
					}
					nestedBlock[v] = boolValue
				default:
					nestedBlock[v] = value
				}
			}
		}
	}

	return nestedBlock
}

func infoBuilder(strBuilder *strings.Builder, promptOrder, selectOrder []string, infoBlock map[string]interface{}) strings.Builder {
	order := append(promptOrder, selectOrder...)
	for _, o := range order {
		v := infoBlock[o]
		if o == "tags" || o == "variables" {
			strBuilder.WriteString("  " + o + " = {\n")
			s := repeatingConfig(v.(string))
			strBuilder.WriteString(s)
			continue
		}
		switch v.(type) {
		case string:
			if v.(string) != "" {
				if govalidator.IsInt(v.(string)) {
					temp, err := strconv.Atoi(v.(string))
					if err != nil {
						log.Fatal(err)
					}
					s := fmt.Sprintf("  "+o+"= %d \n", temp)
					strBuilder.WriteString(s)
				} else if v.(string) == "true" || v.(string) == "false" {
					b, err := strconv.ParseBool(v.(string))
					if err != nil {
						fmt.Println(err)
					}
					s := fmt.Sprintf("  "+o+"= %t \n", b)
					strBuilder.WriteString(s)
				} else if strings.HasPrefix(v.(string), "[") && strings.HasSuffix(v.(string), "]") {
					s := fmt.Sprintf("  " + o + "= " + v.(string) + "\n")
					strBuilder.WriteString(s)
				} else {
					s := fmt.Sprintf("  " + o + "= \"" + v.(string) + "\"\n")
					strBuilder.WriteString(s)
				}
			}
		case map[string]interface{}:
			strBuilder.WriteString("  " + o + " {\n")
			if len(v.(map[string]interface{})) != 0 {
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
						} else if strings.HasPrefix(i.(string), "[") && strings.HasSuffix(i.(string), "]") {
							s := fmt.Sprintf("  " + nestedK + "= " + i.(string) + "\n")
							strBuilder.WriteString(s)
						} else {
							s := fmt.Sprintf("  " + nestedK + " = \"" + i.(string) + "\"\n")
							strBuilder.WriteString(s)
						}
					}
				}
			}
			strBuilder.WriteString("}\n")
		}
	}

	return *strBuilder
}

// tags deal with input such as "k1=v1,k2=v2" and populate it into the tags field of the config
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

func terraformExists() bool {
	_, err := exec.LookPath("terraform")
	return err == nil
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
			case k.String() == "tags" || k.String() == "variables":
				temp := fmt.Sprintf("%s = %s", k, "{\n")
				strBuilder.WriteString(temp)
				strBuilder.WriteString(repeatingConfig(s))
			case strings.HasPrefix(s, "map"):
				strBuilder.WriteString(fmt.Sprintf("%s %s", k, "{\n"))
			case strings.Contains(s, "int"):
				strBuilder.WriteString(fmt.Sprintf("%s = %d\n", k, v.MapIndex(k)))
			case strings.Contains(s, "bool"):
				strBuilder.WriteString(fmt.Sprintf("%s = %t\n", k, v.MapIndex(k)))
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
