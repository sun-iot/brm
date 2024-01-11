package service

import (
	"github.com/fatih/color"
	"github.com/spf13/cast"
	"github.com/sun-iot/brm/internal/config"
	"github.com/sun-iot/brm/internal/util"

	"strings"
)

// ListLocalSource 列举出本地所有的源
// speed: 可以对源进行一个测速
func ListLocalSource(speed bool) {
	allSource := config.GetAllSource()
	color.GreenString("Current Brew Source is %v", allSource.CurrentSource)
	res := []string{"源名称", "Core地址", "Brew地址", "延时", "说明"}
	color.Blue(strings.Join(res, "\t"))

	for _, source := range allSource.Sources {
		res = []string{source.DisplayName, source.CoreGit, source.BrewGit}
		if speed {
			remote, err := util.GetGitLsRemote(source.CoreGit)
			if err != nil {
				res = append(res, "-", err.Error())
			} else {
				res = append(res, cast.ToString(remote.Microseconds())+"ms", "-")
			}
		} else {
			res = append(res, "-", "-")
		}
		color.Blue(strings.Join(res, "\t"))
	}
}
