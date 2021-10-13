package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		panic("Not enough cmd line arguments!\nExiting.")
	}
	file, error := os.Open(os.Args[1]) // open file from cmd args
	if error != nil {
		msg := "Could not open the file: " + os.Args[1] +
			"\n" + "\"" + error.Error() + "\""
		panic(msg)
	}

	// Lexical scan of the opened file.  Saves each lexeme found and it's Token
	tokensLexemePairs, err := ScanFileTokens(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	var parser SyntaxParser
	parser.ParseTokens(tokensLexemePairs, GetGrammarRules())
}
