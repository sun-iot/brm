package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// Source 定义了单个源的结构
type Source struct {
	Name         string `yaml:"name"`
	DisplayName  string `yaml:"displayName"`
	BrewGit      string `yaml:"brewGit"`
	CoreGit      string `yaml:"coreGit"`
	CaskGit      string `yaml:"caskGit"`
	BottleDomain string `yaml:"bottleDomain"`
}

// Config 定义了配置文件的结构
type Config struct {
	Sources       []Source `yaml:"sources"`
	CurrentSource string   `yaml:"currentSource"`
}

var globalConfig Config

func LoadConfig(cfgPath string) error {
	file, err := os.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("加载配置项失败: %v", err)
	}
	if err := yaml.Unmarshal(file, &globalConfig); err != nil {
		return fmt.Errorf("无法解析配置文件: %v", err)
	}

	return nil
}

// InitLoadConfig 用于第一次的配置文件的加载
func InitLoadConfig(defaultPath string) error {
	// 根据默认的配置文件生成当前的结构体
	if err := yaml.Unmarshal([]byte(defaultConfig), &globalConfig); err != nil {
		return fmt.Errorf("无法解析默认的配置文件: %v", err)
	}
	// 将当前的默认的配置文件写入到 默认的文件目录里面去
	if err := os.WriteFile(filepath.Join(defaultPath, ".brm.yaml"), []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("无法去初始化配置: %v", err)
	}

	return nil
}

func GetAllSource() Config {
	return globalConfig
}

// GetCurrentSourceConfig 返回当前源的配置
func GetCurrentSourceConfig() (Source, error) {
	for _, source := range globalConfig.Sources {
		if source.Name == globalConfig.CurrentSource {
			return source, nil
		}
	}
	return Source{}, fmt.Errorf("current source not found")
}

var defaultConfig = `
sources:
  - name: official
    displayName: "官方源"
    brewGit: "https://github.com/Homebrew/brew.git"
    coreGit: "https://github.com/Homebrew/homebrew-core.git"
    caskGit: https://github.com/Homebrew/homebrew-cask.git
    bottleDomain: https://homebrew.bintray.com

  - name: aliyun
    displayName: "阿里云"
    brewGit: "https://mirrors.aliyun.com/homebrew/brew.git"
    coreGit: "https://mirrors.aliyun.com/homebrew/homebrew-core.git"
    caskGit: "https://mirrors.aliyun.com/homebrew/homebrew-cask.git"
    bottleDomain: "https://mirrors.aliyun.com/homebrew/homebrew-bottles"

  - name: tsinghua
    displayName: "清华源"
    brewGit: "https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git"
    coreGit: "https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git"
    caskGit: "https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-cask.git"
    bottleDomain: "https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles"

  - name: ustc
    displayName: "科大源"
    brewGit: "https://mirrors.ustc.edu.cn/brew.git"
    coreGit: "https://mirrors.ustc.edu.cn/homebrew-core.git"
    caskGit: "https://mirrors.ustc.edu.cn/homebrew-cask.git"
    bottleDomain: "https://mirrors.ustc.edu.cn/homebrew-bottles"

currentSource: official
`
