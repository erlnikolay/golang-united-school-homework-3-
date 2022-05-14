package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	for _, value := range input {
		sum = sum + value
	}
	return sum / float32(len(input))
}

//func main() {
//	x := [15]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
//	fmt.Println(average(x))
//}
