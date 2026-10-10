package hamming

import "fmt"

// Distance accepts two strands and returns their hamming distance.
func Distance(a, b string) (distance int, e error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Strands must be of equal length")
	}

	for i := range len(a) {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
