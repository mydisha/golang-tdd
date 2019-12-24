package packages_test

import (
	"encoding/json"
	"github.com/mydisha/tdd/packages"
	"testing"
)

func TestParseFloat64(t *testing.T) {
	t.Run("Nil input", func(t *testing.T) {
		var i interface{}

		convertNum, _ := packages.ParseFloat64(i)

		if convertNum != 1 {
			t.Fatalf("return must be 1 instead of %v", convertNum)
		}
	})
	t.Run("Json number", func(t *testing.T) {
		var i interface{} = json.Number("20.4")

		convertedNum, _ := packages.ParseFloat64(i)

		if convertedNum != 20.4 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
}
