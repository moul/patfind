package main // import "moul.io/patfind"

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"

	"moul.io/patfind/score"
)

func main() {
	sequences := []score.Sequence{}

	if len(os.Args) < 2 { // stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Printf("error: failed to read line from stdin: %v", err)
				break
			} else {
				sequences = append(sequences, score.New(scanner.Text()))
			}
		}
	} else {
		for _, arg := range os.Args[1:] {
			sequences = append(sequences, score.New(arg))
		}
	}

	data := [][]string{}
	for _, sequence := range sequences {
		score, comments := sequence.Score()
		shift := math.Pow(10, float64(5))
		roundedScore := math.Floor(score*shift+.5) / shift
		data = append(data, []string{sequence.String(), fmt.Sprintf("%v", roundedScore), strings.Join(comments, ";\n")})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sequence", "Score", "Comments"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColWidth(100)
	table.AppendBulk(data)
	table.Render()
}
