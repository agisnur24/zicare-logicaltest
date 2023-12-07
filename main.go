package main

import "fmt"

// This function checks if a sequence exists in a larger slice of integers.
func sequenceExists(numbers []int, seq []int) bool {
	// Loop through each element in the larger slice.
	for i := range numbers {
		// Check if the remaining slice is too short to hold the sequence.
		if len(seq) > len(numbers)-i {
			return false
		}

		// Loop through each element in the sequence.
		for j := range seq {
			// Check if the corresponding elements are different.
			if numbers[i+j] != seq[j] {
				break
			}

			// Check if we reached the end of the sequence.
			if j == len(seq)-1 {
				return true
			}
		}
	}

	// No matching sequence found.
	return false
}

func main() {
	// Define the larger slice of integers.
	numbers := []int{20, 7, 8, 10, 2, 5, 6}

	// Define three different sequences to search for.
	seq1 := []int{7, 8}
	seq2 := []int{8, 7}
	seq3 := []int{7, 10}

	// Check if each sequence exists in the larger slice and print the result.
	fmt.Println(sequenceExists(numbers, seq1))
	fmt.Println(sequenceExists(numbers, seq2))
	fmt.Println(sequenceExists(numbers, seq3))
}
