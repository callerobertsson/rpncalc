package rpncalc

import (
	"math"
	"strings"
	"testing"
)

func TestOpsInfo(t *testing.T) {

	for _, o := range OpsInfo() {

		// Static ops should have one or more names
		if o.Type == StaticOp && len(o.Names) < 1 {
			t.Errorf("Empty names for operator %v", o)
			continue
		}

		// Dynamic ops should have a prefix
		if o.Type == DynamicOp && o.Prefix == "" {
			t.Errorf("Empty names for operator %v", o)
			continue
		}

		// Only names or prefix allowed
		if len(o.Names) > 0 && o.Prefix != "" {
			t.Errorf("Not allowed to have both names and prefix%v", o)
			continue
		}

		// Check if any name is empty
		for _, n := range o.Names {
			if strings.TrimSpace(n) == "" {
				t.Errorf("Empty name for operator %v", o)
			}
		}

		// Check that a description exists
		if "" == o.Description {
			t.Errorf("Empty description for operator %v", o)
		}
	}

}

// Helper func for checking that two float64 are almost equal
func almostEqual(x, y float64) bool {
	delta := 0.000000000000001
	return math.Abs(x-y) < delta
}
