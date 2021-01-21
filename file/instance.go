package file

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/g14a/tf/utils"
	"github.com/manifoldco/promptui"
)

var TerraformFile os.File

// TFFileInstance creates a file instance of the provided file name
func TFFileInstance(name string) {
	fileName := ""

	if strings.HasSuffix(name, ".tf") {
		fileName = name
	} else {
		fileName = name + ".tf"
	}

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	TerraformFile = *f
}

// Prompt runs the prompt in which the user
// enter the wanted terraform file name
func Prompt() {

	color.Green("\nEnter terraform file name to save your configuration in this directory\n", "text")

	filePrompt := promptui.Prompt{
		Label:    "",
		Validate: utils.StringValidator,
	}

	fileName, err := filePrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	TFFileInstance(fileName)
}
