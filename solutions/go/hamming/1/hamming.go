package hamming

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands not of equal length")
	}
	h := 0
	for i, r := range a {
		fmt.Printf("%d  ->  %c\n", i, r)
		if r != rune(b[i]) {
			h++
		}
	}
	return h, nil
}
