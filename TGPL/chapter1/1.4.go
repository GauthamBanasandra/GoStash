package main

import (
	"os"
	"bufio"
	"fmt"
)

type Set map[string]int

func main() {
	counts := make(map[string]int)
	fileLine := make(map[string]Set)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "Too few arguments")
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup err:%v", err)
				continue
			}

			input := bufio.NewScanner(file)
			for input.Scan() {
				line := input.Text()
				if len(line) > 0 {
					counts[line]++
					if _, exists := fileLine[line]; !exists {
						fileLine[line] = make(Set)
					}
					fileLine[line][arg] = 0
				}
			}
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\t", count, line)
			for file := range fileLine[line] {
				fmt.Printf("%s\t", file)
			}
			fmt.Println()
		}
	}
}
