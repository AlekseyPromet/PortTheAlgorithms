package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDelete_nth_naive(t *testing.T) {
	with_repetitions := []int{1, 1, 2, 2,  3, 3, 3, 2, 3}

	expected_result := []int{1, 2, 3}
	result := Delete_nth_naive(with_repetitions, 1)

	if !reflect.DeepEqual(expected_result, result) {
		fmt.Println(result)
		t.Error()
	}
}

func TestDelete_nth(t *testing.T) {
	with_repetitions := []int{1, 1, 2, 2,  3, 3, 3, 2, 3}

	expected_result := []int{1, 2, 3}
	result := Delete_nth(with_repetitions, 1)

	if !reflect.DeepEqual(expected_result, result) {
		fmt.Println(result)
		t.Error()
	}
}

