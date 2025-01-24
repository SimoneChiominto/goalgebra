package main

import (
	"fmt"
	"goalgebra"
)

func main() {
	n := goalgebra.Int(3)
	m := goalgebra.Int(4)
	fmt.Println(goalgebra.Quotient(n, m))
	fmt.Println(goalgebra.IntMul(n, m))
	fmt.Println(goalgebra.IntAdd(n, m))
	fmt.Println(goalgebra.Mod(n, m))
	fmt.Println(goalgebra.IntMinus(n))
	fmt.Println(goalgebra.IntAbs(n))
	fmt.Println(goalgebra.IntPow(n, m))
	fmt.Println(goalgebra.IntSubtract(n, m))
	fmt.Println(goalgebra.GCD(n, m))
	u := goalgebra.Frac(goalgebra.Int(2), goalgebra.Int(-4)).SimplifyRational()
	l := goalgebra.Frac(goalgebra.Int(4), goalgebra.Int(2)).SimplifyRational()
	v := goalgebra.Frac(goalgebra.Int(1), goalgebra.Int(3))
	fmt.Println(goalgebra.Inv(l))
	fmt.Println(goalgebra.RatMul(u, v))
	fmt.Println(goalgebra.RatAdd(u, v))
	fmt.Println(goalgebra.Div(u, v))
	fmt.Println(goalgebra.Minus(u))
	fmt.Println(goalgebra.Abs(u))
	fmt.Println(goalgebra.RatPow(u, n))
	fmt.Println(goalgebra.Subtract(u, v))

	fmt.Println(goalgebra.Frac(goalgebra.Int(2), goalgebra.Int(-4)).Simplify())
	fmt.Println(goalgebra.Int(1).Simplify())
	fmt.Println(goalgebra.Log(goalgebra.Int(1)).Simplify())
	fmt.Println(goalgebra.Log(goalgebra.Exp(goalgebra.Int(1))).Simplify())
	fmt.Println(goalgebra.Exp(goalgebra.Log(goalgebra.Int(1))).Simplify())
	fmt.Println(goalgebra.Add(u, goalgebra.Add(l, v)))
	fmt.Println(goalgebra.Mul(u, goalgebra.Mul(l, v)))
	fmt.Println(goalgebra.Pow(u, l))
}
