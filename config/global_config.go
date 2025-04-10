package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const AppDir = ".ccbs"
const GlobalConfigFileName = "global_config.json"

type GlobalConfig struct {
	UseGit      bool   `json:"use_git"`
	DefaultLang string `json:"default_lang"`
	CppStandard string `json:"cpp_std"`
	CStandard   string `json:"c_std"`
}

func (cfg *GlobalConfig) Validate() error {
	if cfg.DefaultLang != "cpp" && cfg.DefaultLang != "c" {
		return fmt.Errorf("value of '%s' is invalid as default language", cfg.DefaultLang)
	}
	return nil
}

func GetGlobalConfig() GlobalConfig {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error: '%s'. Unable to determine user home dir", err)
		return DefaultGlobalConfig
	}
	confPath := path.Join(homeDir, AppDir, GlobalConfigFileName)
	confDir := path.Join(homeDir, AppDir)
	if _, err = os.Stat(confDir); !os.IsNotExist(err) {
		err = os.MkdirAll(confDir, 0644)
		if err != nil {
			fmt.Printf("error: '%s'. Unable to create config dir with path '%s'.", err, confDir)
			return DefaultGlobalConfig
		}
	}
	_, err = os.Stat(confPath)
	if os.IsNotExist(err) {
		f, err := os.Create(confPath)
		if err != nil {
			fmt.Printf("error: '%s'. Global config file was not created", err)
			return DefaultGlobalConfig
		}
		defer f.Close()
		jsonData, err := json.MarshalIndent(DefaultGlobalConfig, "", "  ")
		if err != nil {
			fmt.Printf("error: '%s'. Global config file was not created", err)
			return DefaultGlobalConfig
		}
		_, err = f.Write(jsonData)
		if err != nil {
			fmt.Printf("error: '%s'. Global config was not written", err)
			return DefaultGlobalConfig
		}
		return DefaultGlobalConfig
	}
	f, err := os.Open(confPath)
	if err != nil {
		fmt.Printf("error: '%s'. Unable to open global config file", err)
		return DefaultGlobalConfig
	}
	defer f.Close()
	var globalConfig GlobalConfig
	err = json.NewDecoder(f).Decode(&globalConfig)
	if err != nil {
		fmt.Printf("error: '%s'. Unable to decode global config file", err)
		return DefaultGlobalConfig
	}
	if err = globalConfig.Validate(); err != nil {
		fmt.Printf("error: '%s'. Global config file is invalid", err)
		return DefaultGlobalConfig
	}
	return globalConfig
}

var DefaultGlobalConfig = GlobalConfig{
	UseGit:      true,
	DefaultLang: "cpp",
	CppStandard: "17",
	CStandard:   "99",
}
