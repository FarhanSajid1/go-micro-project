package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "filename", "problems.csv", "name of csv file to open")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening the file %v", err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error opening the file %v", err)
	}
	// contents of the csv file
	csvContents := string(b)
	s := strings.Split(csvContents, "\n")

	questions, answers := QuestionsAnswers(s)
	score := 0
	for i := 0; i < len(questions); i++ {
		fmt.Printf("%s = ", questions[i])
		input := bufio.NewReader(os.Stdin)
		ans, err := input.ReadString('\n')
		if err != nil {
			log.Printf("could not read input! %v", err)
		}
		ansSlice := strings.Split(ans, "\n")
		if string(ansSlice[0]) == answers[i] {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d", score, len(questions))

}

func QuestionsAnswers(s []string) ([]string, []string) {
	var questions []string
	var answers []string

	// splitting the csv file, ensuring that there is data that we can split
	for _, value := range s {
		if value != "" {
			// fmt.Println(value)
			splitting := strings.Split(value, ",")
			questions = append(questions, splitting[0])
			answers = append(answers, splitting[1])
		}
	}
	return questions, answers
}
