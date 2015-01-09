package pp

import (
	"bytes"
	"testing"
)

func TestDefaultOutput(t *testing.T) {
	testOutput := new(bytes.Buffer)
	init := GetDefaultOutput()
	SetDefaultOutput(testOutput)
	if GetDefaultOutput() != testOutput {
		t.Errorf("failed to SetDefaultOutput")
	}
	if len(testOutput.String()) != 0 {
		t.Errorf("testOutput should be initialized")
	}
	Print("abcde")
	if len(testOutput.String()) == 0 {
		t.Errorf("Expected Print output to testOutput, testOutput is %s", testOutput.String())
	}
	if init == GetDefaultOutput() {
		t.Errorf("it should be changed DefaultOutput")
	}
	ResetDefaultOutput()
	if init != GetDefaultOutput() {
		t.Errorf("it should be reset to initial default output")
	}
}
