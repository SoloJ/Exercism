package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
func ConcurrentFrequency(in []string) FreqMap {
	c := make(chan FreqMap)
	m := FreqMap{}
	for _, s := range in {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}
	for i := 0; i < len(in); i++ {
		temp := <-c
		for k, val := range temp {
			m[k] += val
		}
	}
	return m
}
