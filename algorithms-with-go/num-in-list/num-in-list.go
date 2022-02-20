package numInList

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 3, 2, 3, 3, 6}
	fmt.Println(numInList(list, 5))
}

func numInList(list []int, num int) bool {
	for _, value := range list {
		if value == num {
			return true
		}
	}
	return false
}
