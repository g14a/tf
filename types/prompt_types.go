package types

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type TfPrompt struct {
	Label  string
	Prompt promptui.Prompt
}

type TfSelect struct {
	Label  string
	Select promptui.Select
}

func (p TfPrompt) Run() (string, error) {
	color.Green("\n"+p.Label+"\n\n", "text")

	value, err := p.Prompt.Run()

	return value, err
}

func (s TfSelect) Run() (string, error) {
	color.Green("\n"+s.Label+"\n\n", "text")

	_, value, err := s.Select.Run()

	return value, err
}
