package main

import "fmt"

func main() {

	arr, arraySize, err := stdinReadArrayFromUserInput()

	if err != nil {
		fmt.Println("Failed to read array")
		return
	}

	fmt.Println("Original array", arr)

	sortedArray := BubbleSort(arr, arraySize)

	fmt.Println("Sorted array:", sortedArray)

}

func BubbleSort(arr []int, size int) []int {
	var j int

	for i := 0; i < size-1; i++ {
		for j = 0; j < size-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func swap(p_first, p_second *int) {
	temp := *p_first
	*p_first = *p_second
	*p_second = temp
}

func stdinReadArraySize() (arraySize int, err error) {
	fmt.Println("Enter the size of the array:")
	_, err = fmt.Scanf("%d", &arraySize)
	return arraySize, err
}

func stdinReadArrayElements(arraySize int) (arr []int, err error) {
	var buffer int

	fmt.Println("Enter array elements:")
	for i := 0; i < arraySize; i++ {

		_, err = fmt.Scanf("%d", &buffer)
		if err != nil {
			break
		}
		arr = append(arr, buffer)
	}
	return arr, err
}

func stdinReadArrayFromUserInput() (arr []int, N int, err error) {

	N, err = stdinReadArraySize()
	if err != nil {
		return arr, N, err
	}
	arr, err = stdinReadArrayElements(N)
	return arr, N, err
}
