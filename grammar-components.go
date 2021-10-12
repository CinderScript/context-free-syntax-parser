package main

import "strings"

// DEFINES THE POSSIBLE TOKENS OF THE LANGUAGE'S GRAMMAR
type GrammarSymbol string

const (
	POINT      GrammarSymbol = "POINT"
	ID         GrammarSymbol = "ID"
	NUM        GrammarSymbol = "NUM"
	SEMICOLON  GrammarSymbol = "SEMICOLON"
	COMMA      GrammarSymbol = "COMMA"
	PERIOD     GrammarSymbol = "PERIOD"
	LPAREN     GrammarSymbol = "LPAREN"
	RPAREN     GrammarSymbol = "RPAREN"
	ASSIGN     GrammarSymbol = "ASSIGN"
	TRIANGLE   GrammarSymbol = "TRIANGLE"
	SQUARE     GrammarSymbol = "SQUARE"
	TEST       GrammarSymbol = "TEST"
	START      GrammarSymbol = "START"
	STMT_LIST  GrammarSymbol = "STMT_LIST"
	STMT       GrammarSymbol = "STMT"
	POINT_DEF  GrammarSymbol = "POINT_DEF"
	TEST_POINT GrammarSymbol = "TEST_POINT"
	OPTION     GrammarSymbol = "OPTION"
	POINT_LIST GrammarSymbol = "POINT_LIST"
	LETTER     GrammarSymbol = "LETTER"
	DIGIT      GrammarSymbol = "DIGIT"
)

// Struct for saving the grammer token and lexeme pair
type TokenLexemePair struct {
	token  GrammarSymbol
	lexeme string
}

// SymbolDefinition is a list of GrammarSymbols that make up
// a parent Symbol (non-terminal symbol)
type SymbolDefinition struct {
	Symbols []GrammarSymbol
}

// GetGrammarRules generates and returns a map of the grammar's
// symbols and their definitions.
func GetGrammarRules() map[GrammarSymbol][]SymbolDefinition {
	// MAP REPRESENTING THE GRAMMAR RULE DEFINITIONS
	// Each KEY is a GrammarSymbol.
	// Each VALUE is a list of SymbolDefinitions
	grammarRules := make(map[GrammarSymbol][]SymbolDefinition)

	// define STMT_LIST (2 defs)
	grammarRules[STMT_LIST] = []SymbolDefinition{
		{[]GrammarSymbol{STMT}},
		{[]GrammarSymbol{STMT, PERIOD}},
		{[]GrammarSymbol{STMT, SEMICOLON, STMT_LIST}}}

	// define STMT (2 defs)
	grammarRules[STMT] = []SymbolDefinition{
		{[]GrammarSymbol{POINT_DEF}},
		{[]GrammarSymbol{TEST_POINT}}}

	// define POINT_DEF (1 def)
	grammarRules[POINT_DEF] = []SymbolDefinition{
		{[]GrammarSymbol{ID, ASSIGN, POINT, LPAREN, NUM, COMMA, RPAREN}}}

	// define TEST_POINT (1 def)
	grammarRules[TEST_POINT] = []SymbolDefinition{
		{[]GrammarSymbol{TEST, LPAREN, OPTION, COMMA, POINT_LIST, RPAREN}}}

	// define OPTION (2 def)
	grammarRules[OPTION] = []SymbolDefinition{
		{[]GrammarSymbol{TRIANGLE}},
		{[]GrammarSymbol{SQUARE}}}

	// define POINT_LIST (2 def)
	grammarRules[POINT_LIST] = []SymbolDefinition{
		{[]GrammarSymbol{ID}},
		{[]GrammarSymbol{ID, COMMA, POINT_LIST}}}

	//don't need to define NUM or ID - those are found by the lexical scanner
	return grammarRules
}

func GetDefinitionString(symbols ...GrammarSymbol) string {
	var symbolList []string
	for _, s := range symbols {
		symbolList = append(symbolList, string(s))
	}

	return strings.Join(symbolList, " ")
}
