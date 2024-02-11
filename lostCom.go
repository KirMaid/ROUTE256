package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func lostCommissions() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	lostCommissions := make([]float64, t)

	for i := 0; i < t; i++ {
		var n, p int
		fmt.Fscan(reader, &n, &p)

		var totalLostCommission float64

		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(reader, &a)
			correctCommission := float64(a) * (float64(p) / 100) // Учитываем округление в меньшую сторону до копеек
			totalLostCommission += correctCommission - math.Floor(correctCommission)
		}
		lostCommissions[i] = totalLostCommission
	}

	for _, commission := range lostCommissions {
		fmt.Fprintf(writer, "%.2f\n", commission)
	}
}

func main1() {
	lostCommissions()
	//highLoadSystem()
}
