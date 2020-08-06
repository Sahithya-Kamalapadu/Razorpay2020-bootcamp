package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type quiz struct {
	que string
	ans string
}

func main() {
	filename := flag.String("csvfile", "problems.csv", "contains questions and answer to it")
	//flag.Parse()
	timelimit := flag.Int("limit", 15, "the time limit for the quiz in seconds")
	file := openFile(*filename)
	lines := reader(file)
	questions := []quiz{}
	for _, r := range lines {
		questions = append(questions, quiz{r[0], r[1]})
	}
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	score := 0
	for _, line := range questions {
		fmt.Printf(" %s ", line.que)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d", score, len(questions))
			return
		case answer := <-answerCh:
			if answer == line.ans {
				score++
			}
		}
	}

}

//To load the csv file
func openFile(f string) io.Reader {
	file, err := os.Open(f)

	if err != nil {
		fmt.Printf("Failed to open the file: %s", f)
		os.Exit(1)
	}

	return file
}

//TO parse the data in csv file
func reader(r io.Reader) [][]string {
	lines, err := csv.NewReader(r).ReadAll()

	if err != nil {
		fmt.Println("Could not Parse all the problems")
	}
	return lines
}
