package utils

import (
	"errors"
	"strconv"
	"strings"
)

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

func BlockDurationValidator(input string) error {
	if input != "" {
		f, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		if f%60 != 0 {
			return errors.New("60 or multiple of 60 needed")
		}
	} else {
		return nil
	}
	return nil
}
