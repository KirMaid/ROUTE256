package winter2025_3

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var t, n int

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	scanner := bufio.NewScanner(in)

	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, cap(buf))

	scanner.Scan()
	t, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		inputLine := scanner.Text()
		inputNumbers := strings.Fields(inputLine)

		scanner.Scan()
		outputLine := scanner.Text()

		nums := make([]int64, n)
		for j := 0; j < n; j++ {
			num, _ := strconv.ParseInt(inputNumbers[j], 10, 64)
			nums[j] = num
		}

		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})

		var sortedLine strings.Builder
		for j, num := range nums {
			if j > 0 {
				sortedLine.WriteString(" ")
			}
			sortedLine.WriteString(strconv.FormatInt(num, 10))
		}

		if sortedLine.String() == outputLine {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}
