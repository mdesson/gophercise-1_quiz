package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// to hold questions and their answer
type question struct {
	question string
	answer   string
}

// error checking shorthand function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// convert csv file to slice of questions
func csvToQuestions(name string) []question {
	data := make([]question, 0)

	// open the file
	file, err := os.Open(name)
	check(err)

	// create csv reader
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, question{record[0], record[1]})
	}

	return data
}

// Asks a question and parses response
func askQuestion(q question) bool {
	// print the question
	fmt.Printf("%v = ", q.question)
	var userAnswer string

	// get user input for the answer and conver to int
	fmt.Scanln(&userAnswer)

	return userAnswer == q.answer
}

func main() {
	// Set up and parse flags
	filename := flag.String("f", "problems.csv", "csv filename for questions")
	timeSeconds := flag.Int("t", 30, "time, in seconds for quiz to run")
	flag.Parse()

	score := 0
	data := csvToQuestions(*filename)
	finishedQuiz := make(chan bool)

	fmt.Print("Press Enter when you are ready to begin...")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeSeconds) * time.Second)

	// quiz loop
	go func() {
		for _, data := range data {

			if askQuestion(data) {
				score++
			}
		}
		finishedQuiz <- true
	}()

	// break on channel that fills up first
	select {
	case <-timer.C:
	case <-finishedQuiz:
		break
	}

	fmt.Printf("Game over! You scored %v/%v\n", score, len(data))

}
