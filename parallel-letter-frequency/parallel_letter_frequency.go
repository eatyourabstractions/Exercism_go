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

var results = make(chan FreqMap)

func ConcurrentFrequency(input_slice []string) FreqMap {
	var results = make(chan FreqMap)
	fm := FreqMap{}
	for _, text := range input_slice {
		go func(w string) { results <- Frequency(w) }(text)
	}

	for range input_slice {
		for letter, freq := range <-results {
			fm[letter] += freq
		}
	}

	return fm
}
