package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// var lines int
// var numLines = flag.IntVar(&lines, "n", 10, "define how many lines to print")
var numLines = flag.Int("n", 10, "define line num")

func main() {
	flag.Parse()
	var input *os.File
	var err error

	if len(os.Args) == 2 {
		input, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

	} else if len(os.Args) > 2 {
		input, err = os.Open(os.Args[3])
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
	} else {
		input = os.Stdin
	}

	getHead(input, *numLines)

}

func getHead(input *os.File, lines int) {
	scanner := bufio.NewScanner(input)
	i := 0

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		i += 1
		if i == lines {
			break
		}

	}

}
