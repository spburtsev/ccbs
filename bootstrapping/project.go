package bootstrapping

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
		err := ensureDirCreated(root)
		if err != nil {
			return err
		}
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
	err = ensureDirCreated(root)
	if err != nil {
		return err
	}
	bootstrapProject(root)
	return nil
}

const gitignoreTemplate = `# ccbs
/out
CMakeCache.txt
/CMakeFiles
Makefile
cmake_install.cmake

*.vcxproj
*.vcxproj.filters
*.vcxproj.user
*.sln
/x64
/*.dir/Debug
.vs
*.slnLaunch.user
`

func bootstrapProject(root string) error {
	conf, err := config.ReadGlobalConfig()
	if err != nil {
		return err
	}
	if conf.UseGit {
		if !isGitAvailable() {
			fmt.Printf("Git is not available. Skipping Git initialization.\n")
		} else if err := initGitRepo(root); err != nil {
			return err
		}
		gitignorePath := path.Join(root, ".gitignore")
		gitignoreFile, err := os.OpenFile(gitignorePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer gitignoreFile.Close()
		_, err = gitignoreFile.Write([]byte(gitignoreTemplate))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Project created in '%s'!\n", root)
	return nil
}

func initGitRepo(root string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error initializing a git repository in %s: %v\nOutput: %s", root, err, string(output))
	}
	fmt.Printf("Git repository initialized in %s\n", root)
	return nil
}

func isGitAvailable() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func ensureDirCreated(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Mkdir(path, 0755)
	}
	return err
}
