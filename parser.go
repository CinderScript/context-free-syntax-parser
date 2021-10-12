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

	tokensLexemePairs, err := ScanFileTokens(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	// MAP REPRESENTING THE GRAMMAR RULE DEFINITIONS
	grammarDefs := make(map[GrammarSymbols][]GrammarSymbols)
	grammarDefs[STMT_LIST] = []GrammarSymbols{STMT}

	_ = grammarDefs
	_ = tokensLexemePairs

	//WriteLine(grammarDefs[STMT_LIST])

	for _, t := range tokensLexemePairs {
		if t.token == ID || t.token == NUM {
			fmt.Println(t.token, t.lexeme)
		} else {
			fmt.Println(t.token)
		}
	}
}

func WriteLine(s string) {
	fmt.Println(s)
}
