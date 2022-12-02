package pipe

import cns "golang.org/x/exp/constraints"

func Sum[T cns.Float | cns.Integer | cns.Complex](a, b T) T {
	return a + b
}
