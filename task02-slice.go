package homework

//import "fmt"

func reverse(input []int64) (result []int64) {
	//Place your code here
	resSlice := make([]int64, len(input), cap(input))
	for iter := len(input) - 1; iter >= 0; iter-- {
		//fmt.Println(input[iter])
		resSlice[(len(input)-1)-iter] = input[iter]
	}
	return resSlice
}

//func main() {
//	x := []int64{1, 2, 3, 4, 5, 65}
//	fmt.Println(reverse(x))
//}
