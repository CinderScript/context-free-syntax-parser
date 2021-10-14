/*
 * Class:		CSC 3100 - Concepts in Programming Languages
 * Title:		GO: Lexical and Syntax Analyzer
 * Purpose:		The purpose of this assignment is to practice the following concepts:
 * 				Context Free Grammar / BNF
 *				Lexical Analasys (scanner)
 *				Syntax Analasys (parser)
 *
 * 				grammar-components.go contains the elements for defining BNF style grammar rules.
 *
 * Author:		Maynard, Greg
 * Date:		10/13/2022
 */

package main

// DEFINES THE POSSIBLE TOKENS OF THE LANGUAGE'S GRAMMAR
type GrammarSymbol string

const (
	POINT         GrammarSymbol = "POINT"
	ID            GrammarSymbol = "ID"
	NUM           GrammarSymbol = "NUM"
	SEMICOLON     GrammarSymbol = "SEMICOLON"
	COMMA         GrammarSymbol = "COMMA"
	PERIOD        GrammarSymbol = "PERIOD"
	LPAREN        GrammarSymbol = "LPAREN"
	RPAREN        GrammarSymbol = "RPAREN"
	ASSIGN        GrammarSymbol = "ASSIGN"
	TRIANGLE      GrammarSymbol = "TRIANGLE"
	SQUARE        GrammarSymbol = "SQUARE"
	TEST          GrammarSymbol = "TEST"
	START         GrammarSymbol = "START"
	STMT_LIST     GrammarSymbol = "STMT_LIST"
	STMT          GrammarSymbol = "STMT"
	POINT_DEF     GrammarSymbol = "POINT_DEF"
	TEST_SHAPE    GrammarSymbol = "TEST_POINT"
	OPTION        GrammarSymbol = "OPTION"
	TEST_SQUARE   GrammarSymbol = "TEST_SQUARE"
	TEST_TRIANGLE GrammarSymbol = "TEST_TRIANGLE"
	POINT_LIST    GrammarSymbol = "POINT_LIST"
	LETTER        GrammarSymbol = "LETTER"
	DIGIT         GrammarSymbol = "DIGIT"
)

// returns the string version of the given GrammarSymbol
func (s GrammarSymbol) String() string {
	return string(s)
}

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

	// define START (1 defs)
	grammarRules[START] = []SymbolDefinition{
		{[]GrammarSymbol{STMT_LIST}}}

	// define STMT_LIST (2 defs)
	grammarRules[STMT_LIST] = []SymbolDefinition{
		{[]GrammarSymbol{STMT, PERIOD}},
		{[]GrammarSymbol{STMT, SEMICOLON, STMT_LIST}}}

	// define STMT (2 defs)
	grammarRules[STMT] = []SymbolDefinition{
		{[]GrammarSymbol{POINT_DEF}},
		{[]GrammarSymbol{TEST_SHAPE}}}

	// define POINT_DEF (1 def)
	grammarRules[POINT_DEF] = []SymbolDefinition{
		{[]GrammarSymbol{ID, ASSIGN, POINT, LPAREN, NUM, COMMA, NUM, RPAREN}}}

	// // define TEST_POINT (1 def)
	// grammarRules[TEST_SHAPE] = []SymbolDefinition{
	// 	{[]GrammarSymbol{TEST, LPAREN, OPTION, COMMA, POINT_LIST, RPAREN}}}

	// define TEST_POINT (1 def)
	grammarRules[TEST_SHAPE] = []SymbolDefinition{
		{[]GrammarSymbol{TEST_SQUARE}},
		{[]GrammarSymbol{TEST_TRIANGLE}}}

	// define TEST_SQUARE
	grammarRules[TEST_SQUARE] = []SymbolDefinition{
		{[]GrammarSymbol{TEST, LPAREN, SQUARE, COMMA,
			ID, COMMA, ID, COMMA, ID, COMMA, ID, RPAREN}}}

	// define TEST_TRIANGLE
	grammarRules[TEST_TRIANGLE] = []SymbolDefinition{
		{[]GrammarSymbol{TEST, LPAREN, TRIANGLE,
			COMMA, ID, COMMA, ID, COMMA, ID, RPAREN}}}

	// define OPTION (2 def)
	grammarRules[OPTION] = []SymbolDefinition{
		{[]GrammarSymbol{TRIANGLE}},
		{[]GrammarSymbol{SQUARE}}}

	// define POINT_LIST (2 def)
	grammarRules[POINT_LIST] = []SymbolDefinition{
		{[]GrammarSymbol{ID}},
		{[]GrammarSymbol{POINT_LIST, COMMA, POINT_LIST}}}

	//don't need to define NUM or ID - those are found by the lexical scanner
	return grammarRules
}
