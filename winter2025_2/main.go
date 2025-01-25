package winter2025_2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	actions := make(map[byte][]byte, 4)
	actions['M'] = []byte{'C', 'R', 'D'}
	actions['R'] = []byte{'C'}
	actions['C'] = []byte{'M'}
	actions['D'] = []byte{'M'}
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var s string
	fmt.Fscan(in, &t)
	in.ReadString('\n')

	for i := 0; i < t; i++ {
		s, _ = in.ReadString('\n')
		s = s[:len(s)-1]

		if s[0] != 'M' || s[len(s)-1] != 'D' {
			fmt.Fprintln(out, "NO")
			continue
		}

		predElem := s[0]
		valid := true
		for j := 1; j < len(s); j++ {
			if slices.Contains(actions[predElem], s[j]) == false {
				fmt.Fprintln(out, "NO")
				valid = false
				break
			}
			predElem = s[j]
		}
		if valid {
			fmt.Fprintln(out, "YES")
		}
	}
}
