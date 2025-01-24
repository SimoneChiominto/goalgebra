package goalgebra

import "fmt"

type rational interface {
	real
	Numerator() integer
	Denominator() integer
	SimplifyRational() rational
}
type fraction struct {
	rational
	num integer
	den integer
}

var EmptyFraction = fraction{}

func Frac(a integer, b integer) fraction {
	if b == Int(0) {
		return EmptyFraction
	}
	if IntMul(a, b).value >= 0 {
		return fraction{num: IntAbs(a), den: IntAbs(b)}
	}
	return fraction{num: IntMinus(IntAbs(a)), den: IntAbs(b)}
}

func (u integer) String() string {
	return fmt.Sprintf("%d", u.value)
}

func (u fraction) String() string {
	return fmt.Sprintf("%s/%s", u.num.String(), u.den.String())
}

func (u integer) Numerator() integer {
	return u
}

func (u fraction) Numerator() integer {
	if u == EmptyFraction {
		return Int(0)
	}
	return u.num
}

func (u integer) Denominator() integer {
	return Int(1)
}

func (u fraction) Denominator() integer {
	if u == EmptyFraction {
		return Int(0)
	}
	return u.den
}

func Inv(u rational) rational {
	return Frac(u.Denominator(), u.Numerator()).SimplifyRational()
}

func (u integer) SimplifyRational() rational {
	return u
}

func (u fraction) SimplifyRational() rational {
	if u == EmptyFraction {
		return EmptyFraction
	}
	r, err := Mod(u.num, u.den)
	if err != nil {
		return EmptyFraction
	}
	if r == Int(0) {
		q, err := Quotient(u.num, u.den)
		if err != nil {
			return EmptyFraction
		}
		return q
	}
	gcd, err := GCD(u.num, u.den)
	if err != nil {
		return EmptyFraction
	}
	num, err1 := Quotient(u.num, gcd)
	den, err2 := Quotient(u.den, gcd)
	if err1 != nil || err2 != nil {
		return EmptyFraction
	}
	return Frac(num, den)
}

func RatMul(u rational, w rational) rational {
	if w.Numerator() == Int(0) {
		return EmptyFraction
	}
	return Frac(
		IntMul(u.Numerator(), w.Numerator()),
		IntMul(u.Denominator(), w.Denominator()),
	).SimplifyRational()
}
func Div(u rational, w rational) rational {
	return RatMul(u, Inv(w))
}

func RatPow(u rational, n integer) rational {
	u = u.SimplifyRational()
	switch v := u.(type) {
	case integer:
		pow, err := IntPow(v, n)
		if err != nil {
			return EmptyFraction
		}
		return pow
	case fraction:
		num, err1 := IntPow(v.num, n)
		den, err2 := IntPow(v.den, n)
		if err1 != nil || err2 != nil {
			return EmptyFraction
		}
		return Frac(num, den)
	}
	return nil
}

func RatAdd(u rational, w rational) rational {
	return Frac(
		IntAdd(
			IntMul(u.Numerator(), w.Denominator()),
			IntMul(u.Denominator(), w.Numerator()),
		),
		IntMul(u.Denominator(), w.Denominator()),
	).SimplifyRational()
}

func Minus(u rational) rational {
	switch v := u.(type) {
	case integer:
		return IntMinus(v)
	case fraction:
		return Frac(IntMinus(u.Numerator()), u.Denominator())
	}
	return nil
}

func Subtract(u rational, w rational) rational {
	return RatAdd(u, Minus(w))
}

func Abs(u rational) rational {
	switch v := u.(type) {
	case integer:
		return IntAbs(v)
	case fraction:
		return Frac(IntAbs(u.Numerator()), u.Denominator())
	}
	return nil
}
