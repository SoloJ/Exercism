package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
var out = FreqMap{}
var mutex = &sync.Mutex{}
var wg sync.WaitGroup

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
func ConcurrentFrequency(in []string) FreqMap {
	wg.Add(len(in))
	for _, i := range in {
		go NotFrequency(i)
	}
	wg.Wait()
	return out
}
func NotFrequency(s string) {
	defer wg.Done()
	for _, r := range s {
		mutex.Lock()
		out[r]++
		mutex.Unlock()
	}
}
