package util

import (
	"github.com/pkg/errors"
	"os"
)

func IsDirExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

func GetHomePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.WithMessage(err, "unable to obtain user directory")
	}
	return dir, nil
}
