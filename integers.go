package goalgebra

import "errors"

type integer struct {
	rational
	value int64
}

func Int(value int64) integer {
	return integer{value: value}
}

func Quotient(a integer, b integer) (integer, error) {
	if b == Int(0) {
		return integer{}, errors.New("Division by 0")
	}
	return Int(a.value / b.value), nil
}

func IntMul(a integer, b integer) integer {
	return Int(a.value * b.value)
}

func IntAdd(a integer, b integer) integer {
	return Int(a.value + b.value)
}

func Mod(a integer, b integer) (integer, error) {
	if b == Int(0) {
		return integer{}, errors.New("Division by 0")
	}
	return Int(a.value % b.value), nil
}

func IntMinus(a integer) integer {
	return Int(-a.value)
}

func IntAbs(a integer) integer {
	if a.value < 0 {
		return IntMinus(a)
	}
	return a
}

func IntPow(a integer, b integer) (integer, error) {
	if a == Int(0) && b == Int(0) {
		return integer{}, errors.New("Undefined 0^0")
	}
	if b.value < 0 {
		return integer{}, errors.New("exponent in IntPow must be non negative")
	}
	if a == Int(0) {
		return Int(0), nil
	}
	if b == Int(0) {
		return Int(1), nil
	}
	partialPow, err := IntPow(a, IntSubtract(b, Int(1)))
	if err != nil {
		return integer{}, err
	}
	return IntMul(partialPow, a), nil
}

func IntSubtract(a integer, b integer) integer {
	return IntAdd(a, IntMinus(b))
}

func GCD(a integer, b integer) (integer, error) {
	for b != Int(0) {
		r, err := Mod(a, b)
		if err != nil {
			return integer{}, nil
		}
		a, b = b, r
	}
	return IntAbs(a), nil
}
