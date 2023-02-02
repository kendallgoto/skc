package lang

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ParseError struct {
	Line    int    `json:"line"`
	Column  int    `json:"col"`
	Message string `json:"msg"`
}
type errorListener struct {
	Errors []ParseError
}

func NewErrorListener() *errorListener {
	return &errorListener{
		Errors: []ParseError{},
	}
}
func (d *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("ParseError: line %d:%d %s\n", line, column, msg)
	d.Errors = append(d.Errors, ParseError{
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

func (d *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	fmt.Println("Got ReportAmbiguity")
}

func (d *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	fmt.Println("Got ReportAttemptingFullContext")
}

func (d *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	fmt.Println("Got ReportContextSensitivity")

}
