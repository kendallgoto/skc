package lang_test

import (
	"testing"

	"github.com/kendallgoto/skc/pkg/lang"
	"github.com/stretchr/testify/assert"
)

const (
	basicPrintInput        = `Say "hello" to me.`
	basicPrintOutput       = `print("hello")`
	basicConditionalInput  = `If "hello" is "hello", then say "hello" to me.`
	basicConditionalOutput = `if("hello" == "hello"):` + "\n\t" + `print("hello")`
)

func TestBasicPrint(t *testing.T) {
	testInput := basicPrintInput
	result, parseerrors, err := lang.Parse(testInput)
	assert.Empty(t, parseerrors, "Parsing had errors %+v", parseerrors)
	assert.Nil(t, err, "lang returned error %s", err)
	assert.NotEmpty(t, result, "lang returned no result")
	assert.Equal(t, basicPrintOutput, result, "lang returned invalid response %s", result)
}
func TestConditional(t *testing.T) {
	testInput := basicConditionalInput
	result, parseerrors, err := lang.Parse(testInput)
	assert.Empty(t, parseerrors, "Parsing had errors %+v", parseerrors)
	assert.Nil(t, err, "lang returned error %s", err)
	assert.NotEmpty(t, result, "lang returned no result")
	assert.Equal(t, basicConditionalOutput, result, "lang returned invalid response %s", result)
}
