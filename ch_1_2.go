package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, str := range files {
			file, err := os.Open(str)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%s\t%d\n", line, n)
	}
}

func countLines(f *os.File, counts map[string]int) {
	in := bufio.NewScanner(f)
	for in.Scan() {
		l := in.Text()
		counts[l]++
		c := counts[l]
		if c > 1 {
			fmt.Printf("Found duplicate line in %s, line => %s, counts => %d.\n", f.Name(), l, counts[l])
		}
	}
}

func main() {
	dup2()
}
