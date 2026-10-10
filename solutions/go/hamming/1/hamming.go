package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) == 0 && len(b) == 0 {
		return 0, nil
	}

	if len(a) == 0 {
		return 0, fmt.Errorf("First string cannot be empty")
	}

	if len(b) == 0 {
		return 0, fmt.Errorf("Second string cannot be empty")
	}

	if len(a) != len(b) {
		return 0, fmt.Errorf("Strings must have equal length")
	}

	hammingDistance := 0

	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
