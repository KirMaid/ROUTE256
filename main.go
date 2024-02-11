package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Directory struct {
	Dir     string      `json:"dir"`
	Files   []string    `json:"files"`
	Folders []Directory `json:"folders"`
}

func countHackFiles(dir Directory) int {
	count := 0
	for _, file := range dir.Files {
		if len(file) > 4 && file[len(file)-4:] == ".hack" {
			count++
		}
	}
	for _, subdir := range dir.Folders {
		count += countHackFiles(subdir)
	}
	return count
}

func hackFiles() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var rootDir Directory
		decoder := json.NewDecoder(os.Stdin)
		err := decoder.Decode(&rootDir)
		if err != nil {
			fmt.Println("Error decoding input:", err)
			continue
		}

		infectedCount := countHackFiles(rootDir)
		fmt.Println(infectedCount)
	}
}

func canRepresentAsPairs(queue []byte) bool {
	// Определение правильных пар
	pairs := [][]byte{[]byte{'X', 'Y'}, []byte{'X', 'Z'}, []byte{'Y', 'Z'}}

	// Проверка каждой возможной пары в очереди
	for _, pair := range pairs {
		for i := 0; i <= len(queue)-2; i++ {
			if queue[i] == pair[0] && queue[i+1] == pair[1] {
				// Удаление найденной пары из очереди
				queue = append(queue[:i], queue[i+2:]...)
				i-- // Сброс индекса, так как строка укорачивается
			}
		}
	}

	// Если очередь пуста, то все пары были найдены
	return len(queue)%2 == 0
}

func threeQuery() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)

		var queue []byte
		queue, _ = reader.ReadBytes('\n')
		queue = queue[:len(queue)-1] // Убираем символ новой строки

		if canRepresentAsPairs(queue) {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
}

func longestSeasonal(prices []int) []int {
	length := len(prices)
	longest := make([]int, length/2)

	for k := 1; k <= length/2; k++ {
		seasonLength := 0
		for i := 0; i+2*k <= length; i++ {
			seasonIncrease := true
			seasonDecrease := true
			for j := 0; j < k; j++ {
				if prices[i+j] >= prices[i+j+k] {
					seasonIncrease = false
				}
			}
			for j := 0; j < k; j++ {
				if prices[i+k+j] <= prices[i+k+j+k] {
					seasonDecrease = false
				}
			}
			if seasonIncrease && seasonDecrease {
				currentSeasonLength := 2 * k
				if currentSeasonLength > seasonLength {
					seasonLength = currentSeasonLength
				}
			}
		}
		longest[k-1] = seasonLength
	}
	return longest
}

func kPila() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		prices := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &prices[j])
		}

		result := longestSeasonal(prices)
		for j := 0; j < len(result); j++ {
			fmt.Fprint(writer, result[j], " ")
		}
		fmt.Fprintln(writer)
	}
}
