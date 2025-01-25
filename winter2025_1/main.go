package winter2025_1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	var t, minValueIndex int
	var s string
	var minFound bool
	defer out.Flush()

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &t)
	in.ReadString('\n')

	for i := 0; i < t; i++ {
		s, _ = in.ReadString('\n')
		s = s[:len(s)-1]

		if len(s) == 1 {
			fmt.Fprintln(out, 0)
			continue
		}

		minValueIndex = 0
		minFound = false
		for j := 1; j < len(s); j++ {
			if s[j] > s[minValueIndex] {
				var result strings.Builder
				result.WriteString(s[:minValueIndex])
				result.WriteString(s[minValueIndex+1:])
				fmt.Fprintln(out, result.String())
				minFound = true
				break
			}
			minValueIndex = j
		}

		if minFound == false {
			var result strings.Builder
			result.WriteString(s[:len(s)-1])
			fmt.Fprintln(out, result.String())
		}
	}
}
