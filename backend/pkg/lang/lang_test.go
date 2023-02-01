package lang_test

import (
	"testing"

	"github.com/kendallgoto/skc/pkg/lang"
	"github.com/stretchr/testify/assert"
)

const (
	basicPrintInput  = `Say "hello" to me.`
	basicPrintOutput = `print("hello")`
)

func TestBasicPrint(t *testing.T) {
	testInput := basicPrintInput
	result, err := lang.Parse(testInput)
	assert.Nil(t, err, "lang returned error %s", err)
	assert.NotEmpty(t, result, "lang returned no result")
	assert.Equal(t, result, basicPrintOutput, "lang returned invalid response %s", result)
}
