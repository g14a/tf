package types

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// TfPrompt is a custom prompt
type TfPrompt struct {
	Label  string
	Prompt promptui.Prompt
}

// TfSelect is a custom select
type TfSelect struct {
	Label  string
	Select promptui.Select
}

// Run runs the custom TfPrompt
func (p TfPrompt) Run() (string, error) {
	if p.Label != "" {
		color.Green("\n"+p.Label+"\n\n", "text")
	}

	value, err := p.Prompt.Run()

	return value, err
}

// Run runs the custom TfSelect
func (s TfSelect) Run() (string, error) {
	if s.Label != "" {
		color.Green("\n"+s.Label+"\n\n", "text")
	}

	_, value, err := s.Select.Run()

	return value, err
}

type Schema struct {
	Type      string
	Field     string
	Doc       string
	Ex        string
	Validator func(string) error
	Items     []string
}

func ProvidePS(schemas []Schema) ([]string, []string, map[string]TfPrompt, map[string]TfSelect) {
	prompts := map[string]TfPrompt{}
	selects := map[string]TfSelect{}

	var promptOrder, selectOrder []string

	for _, v := range schemas {
		switch v.Type {
		case "select":
			s := TfSelect{
				Label: "Enter " + v.Field + "\n" + v.Doc,
				Select: promptui.Select{
					Label: "",
					Items: v.Items,
				},
			}
			selects[v.Field] = s
			selectOrder = append(selectOrder, v.Field)
		default:
			var p TfPrompt
			if v.Ex == "" {
				p = TfPrompt{
					Label: "Enter " + v.Field + "\n" + v.Doc,
					Prompt: promptui.Prompt{
						Label:    "",
						Validate: v.Validator,
					},
				}
			} else {
				p = TfPrompt{
					Label: "Enter " + v.Field + ": e.g. " + v.Ex + "\n" + v.Doc,
					Prompt: promptui.Prompt{
						Label:    "",
						Validate: v.Validator,
					},
				}
			}
			prompts[v.Field] = p
			promptOrder = append(promptOrder, v.Field)
		}
	}

	return promptOrder, selectOrder, prompts, selects
}
