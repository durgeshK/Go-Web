package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := make(map[string][]Hotel)

	hotels["Southern"] = []Hotel{
		Hotel{"Bay Watch", "Juhu Beach", "Mumbai", "100001"},
		Hotel{"Taj Resort", "Central Secretariat", "New Delhi", "100002"},
		Hotel{"Happy Stay", "Juhu Chowpaty", "Chennai", "100003"},
		Hotel{"Hotel Jhakas", "M.G.Road", "Kolkata", "100004"},
	}
	hotels["Central"] = []Hotel{
		Hotel{"Hotel Rajdhani", "Juhu Beach", "Bhopal", "100001"},
		Hotel{"Taj Resort", "Central Secretariat", "Nagpur", "100002"},
		Hotel{"Happy Stay", "Juhu Chowpaty", "Raipur", "100003"},
		Hotel{"Hotel Jhakas", "M.G.Road", "Ranchi", "100004"},
	}
	hotels["Northern"] = []Hotel{
		Hotel{"Bay Watch", "Juhu Beach", "Mumbai", "100001"},
		Hotel{"Taj Resort", "Central Secretariat", "New Delhi", "100002"},
		Hotel{"Happy Stay", "Juhu Chowpaty", "Chennai", "100003"},
		Hotel{"Hotel Jhakas", "M.G.Road", "Kolkata", "100004"},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
