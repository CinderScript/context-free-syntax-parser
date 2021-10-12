package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

// ScanFileTokens generates a slice of Token Lexeme pairs found
// in the given file. If a lexical error is found, this application
// panics with the message "Lexical error [symbol] not recognised"
func ScanFileTokens(file *os.File) ([]TokenLexemePair, error) {
	var tokens []TokenLexemePair // list of each token/lexeme pair found in the file

	regexId, _ := regexp.Compile(`^[a-zA-Z]+`)
	regexNum, _ := regexp.Compile(`^\d+`)

	// gets each token separated by spaces.

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() { // for each word
		line := scanner.Text()

		matchEndPosition := 0

		for len(line) != 0 { // for each lexeme in line
			firstChar := line[0]
			isSingleChar := true
			isKeyword := true
			isIdOrNum := true

			// CHECK  FOR A SPECIAL CHAR
			switch firstChar {
			case ';':
				t := TokenLexemePair{SEMICOLON, string(firstChar)}
				tokens = append(tokens, t)

			case ',':
				t := TokenLexemePair{COMMA, string(firstChar)}
				tokens = append(tokens, t)

			case '.':
				t := TokenLexemePair{PERIOD, string(firstChar)}
				tokens = append(tokens, t)

			case '(':
				t := TokenLexemePair{LPAREN, string(firstChar)}
				tokens = append(tokens, t)

			case ')':
				t := TokenLexemePair{RPAREN, string(firstChar)}
				tokens = append(tokens, t)

			case '=':
				t := TokenLexemePair{ASSIGN, string(firstChar)}
				tokens = append(tokens, t)

			default:
				isSingleChar = false
			}

			if !isSingleChar {

				// CHECK FOR A SPECIAL KEYWORD

				switch 0 {

				case strings.Index(line, "point"):
					matchEndPosition = 5 // point has 5 letters
					t := TokenLexemePair{POINT, line[:matchEndPosition]}
					tokens = append(tokens, t)

				case strings.Index(line, "triangle"):
					matchEndPosition = 8
					t := TokenLexemePair{TRIANGLE, line[:matchEndPosition]}
					tokens = append(tokens, t)

				case strings.Index(line, "square"):
					matchEndPosition = 6
					t := TokenLexemePair{SQUARE, line[:matchEndPosition]}
					tokens = append(tokens, t)

				case strings.Index(line, "test"):
					matchEndPosition = 4
					t := TokenLexemePair{TEST, line[:matchEndPosition]}
					tokens = append(tokens, t)

				default:
					isKeyword = false
				}

				if !isKeyword {

					// CHECK FOR A USER MADE ID OR NUMBER

					if regexId.MatchString(line) {
						matchLocation := regexId.FindStringIndex(line)
						matchEndPosition = matchLocation[1] // finds last position of the match

						t := TokenLexemePair{ID, line[:matchEndPosition]}
						tokens = append(tokens, t)

					} else if regexNum.MatchString(line) {
						matchLocation := regexNum.FindStringIndex(line)
						matchEndPosition = matchLocation[1] // finds last position of the match

						t := TokenLexemePair{NUM, line[:matchEndPosition]}
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
				return tokens, errors.New("Lexical Error " + line + " not recognized")
			}

			// ADVANCE LINE FOR NEXT FOR LOOP
			line = line[matchEndPosition:]

		} // for each lexeme in word

	} // for each word in file

	return tokens, nil
}
