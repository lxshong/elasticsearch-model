package src

func Sum(params... int) int {
	var sum = 0
	for _, param := range params {
		sum += param
	}
	return sum
}
