package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./everybody_codes_e2025_q06_p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	combinations := 0
	mentors := 0
	novices := 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRaw := scanner.Text()
	for _, character := range inputRaw {
		if character == 'A' || character == 'a' {
			if character == 'A' {
				mentors += 1
			} else {
				novices += 1
				combinations += mentors
			}
		}
	}
	fmt.Printf("Output: %d", combinations)
}
