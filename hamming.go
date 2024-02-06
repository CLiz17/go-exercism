package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("Not the same length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
