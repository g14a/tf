package file

import (
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"strings"
)

var TerraformFile os.File

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

func FilePrompt() {

	filePrompt := promptui.Prompt{
		Label: "Enter terraform file name to save your configuration in this directory",
	}

	fileName, err := filePrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	TFFileInstance(fileName)
}