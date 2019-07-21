// A tentative of rewriting https://github.com/keon/algorithms
// for learning purposes.

//Given a list lst and a number N, create a new list
//that contains each number of the list at most N times without reordering.
//
//For example if N = 2, and the input is [1,2,3,1,2,1,2,3], you take [1,2,3,1,2],
//drop the next [1,2] since this would lead to 1 and 2 being in the result 3 times, and then take 3,
//which leads to [1,2,3,1,2,3]

package arrays


func Delete_nth_naive(array []int, n int) (ans []int){
	for _, num := range array {
		if n_occurs := count_occurrences(num, ans) ; n_occurs < n {
			ans = append(ans, num)
		}
	}
	return ans
}


func Delete_nth (array []int, n int) (result []int){

	counter := make(map[int]int)

	for _, num := range array {
		if counter[num] <  n{
			result = append(result, num)
			counter[num] ++
		}
	}
	return result
}

func count_occurrences (element int, a []int) (occurrences int){
	for _, el := range a {
		if el == element {
			occurrences ++
		}
	}
	return occurrences
}

