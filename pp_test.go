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
		t.Errorf("failed to SetOutput")
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

func TestColorScheme(t *testing.T) {
	SetColorScheme(ColorScheme{})

	if defaultPrettyPrinter.currentScheme.FieldName == 0 {
		t.FailNow()
	}
}

func TestWithLineInfo(t *testing.T) {
	outputWithoutLineInfo := new(bytes.Buffer)
	SetDefaultOutput(outputWithoutLineInfo)
	Print("abcde")

	outputWithLineInfo := new(bytes.Buffer)
	SetDefaultOutput(outputWithLineInfo)
	WithLineInfo = true
	Print("abcde")

	ResetDefaultOutput()

	if bytes.Equal(outputWithLineInfo.Bytes(), outputWithoutLineInfo.Bytes()) {
		t.Errorf("outputWithLineInfo should not have the same contents than outputWithoutLineInfo")
	}
}

func TestWithLineInfoBackwardsCompatible(t *testing.T) {
	// Test that the global accessible field `WithLineInfo` does not mutate other instances

	outputWithLineInfo := new(bytes.Buffer)
	SetDefaultOutput(outputWithLineInfo)
	WithLineInfo = true
	Print("abcde")

	outputWithoutLineInfo := new(bytes.Buffer)
	pp := New()
	pp.SetOutput(outputWithoutLineInfo)
	pp.Print("abcde")

	if bytes.Equal(outputWithLineInfo.Bytes(), outputWithoutLineInfo.Bytes()) {
		t.Errorf("outputWithLineInfo should not have the same contents than outputWithoutLineInfo")
	}

	ResetDefaultOutput()
}
