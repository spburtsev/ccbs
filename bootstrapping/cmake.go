package bootstrapping

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/spburtsev/ccbs/config"
)

const cmakeTemplate = `cmake_minimum_required(VERSION {{ .CmakeVersion }})

project({{ .ProjectName }})

set(CMAKE_CXX_STANDARD {{ .CppStandard }})
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Set base directory for outputs
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/out) # Executables
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/out/lib) # Libraries
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/out/lib) # Static libraries

add_executable( 
    {{ .ProjectName }}
    {{ .ProjectName }}_main.cpp
)
`
const mainContent = `#include <cstdio>

int main() {
	printf("Hello, World!\n");
	return 0;
}
`

func createCMakeLists(root string, config *config.GlobalConfig) error {
	t, err := template.New("cmake").Parse(cmakeTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err) // Use fmt.Errorf for better error wrapping
	}
	listsPath := path.Join(root, "CMakeLists.txt")
	file, err := os.Create(listsPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	projectName := path.Base(root)
	data := struct {
		CmakeVersion string
		ProjectName  string
		CppStandard  string
	}{
		CmakeVersion: config.CmakeVersion,
		ProjectName:  projectName,
		CppStandard:  config.CppStandard,
	}
	err = t.Execute(file, data)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}
	return nil
}

// func createSrcDir(root string) error {
// 	srcDir := path.Join(root, "src")
// 	// if directory exists, do not create it
// 	if _, err := os.Stat(srcDir); !os.IsNotExist(err) {
// 		err = os.Mkdir(srcDir, 0755)
// 		if err != nil {
// 			return fmt.Errorf("error creating src directory: %w", err)
// 		}
// 	} else if err != nil {
// 		return fmt.Errorf("error checking src directory: %w", err)
// 	}
// 	// Create main.cpp in src directory
// 	err := createMainFile(srcDir)
// 	if err != nil {
// 		return fmt.Errorf("error creating main.cpp: %w", err)
// 	}
// 	return nil
// }

func createMainFile(root string) error {
	projectName := path.Base(root)
	mainName := fmt.Sprintf("%s_main.cpp", projectName)
	mainFilePath := path.Join(root, mainName)
	file, err := os.Create(mainFilePath)
	if err != nil {
		return fmt.Errorf("error creating main.cpp: %w", err)
	}
	defer file.Close()
	_, err = file.WriteString(mainContent)
	if err != nil {
		return fmt.Errorf("error writing to main.cpp: %w", err)
	}
	return nil
}
