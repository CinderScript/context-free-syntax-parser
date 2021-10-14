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
 *				the grammar rules are defined in the grammar-components.go file (change the
*				BNF style rules here to parse as a different language)
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
	}

	// parse the list of tokens generated by the lexical scan
	var parser SyntaxParser
	parser, err = parser.ParseTokens(tokensLexemePairs, GetGrammarRules())
	if err != nil {
		fmt.Println("Syntax Error - " + err.Error())
		os.Exit(4)
	}

	// Format code from parser's data
	if os.Args[2] == "-s" {
		PrintScheme(parser, file.Name())
	} else if os.Args[2] == "-p" {
		PrintProlog(parser, file.Name())
	} else {
		fmt.Println("Must specify -s or -p in command arguments.")
	}

}

// prints out scheme version of code
func PrintScheme(parser SyntaxParser, filename string) {
	fmt.Println(";  processing input file", filename)
	fmt.Println(";  Lexical and Syntax analysis passed")
	fmt.Println(";  Generating Scheme Code...")
	fmt.Println()

	for _, operation := range parser.Operations {
		fmt.Print("(process " + operation.name + " ")
		for _, argument := range operation.arguments {

			// on error - no definition
			if !IsIdDefined(argument, parser.PointTable) {
				fmt.Println("... \nError: ID ", argument, " is undefined.")
				os.Exit(5)

				// print list of points
			} else {
				fmt.Print(" (make-point ", parser.PointTable[argument][0]+" "+parser.PointTable[argument][1]+") ")
			}
		}
		fmt.Println(")")
	}
}

// prints out Prolog version of code
func PrintProlog(parser SyntaxParser, filename string) {
	fmt.Println("/*  processing input file", filename)
	fmt.Println("    Lexical and Syntax analysis passed")
	fmt.Println("    Generating Prolog Code...            */")
	fmt.Println()

	for _, operation := range parser.Operations {
		fmt.Print("(query(" + operation.name + "(")
		for i, argument := range operation.arguments {

			// on error - no definition
			if !IsIdDefined(argument, parser.PointTable) {
				fmt.Println("... \nError: ID ", argument, " is undefined.")
				os.Exit(5)

				// print list of points
			} else {
				fmt.Print("(point2d(" + parser.PointTable[argument][0] + "," + parser.PointTable[argument][1] + ")")
				if i < len(operation.arguments)-1 {
					fmt.Print(", ")
				}
			}
		}

		fmt.Println(")")
	}
}

func IsIdDefined(name string, table map[string][]string) bool {
	_, exists := table[name]
	return exists
}
