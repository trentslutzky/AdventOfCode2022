package main

import (
  "fmt"
  "bufio"
  "strconv"
  "math"
  "container/heap"
  "os"
)

type calorieHeap []int

func (h calorieHeap) Len() int           { return len(h) }
func (h calorieHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h calorieHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *calorieHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *calorieHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

  data, _ := os.Open("./day1.txt")
  scanner := bufio.NewScanner(data)

  h := &calorieHeap{0, 0, 0}
  heap.Init(h)

  var max_cal = 0
  var running_cal = 0

  for scanner.Scan(){
    var line = scanner.Text()
    var line_int, _ = strconv.Atoi(line)
    if line == "" {
      max_cal = int(math.Max(float64(max_cal), float64(running_cal)))
      heap.Push(h, -running_cal)
      running_cal = 0
    } else {
      running_cal += line_int
    }
  }

  fmt.Println(*h)
}
