package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calcSpine(input []int) int64 {
	spine := []int{}
	left := make(map[int]int)
	right := make(map[int]int)

	spine = append(spine, input[0])

	for i := 1; i < len(input); i++ {
		toContinue := false
		for j, spineVal := range spine {
			if input[i] < spineVal {
				if _, contains := left[j]; !contains {
					left[j] = input[i]
					toContinue = true
					break
				}
			} else if input[i] > spineVal {
				if _, contains := right[j]; !contains {
					right[j] = input[i]
					toContinue = true
					break
				}
			}
		}
		if toContinue {
			continue
		}

		spine = append(spine, input[i])
	}

	spineString := ""
	for _, value := range spine {
		spineString += fmt.Sprintf("%d", value)
	}
	swordValue, _ := strconv.ParseInt(spineString, 10, 64)
	return int64(swordValue)
}

func main() {
	file, err := os.Open("./everybody_codes_e2025_q05_p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	spines := []int64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		swordText := scanner.Text()
		inputFishboneRaw := strings.Split(swordText, ":")[1]
		swordValuesRaw := strings.Split(inputFishboneRaw, ",")
		swordValues := []int{}
		for _, rawValue := range swordValuesRaw {
			value, _ := strconv.ParseInt(rawValue, 10, 32)
			swordValues = append(swordValues, int(value))
		}
		spines = append(spines, calcSpine(swordValues))
	}

	max := int64(0)
	min := spines[0]

	for _, spine := range spines {
		if spine > max {
			max = spine
		}
		if spine < min {
			min = spine
		}
	}

	fmt.Printf("Output: %d\n", max-min)
	fmt.Printf("Spines: %v\n", spines)
	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Min: %d\n", min)
}
