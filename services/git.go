package services

import (
	"os/exec"

	"github.com/spf13/viper"
)

// GetGitUsername tries to find the git username inside cli config or in the git config
func GetGitUsername() string {
	var userName string

	// try to get from viper
	userName = viper.GetString("git-username")
	if userName != "" {
		return userName
	}

	// Try to get from git config
	userName = getFromGit()
	storeToConfig(userName)

	return userName
}

func getFromGit() string {
	cmd := exec.Command("git", "config", "user.name")
	stdout, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(stdout)
}

func storeToConfig(userName string) {
	if userName != "" {
		viper.Set("git-username", userName)
		viper.WriteConfig()
	}
}

// GitInit initializes a git repository to the specified path
func GitInit(path string) error {
	return exec.Command("git", "init", path).Run()
}
