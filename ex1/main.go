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

func getCsvReader(filename string) (*csv.Reader, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not open file, %s", filename))
	}
	// defer file.Close()


	csvReader := csv.NewReader(file)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not read file, %s", filename))
	}
	return csvReader, file
}

func promptQuestion(problem []string) bool {
	var answer string
	fmt.Printf("What is %s?: ", problem[0])
	fmt.Scan(&answer)

	if answer == problem[1] {
		fmt.Println("CORRECT!")
		return true
	}
	fmt.Println("WRONG!")
	return false	
}

func main() {
	settings := checkFlags()
	csvReader, file := getCsvReader(settings.file)
	defer file.Close()

	for problem, err := csvReader.Read(); problem != nil; {
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not parse csv, %s", settings.file))
		}

		promptQuestion(problem)
		problem, err = csvReader.Read()
	}

}
