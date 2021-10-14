package main

import (
	"errors"
)

type SyntaxParser struct {
	rules                 map[GrammarSymbol][]SymbolDefinition
	idTable               map[string]string
	expectedToken         GrammarSymbol
	foundToken            GrammarSymbol
	ranOutOfParsedSymbols bool
	nextID                GrammarSymbol
}

func (s SyntaxParser) ParseTokens(codeTokens []TokenLexemePair, rules map[GrammarSymbol][]SymbolDefinition) error {
	var codeSymbols []GrammarSymbol
	for _, symbol := range codeTokens {
		codeSymbols = append(codeSymbols, symbol.token)
	}

	s.rules = rules

	_, _, _, err := ParseSymbols(codeSymbols, START, s)
	return err
}

func ParseSymbols(parsedSymbols []GrammarSymbol,
	parentSymbol GrammarSymbol, parser SyntaxParser) ([]GrammarSymbol, bool, SyntaxParser, error) {

	if len(parsedSymbols) > 0 && parsedSymbols[0] == START {
		if len(parsedSymbols) > 1 {
			return parsedSymbols, true, parser, errors.New("Syntax error: found " + string(parsedSymbols[1]) + " after program termination.")
		} else {
			return parsedSymbols, true, parser, nil
		}
	}

	// keep track of expected tokens
	var currentCodeSymbol GrammarSymbol

	// get each possible definition for the passed symbol
	symbolDefinitions := parser.rules[parentSymbol]

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
				parser.expectedToken = parser.FindFirstToken(defSymbol)
				parser.ranOutOfParsedSymbols = true
				break // not enough symbols in code for this definition, skip
			}
			parser.ranOutOfParsedSymbols = false

			// if the definition symbol doesn't match with the parsed code - try to drill down
			currentCodeSymbol = parsedSymbols[defSymbolIndex]

			// if currentCodeSymbol == START {
			// 	asdf := 0
			// 	_ = asdf
			// }

			if defSymbol != currentCodeSymbol {

				// if definition symbol is non-terminal, drill down / try to find a symbol match in definition
				_, isNonTerminal := parser.rules[defSymbol]
				if isNonTerminal {

					//parse code for current definition symbol
					updatedSymbols, complete, updatedParser, err := ParseSymbols(parsedSymbols[defSymbolIndex:], defSymbol, parser)
					parser = updatedParser

					if err != nil {
						return nil, true, parser, err // propagate caught error to top

					} else if complete {
						// END CASE - SUCCESS!!!
						return []GrammarSymbol{START}, true, parser, nil

					} else if updatedSymbols != nil { // if match was found
						//Parse the newley parsed symbols from start
						return ParseSymbols(updatedSymbols, START, parser)

					} else if updatedSymbols == nil && !complete && err == nil { // if no derivation was found for definition
						break // try next definition
					}

				} else { // definition symbol is terminal and didn't match
					parser.expectedToken = defSymbol
					parser.foundToken = currentCodeSymbol
					break // code doens't match this definition, try next
				}

			}
			// capture ID values
			// current symbol matches definition symbol
			if currentCodeSymbol == ID {
				parser.nextID = currentCodeSymbol
			}

			// will there be a new definition loop, false?
			// check if all definition symbols were matched
			moreSymbolsInDefinition := defSymbolIndex != len(definition.Symbols)-1
			if !moreSymbolsInDefinition {
				derivationFound = true
				break
			}

		} // (for each symbol in the definition)

		if derivationFound {
			replaceCount := symbolDefinitionLength
			updatedSymbols := []GrammarSymbol{parentSymbol}
			updatedSymbols = append(updatedSymbols, parsedSymbols[replaceCount:]...)
			isComplete := (len(updatedSymbols) == 1) && (updatedSymbols[0] == START)

			return updatedSymbols, isComplete, parser, nil

		} else if defIndex+1 == len(symbolDefinitions) { // if on the last definition AND no match found
			if parentSymbol == START { // if failed last definition is START, no more definitions to try -fail
				return nil, true, parser, errors.New("Found " + parser.foundToken.String() + ", expected: " + parser.expectedToken.String())
			} else {
				return nil, false, parser, nil
			}
		}
	} // for each definition

	msg := ""
	if parser.ranOutOfParsedSymbols {
		msg = "Expected " + parser.expectedToken.String() + " after " + parser.foundToken.String()
	} else {
		msg = "Found " + parser.foundToken.String() + ", expected: " + parser.expectedToken.String()
	}
	// COULD NOT PARSE ALL OF THE SYMBOLS - SYNTAX ERROR
	return nil, true, parser, errors.New(msg)
}

// Gets the first token from the given symbol's first definition.
func (s SyntaxParser) FindFirstToken(definition GrammarSymbol) GrammarSymbol {
	thisDefinitionsSymbols, isNonTerminal := s.rules[definition]

	if isNonTerminal {
		return s.FindFirstToken(thisDefinitionsSymbols[0].Symbols[0])
	}

	return definition
}
