package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func highLoad() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var process string
		fmt.Fscan(reader, &process)

		state := 0
		isValid := true

		for _, action := range strings.Split(process, "") {
			switch action {
			case "M":
				if state == 1 || state == 3 {
					isValid = false
					break
				}
				state = 1
			case "R":
				if state != 1 && state != 2 {
					isValid = false
					break
				}
				state = 3
			case "C":
				if state != 1 && state != 3 {
					isValid = false
					break
				}
				state = 2
			case "D":
				if state != 1 && state != 3 {
					isValid = false
					break
				}
				state = 4
			default:
				isValid = false
				break
			}
		}

		if !isValid || state != 4 {
			fmt.Fprintln(writer, "NO")
		} else {
			fmt.Fprintln(writer, "YES")
		}
	}
}
