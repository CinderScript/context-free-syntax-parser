package main

// DEFINES THE POSSIBLE TOKENS OF THE LANGUAGE'S GRAMMAR
type GrammarSymbols string

const (
	POINT      GrammarSymbols = "POINT"
	ID         GrammarSymbols = "ID"
	NUM        GrammarSymbols = "NUM"
	SEMICOLON  GrammarSymbols = "SEMICOLON"
	COMMA      GrammarSymbols = "COMMA"
	PERIOD     GrammarSymbols = "PERIOD"
	LPAREN     GrammarSymbols = "LPAREN"
	RPAREN     GrammarSymbols = "RPAREN"
	ASSIGN     GrammarSymbols = "ASSIGN"
	TRIANGLE   GrammarSymbols = "TRIANGLE"
	SQUARE     GrammarSymbols = "SQUARE"
	TEST       GrammarSymbols = "TEST"
	STMT_LIST  GrammarSymbols = "STMT_LIST"
	STMT       GrammarSymbols = "STMT"
	POINT_DEF  GrammarSymbols = "POINT_DEF"
	TEST_POINT GrammarSymbols = "TEST_POINT"
	OPTION     GrammarSymbols = "OPTION"
	POINT_LIST GrammarSymbols = "POINT_LIST"
	LETTER     GrammarSymbols = "LETTER"
	DIGIT      GrammarSymbols = "DIGIT"
)

// Struct for saving the grammer token and lexeme pair
type TokenLexemePair struct {
	token  GrammarSymbols
	lexeme string
}

// SymbolDefinition is a list of GrammarSymbols that make up
// a parent Symbol (non-terminal symbol)
type SymbolDefinition struct {
	definition []GrammarSymbols
}
