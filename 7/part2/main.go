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
	file, err := os.Open("./everybody_codes_e2025_q07_p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	names := strings.Split(inputRaw, ",")
	nameIndexes := make([]int, len(names))
	for x := range names {
		nameIndexes[x] = x
	}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		rule := line[0:1]
		test := line[4:]
		var toRemove []int
		for i, index := range nameIndexes {
			if strings.Contains(names[index], rule) {
				splitName := strings.Split(names[index], rule)
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
			nameIndexes[i] = nameIndexes[len(nameIndexes)-1]
			nameIndexes = nameIndexes[:len(nameIndexes)-1]
		}
	}

	total := 0
	for _, index := range nameIndexes {
		total += (index + 1)
	}

	fmt.Printf("Output: %v\n", names)
	fmt.Printf("Index total: %d\n", total)
}
