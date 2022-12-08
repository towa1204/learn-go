package main

// based https://pkg.go.dev/encoding/csv#Writer example

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	file, err := os.Create("q2.txt")
	if err != nil {
		panic(err)
	}

	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(io.MultiWriter(os.Stdout, file))

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
