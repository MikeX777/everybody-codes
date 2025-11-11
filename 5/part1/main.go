package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./everybody_codes_e2025_q05_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	inputFishboneRaw := strings.Split(inputRaw, ":")[1]
	swordValuesRaw := strings.Split(inputFishboneRaw, ",")
	swordValues := []int{}
	for _, rawValue := range swordValuesRaw {
		value, _ := strconv.ParseInt(rawValue, 10, 32)
		swordValues = append(swordValues, int(value))
	}

	spine := []int{}
	left := make(map[int]int)
	right := make(map[int]int)

	spine = append(spine, swordValues[0])

	for i := 1; i < len(swordValues); i++ {
		toContinue := false
		for j, spineVal := range spine {
			if swordValues[i] < spineVal {
				if _, contains := left[j]; !contains {
					left[j] = swordValues[i]
					toContinue = true
					break
				}
			} else if swordValues[i] > spineVal {
				if _, contains := right[j]; !contains {
					right[j] = swordValues[i]
					toContinue = true
					break
				}
			}
		}
		if toContinue {
			continue
		}

		spine = append(spine, swordValues[i])
	}

	// fmt.Printf("Output: ")

	for i, value := range spine {
		leftVal, _ := left[i]
		rightVal, _ := right[i]
		fmt.Printf("%d-%d-%d\n", leftVal, value, rightVal)
		fmt.Println("  |")
	}

	fmt.Println("")
	fmt.Printf("Spine: %v\n", spine)
	fmt.Printf("left: %v\n", left)
	fmt.Printf("right: %v\n", right)
	fmt.Printf("Output: ")
	for _, value := range spine {
		fmt.Printf("%d", value)
	}
}
