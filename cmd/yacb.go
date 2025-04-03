package cmd

import (
	"fmt"
	"os"
	"path"
)

func bootstrapProject(root string) {
	fmt.Printf("Project created in '%s'!\n", root)
}

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
