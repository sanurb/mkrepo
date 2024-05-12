package github

import (
	"fmt"
	"github.com/sanurb/mkrepo/internal/config"
	"os"
	"os/exec"
	"strings"
	"time"
)

func CheckGHCLI() error {
	if _, err := exec.LookPath("gh"); err != nil {
		return fmt.Errorf("error: The 'gh' command is not available. Make sure GitHub CLI is installed")
	}
	return nil
}

func CreateRepository(cfg config.RepoConfig) error {
	visibility := "--private"
	if cfg.IsPublic {
		visibility = "--public"
	}

	repoPath := cfg.RepoName
	if cfg.OrgName != "" {
		repoPath = fmt.Sprintf("%s/%s", cfg.OrgName, cfg.RepoName)
	}

	cmd := exec.Command("gh", "repo", "create", repoPath, visibility, "-p", cfg.TemplateName, "-d", cfg.Description)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error creating repository: %v", err)
	}
	return nil
}

func CloneRepository(repoName string) error {
	userName, err := getGitHubUserName()
	if err != nil {
		return err
	}

	fullRepoPath := fmt.Sprintf("%s/%s", userName, repoName)
	for i := 0; i < 3; i++ {
		cmd := exec.Command("gh", "repo", "clone", fullRepoPath)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err == nil {
			return nil
		}
		if i < 2 { // Wait for 2 seconds before retrying
			time.Sleep(2 * time.Second)
		}
	}
	return fmt.Errorf("failed to clone repository after several attempts")
}

func getGitHubUserName() (string, error) {
	cmd := exec.Command("gh", "api", "user", "-q", ".login")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func ProcessRepository(cfg config.RepoConfig) error {
	if err := CheckGHCLI(); err != nil {
		return err
	}

	if err := CreateRepository(cfg); err != nil {
		return err
	}

	return CloneRepository(cfg.RepoName)
}
