package lang

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kendallgoto/skc/pkg/lang/parser"
)

func (v *skcVisitor) VisitSayStatement(ctx *parser.SayStatementContext) interface{} {
	endTack := ""
	if ctx.OutputFile() != nil {
		targetFile := ctx.StringLiteral()
		v.result += "f = open("
		if targetFile != nil {
			v.result += targetFile.GetText()
		}
		v.result += `, "w")` + "\n" + `f.write(`
		endTack = ")\nf.close()"
	} else {
		v.result += "print("
		endTack = ")"
	}
	if ctx.Literal() != nil {
		v.result += ctx.Literal().GetText()
	}
	v.result += endTack
	return nil
}
func (v *skcVisitor) VisitLine(ctx *parser.LineContext) interface{} {
	v.VisitChildren(ctx)
	v.result += "\n"
	return nil
}
func (v *skcVisitor) VisitEquality(ctx *parser.EqualityContext) interface{} {
	hasNot := (ctx.GetParent().(*parser.ConditionContext).Not() != nil)
	if ctx.Equal() != nil {
		if hasNot {
			v.result += " != "
		} else {
			v.result += " == "
		}
	} else if ctx.GreaterThan() != nil {
		if hasNot {
			v.result += " <= "
		} else {
			v.result += " > "
		}
	} else if ctx.LessThan() != nil {
		if hasNot {
			v.result += " >= "
		} else {
			v.result += " < "
		}
	} else if ctx.LessThanOrEqual() != nil {
		if hasNot {
			v.result += " > "
		} else {
			v.result += " <= "
		}
	} else if ctx.GreaterThanOrEqual() != nil {
		if hasNot {
			v.result += " < "
		} else {
			v.result += " >= "
		}
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

func (v *skcVisitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	v.result += "while("
	ctx.Condition().Accept(v)
	v.result += "):\n"

	currentResult := strings.Clone(v.result)
	v.result = ""
	ctx.Statement().Accept(v)

	scanner := bufio.NewScanner(strings.NewReader(v.result))
	v.result = ""
	for scanner.Scan() {
		v.result += "\t" + scanner.Text() + "\n"
	}

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
