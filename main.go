package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// to hold questions and their answer
type question struct {
	question string
	answer   int
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

		// convert answer to int
		answerInt, errInt := strconv.Atoi(record[1])
		check(errInt)

		data = append(data, question{record[0], answerInt})
	}

	return data
}

func askQuestion(q question) bool {
	// print the question
	fmt.Printf("%v = ", q.question)
	var userString string

	// get user input for the answer and conver to int
	fmt.Scanln(&userString)
	userAnswer, err := strconv.Atoi(userString)
	if err != nil {
		return false
	}

	return userAnswer == q.answer
}

func main() {

	score := 0
	data := csvToQuestions("problems.csv")

	// will always run after quiz or timer
	defer fmt.Printf("Game over! You scored %v/%v\n", score, len(data))

	// quiz loop
	for _, data := range data {
		if askQuestion(data) {
			score++
		}
	}

}
