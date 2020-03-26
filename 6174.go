package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("A program to demonstrate Kaprekar's Constant (6174) ")
	fmt.Print("Enter any non-repeating 4 digit number:")
	var input int64
	fmt.Scanln(&input)
	var iterations int64 = 0
	compute(input, &iterations)

}

/* function that keeps recursively calling until Kaprekar's Constant is found */
func compute(input int64, iterations *int64) {
	//ascending
	var asc = Int2Slice(input)
	sort.Slice(asc, func(i, j int) bool { return asc[i] < asc[j] })
	//decending
	var dsc = Int2Slice(input)
	sort.Slice(dsc, func(i, j int) bool { return dsc[i] > dsc[j] })

	var ascInt = Slice2Int(asc)
	var dscInt = Slice2Int(dsc)

	if dscInt - ascInt == 6174 {
		/* if number found exit */
		fmt.Printf("Found Kaprekar's Constant after %v iterations \n", *iterations)
	} else {
		/* keep looping if not found */
		*iterations++
		if *iterations > 50000 {
			fmt.Printf("Safety Quit: Iterations %v \n", *iterations)
		} else {
			compute(dscInt - ascInt, iterations)
		}
	}
}

/* A function that converts int to slice */
func Int2Slice(input int64) []int64 {
	var digit = input % 10
	var num = input / 10
	var seq []int64
	for num > 0 {
		seq = append(seq, digit)
		digit = num % 10
		num = num / 10
	}
	seq = append(seq, digit)
	return seq
}

/* A function that converts slice to int */
func Slice2Int(slice []int64) int64 {
	var result int64 = 0
	for i := 0; i < len(slice); i++ {
		var sls = slice[i] * int64(math.Pow(10, float64(len(slice) - i - 1)))
		result = result + sls
	}
	return result
}
