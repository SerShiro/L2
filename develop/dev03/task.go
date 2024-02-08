package dev

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mySort() {
	filePath := flag.String("file", "", "Путь к файлу для сортировки")
	keyColumn := flag.Int("k", -1, "Номер колонки для сортировки (по умолчанию разделитель - пробел)")
	numericSort := flag.Bool("n", false, "Сортировка по числовому значению")
	reverseSort := flag.Bool("r", false, "Сортировка в обратном порядке")
	uniqueSort := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Необходимо указать путь к файлу")
		return
	}

	lines, err := readLines(*filePath)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	sortLines(lines, *keyColumn, *numericSort, *reverseSort, *uniqueSort)

	err = writeLines(*filePath, lines)
	if err != nil {
		fmt.Println("Ошибка при записи отсортированных строк в файл:", err)
		return
	}

	fmt.Println("Файл успешно отсортирован.")
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func sortLines(lines []string, keyColumn int, numericSort, reverseSort, uniqueSort bool) {
	sort.SliceStable(lines, func(i, j int) bool {
		var keyI, keyJ string

		columnsI := strings.Fields(lines[i])
		columnsJ := strings.Fields(lines[j])

		if keyColumn >= 0 && keyColumn < len(columnsI) {
			keyI = columnsI[keyColumn]
			keyJ = columnsJ[keyColumn]
		} else {
			keyI = lines[i]
			keyJ = lines[j]
		}

		if numericSort {
			numI, errI := strconv.Atoi(keyI)
			numJ, errJ := strconv.Atoi(keyJ)

			if errI == nil && errJ == nil {
				return numI < numJ
			}
		}

		if keyI < keyJ {
			return !reverseSort
		}
		if keyI > keyJ {
			return reverseSort
		}
		return false
	})

	if uniqueSort {
		j := 0
		for i := 1; i < len(lines); i++ {
			if lines[j] != lines[i] {
				j++
				lines[j] = lines[i]
			}
		}
		lines = lines[:j+1]
	}
}
