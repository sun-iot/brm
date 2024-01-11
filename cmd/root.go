/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sun-iot/brm/internal/config"
	"github.com/sun-iot/brm/internal/util"
	"os"
	"path/filepath"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "brm",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initBrm)
}

// 这里需要对brm 进行一个初始化的配置，且每次使用时，都要去检测这个配置
func initBrm() {
	// 首先，需要初始化一个文件目录
	path, err := util.GetHomePath()
	if err != nil {
		os.Exit(1)
	}

	brmPath := filepath.Join(path, ".brm")

	if !util.IsDirExist(brmPath) {
		// 不存在这个目录，就需要去创建
		if err := os.MkdirAll(brmPath, 0755); err != nil {
			os.Exit(1)
		}
		// 这里需要去构建出,将程序内置的配置提供出来 config.yaml
		if err := config.InitLoadConfig(brmPath); err != nil {
			os.Exit(1)
		}

	} else {
		// 文件目录存在，需要在次判断是否存在 .brm.yaml
		cfgPath := filepath.Join(brmPath, ".brm.yaml")
		if !util.IsDirExist(cfgPath) {
			// 这里需要去构建出,将程序内置的配置提供出来 config.yaml
			if err := config.InitLoadConfig(brmPath); err != nil {
				os.Exit(1)
			}
		} else {
			// 这里直接根据当前的配置项去获取到配置
			if err := config.LoadConfig(cfgPath); err != nil {
				os.Exit(1)
			}
		}
	}
}
