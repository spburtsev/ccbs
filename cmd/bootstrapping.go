package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/spburtsev/ccbs/config"
)

func ExecInit() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	bootstrapProject(currentDir)
	return nil
}

func ExecNew(root string) error {
	if path.IsAbs(root) {
		bootstrapProject(root)
		return nil
	}
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	root = path.Join(currentDir, root)
	if _, err = os.Stat(root); !os.IsNotExist(err) {
		return fmt.Errorf("destination '%s' already exists", root)
	}
	bootstrapProject(root)
	return nil
}

func bootstrapProject(root string) {
	conf := config.GetGlobalConfig()
	if conf.UseGit {
		if !isGitAvailable() {
			fmt.Println("Git is not available. Skipping Git initialization.")
		} else if err := initGitRepo(root); err != nil {
			fmt.Println("Error initializing Git repository:", err)
		}
	}
	fmt.Printf("Project created in '%s'!\n", root)
}

func initGitRepo(root string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error initializing Git repository in %s: %v\nOutput: %s", root, err, string(output))
	}
	fmt.Printf("Successfully initialized Git repository in %s\nOutput: %s", root, string(output))
	return nil
}

func isGitAvailable() bool {
	_, err := exec.LookPath("git")
	return err == nil
}
