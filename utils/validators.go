package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var r *regexp.Regexp

func IntValidator(input string) error {
	if input != "" {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
	} else {
		return nil
	}
	return nil
}

func StringValidator(input string) error {
	if input == "" {
		_, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return errors.New("not a string")
		} else {
			return nil
		}
	}
	return nil
}

func RCValidator(input string) error {
	if input != "" {
		if !strings.Contains(input, "=") {
			return errors.New("should be of the format k1=v1")
		}
	}
	return nil
}

func BoolValidator(input string) error {
	if input != "" {
		if input == "true" || input == "false" {
			return nil
		} else {
			return errors.New("only true/false allowed")
		}
	}
	return nil
}

func BlockNameValidator(input string) error {
	if !r.MatchString(input) {
		return errors.New("a block name must start with a letter or underscore and may contain only letters, digits, underscore and dashes")
	}
	return nil
}

//
//func init() {
//	r, _ = regexp.Compile("^[a-zA-Z0-9_]*$")
//}