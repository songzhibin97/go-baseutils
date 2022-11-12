package bternaryexpr

func TernaryExpr[T any](boolExpr bool, trueReturn T, falseReturn T) T {
	if boolExpr {
		return trueReturn
	}
	return falseReturn
}
