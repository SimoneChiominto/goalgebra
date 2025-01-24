package goalgebra

import (
	"fmt"
	"github.com/victorbrun/gosymbol"
	"slices"
	"strings"
)

type real interface {
	String() string
	Simplify() real
	toGoSymbol() gosymbol.Expr
}

func (n integer) Simplify() real {
	return n
}

func (n integer) toGoSymbol() gosymbol.Expr {
	return gosymbol.Const(float64(n.value))
}

func (u fraction) Simplify() real {
	return u.SimplifyRational()
}

func (u fraction) toGoSymbol() gosymbol.Expr {
	return gosymbol.Const(float64(u.num.value) / float64(u.den.value))
}

type logarithm struct {
	real
	arg real
}

func Log(x real) logarithm {
	// TODO: condition x > 0
	return logarithm{arg: x}
}

func (x logarithm) String() string {
	return fmt.Sprintf("log(%s)", x.arg.String())
}
func (x logarithm) Simplify() real {
	switch y := x.arg.(type) {
	case integer:
		if y == Int(1) {
			return Int(0)
		}
	case fraction:
		if y.SimplifyRational() == Int(1) {
			return Int(0)
		}
	case exponential:
		return y.arg.Simplify()
	}
	return Log(x.arg.Simplify())
}

func (x logarithm) toGoSymbol() gosymbol.Expr {
	return gosymbol.Log(x.arg.toGoSymbol())
}

type exponential struct {
	real
	arg real
}

func Exp(x real) exponential {
	// TODO: condition x > 0
	return exponential{arg: x}
}

func (x exponential) toGoSymbol() gosymbol.Expr {
	return gosymbol.Exp(x.arg.toGoSymbol())
}

func (x exponential) String() string {
	return fmt.Sprintf("exp(%s)", x.arg.String())
}

func (x exponential) Simplify() real {
	switch y := x.arg.(type) {
	case integer:
		if y == Int(0) {
			return Int(1)
		}
	case fraction:
		if y.SimplifyRational() == Int(0) {
			return Int(1)
		}
	case logarithm:
		return y.arg.Simplify()
	}
	return Log(x.arg.Simplify())
}

type sum struct {
	real
	operands []real
}

func Add(x real, y real) real {
	var xOperands []real
	var yOperands []real
	switch x := x.(type) {
	case sum:
		xOperands = x.operands
	default:
		xOperands = []real{x}
	}
	switch y := y.(type) {
	case sum:
		yOperands = y.operands
	default:
		yOperands = []real{y}
	}
	operands := slices.Concat(xOperands, yOperands)
	return sum{operands: operands}.Simplify()
}
func (s sum) toGoSymbol() gosymbol.Expr {
	var convOperands []gosymbol.Expr
	for _, operand := range s.operands {
		convOperands = append(convOperands, operand.toGoSymbol())
	}
	return gosymbol.Add(convOperands...)
}
func (s sum) Simplify() real {
	if len(s.operands) == 1 {
		return s.operands[0].Simplify()
	}
	var simplifiedOperands []real
	for _, operand := range s.operands {
		simplifiedOperands = append(simplifiedOperands, operand.Simplify())
	}
	var partialSum real
	switch x := simplifiedOperands[0].(type) {
	case integer:
		switch y := simplifiedOperands[1].(type) {
		case integer:
			partialSum = IntAdd(x, y)
		case fraction:
			partialSum = RatAdd(x, y)
		default:
			partialSum = sum{operands: []real{x, y}}
		}
	case fraction:
		switch y := simplifiedOperands[1].(type) {
		case integer:
			partialSum = RatAdd(x, y)
		case fraction:
			partialSum = RatAdd(x, y)
		default:
			partialSum = sum{operands: []real{x, y}}
		}
	case logarithm:
		switch y := simplifiedOperands[1].(type) {
		case logarithm:
			partialSum = Log(Mul(x, y))
		default:
			partialSum = sum{operands: []real{x, y}}
		}
	}

	return Add(partialSum, sum{operands: simplifiedOperands[2:]})
}

func (s sum) String() string {
	var stringS []string
	for _, operand := range s.operands {
		stringS = append(stringS, operand.String())
	}
	return strings.Join(stringS, " + ")
}

type product struct {
	real
	operands []real
}

func Mul(x real, y real) real {
	var xOperands []real
	var yOperands []real
	switch x := x.(type) {
	case product:
		xOperands = x.operands
	default:
		xOperands = []real{x}
	}
	switch y := y.(type) {
	case product:
		yOperands = y.operands
	default:
		yOperands = []real{y}
	}
	operands := slices.Concat(xOperands, yOperands)
	return product{operands: operands}.Simplify()
}

func (p product) toGoSymbol() gosymbol.Expr {
	var convOperands []gosymbol.Expr
	for _, operand := range p.operands {
		convOperands = append(convOperands, operand.toGoSymbol())
	}
	return gosymbol.Mul(convOperands...)
}
func (p product) Simplify() real {
	var simplifiedOperands []real
	for _, operand := range p.operands {
		simplifiedOperands = append(simplifiedOperands, operand.Simplify())
	}
	return sum{operands: simplifiedOperands}
}

func (p product) String() string {
	var stringS []string
	for _, operand := range p.operands {
		var strOperand string
		switch o := operand.(type) {
		case sum:
			strOperand = fmt.Sprintf("(%s)", o.String())
		default:
			strOperand = o.String()
		}
		stringS = append(stringS, strOperand)
	}
	return strings.Join(stringS, " ")
}

type power struct {
	real
	base     real
	exponent real
}

func Pow(b real, e real) real {
	return power{base: b, exponent: e}.Simplify()
}

func (p power) toGoSymbol() gosymbol.Expr {
	return gosymbol.Pow(
		p.base.toGoSymbol(),
		p.exponent.toGoSymbol(),
	)
}

func (p power) Simplify() real {
	return power{base: p.base.Simplify(), exponent: p.exponent.Simplify()}
}

func (p power) String() string {
	var stringBase string
	switch b := p.base.(type) {
	case sum:
		stringBase = fmt.Sprintf("(%s)", b.String())
	case product:
		stringBase = fmt.Sprintf("(%s)", b.String())
	default:
		stringBase = b.String()
	}
	var stringExp string
	switch e := p.exponent.(type) {
	case sum:
		stringExp = fmt.Sprintf("(%s)", e.String())
	case product:
		stringExp = fmt.Sprintf("(%s)", e.String())
	default:
		stringExp = e.String()
	}
	return fmt.Sprintf("%s^%s", stringBase, stringExp)
}

type undefined struct {
	real
}

func (p undefined) toGoSymbol() gosymbol.Expr {
	return gosymbol.Undefined()
}
func (u undefined) Simplify() real {
	return undefined{}
}

func (p undefined) String() string {
	return "UNDEFINED"
}
