/*
 * Class:		CSC 3310 - Concepts in Programming Languages
 * Title:		GO: Lexical and Syntax Analyzer
 * Purpose:		The purpose of this assignment is to practice the following concepts:
 * 				Context Free Grammar / BNF
 *				Lexical Analasys (scanner)
 *				Syntax Analasys (parser)
 *
 * 				syntax-parser.go contains the logic required to perfom a top-down, recursive
 *				syntax analasys on a given set of Lexical tokens.
 *
 * Author:		Maynard, Greg
 * Date:		10/13/2022
 */

package main

import (
	"errors"
)

// SyntaxParser is a struct that has a method for performing syntax analasys.
type SyntaxParser struct {

	//contains parsed identifiers
	PointTable map[string][]string
	Operations []OperationStatement

	// used internally by ParseTokens
	_rules map[GrammarSymbol][]SymbolDefinition

	_nextID       string
	_nextIdValues []string

	_nextOperation     string
	_nextOperationArgs []string

	_expectedToken                GrammarSymbol
	_foundToken                   GrammarSymbol
	_ranOutOfParsedSymbols        bool
	_currentDefinitionSymbolCount int
}

type OperationStatement struct {
	name      string
	arguments []string
}

// ParseTokens takes in a slice of TokenLexemePairs and then performs a syntax analasys
// using the given context fre grammar rules.  Returns the SyntaxParser object that has
// data relating to the analasys, including the identifiers found and the option statements
func (s SyntaxParser) ParseTokens(codeTokens []TokenLexemePair, rules map[GrammarSymbol][]SymbolDefinition) (SyntaxParser, error) {
	s._rules = rules

	s.PointTable = make(map[string][]string)
	_, _, completedParser, err := ParseSymbols(codeTokens, START, s)
	return completedParser, err
}

func ParseSymbols(parsedSymbols []TokenLexemePair,
	parentSymbol GrammarSymbol, parser SyntaxParser) (TokenLexemePair, bool, SyntaxParser, error) {

	nullToken := TokenLexemePair{"", ""}
	startToken := TokenLexemePair{START, ""}

	if len(parsedSymbols) > 0 && parsedSymbols[0].token == START {
		if len(parsedSymbols) > 1 {
			return nullToken, true, parser, errors.New("Syntax error: found " + string(parsedSymbols[1].token) + " after program termination.")
		} else {
			return nullToken, true, parser, nil
		}
	}

	// keep track of expected tokens
	var currentCodeSymbol TokenLexemePair

	// get each possible definition for the passed symbol
	symbolDefinitions := parser._rules[parentSymbol]

	// for each definition
	for defIndex, definition := range symbolDefinitions {

		// search tokens for a match to the current definition:
		// for each symbol in the definition
		derivationFound := false
		symbolDefinitionLength := len(definition.Symbols)
		for defSymbolIndex, defSymbol := range definition.Symbols {

			// ran out of parsed symbols - find what the next token should have been
			if defSymbolIndex >= len(parsedSymbols) {
				// find next token in this definition
				parser._expectedToken = parser.FindFirstToken(defSymbol)
				parser._ranOutOfParsedSymbols = true
				break // not enough symbols in code for this definition, skip
			}
			parser._ranOutOfParsedSymbols = false

			// if the definition symbol doesn't match with the parsed code - try to drill down
			currentCodeSymbol = parsedSymbols[defSymbolIndex]

			// if currentCodeSymbol == START {
			// 	asdf := 0
			// 	_ = asdf
			// }

			if defSymbol != currentCodeSymbol.token {

				// if definition symbol is non-terminal, drill down / try to find a symbol match in definition
				_, isNonTerminal := parser._rules[defSymbol]
				if isNonTerminal {

					//parse code for current definition symbol
					matchedSymbol, complete, updatedParser, err := ParseSymbols(parsedSymbols[defSymbolIndex:], defSymbol, parser)
					parser = updatedParser

					if err != nil {
						return nullToken, true, parser, err // propagate caught error to top

					} else if complete {
						// END CASE - SUCCESS!!!
						return startToken, true, parser, nil

					} else if matchedSymbol.token != "" { // if match was found
						//Add matched symbol to the parse and

						// add symbols before match to list, then add the replacement, then add the rest back
						var updated []TokenLexemePair
						updated = append(updated, matchedSymbol)
						indexAfterReplacement := parser._currentDefinitionSymbolCount + 1
						updated = append(updated, parsedSymbols[indexAfterReplacement:]...)

						return ParseSymbols(updated, START, parser)

					} else if matchedSymbol.token == "" && !complete && err == nil { // if no derivation was found for definition
						break // try next definition
					}

				} else { // definition symbol is terminal and didn't match
					parser._expectedToken = defSymbol
					parser._foundToken = currentCodeSymbol.token
					break // code doens't match this definition, try next
				}

			}
			// capture ID values and Operations
			if currentCodeSymbol.token == ID {
				parser._nextID = currentCodeSymbol.lexeme
				parser._nextOperationArgs = append(parser._nextOperationArgs, currentCodeSymbol.lexeme)

				if parser._nextOperation == SQUARE.String() {
					if len(parser._nextOperationArgs) == 4 {
						operation := OperationStatement{SQUARE.String(), parser._nextOperationArgs}
						parser.Operations = append(parser.Operations, operation)
					}
				} else if parser._nextOperation == TRIANGLE.String() {
					if len(parser._nextOperationArgs) == 3 {
						operation := OperationStatement{TRIANGLE.String(), parser._nextOperationArgs}
						parser.Operations = append(parser.Operations, operation)
					}
				}

				parser._nextIdValues = nil // clear - prepair for next id values

			} else if currentCodeSymbol.token == NUM {
				parser._nextIdValues = append(parser._nextIdValues, currentCodeSymbol.lexeme)
				if len(parser._nextIdValues) == 2 {
					// collected all values - save
					parser.PointTable[parser._nextID] = parser._nextIdValues
				}

			} else if currentCodeSymbol.token == SQUARE {
				parser._nextOperation = SQUARE.String()
				parser._nextOperationArgs = nil

			} else if currentCodeSymbol.token == TRIANGLE {
				parser._nextOperation = TRIANGLE.String()
				parser._nextOperationArgs = nil

			}

			// will there be a new definition loop, false?
			// check if all definition symbols were matched
			moreSymbolsInDefinition := defSymbolIndex != len(definition.Symbols)-1
			if !moreSymbolsInDefinition {
				derivationFound = true
				parser._currentDefinitionSymbolCount = defSymbolIndex
				break
			}

		} // (for each symbol in the definition)

		if derivationFound {
			parentTokenPair := TokenLexemePair{parentSymbol, ""}
			replaceCount := symbolDefinitionLength
			updatedSymbols := []TokenLexemePair{parentTokenPair}
			updatedSymbols = append(updatedSymbols, parsedSymbols[replaceCount:]...)
			isComplete := (len(updatedSymbols) == 1) && (updatedSymbols[0].token == START)

			return parentTokenPair, isComplete, parser, nil

		} else if defIndex+1 == len(symbolDefinitions) { // if on the last definition AND no match found
			if parentSymbol == START { // if failed last definition is START, no more definitions to try -fail
				return nullToken, true, parser, errors.New("Found " + parser._foundToken.String() + ", expected: " + parser._expectedToken.String())
			} else {
				return nullToken, false, parser, nil
			}
		}
	} // for each definition

	msg := ""
	if parser._ranOutOfParsedSymbols {
		msg = "Expected " + parser._expectedToken.String() + " after " + parser._foundToken.String()
	} else {
		msg = "Found " + parser._foundToken.String() + ", expected: " + parser._expectedToken.String()
	}
	// COULD NOT PARSE ALL OF THE SYMBOLS - SYNTAX ERROR
	return nullToken, true, parser, errors.New(msg)
}

// Gets the first token from the given symbol's first definition. this is a helper function
// used by the parser
func (s SyntaxParser) FindFirstToken(definition GrammarSymbol) GrammarSymbol {
	thisDefinitionsSymbols, isNonTerminal := s._rules[definition]

	if isNonTerminal {
		return s.FindFirstToken(thisDefinitionsSymbols[0].Symbols[0])
	}

	return definition
}
