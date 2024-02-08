package dev

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cut(input *os.File, fields string, delimiter string, separated bool) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		parts := strings.Split(line, delimiter)
		selectedFields := make([]string, 0)

		for _, field := range strings.Split(fields, ",") {
			fieldNum := parseFieldNumber(field)
			if fieldNum > 0 && fieldNum <= len(parts) {
				selectedFields = append(selectedFields, parts[fieldNum-1])
			}
		}

		fmt.Println(strings.Join(selectedFields, delimiter))
	}
}

func parseFieldNumber(field string) int {
	if num, err := strconv.Atoi(field); err == nil {
		return num
	}
	return -1
}
