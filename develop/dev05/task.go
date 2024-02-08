package dev

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func grep(filename string, pattern string, options Options) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	found := false
	contextBuffer := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		match := false
		if options.ignoreCase {
			match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		} else {
			match = strings.Contains(line, pattern)
		}

		if (options.invert && !match) || (!options.invert && match) {
			found = true

			if options.lineNumber {
				fmt.Printf("%d:", lineNumber)
			}

			fmt.Println(line)

			if options.count {
				continue
			}

			if options.after > 0 {
				contextBuffer = append(contextBuffer, line)
				if len(contextBuffer) > options.after {
					contextBuffer = contextBuffer[1:]
				}
			}

			if options.before > 0 {
				contextBuffer = append(contextBuffer, line)
				if len(contextBuffer) > options.before {
					contextBuffer = contextBuffer[:len(contextBuffer)-1]
				}
			}

			if options.context > 0 {
				contextBuffer = append(contextBuffer, line)
				if len(contextBuffer) > options.context*2 {
					contextBuffer = contextBuffer[1 : len(contextBuffer)-1]
				}
			}
		} else {
			found = false
		}
	}

	if options.count && found {
		fmt.Printf("Total matches: %d\n", 1)
	}
}

type Options struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	lineNumber bool
}
