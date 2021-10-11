package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Not enough cmd line arguments!\nExiting.")
		os.Exit(1)
	}
	file, error := os.Open(os.Args[1]) // open file from cmd args
	if error != nil {
		fmt.Println("Could not open the file: ", os.Args[1])
		fmt.Println("\"", error.Error(), "\"")
		fmt.Println("Exiting.")
		os.Exit(2)
	}

	regexId, _ := regexp.Compile(`^[a-zA-Z]+`)
	regexNum, _ := regexp.Compile(`^\d+`)

	// gets each token separated by spaces.

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		matchEndPosition := 0

		for len(line) != 0 {
			firstChar := line[0]
			isSingleChar := true
			isKeyword := true
			isIdOrNum := true

			// CHECK  FOR A SPECIAL CHAR
			switch firstChar {
			case ';':
				fmt.Println("fount a ;")
			case ',':
				fmt.Println("fount a ,")
			case '.':
				fmt.Println("fount a .")
			case '(':
				fmt.Println("fount a (")
			case ')':
				fmt.Println("fount a )")
			case '=':
				fmt.Println("fount a =")
			default:
				isSingleChar = false
			}

			if !isSingleChar {

				// CHECK FOR A SPECIAL KEYWORD

				switch 0 {

				case strings.Index(line, "point"):
					fmt.Println("Found point")
					matchEndPosition = 5 // point has 5 letters

				case strings.Index(line, "triangle"):
					fmt.Println("Found triangle")
					matchEndPosition = 8

				case strings.Index(line, "square"):
					fmt.Println("Found square")
					matchEndPosition = 6

				case strings.Index(line, "test"):
					fmt.Println("Found test")
					matchEndPosition = 4

				default:
					isKeyword = false
				}

				if !isKeyword {

					// CHECK FOR A USER MADE ID OR NUMBER

					if regexId.MatchString(line) {
						matchLocation := regexId.FindStringIndex(line)
						matchEndPosition = matchLocation[1] // finds last position of the match
						fmt.Println("Found user made ID: ", line[:matchEndPosition])

					} else if regexNum.MatchString(line) {
						matchLocation := regexNum.FindStringIndex(line)
						matchEndPosition = matchLocation[1] // finds last position of the match
						fmt.Println("Found user made NUMBER: ", line[:matchEndPosition])

					} else {
						isIdOrNum = false
					}

				}

			} else {
				isKeyword = false
				isIdOrNum = false
				matchEndPosition = 1 // single char read
			}

			// LEXICAL ERROR IF NO MATCH
			if !isSingleChar && !isKeyword && !isIdOrNum {
				fmt.Println("Lexical Error", line, "not recognized")
				os.Exit(1)
			}

			// ADVANCE LINE FOR NEXT FOR LOOP
			line = line[matchEndPosition:]
		}

	}

	fmt.Println()
	fmt.Println("Done.")

}
