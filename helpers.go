package vm

import (
	"fmt"
)

func UnsupportedOperandType(op1, op2 ZVal, op string) error {
	if op2 == nil {
		return fmt.Errorf("Unsupported operand types: %s %s", op, op1.String())
	}

	return fmt.Errorf("Unsupported operand types: %s %s %s", op1.String(), op, op2.String())
}

func UnwrapVariable(variable ZVal) ZVal {
	if variable, ok := variable.(*ZVariable); ok {
		return variable.ZVal
	}

	return variable
}

func NumMax[T ~int | ~float64](x, y T) T {
	if x > y {
		return x
	}

	return y
}
