// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func isSimilar(s1, s2 string) bool {
// 	if s1 == s2 {
// 		return true
// 	}
// 	if len(s1) != len(s2) {
// 		return false
// 	}
// 	diff := 0
// 	for i := 0; i < len(s1); i++ {
// 		if s1[i] != s2[i] {
// 			diff++
// 		}
// 	}
// 	return diff == 2
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()

// 	var t int
// 	fmt.Fscan(reader, &t)

// 	for i := 0; i < t; i++ {
// 		var n int
// 		fmt.Fscan(reader, &n)

// 		existingLogins := make(map[string]bool)
// 		for j := 0; j < n; j++ {
// 			var login string
// 			fmt.Fscan(reader, &login)
// 			existingLogins[login] = true
// 		}

// 		var m int
// 		fmt.Fscan(reader, &m)

// 		for j := 0; j < m; j++ {
// 			var desiredLogin string
// 			fmt.Fscan(reader, &desiredLogin)

// 			found := false
// 			for existingLogin := range existingLogins {
// 				if isSimilar(existingLogin, desiredLogin) {
// 					found = true
// 					break
// 				}
// 			}

// 			if found {
// 				fmt.Fprintln(writer, 1)
// 			} else {
// 				fmt.Fprintln(writer, 0)
// 			}
// 		}
// 	}
// }

package main

import (
	"fmt"
)

func areStringsSimilar(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	diffCount := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount++
			if diffCount > 2 {
				return false
			}
		}
	}

	return diffCount == 2
}

func main() {
	// Example usage:
	logins := []string{"hello", "ozoner", "roma", "anykey"}                            // Current employee logins
	newLogins := []string{"roma", "ello", "zooner", "ankyey", "ynakey", "amor", "rom"} // New employee desired logins

	for _, newLogin := range newLogins {
		found := false
		for _, login := range logins {
			if areStringsSimilar(login, newLogin) {
				found = true
				break
			}
		}
		fmt.Println(map[bool]string{true: "1", false: "0"}[found])
	}
}
