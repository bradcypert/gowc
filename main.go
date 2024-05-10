package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	c := flag.Bool("c", false, "output number of bytes")
	l := flag.Bool("l", false, "output number of lines")
	w := flag.Bool("w", false, "output number of words")
	flag.Parse()

	if len(os.Args) < 2 {
		panic("Please pass in the file path to parse")
	}

	filePath := os.Args[len(os.Args)-1]
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cant open provided file path!")
	}

	defer file.Close()

	var bytes, lines, words int

	if *c == false && *w == false && *l == false {
		*c = true
		*w = true
		*l = true
	}

	if *c == true {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			bytes++
		}
	}

	if *w == true {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	if *l == true {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines++
		}
	}

	if bytes != 0 {
		fmt.Printf("bytes: %d\n", bytes)
	}

	if words != 0 {
		fmt.Printf("words: %d\n", words)
	}

	if lines != 0 {
		fmt.Printf("lines: %d\n", lines)
	}

}
