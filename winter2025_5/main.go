package winter2025_5

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type Car struct {
	start    int
	end      int
	capacity int
	index    int
	load     int
}

type PriorityQueue []*Car

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].start == pq[j].start {
		return pq[i].index < pq[j].index
	}
	return pq[i].start < pq[j].start
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Car)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, m int
	fmt.Fscan(in, &t)

	for k := 0; k < t; k++ {
		fmt.Fscan(in, &n)
		arrival := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arrival[i])
		}

		fmt.Fscan(in, &m)
		cars := make([]Car, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &cars[j].start, &cars[j].end, &cars[j].capacity)
			cars[j].index = j + 1
			cars[j].load = 0
		}

		sort.Slice(cars, func(i, j int) bool {
			return cars[i].start < cars[j].start
		})

		orders := make([]int, n)
		for i := 0; i < n; i++ {
			orders[i] = i
		}
		sort.Slice(orders, func(i, j int) bool {
			return arrival[orders[i]] < arrival[orders[j]]
		})

		pq := make(PriorityQueue, 0)
		heap.Init(&pq)

		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = -1
		}

		carPtr := 0
		for _, orderIdx := range orders {
			arrivalTime := arrival[orderIdx]

			for carPtr < m && cars[carPtr].start <= arrivalTime {
				heap.Push(&pq, &cars[carPtr])
				carPtr++
			}

			for pq.Len() > 0 && pq[0].end < arrivalTime {
				heap.Pop(&pq)
			}

			if pq.Len() == 0 {
				res[orderIdx] = -1
				continue
			}

			car := heap.Pop(&pq).(*Car)
			if car.load < car.capacity {
				res[orderIdx] = car.index
				car.load++
				if car.load < car.capacity {
					heap.Push(&pq, car)
				}
			} else {
				res[orderIdx] = -1
			}
		}

		for i := 0; i < n; i++ {
			fmt.Fprint(out, res[i], " ")
		}
		fmt.Fprintln(out)
	}
}
