// A tentative of rewriting https://github.com/TheAlgorithms/C
// for learning purposes.

package main

import "fmt"

func main() {

	arraySize, err := stdinReadArraysize()

	if err != nil {
		fmt.Println("Error while reading the size of the array")
		return
	}

	arr, err := arrayElements(arraySize)

	if err != nil {
		fmt.Println("Error while reading elements of the array")
		return
	}

	fmt.Println("Original array:", arr)

	arr = QuickSort(arr, 0, arraySize-1)

	fmt.Println("Sorted array:", arr)

}

func stdinReadArraysize() (arraySize int, err error) {
	fmt.Printf("Enter the size of array:\n")
	_, err = fmt.Scanf("%d", &arraySize)
	return arraySize, err
}

func arrayElements(size int)(array []int, err error ) {
	var temp int

	fmt.Printf("Enter the elements of the array\n")

	for i := 0; i < size; i++ {
		_, err :=fmt.Scanf("%d", &temp)

		if err != nil {
			break
		}

		array = append(array, temp)
	}
	return array, err
}

func QuickSort(arr []int, startIndex int, endIndex int) []int {
	if endIndex > startIndex {
		partitionIndex := partition(arr, startIndex, endIndex)
		arr = QuickSort(arr, startIndex, partitionIndex-1)
		arr = QuickSort(arr, partitionIndex+1, endIndex)
	}
	return arr
}

func partition(arr []int, lower int, upper int) int {
	i := lower - 1
	pivot := arr[upper]
	for j := lower; j < upper; j++ {
		if arr[j] <= pivot {
			i++
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[i+1], &arr[upper])
	return i + 1
}

func swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
}
