package basic

import "fmt"

func forloop() []int {
	test := []int{2, 3, 4, 5, 6, 7, 8, 22, 33, 45, 66}
	target := 40

	for i := range test {
		for j := range test {
			if i == j {
				continue
			}

			if test[i]+test[j] == target {
				if i > j {
					return []int{j, i}
				} else {
					return []int{i, j}
				}
			} else {
				continue
			}
		}
	}

	return []int{}

}

func hashmap() []int {
	test := []int{2, 3, 4, 5, 6, 7, 8, 22, 33, 45, 66}
	target := 40

	var hmap = make(map[int]int)

	for i := range test {
		diff := target - test[i]

		if value, ok := hmap[diff]; ok {
			if value > i {
				return []int{i, value}
			} else {
				return []int{value, i}
			}
		}

		hmap[test[i]] = i

	}

	return []int{}

}

func twopointers() []int {
	test := []int{2, 3, 4, 5, 6, 7, 8, 22, 33, 45, 66}
	target := 40

	left := 0
	right := len(test) - 1

	for left < right {
		sum := test[left] + test[right]
		if sum == target {
			return []int{left, right}
		}

		if sum > target {
			right = right - 1
		} else {
			left += 1
		}
	}

	return []int{}
}

func main() {

	fmt.Println(forloop())

	fmt.Println(hashmap())

	fmt.Println(twopointers())
}
