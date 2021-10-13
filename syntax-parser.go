package main

import (
	"errors"
	"fmt"
)

func ParseTokens(codeTokens []TokenLexemePair, rules map[GrammarSymbol][]SymbolDefinition) {
	var codeSymbols []GrammarSymbol
	for _, symbol := range codeTokens {
		codeSymbols = append(codeSymbols, symbol.token)
	}

	_, _, err := ParseSymbols(codeSymbols, rules, START)
	if err != nil {
		fmt.Println(string(err.Error()))
	}
}

func ParseSymbols(parsedSymbols []GrammarSymbol,
	rules map[GrammarSymbol][]SymbolDefinition,
	currentSymbol GrammarSymbol) ([]GrammarSymbol, bool, error) {

	fmt.Println("Parsing: ", currentSymbol)

	// get each possible definition for the passed symbol
	symbolDefinitions := rules[currentSymbol]

	// for each definition
	for _, definition := range symbolDefinitions {

		// search tokens for a match to the current definition:
		// for each symbol in the definition
		derivationFound := false
		symbolDefinitionLength := len(definition.Symbols)
		for i, defSymbol := range definition.Symbols {

			if i >= len(parsedSymbols) {
				break // not enough symbols in code for this definition, skip
			}

			// if the definition symbol doesn't match with the parsed code - try to drill down
			if defSymbol != parsedSymbols[i] {

				// if definition symbol is non-terminal, drill down / try to parse code against derivation of symbol
				_, isSymbolNonTerminal := rules[defSymbol]
				if isSymbolNonTerminal {
					//parse code for current definition symbol
					updatedSymbols, complete, err := ParseSymbols(parsedSymbols[i:], rules, defSymbol)

					if err != nil {
						return nil, true, err // propagate error to top

					} else if complete {
						// END CASE - SUCCESS!!!
						return []GrammarSymbol{START}, true, nil

					} else if updatedSymbols != nil { // if match was found
						//Parse the newley parsed symbols from start
						return ParseSymbols(updatedSymbols, rules, START)
					}

				} else { // definition symbol is terminal
					break // code doens't match this definition, try next
				}
			}
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

	// COULD NOT PARSE ALL OF THE SYMBOLS - SYNTAX ERROR
	return nil, true, errors.New("SYNTAX ERROR: " + string(currentSymbol))
}
