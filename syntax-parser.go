package main

// SymbolDefinition is a list of GrammarSymbols that make up
// a parent Symbol (non-terminal symbol)
type SymbolDefinition struct {
	definition []GrammarSymbols
}

func GetGrammarRules() map[GrammarSymbols][]SymbolDefinition {
	// MAP REPRESENTING THE GRAMMAR RULE DEFINITIONS
	// Each KEY is a GrammarSymbol.
	// Each VALUE is a list of SymbolDefinitions
	grammarRules := make(map[GrammarSymbols][]SymbolDefinition)

	// define STMT_LIST (2 defs)
	grammarRules[STMT_LIST] = []SymbolDefinition{
		{[]GrammarSymbols{STMT, PERIOD}},
		{[]GrammarSymbols{STMT, SEMICOLON, STMT_LIST}}}

	// define STMT (2 defs)
	grammarRules[STMT] = []SymbolDefinition{
		{[]GrammarSymbols{POINT_DEF}},
		{[]GrammarSymbols{TEST_POINT}}}

	// define POINT_DEF (1 def)
	grammarRules[POINT_DEF] = []SymbolDefinition{
		{[]GrammarSymbols{ID, ASSIGN, POINT, LPAREN, NUM, COMMA, RPAREN}}}

	// define TEST_POINT (1 def)
	grammarRules[TEST_POINT] = []SymbolDefinition{
		{[]GrammarSymbols{TEST, LPAREN, OPTION, COMMA, POINT_LIST, RPAREN}}}

	// define ID (1 def)
	grammarRules[ID] = []SymbolDefinition{
		{[]GrammarSymbols{LETTER}}}

	// define NUM (1 def)
	grammarRules[NUM] = []SymbolDefinition{
		{[]GrammarSymbols{DIGIT}}}

	// define OPTION (2 def)
	grammarRules[OPTION] = []SymbolDefinition{
		{[]GrammarSymbols{TRIANGLE}},
		{[]GrammarSymbols{SQUARE}}}

	// define POINT_LIST (2 def)
	grammarRules[POINT_LIST] = []SymbolDefinition{
		{[]GrammarSymbols{ID}},
		{[]GrammarSymbols{ID, COMMA, POINT_LIST}}}

	return grammarRules
}
