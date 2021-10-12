package main

func ParseSyntax(tokenLexemPair []TokenLexemePair,
	rules map[GrammarSymbol][]SymbolDefinition,
	currentSymbol GrammarSymbol) {

	symbolDefinitions := rules[currentSymbol]

	// for each definition of the current symbol
	for _, definition := range symbolDefinitions {

		// search tokens for a match to the current definition:
		// for each symbol in the definition
		for _, symbol := range definition.Symbols {
			// if symbol is not terminal, i.e., it has a definition, look for a match
			// find symbol match

		}

		// was there a definition match?
	}

}

// FindDerivation check to see if the
func FindDerivation(parsedSymbols []GrammarSymbol, def SymbolDefinition) (int, bool) {

	return 1, false
}

//func FindDerivation
