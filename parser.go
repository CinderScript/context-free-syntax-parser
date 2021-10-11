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

	var tokens []TokenLexeme // list of each token/lexeme pair found in the file

	regexId, _ := regexp.Compile(`^[a-zA-Z]+`)
	regexNum, _ := regexp.Compile(`^\d+`)

	// gets each token separated by spaces.

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()

		matchEndPosition := 0

		for len(line) != 0 {
			firstChar := line[0]
			isSingleChar := true
			isKeyword := true
			isIdOrNum := true

			// CHECK  FOR A SPECIAL CHAR
			switch firstChar {
			case ';':
				t := TokenLexeme{Semicolon, string(firstChar)}
				tokens = append(tokens, t)

			case ',':
				t := TokenLexeme{Comma, string(firstChar)}
				tokens = append(tokens, t)

			case '.':
				t := TokenLexeme{Period, string(firstChar)}
				tokens = append(tokens, t)

			case '(':
				t := TokenLexeme{LParen, string(firstChar)}
				tokens = append(tokens, t)

			case ')':
				t := TokenLexeme{RParen, string(firstChar)}
				tokens = append(tokens, t)

			case '=':
				t := TokenLexeme{Assign, string(firstChar)}
				tokens = append(tokens, t)

			default:
				isSingleChar = false
			}

			if !isSingleChar {

				// CHECK FOR A SPECIAL KEYWORD

				switch 0 {

				case strings.Index(line, "point"):
					matchEndPosition = 5 // point has 5 letters
					t := TokenLexeme{Point, line[:matchEndPosition]}
					tokens = append(tokens, t)

				case strings.Index(line, "triangle"):
					matchEndPosition = 8
					t := TokenLexeme{Triangle, line[:matchEndPosition]}
					tokens = append(tokens, t)

				case strings.Index(line, "square"):
					matchEndPosition = 6
					t := TokenLexeme{Square, line[:matchEndPosition]}
					tokens = append(tokens, t)

				case strings.Index(line, "test"):
					matchEndPosition = 4
					t := TokenLexeme{Test, line[:matchEndPosition]}
					tokens = append(tokens, t)

				default:
					isKeyword = false
				}

				if !isKeyword {

					// CHECK FOR A USER MADE ID OR NUMBER

					if regexId.MatchString(line) {
						matchLocation := regexId.FindStringIndex(line)
						matchEndPosition = matchLocation[1] // finds last position of the match

						t := TokenLexeme{Id, line[:matchEndPosition]}
						tokens = append(tokens, t)

					} else if regexNum.MatchString(line) {
						matchLocation := regexNum.FindStringIndex(line)
						matchEndPosition = matchLocation[1] // finds last position of the match

						t := TokenLexeme{Num, line[:matchEndPosition]}
						tokens = append(tokens, t)
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

	for _, t := range tokens {
		if t.token == Id || t.token == Num {
			fmt.Println(TokenCatagoryToString(t.token), t.lexeme)
		} else {
			fmt.Println(TokenCatagoryToString(t.token))
		}
	}
}

type TokenCatagory int

const (
	Point TokenCatagory = iota
	Id
	Num
	Semicolon
	Comma
	Period
	LParen
	RParen
	Assign
	Triangle
	Square
	Test
)

func TokenCatagoryToString(n TokenCatagory) string {
	tokenNames := []string{
		"POINT", "ID", "NUM", "SEMICOLON", "COMMA",
		"PERIOD", "LPAREN", "RPAREN", "ASSIGN",
		"TRIANGLE", "SQUARE", "TEST"}
	return tokenNames[n]
}

type TokenLexeme struct {
	token  TokenCatagory
	lexeme string
}
