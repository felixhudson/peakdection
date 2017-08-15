package peakdection

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func Findpeaks(data []int) ([]int, error) {
	window := 5
	// find all peaks in an array of data
	result := make([]int,0)
	past:= make([]int,5)

	// calculate a running ave
	smoothed := make([]float64,0)
	for k, _ := range past {
		past[k] = data[0]
	}
	sum :=0
	var value float64
	for _, v := range data {
		sum = 0
		for i := 1; i < window; i++ {
			sum = sum + past[i]
		}
		value = float64(sum + v) / float64(window)
		smoothed = append(smoothed, value)
	}
	// find where the data is higher than the ave!
	for k, v := range data {
		if float64(v) > smoothed[k] {
			result = append(result,k)
		}
	}
	return result, nil
}
