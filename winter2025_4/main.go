package winter2025_4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	in.ReadString('\n')

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadString('\n')

		var jsonBuilder strings.Builder
		for j := 0; j < n; j++ {
			line, _ := in.ReadString('\n')
			jsonBuilder.WriteString(line)
		}

		var folder Folder
		json.Unmarshal([]byte(jsonBuilder.String()), &folder)

		totalCountFiles := countHackedFiles(&folder, false)
		fmt.Fprintln(out, totalCountFiles)
	}
}

func countHackedFiles(folder *Folder, isParentInfected bool) int {
	total := 0
	isInfected := isParentInfected

	if !isInfected {
		for _, file := range folder.Files {
			if len(file) >= 5 && file[len(file)-5:] == ".hack" {
				isInfected = true
				break
			}
		}
	}

	if isInfected {
		total += len(folder.Files)
		for _, fldr := range folder.Folders {
			total += countAllFiles(&fldr)
		}
	} else {
		for _, fldr := range folder.Folders {
			total += countHackedFiles(&fldr, false)
		}
	}

	return total
}

func countAllFiles(folder *Folder) int {
	total := len(folder.Files)
	for _, fldr := range folder.Folders {
		total += countAllFiles(&fldr)
	}
	return total
}
