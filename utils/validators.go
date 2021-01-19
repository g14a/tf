package utils

import (
	"errors"
	"strconv"
	"strings"
)

// IntValidator validates whether input is an integer
func IntValidator(input string) error {
	if input != "" {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
	}
	return nil
}

// StringValidator validates whether input is actually a string
func StringValidator(input string) error {
	if input == "" {
		_, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return errors.New("not a string")
		}
	}
	return nil
}

// RCValidator validates whether input is of the form k1=v1,k2=v2
func RCValidator(input string) error {
	if input != "" {
		if !strings.Contains(input, "=") {
			return errors.New("should be of the format k1=v1")
		}
	}
	return nil
}

// BoolValidator validates whether input is true/false
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

// BlockDurationValidator validates whether input is a multiple of 60
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

func MinMaxIntValidator(min, max int64) func(input string) error {
	return func(input string) error {
		if input != "" {
			f, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				return errors.New("invalid number")
			}
			if f >= min && f <= max {
				return nil
			} else {

				return errors.New("value has to be within " + strconv.FormatInt(min, 10) + " and " + strconv.FormatInt(max, 10))
			}
		}
		return nil
	}
}
