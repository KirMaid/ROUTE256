package winter2025_6

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for testCase := 0; testCase < t; testCase++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		levers := make([][2]int, m)
		for i := 0; i < m; i++ {
			var a, b int
			fmt.Fscan(in, &a, &b)
			levers[i] = [2]int{a - 1, b - 1}
		}

		maxMask := 1 << n
		dp := make([]int, maxMask)
		prev := make([]int, maxMask)
		leverUsed := make([]int, maxMask)

		for mask := 0; mask < maxMask; mask++ {
			dp[mask] = -1
		}
		dp[0] = 0

		for mask := 0; mask < maxMask; mask++ {
			if dp[mask] == -1 {
				continue
			}
			for i, lever := range levers {
				a, b := lever[0], lever[1]
				newMask := mask ^ (1 << a) ^ (1 << b)
				newCount := countBits(newMask)
				if newCount > dp[newMask] {
					dp[newMask] = newCount
					prev[newMask] = mask
					leverUsed[newMask] = i + 1
				}
			}
		}

		maxCount := 0
		bestMask := 0
		for mask := 0; mask < maxMask; mask++ {
			if dp[mask] > maxCount {
				maxCount = dp[mask]
				bestMask = mask
			}
		}

		usedLevers := make([]int, 0)
		currentMask := bestMask
		for currentMask != 0 {
			lever := leverUsed[currentMask]
			usedLevers = append(usedLevers, lever)
			currentMask = prev[currentMask]
		}

		for i, j := 0, len(usedLevers)-1; i < j; i, j = i+1, j-1 {
			usedLevers[i], usedLevers[j] = usedLevers[j], usedLevers[i]
		}

		fmt.Fprintln(out, maxCount)
		fmt.Fprintln(out, len(usedLevers))
		for _, lever := range usedLevers {
			fmt.Fprint(out, lever, " ")
		}
		fmt.Fprintln(out)
	}
}

func countBits(mask int) int {
	count := 0
	for mask > 0 {
		count += mask & 1
		mask >>= 1
	}
	return count
}
