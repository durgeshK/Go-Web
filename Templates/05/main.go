package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"strconv"
)

type Record struct {
	Index int64
	Date  string
	Open  float64
	Close float64
}

func readFile(fileName string) []Record {
	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	records := make([]Record, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}
		opening, _ := strconv.ParseFloat(row[1], 64)
		closing, _ := strconv.ParseFloat(row[4], 64)
		records = append(records, Record{
			int64(i),
			row[0],
			opening,
			closing,
		})
	}

	return records

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	records := readFile("table.csv")
	er := tpl.Execute(os.Stdout, records)
	if er != nil {
		log.Fatalln(er)
	}
}
