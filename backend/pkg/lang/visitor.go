package lang

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kendallgoto/skc/pkg/lang/parser"
)

type skcVisitor struct {
	parser.BaseSkcParserVisitor
	result string
}

func (v *skcVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}
func (v *skcVisitor) VisitEnter(ctx *parser.EnterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *skcVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *skcVisitor) VisitOutputTo(ctx *parser.OutputToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *skcVisitor) VisitOutputType(ctx *parser.OutputTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *skcVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *skcVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
