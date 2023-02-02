package lang

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kendallgoto/skc/pkg/lang/parser"
)

func (v *skcVisitor) VisitSayStatement(ctx *parser.SayStatementContext) interface{} {
	v.result += "print("
	if ctx.Literal() != nil {
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
func (v *skcVisitor) VisitEquality(ctx *parser.EqualityContext) interface{} {
	if ctx.Equal() != nil {
		v.result += " == "
	} else if ctx.GreaterThan() != nil {
		v.result += " > "
	} else if ctx.LessThan() != nil {
		v.result += " < "
	}
	return nil
}
func (v *skcVisitor) VisitCondition(ctx *parser.ConditionContext) interface{} {
	if ctx.Literal(0) != nil {
		v.result += ctx.Literal(0).GetText()
	}
	// do we have an equality?
	if ctx.Equality() != nil {
		ctx.Equality().Accept(v)
	} else if ctx.Is() != nil {
		// no equality, but we have `is`, assume ==
		v.result += " == "
	}
	if ctx.Literal(1) != nil {
		v.result += ctx.Literal(1).GetText()
	}
	return nil
}

func (v *skcVisitor) VisitConditionalStatement(ctx *parser.ConditionalStatementContext) interface{} {
	v.result += "if("
	ctx.Condition().Accept(v)
	v.result += "):\n"
	// handle block indent:

	// cache the result and clear it
	currentResult := strings.Clone(v.result)
	v.result = ""
	for i := 3; i < len(ctx.GetChildren()); i++ {
		ctx.GetChild(i).(antlr.ParseTree).Accept(v)
	}

	// now indent everything
	scanner := bufio.NewScanner(strings.NewReader(v.result))
	v.result = ""
	for scanner.Scan() {
		fmt.Println("add line", scanner.Text())
		v.result += "\t" + scanner.Text() + "\n"
	}
	// now rebuild
	v.result = currentResult + v.result

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

	return strings.TrimSpace(visitor.result), listener.Errors, nil
}
