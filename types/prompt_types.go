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
	color.Green("\n"+p.Label+"\n\n", "text")

	value, err := p.Prompt.Run()

	return value, err
}

// Run runs the custom TfSelect
func (s TfSelect) Run() (string, error) {
	color.Green("\n"+s.Label+"\n\n", "text")

	_, value, err := s.Select.Run()

	return value, err
}
