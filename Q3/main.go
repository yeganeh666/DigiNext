// Minimum Number of Swaps Required to Sort an Array
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() []int32 {
	fmt.Println("Enter numbers (separate with spaces): ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Trim(input, "\n")
	strArr := strings.Split(input, " ")
	arr := make([]int32, len(strArr))

	for i, number := range strArr {
		val, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		arr[i] = int32(val)

	}
	return arr
}

type Arr struct {
	value int32
	index int
}

type ByValue []Arr

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

func minimumSwaps(input []int32) int32 {
	arr := make([]Arr, 0, len(input))

	for i, v := range input {
		arr = append(arr, Arr{v, i})
	}

	sort.Sort(ByValue(arr))

	idx := make([]int, 0, len(input))
	for _, ar := range arr {
		idx = append(idx, ar.index)
	}

	var result int32
	for i := 0; i < len(input); i++ {
		if i == idx[i] {
			continue
		}

		input[i], input[idx[i]] = input[idx[i]], input[i]
		idx[i], idx[input[idx[i]]-1] = i, idx[i]
		result++
	}
	return result
}

func main() {
	arr := getInput()
	fmt.Println("Minimum number of 	swap to sort: ", minimumSwaps(arr))
}
