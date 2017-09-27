package rpncalc

import (
	"math"
	"strings"
	"testing"
)

func TestOpsInfo(t *testing.T) {

	for _, o := range OpsInfo() {

		if o.Type == StaticOp && len(o.Names) < 1 {
			t.Errorf("Empty names for operator %v", o)
			continue
		}
		if o.Type == DynamicOp && o.Prefix == "" {
			t.Errorf("Empty names for operator %v", o)
			continue
		}
		for _, n := range o.Names {
			if strings.TrimSpace(n) == "" {
				t.Errorf("Empty name for operator %v", o)
			}
		}

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
