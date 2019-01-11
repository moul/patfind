package main // import "moul.io/patfind"

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"

	"moul.io/patfind/score"
)

func main() {
	data := [][]string{}

	for _, arg := range os.Args[1:] {
		sequence := score.New(arg)
		score, comments := sequence.Score()
		shift := math.Pow(10, float64(5))
		roundedScore := math.Floor(score*shift+.5) / shift
		data = append(data, []string{sequence.String(), fmt.Sprintf("%v", roundedScore), strings.Join(comments, ";\n")})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sequence", "Score", "Comments"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColWidth(50)
	table.AppendBulk(data)
	table.Render()
}
