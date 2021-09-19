package exercises

/*Q5. (1) Average
1. Write code to calculate the average of a float64 slice. In a later exercise (Q6 you
will make it into a function
*/

func Average(input []int) float64 {
	var avg float64 = 0.0
	sum := 0
	total := len(input)
	for _, v := range input {
		sum = sum + v
		avg = float64(sum / total)
	}
	return avg
}
