package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func all[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q07_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	names := strings.Split(inputRaw, ",")
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		rule := line[0:1]
		test := line[4:]
		var toRemove []int
		for i, name := range names {
			if strings.Contains(name, rule) {
				splitName := strings.Split(name, rule)
				matchTests := make([]bool, len(splitName))
				for tests := range strings.SplitSeq(test, ",") {
					for x, split := range splitName {
						if x == 0 || len(split) == 0 {
							matchTests[x] = true
							continue
						}
						if matchTests[x] {
							continue
						}
						if string(split[0]) == tests {
							matchTests[x] = true
						}
					}
					if all(matchTests, func(b bool) bool { return b }) {
						break
					}
				}
				if !all(matchTests, func(b bool) bool { return b }) {
					toRemove = append([]int{i}, toRemove...)
				}
			}
		}
		for _, i := range toRemove {
			names[i] = names[len(names)-1]
			names = names[:len(names)-1]
		}
	}

	fmt.Printf("Output: %v\n", names)
}
