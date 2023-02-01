%{

package yacc
import (
	"fmt"
	"math/big"
	"bytes"
	"fmt"
	"log"
	"unicode/utf8"
)
%}

%union {
	num *big.Rat
}
%type <num> expr expr1
%token '+' '-' '*' '/' '(' ')'

%token <num> NUM

%%
top:
	expr
	{
		if $1.IsInt() {
			fmt.Println($1.Num().String())
		} else {
			fmt.Println($1.String())
		}
	}
expr:
	expr1
	|	'+' expr
		{
			$$ = $2
		}
expr1:
	NUM
	| expr1 '+' NUM
	{
		$$ = $1.Add($1, $3)
	}
%%
const eof = 0
type skcLex struct {
	line []byte
	peek rune
}
func (x *skcLex) Lex(skclval *skcSymType) int {
	for {
		c := x.next()
		switch c {
		case eof:
			return eof
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.num(c, skclval)
		case '+', '-', '*', '/', '(', ')':
			return int(c)

		// Recognize Unicode multiplication and division
		// symbols, returning what the parser expects.
		case '×':
			return '*'
		case '÷':
			return '/'

		case ' ', '\t', '\n', '\r':
		default:
			log.Printf("unrecognized character %q", c)
		}
	}
}

// Lex a number.
func (x *skcLex) num(c rune, skclval *skcSymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			log.Fatalf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
	L: for {
		c = x.next()
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E':
			add(&b, c)
		default:
			break L
		}
	}
	if c != eof {
		x.peek = c
	}
	skclval.num = &big.Rat{}
	_, ok := skclval.num.SetString(b.String())
	if !ok {
		log.Printf("bad number %q", b.String())
		return eof
	}
	return NUM
}

// Return the next rune for the lexer.
func (x *skcLex) next() rune {
	if x.peek != eof {
		r := x.peek
		x.peek = eof
		return r
	}
	if len(x.line) == 0 {
		return eof
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError && size == 1 {
		log.Print("invalid utf8")
		return x.next()
	}
	return c
}
func (x *skcLex) Error(s string) {
	log.Printf("parse error: %s", s)
}

func Execute(line string) {
	skcParse(&skcLex{
		line: []byte(line),
	})
}
