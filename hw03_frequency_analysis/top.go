package hw03frequencyanalysis

import (
	"container/heap"
	"strings"
)

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority != pq[j].priority {
		return pq[i].priority > pq[j].priority
	}

	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Top10(s string) []string {
	items := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		val, ok := items[word]
		if !ok {
			items[word] = 1
		} else if word != "" {
			items[word] = val + 1
		}
	}

	heapq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		heapq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&heapq)
	j := min(10, heapq.Len())
	result := make([]string, 0, j)
	for j > 0 {
		item := heap.Pop(&heapq).(*Item)
		result = append(result, item.value)
		j--
	}

	return result
}
