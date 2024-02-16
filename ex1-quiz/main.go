package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Settings struct {
	file string
	timer time.Duration
}

func checkFlags() Settings {
	var csvFlag = flag.String("f", "problems.csv", "CSV file with input problems and answers")
	var timerFlag = flag.String("t", "30", "How long should the timer be in seconds (default: 30)")
	flag.Parse()
	timer, err := strconv.Atoi(*timerFlag)

	if err != nil {
		log.Fatal("Invalid time value provided")
	}

	return Settings{
		file: *csvFlag,
		timer: time.Duration(timer) * time.Second,
	}
}

func getCsvReader(filename string) (*csv.Reader, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not open file, %s", filename))
	}

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
		return true
	}
	return false	
}

func askQuestions(correctAnswers *int, csvReader *csv.Reader, done chan bool) {
	for problem, err := csvReader.Read(); problem != nil; {
		if err != nil {
			log.Fatal("Could not parse csv file")
		}

		answer := promptQuestion(problem)

		if answer {
			*correctAnswers++
		}
		problem, err = csvReader.Read()
	}
	done <- true
}

func main() {
	var done chan bool
	settings := checkFlags()
	csvReader, file := getCsvReader(settings.file)
	defer file.Close()
	
	fmt.Printf("Hit enter, when ready to start")
	fmt.Scanf("\n")

	correctAnswers := 0
	done = make(chan bool)
	go askQuestions(&correctAnswers, csvReader, done)

	go func() {
		time.Sleep(settings.timer)
		done <- true
	}()

	if(<-done) {
		fmt.Printf("\nYou got %d correct.\n", correctAnswers)
	}
}
