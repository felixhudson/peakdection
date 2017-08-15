package peakdection

import (
	"fmt"
	"testing"
)

func Test_peak(t *testing.T) {
	data := []int{1, 5, 1, 5, 1, 5, 1, 5}
	result, err := Findpeaks(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) <= 0 {
		t.Fatal("no peak detected")
	}
	fmt.Printf("result = %+v\n", result)
}
