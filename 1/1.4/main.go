package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files_count := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, files_count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files_count)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v\t%d\t%s\n", files_count[line], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, files_count map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		if len(files_count[t]) == 0 || files_count[t][len(files_count[t])-1] != f.Name() {
			files_count[t] = append(files_count[t], f.Name())
		}
	}
}
