package letter

import (
	"log"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	log.Printf("Frequency: %v...", string([]rune(text)[:20]))
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) FreqMap {
	resultChannel := make(chan FreqMap)
	textLen := len(texts)
	log.Printf("Starting %d goroutines", textLen)
	for _, t := range texts {
		go func() {
			resultChannel <- Frequency(t)
		}()
	}
	result := FreqMap{}
	for i := 0; i < textLen; i++ {
		res := <-resultChannel
		//log.Printf("Counting the result: %v", res)
		for r, count := range res {
			result[r] += count
		}
	}

	log.Printf("Returning result: %v", result)
	return result
}
