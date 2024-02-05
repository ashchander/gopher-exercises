package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Settings struct {
	file string
}

func checkFlags() Settings {
	var csvFlag = flag.String("f", "problems.csv", "CSV file with input problems and answers")
	flag.Parse()
	return Settings{
		file: *csvFlag,
	}
}

func main() {
	settings := checkFlags()
	data, err := os.ReadFile(settings.file)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not read file, %s", settings.file))
	}

	csvReader := csv.NewReader(data)

	os.Stdout.Write(data)
}
