package config

import "fmt"

type GlobalConfig struct {
	UseGit      bool   `json:"use_git"`
	DefaultLang string `json:"default_lang"`
	CppStandard string `json:"cpp_std"`
	CStandard   string `json:"c_std"`
}

func (cfg *GlobalConfig) Validate() error {
	if cfg.DefaultLang != "cpp" && cfg.DefaultLang != "c" {
		return fmt.Errorf("Value of '%s' is invalid as default language", cfg.DefaultLang)
	}
	return nil
}

func GetGlobalConfig() GlobalConfig {
	return DefaultGlobalConfig
}

var DefaultGlobalConfig = GlobalConfig{
	UseGit:      true,
	DefaultLang: "cpp",
	CppStandard: "17",
	CStandard:   "99",
}
