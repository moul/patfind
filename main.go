package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{}

	for _, arg := range os.Args[1:] {
		number := New(arg)
		score, comments := number.Score()
		data = append(data, []string{number.String(), fmt.Sprintf("%5f", score), strings.Join(comments, ";\n")})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Number", "Score", "Comments"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColWidth(50)
	table.AppendBulk(data)
	table.Render()
}
