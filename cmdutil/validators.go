package cmdutil

import (
	"errors"
	"regexp"
)

func IsValidNameString(s string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !regex.MatchString(s) {
		return errors.New("use of special characters or spaces not allowed")
	}

	return nil
}
