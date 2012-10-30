package main
import (
	"strings"
	"tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	work := strings.Fields(s)
	for _,v := range work {
		_, ok := result[v]
		if ok == false {
			result[v] = 1
		} else {
			result[v]++
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}