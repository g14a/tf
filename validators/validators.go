package validators

import (
	"errors"
	"fmt"
	"net"
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

func CIDRValidator(input string) error {
	_, ipnet, err := net.ParseCIDR(input)
	if err != nil {
		return fmt.Errorf("%q is not a valid CIDR block: %w", input, err)
	}

	if !cidrBlocksEqual(input, ipnet.String()) {
		return fmt.Errorf("%q is not a valid CIDR block; did you mean %q?", input, ipnet)
	}

	return nil
}

func cidrBlocksEqual(cidr1, cidr2 string) bool {
	ip1, ipnet1, err := net.ParseCIDR(cidr1)
	if err != nil {
		return false
	}
	ip2, ipnet2, err := net.ParseCIDR(cidr2)
	if err != nil {
		return false
	}

	return ip2.String() == ip1.String() && ipnet2.String() == ipnet1.String()
}

func Ipv4CIDRBlockValidator(input string) error {
	ip, ipnet, err := net.ParseCIDR(input)
	if err != nil {
		return fmt.Errorf("%q is not a valid CIDR block: %w", input, err)
	}

	ipv4 := ip.To4()
	if ipv4 == nil {
		return fmt.Errorf("%q is not a valid IPv4 CIDR block", input)
	}

	if !cidrBlocksEqual(input, ipnet.String()) {
		return fmt.Errorf("%q is not a valid IPv4 CIDR block; did you mean %q?", input, ipnet)
	}

	return nil
}

func Ipv6CIDRBlockValidator(input string) error {
	ip, ipnet, err := net.ParseCIDR(input)
	if err != nil {
		return fmt.Errorf("%q is not a valid CIDR block: %w", input, err)
	}

	ipv4 := ip.To4()
	if ipv4 != nil {
		return fmt.Errorf("%q is not a valid IPv6 CIDR block", input)
	}

	if !cidrBlocksEqual(input, ipnet.String()) {
		return fmt.Errorf("%q is not a valid IPv6 CIDR block; did you mean %q?", input, ipnet)
	}

	return nil
}
