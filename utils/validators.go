package utils

import (
	"errors"
	"strconv"
)

func IntValidator(input string) error {
	if input != "" {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
	} else {
		return nil
	}
	return nil
}

