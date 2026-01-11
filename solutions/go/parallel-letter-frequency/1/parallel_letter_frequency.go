package letter

import (
	"log"
	"sync"
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
	var wg sync.WaitGroup
	textLen := len(texts)
	log.Printf("Starting %d goroutines", textLen)
	wg.Add(textLen)
	for _, t := range texts {
		go func() {
			defer wg.Done()
			resultChannel <- Frequency(t)
		}()
	}
	go func() {
		log.Println("Wait until all go routines are done THEN close the channel...")
		wg.Wait()
		log.Println("go routines are done. Now Closing the channel.")
		close(resultChannel)
	}()

	result := count(resultChannel)
	log.Printf("Returning result: %v", result)
	return result
}

func count(resultChannel <-chan FreqMap) FreqMap {
	result := FreqMap{}
	for res := range resultChannel {
		//log.Printf("Counting the result: %v", res)
		for r, count := range res {
			result[r] += count
		}
	}
	return result
}
