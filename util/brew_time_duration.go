package util

import (
	"os/exec"
	"time"
)

func GetGitLsRemote(url string) (time.Duration, error) {
	startTime := time.Now()

	cmd := exec.Command("git", "ls-remote", url)
	if err := cmd.Run(); err != nil {
		return 0, err
	}

	duration := time.Since(startTime)
	return duration, nil
}

func GetBrewFormulaTime(formula string) (time.Duration, error) {
	startTime := time.Now()

	cmd := exec.Command("brew", "fetch", "--dry-run", formula)
	if err := cmd.Run(); err != nil {
		return 0, err
	}

	duration := time.Since(startTime)
	return duration, nil
}
