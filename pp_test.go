package pp

import (
	"bytes"
	"strings"
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

	if Default.currentScheme.FieldName == 0 {
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

func TestStructPrintingWithTags(t *testing.T) {
	type Foo struct {
		IgnoreMe     interface{} `pp:"-"`
		ChangeMyName string      `pp:"NewName"`
		OmitIfEmpty  string      `pp:",omitempty"`
		Full         string      `pp:"full,omitempty"`
	}

	testCases := []struct {
		name               string
		foo                Foo
		omitIfEmptyOmitted bool
		fullOmitted        bool
	}{
		{
			name: "all set",
			foo: Foo{
				IgnoreMe:     "i'm a secret",
				ChangeMyName: "i'm an alias",
				OmitIfEmpty:  "i'm not empty",
				Full:         "hello",
			},
			omitIfEmptyOmitted: false,
			fullOmitted:        false,
		},
		{
			name: "omit if empty not set",
			foo: Foo{
				IgnoreMe:     "i'm a secret",
				ChangeMyName: "i'm an alias",
				OmitIfEmpty:  "",
				Full:         "hello",
			},
			omitIfEmptyOmitted: true,
			fullOmitted:        false,
		},
		{
			name: "both omitted",
			foo: Foo{
				IgnoreMe:     "i'm a secret",
				ChangeMyName: "i'm an alias",
				OmitIfEmpty:  "",
				Full:         "",
			},
			omitIfEmptyOmitted: true,
			fullOmitted:        true,
		},
		{
			name:               "zero",
			foo:                Foo{},
			omitIfEmptyOmitted: true,
			fullOmitted:        true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			output := new(bytes.Buffer)
			pp := New()
			pp.SetOutput(output)

			pp.Print(tc.foo)

			result := output.String()

			if strings.Contains(result, "IgnoreMe") {
				t.Error("result should not contain IgnoreMe")
			}

			if strings.Contains(result, "OmitIfEmpty") && tc.omitIfEmptyOmitted {
				t.Error("result should not contain OmitIfEmpty")
			} else if !strings.Contains(result, "OmitIfEmpty") && !tc.omitIfEmptyOmitted {
				t.Error("result should contain OmitIfEmpty")
			}

			// field Full is renamed to full by the tag
			if strings.Contains(result, "full") && tc.fullOmitted {
				t.Error("result should not contain full")
			} else if !strings.Contains(result, "full") && !tc.fullOmitted {
				t.Error("result should contain full")
			}
		})
	}

}
