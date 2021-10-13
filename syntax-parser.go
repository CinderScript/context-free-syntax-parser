package main

import (
	"errors"
	"fmt"
)

type SyntaxParser struct {
	rules                 map[GrammarSymbol][]SymbolDefinition
	idTable               map[string]string
	expectedToken         GrammarSymbol
	foundToken            GrammarSymbol
	ranOutOfParsedSymbols bool
}

func (s SyntaxParser) ParseTokens(codeTokens []TokenLexemePair, rules map[GrammarSymbol][]SymbolDefinition) {
	var codeSymbols []GrammarSymbol
	for _, symbol := range codeTokens {
		codeSymbols = append(codeSymbols, symbol.token)
	}

	s.rules = rules

	_, _, err := s.ParseSymbols(codeSymbols, START)
	if err != nil {
		fmt.Println(string(err.Error()))
	}
}

func (s SyntaxParser) ParseSymbols(parsedSymbols []GrammarSymbol,
	currentSymbol GrammarSymbol) ([]GrammarSymbol, bool, error) {

	// keep track of expected tokens
	var currentCodeSymbol GrammarSymbol

	// get each possible definition for the passed symbol
	symbolDefinitions := s.rules[currentSymbol]

	// for each definition
	for _, definition := range symbolDefinitions {

		// search tokens for a match to the current definition:
		// for each symbol in the definition
		derivationFound := false
		symbolDefinitionLength := len(definition.Symbols)
		for i, defSymbol := range definition.Symbols {

			// ran out of parsed symbols - find what the next token should have been
			if i >= len(parsedSymbols) {
				// find next token in this definition
				//expectedToken = FindFirstToken(defSymbol, rules)
				s.ranOutOfParsedSymbols = true
				break // not enough symbols in code for this definition, skip
			}
			s.ranOutOfParsedSymbols = false

			// if the definition symbol doesn't match with the parsed code - try to drill down
			currentCodeSymbol = parsedSymbols[i]
			if defSymbol != currentCodeSymbol {

				// if definition symbol is non-terminal, drill down / try to parse code against derivation of symbol
				_, isNonTerminal := s.rules[defSymbol]
				if isNonTerminal {

					//parse code for current definition symbol
					updatedSymbols, complete, err := s.ParseSymbols(parsedSymbols[i:], defSymbol)

					if err != nil {
						return nil, true, err // propagate caught error to top

					} else if complete {
						// END CASE - SUCCESS!!!
						return []GrammarSymbol{START}, true, nil

					} else if updatedSymbols != nil { // if match was found
						//Parse the newley parsed symbols from start
						return s.ParseSymbols(updatedSymbols, START)
					}

				} else { // definition symbol is terminal and didn't match
					s.expectedToken = defSymbol
					s.foundToken = currentCodeSymbol
					break // code doens't match this definition, try next
				}
			}
			// will there be a new definition loop?
			// check if all definition symbols were matched
			if i == len(definition.Symbols)-1 {
				derivationFound = true
			}

		} // (for each symbol in the definition)

		if derivationFound {
			replaceCount := symbolDefinitionLength
			updatedSymbols := []GrammarSymbol{currentSymbol}
			updatedSymbols = append(updatedSymbols, parsedSymbols[replaceCount:]...)
			isComplete := (len(updatedSymbols) == 1) && (updatedSymbols[0] == START)

			return updatedSymbols, isComplete, nil
		}
	} // for each definition

	msg := ""
	if s.ranOutOfParsedSymbols {
		msg = "Expected: " + s.expectedToken.String()
	} else {
		msg = "Expected: " + s.expectedToken.String() + "\nFound: " + s.foundToken.String()
	}
	// COULD NOT PARSE ALL OF THE SYMBOLS - SYNTAX ERROR
	return nil, true, errors.New(msg)
}

// Gets the first token from the given symbol's first definition.
func FindFirstToken(definition GrammarSymbol, rules map[GrammarSymbol][]SymbolDefinition) GrammarSymbol {
	thisDefinitionsSymbols, isNonTerminal := rules[definition]

	if isNonTerminal {
		return FindFirstToken(thisDefinitionsSymbols[0].Symbols[0], rules)
	}

	return definition
}
