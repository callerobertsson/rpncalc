// Package rpncalc operators
package rpncalc

// OpInfo contains displayable operation information
type OpInfo struct {
	Names       []string
	Description string
}

// OpsInfo returns an OpInfo slice with all supported static operators
func OpsInfo() []OpInfo {
	ois := []OpInfo{}
	for _, o := range unaryOps {
		ois = append(ois, OpInfo{o.names, o.description})
	}
	for _, o := range binaryOps {
		ois = append(ois, OpInfo{o.names, o.description})
	}
	for _, o := range registerOps {
		ois = append(ois, OpInfo{[]string{o.prefix + "X"}, o.description})
	}
	return ois
}

func opClearStack(r *RpnCalc) error {
	r.ClearStack()
	return nil
}

func opClearRegs(r *RpnCalc) error {
	r.ClearRegs()
	return nil
}
