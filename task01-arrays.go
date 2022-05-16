package homework

func average(input [15]float32) (result float32) {
	var s float32
	for i := 0; i < len(input); i++ {
		s += input[i]
	}
	return s / float32(len(input))
}
