package rpncalc

import (
	"math"
	"strings"
	"testing"
)

func TestOpsInfo(t *testing.T) {

	cases := [][]OpInfo{StaticOpsInfo(), DynamicOpsInfo()}

	for i, c := range cases {

		if len(c) < 1 {
			t.Errorf("Case %d: Empty OpInfo list", i)
		}

		for _, o := range c {
			if len(o.Names) < 1 {
				t.Errorf("Case %d: Empty names for operator %v", i, o)
				continue
			}
			for _, n := range o.Names {
				if strings.TrimSpace(n) == "" {
					t.Errorf("Case %d: Empty name for operator %v", i, o)
				}
			}

			if "" == o.Description {
				t.Errorf("Case %d: Empty description for operator %v", i, o)
			}
		}
	}

}

// Helper func for checking that two float64 are almost equal
func almostEqual(x, y float64) bool {
	delta := 0.000000000000001
	return math.Abs(x-y) < delta
}
