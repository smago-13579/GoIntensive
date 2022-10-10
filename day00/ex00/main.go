package main

import (
	"fmt"
	"io"
	"math"
	"sort"
)

func main() {
	var mean, median, sd float32
	values := make(map[int]int)
	var a, sum, mode, length int
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
	length = len(numbers)
	sort.Ints(numbers)

	if length%2 == 1 {
		l := length / 2
		median = float32(numbers[l])
	} else {
		l := length / 2
		median = float32(numbers[l]+numbers[l-1]) / 2
	}

	for _, val := range numbers {
		sum += val

		if v, ok := values[val]; ok {
			values[val] = v + 1
		} else {
			values[val] = 1
		}
	}
	mean = float32(sum) / float32(length)
	var res float32

	for _, v := range numbers {
		res += (mean - float32(v)) * (mean - float32(v))
	}
	sd = float32(math.Sqrt(float64(res) / float64(length)))
	i := 1000000

	for key, val := range values {
		if i == 1000000 || values[i] < val || (values[i] == val && key < i) {
			i = key
		}
	}
	mode = i

	fmt.Println("Mean:", mean)
	fmt.Println("Median:", median)
	fmt.Println("Mode:", mode)
	fmt.Printf("SD: %.2f \n", sd)
}
