/*
 * Class:		CSC 3100 - Concepts in Programming Languages
 * Title:		GO: Lexical and Syntax Analyzer
 * Purpose:		The purpose of this assignment is to practice the following concepts:
 * 				Context Free Grammar / BNF
 *				Lexical Analasys (scanner)
 *				Syntax Analasys (parser)
 *
 *				Files: parser.go, syntax-parser.go, lexical-scanner.go, grammar-components.go
 *
 * 				parser.go contains the main entry point to this program.  This program
 *				reads text from a file, then scans and parses the input for errors using
 *				the grammar rules are defined in the grammar-components.go file.
 *
 * Author:		Maynard, Greg
 * Date:		10/13/2022
 */

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
	} else {
		fmt.Println("No lexical error detected.")
	}

	var parser SyntaxParser
	parser, err = parser.ParseTokens(tokensLexemePairs, GetGrammarRules())
	if err != nil {
		fmt.Println("Syntax Error - " + err.Error())
		os.Exit(4)
	} else {
		fmt.Println("No syntax error detected.")

		fmt.Println("points found:")

		for k, v := range parser.idTable {
			fmt.Println(k, "= (", v[0], ", ", v[1], ")")
		}
	}

}
