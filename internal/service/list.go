package service

import (
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cast"
	"github.com/sun-iot/brm/internal/config"
	"github.com/sun-iot/brm/util"
	"sync"
	"time"
)

// ListLocalSource 列举出本地所有的源
// speed: 可以对源进行一个测速
func ListLocalSource(speed bool) {
	allSource := config.GetAllSource()
	color.GreenString("Current Brew Source is %v", allSource.CurrentSource)
	header := []string{"源名称", "地址", "延时", "说明"}

	data := make([][]string, 4)
	bar := progressbar.Default(int64(len(allSource.Sources)), "Processing...")
	wg := sync.WaitGroup{}
	for _, source := range allSource.Sources {
		wg.Add(1)
		go func(source config.Source) {
			defer wg.Done()
			res := []string{
				color.BlueString(source.DisplayName),
				color.BlueString(source.CoreGit),
			}
			if speed {
				remote, err := util.GetGitLsRemote(source.CoreGit)
				if err != nil {
					res = append(res, "-", color.RedString(err.Error()))
				} else {
					res = append(res, color.GreenString(cast.ToString(remote.Milliseconds())+"ms"), "-")
				}
			} else {
				res = append(res, "-", "-")
			}
			bar.Add(1)
			data = append(data, res)
		}(source)

		time.Sleep(time.Millisecond * 10)
	}
	wg.Wait()
	util.PrintTable(header, data)
}
