package util

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // 设置对齐方式
	table.AppendBulk(data)                     // 添加数据
	table.Render()                             // 渲染表格
}
