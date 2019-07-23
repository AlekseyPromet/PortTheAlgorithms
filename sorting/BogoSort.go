package main

import "fmt"
import "math/rand"

func main () {
	arr, err := stdinReadNumbersUnsorted()
	if err != nil {
		fmt.Println("Failed to read user input")
		return
	}

	fmt.Println("Original array: ", arr)
	sort(&arr, len(arr))
	fmt.Println("Sorted array: ", arr)

}

func sort(arr *[]int, size int) {

	for !checkSorted(arr, size) {
		shuffle(arr, size)
	}
}

func checkSorted(arr *[]int, size int) bool {
	for i := size-1 ; i >= 1; i-- {
		if (*arr)[i] < (*arr)[i-1] {
			return false
		}
	}
	return true
}

func shuffle(arr *[]int, size int) {
	var i, t, r int
	for i = 0; i < size; i++ {
		t = (*arr)[i]
		r = rand.Intn(size)
		(*arr)[i] = (*arr)[r]
		(*arr)[r] = t
	}
}

func stdinReadNumbersUnsorted() (arr []int, err error) {
	size := 6
	fmt.Printf("Enter %d numbers unsorted:", size)

	var buffer int

	for i := 0; i < size; i ++ {
		_, err = fmt.Scanf("%d", &buffer)
		if err != nil {
			break
		}
		arr = append(arr, buffer)
	}

	return
}
