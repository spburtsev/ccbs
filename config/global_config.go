package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const appDir = ".ccbs"
const globalConfigFileName = "global_config.json"

type GlobalConfig struct {
	UseGit        bool   `json:"use_git"`
	CmakeVersion  string `json:"cmake_version"`
	DefaultLang   string `json:"default_lang"`
	CppStandard   string `json:"cpp_std"`
	CStandard     string `json:"c_std"`
	AutoExecCmake bool   `json:"auto_exec_cmake"`
}

func GlobalConfigFromFile(f *os.File) (*GlobalConfig, error) {
	var conf GlobalConfig
	err := json.NewDecoder(f).Decode(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func (cfg *GlobalConfig) Validate() error {
	if cfg.DefaultLang != "cpp" && cfg.DefaultLang != "c" {
		return fmt.Errorf("value of '%s' is invalid as default language", cfg.DefaultLang)
	}
	return nil
}

func (cfg *GlobalConfig) Serialize() ([]byte, error) {
	return json.MarshalIndent(cfg, "", "  ")
}

var DefaultGlobalConfig = GlobalConfig{
	UseGit:        true,
	CmakeVersion:  "3.10",
	DefaultLang:   "cpp",
	CppStandard:   "17",
	CStandard:     "99",
	AutoExecCmake: true,
}

func ReadGlobalConfig() (*GlobalConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	confPath := path.Join(homeDir, appDir, globalConfigFileName)
	configFile, err := os.Open(confPath)
	if os.IsNotExist(err) {
		return createGlobalConfig(confPath)
	}
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	return GlobalConfigFromFile(configFile)
}

func ResetGlobalConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	confPath := path.Join(homeDir, appDir, globalConfigFileName)
	_, err = os.Stat(confPath)
	if os.IsNotExist(err) {
		_, err = createGlobalConfig(confPath)
		return err
	}
	if err != nil {
		return err
	}
	newContent, err := DefaultGlobalConfig.Serialize()
	if err != nil {
		return err
	}
	oldConfig, err := os.OpenFile(confPath, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer oldConfig.Close()
	_, err = oldConfig.Write(newContent)
	return err
}

func createGlobalConfig(confPath string) (*GlobalConfig, error) {
	confDir := filepath.Dir(confPath)
	err := os.MkdirAll(confDir, 0755)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(confPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	jsonData, err := DefaultGlobalConfig.Serialize()
	if err != nil {
		return nil, err
	}
	_, err = f.Write(jsonData)
	if err != nil {
		return nil, err
	}
	return &DefaultGlobalConfig, nil
}
