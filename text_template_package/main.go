package main

import (
	"os"
	"text/template"
)

// Entry is a structure
type Entry struct {
	Number, Square int
}

func main() {
	data := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}
	var Entries []Entry

	for _, value := range data {
		Entries = append(Entries, Entry{value[0], value[1]})
	}

	t := template.Must(template.ParseGlob("data.txt"))
	t.Execute(os.Stdout, Entries)
}
