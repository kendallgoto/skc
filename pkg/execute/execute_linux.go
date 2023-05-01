//go:build false
// +build false

package execute

import python3 "github.com/go-python/cpy3"

const stdoutCapture = `
import sys
class CatchOutput:
	def __init__(self):
		self.value = ''
	def write(self, txt):
		self.value += txt
catchOut = CatchOutput()
sys.stdout = catchOut
sys.stderr = catchOut
`

func Run(code string) (string, error) {
	// todo: you should sandbox this!
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	pModule := python3.PyImport_AddModule("__main__") //main module to capture scope
	python3.PyRun_SimpleString(stdoutCapture)
	python3.PyRun_SimpleString(code)
	catch := pModule.GetAttrString("catchOut")
	value := catch.GetAttrString("value")
	return python3.PyUnicode_AsUTF8(value), nil
}
