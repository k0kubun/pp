package pp

import (
	"math"
	"strings"
	"testing"
)

func TestMapWithNaNKeysPreservesValues(t *testing.T) {
	printer := New()
	printer.SetColoringEnabled(false)
	for _, object := range []interface{}{
		map[float64]string{math.NaN(): "float value"},
		map[complex128]string{complex(math.NaN(), 1): "complex value"},
		map[[1]float64]string{{math.NaN()}: "array value"},
	} {
		got := printer.Sprint(object)
		if !strings.Contains(got, "value") {
			t.Errorf("map values were lost: %s", got)
		}
	}
}
