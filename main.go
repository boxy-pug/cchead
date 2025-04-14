package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type headObj struct {
	inputFiles    []*os.File
	numLines      int
	numBytes      int
	lineFlag      bool
	byteFlag      bool
	multipleFiles bool
}

func main() {

	var numLines int
	var numBytes int
	flag.IntVar(&numLines, "n", 10, "number of lines to print")
	flag.IntVar(&numBytes, "c", 0, "number of bytes to print")

	flag.Parse()
	remainingArgs := flag.Args()

	ho := headObj{
		numLines:      numLines,
		numBytes:      numBytes,
		lineFlag:      true,
		multipleFiles: false,
	}

	if len(remainingArgs) > 1 {
		ho.multipleFiles = true
	}

	if ho.numBytes > 0 && ho.numLines == 10 {
		ho.lineFlag = false
		ho.byteFlag = true
	}

	for _, file := range remainingArgs {
		openFile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ho.inputFiles = append(ho.inputFiles, openFile)
	}

	ho.getHead()

}

func (ho *headObj) getHead() {
	for i, file := range ho.inputFiles {
		scanner := bufio.NewScanner(file)
		num := 0
		if ho.multipleFiles {
			fmt.Printf("==> %s <==\n", file.Name())
		}

		scanner.Split(bufio.ScanBytes)

		for scanner.Scan() {
			curByte := scanner.Text()
			fmt.Printf("%s", curByte)

			if ho.lineFlag && (curByte == "\n" || curByte == "\r\n") {
				num += 1
			}
			if ho.byteFlag {
				num += 1
			}

			if (ho.lineFlag && num == ho.numLines) || (ho.byteFlag && num == ho.numBytes) {
				break
			}

		}
		if ho.multipleFiles && i < len(ho.inputFiles)-1 {
			fmt.Println()
		}
	}
}
