package main

import (
	"fmt"
	"io"
	"math"
	"sort"
)

func main() {
	var mean, median, sd float32
	var a, mode int
	var numbers []int
	_, err := fmt.Scan(&a)

	for err == nil {
		if a < -100000 || a > 100000 {
			fmt.Println("Error number:", a)
			return
		}
		numbers = append(numbers, a)
		_, err = fmt.Scan(&a)
	}

	if err != nil && err != io.EOF {
		fmt.Println("Error number:", err)
		return
	}
	sort.Ints(numbers)

	median = findMedian(numbers)
	mean, sd, mode = findMeanSdMode(numbers)

	fmt.Println("Mean:", mean)
	fmt.Println("Median:", median)
	fmt.Println("Mode:", mode)
	fmt.Printf("SD: %.2f \n", sd)
}

func findMedian(numbers []int) float32 {
	length := len(numbers)

	if length%2 == 1 {
		l := length / 2
		return float32(numbers[l])
	} else {
		l := length / 2
		return float32(numbers[l]+numbers[l-1]) / 2
	}
}

func findMeanSdMode(numbers []int) (float32, float32, int) {
	var mean, sd float32
	var mode = 1000000
	values := make(map[int]int)
	length := len(numbers)
	sum := 0

	for _, val := range numbers {
		sum += val

		if v, ok := values[val]; ok {
			values[val] = v + 1
		} else {
			values[val] = 1
		}
	}
	mean = float32(sum) / float32(length)

	for key, val := range values {
		if mode == 1000000 || values[mode] < val || (values[mode] == val && key < mode) {
			mode = key
		}
	}

	for _, val := range numbers {
		sd += (mean - float32(val)) * (mean - float32(val))
	}
	sd = float32(math.Sqrt(float64(sd) / float64(length)))

	return mean, sd, mode
}
