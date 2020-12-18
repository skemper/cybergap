package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var dateRegexp = regexp.MustCompile(`^(\w+\s+\d{4})\.`)

var theMap = make(map[string]int, 200)

func main() {
	fil, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening %s: %s", os.Args[1], err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		line := scanner.Text()
		dateText := dateRegexp.FindStringSubmatch(line)
		if dateText == nil {
			continue
		}
		month := dateText[1]
		// fmt.Println(month)

		if _, ok := theMap[month]; !ok {
			theMap[month] = 0
		}
		theMap[month]++

		// date := time.Parse("January 2006", dateText)
	}

	for k, v := range theMap {
		fmt.Printf("%s,%v\n", k, v)
	}
}
