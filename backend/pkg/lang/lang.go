package lang

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kendallgoto/skc/pkg/lang/yacc"
)

// TODO: figure out an extensible way to get grammars working ...
type Token struct {
	TokenName         string
	TokenAlternatives []string // this should prob get a map for faster lookup
	TokenFn           func(string) (string, error)
}

var (
	KnownTokens = map[string]*Token{
		"say": {
			TokenName:         "say",
			TokenAlternatives: []string{"print", "output"},
			TokenFn:           printString,
		},
	}
)

func printString(input string) (string, error) {
	// expect string after token
	r, _ := regexp.Compile(`^"(\S*)".*`)
	match := r.FindSubmatch([]byte(input))
	// if match is a quote-defined string literal
	if len(match) < 2 {
		return "", fmt.Errorf("expected string, got %s", input)
	}
	return fmt.Sprintf(`print("%s")`, match[1]), nil
}

type ParseContext struct {
	ParseLines   []*ParseContext
	ParseContent string
	ParseTokens  []string
	Result       string
}

func (ctx *ParseContext) String() string {
	res := ""
	for _, nestedCtx := range ctx.ParseLines {
		res += nestedCtx.String()
	}
	if len(res) == 0 {
		return "\n" + ctx.Result
	}
	return res
}
func parseContexts(input string) *ParseContext {
	// for now, just split on the period ...
	subcontexts := []*ParseContext{}
	segments := strings.Split(input, ".")
	if len(segments) > 1 {
		for _, segment := range segments {
			subcontexts = append(subcontexts, parseContexts(strings.TrimSpace(segment)))
		}
	}
	// find relevant tokens for this ...
	result := ""
	words := strings.Fields(input)
	for i, word := range words {
		if len(result) > 0 {
			break
		}
		for _, token := range KnownTokens {
			if strings.EqualFold(word, token.TokenName) {
				// concat following words and send
				rest := strings.Join(words[(i+1):], " ")
				result, _ = token.TokenFn(rest)
				break
			}
		}

	}
	context := &ParseContext{
		ParseLines:   subcontexts,
		ParseContent: strings.TrimSpace(input),
		ParseTokens:  []string{},
		Result:       result,
	}
	return context
}
func Parse(input string) (string, error) {
	//context := parseContexts(input)
	//return strings.TrimSpace(context.String()), nil
	yacc.Execute(input)
	return "", nil
}
