package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("./everybody_codes_e2025_q06_p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	combinations := 0
	mentors := make(map[rune]int)
	novices := make(map[rune]int)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	for _, character := range inputRaw {
		if unicode.IsUpper(character) {
			if _, exists := mentors[character]; !exists {
				mentors[character] = 0
			}
			mentors[character] += 1
		} else {
			if _, exists := novices[character]; !exists {
				novices[character] = 0
			}
			novices[character] += 1
			combinations += mentors[unicode.ToUpper(character)]
		}
	}
	fmt.Printf("Output: %d\n", combinations)
}
