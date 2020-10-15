package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

type config struct {
	Mean          bool
	Median        bool
	Mode          bool
	SD            bool
	DecimalPlaces uint
}

func main() {
	flags := prepareConfig()
	nums, err := readNums(os.Stdin)
	if err != nil {
		println("read nums from stdin error:", err.Error())
		return
	}
	sort.Ints(nums)

	floatFormat := fmt.Sprintf("%%.%df", flags.DecimalPlaces)
	if flags.Mean {
		fmt.Println("Mean:", fmt.Sprintf(floatFormat, mean(nums)))
	}
	if flags.Median {
		fmt.Println("Medium:", fmt.Sprintf(floatFormat, median(nums)))
	}
	if flags.Mode {
		fmt.Printf("Mode: %d\n", mode(nums))
	}
	if flags.SD {
		fmt.Println("SD:", fmt.Sprintf(floatFormat, sd(nums)))
	}
}

func mode(nums []int) (mode int) {
	mode = -1
	currentTimes := 0
	maxTimes := 0
	currentValue := nums[0]
	for _, v := range nums {
		if v == currentValue {
			currentTimes++
		} else {
			if currentTimes > maxTimes || (currentTimes == maxTimes && currentValue < mode) {
				mode = currentValue
				maxTimes = currentTimes
			}
			currentValue = v
			currentTimes = 1
		}
	}
	return
}

func median(nums []int) float64 {
	length := len(nums)
	if length%2 != 0 {
		return float64(nums[length/2])
	}
	return mean(nums[length/2-1 : length/2+1])
}

func mean(nums []int) float64 {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}

func sd(nums []int) float64 {
	mean := mean(nums)
	sum := 0.0
	for _, v := range nums {
		sum += math.Pow(float64(v) - mean, 2)
	}
	return math.Sqrt(sum / float64(len(nums)))
}

func readNums(r io.Reader) (arr []int, err error) {
	arr = make([]int, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var num int
		if line := scanner.Text(); line != "" {
			if num, err = strconv.Atoi(line); err != nil {
				return
			}
			arr = append(arr, num)
		}
	}
	if err = scanner.Err(); err != nil {
		return
	}
	if len(arr) < 1 {
		err = fmt.Errorf("no numbers")
	}
	return
}

func prepareConfig() (flags config) {
	flag.UintVar(&flags.DecimalPlaces, "decimal-places", 2, "quantity(positive) of decimal places should results rounded to")
	flag.BoolVar(&flags.Mean, "mean", false, "count mean")
	flag.BoolVar(&flags.Median, "median", false, "count median")
	flag.BoolVar(&flags.Mode, "mode", false, "count mode")
	flag.BoolVar(&flags.SD, "sd", false, "count sd")

	flag.Parse()
	return
}
