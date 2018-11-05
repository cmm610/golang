package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(ss []string) FreqMap {
	var c = make(chan FreqMap)
	var totalMap = make(FreqMap)
	for _, s := range ss {
		go func() {
			c <- Frequency(s)
		}()
		for k,v := range <- c {
			totalMap[k] += v
		}
	}
	return totalMap
}