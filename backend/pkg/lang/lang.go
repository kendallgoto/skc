package lang

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/davecgh/go-spew/spew"
	"github.com/kendallgoto/skc/pkg/lang/parser"
)

func (v *skcVisitor) VisitSayStatement(ctx *parser.SayStatementContext) interface{} {
	v.result += "print("
	if ctx.Literal() != nil {
		spew.Dump(ctx.Literal())
		v.result += ctx.Literal().GetText()
	}
	v.result += ")"
	return nil
}
func (v *skcVisitor) VisitLine(ctx *parser.LineContext) interface{} {
	v.VisitChildren(ctx)
	v.result += "\n"
	return nil
}

func Parse(input string) (string, []ParseError, error) {
	fmt.Println(input)
	visitor := &skcVisitor{}
	listener := &errorListener{}
	is := antlr.NewInputStream(input)
	lexer := parser.NewSkcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSkcParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(listener)
	ctx := p.Enter()
	visitor.Visit(ctx)

	return visitor.result, listener.Errors, nil
}
