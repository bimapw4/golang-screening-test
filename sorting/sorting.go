package main

import "fmt"

/*
	Run the program with go run sorting\sorting.go ,This command will automatic call function main.
	Comment function inside function main if not used
*/
func main() {
	/*
		The bar chart function is used to represent the length of the number in question.
		Command this function if you not used
	*/
	barChart([6]int{1, 4, 5, 6, 8, 2})

	/*
		The isertion sort used for sorting value in array,
		This function will automatic call function bar chart for showing basic chart.
		Command this function if you not used
	*/
	insertionSort([6]int{1, 4, 5, 6, 8, 2})
}

/*
	This function is a implementation for insertion sort,
	if I'm not mistaken isertion sort is a method that compares 2 values at once,
	namely the value of x with a value of x -1, that is,
	if the value of x-1 is greater than the value of x, the two values will swap places.
*/
func insertionSort(sort [6]int) {
	for i := 1; i < len(sort); i++ {

		j := i

		for j > 0 {
			if sort[j-1] > sort[j] {
				sort[j-1], sort[j] = sort[j], sort[j-1]
			}
			j = j - 1
			fmt.Println("")
			barChart(sort)
		}
	}

}

/*
	This function is a function to represent the length of a number.
	I take the difference of each number to the maximum value of the array used.
	then I loop until the value of the result of the subtraction is greater than or equal to 0
*/
func barChart(sort [6]int) {
	minus := [6]int{}
	_, max := findMinAndMax(sort)
	for i, v := range sort {
		minus[i] = v - max
	}

	countMoreZero := countMoreThanZero(minus)
	// fmt.Print(minus)
	for countMoreZero <= 6 {
		// for Max9 == 9 {
		for i, v := range minus {
			p := v + 1
			if v >= 0 {
				fmt.Print(" | ")
			} else {
				fmt.Print("   ")
			}
			minus[i] = p
			// fmt.Print(p)
		}
		fmt.Println(" ")
		countMoreZero = countMoreThanZero(minus)
		// Max9 = countMoreThanZero(minus)
		if countMoreZero == 7 {
			for _, v := range sort {
				fmt.Printf(" %v ", v)
			}
		}
	}
}

/*
	function to find the number of values greater than 0
*/
func countMoreThanZero(v [6]int) (out int) {
	out = 1
	for _, v := range v {
		if v > 0 {
			out = out + 1
		}
	}

	return
}

/*
	function for search maximum and minimun value in array
*/
func findMinAndMax(a [6]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
