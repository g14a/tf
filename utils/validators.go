package utils

import (
	"errors"
	"strconv"
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
		return errors.New("empty file path")
	}
	return nil
}
