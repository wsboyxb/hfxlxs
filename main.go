package main

import (
	"fmt"
	"github.com/tealeg/xlsx/v2"
	"regexp"
	"strings"
)

func main() {
	excelFileName := "abc.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	m := make(map[string][]string)
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var to string
			var from = make([]string, 0, 1)
			for idx, cell := range row.Cells {
				colLetter := xlsx.ColIndexToLetters(idx)
				if colLetter == "A" {
					continue
				}
				text := cell.String()
				if len(text) == 0 {
					continue
				}
				if colLetter == "B" {
					to = trim(text)
				} else {
					from = append(from, trim(text))
				}
				//fmt.Printf("%s\t", trim(text))
			}
			m[to] = from
			//fmt.Println()
		}
	}

	for to, from := range m {
		fmt.Print("\n", to, "\t")
		for _, f := range from {
			fmt.Print(f, "\t")
		}
	}

}

//str := "M46-装甲风暴[6046]" --> 6046
func trim(s string) string {
	//str := "M46-装甲风暴[6046]"
	re := regexp.MustCompile("\\[[0-9]+]")
	allString := re.FindAllString(s, 1)
	if len(allString) > 0 {
		s2 := allString[0]
		s2 = strings.TrimPrefix(s2, "[")
		s2 = strings.TrimSuffix(s2, "]")

		return s2
	}
	return ""
}
