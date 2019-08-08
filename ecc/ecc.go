package ecc

import (
	"errors"
	"math/big"
)

//
// General ECC Information -
//
// Curve: Curve25519
// Curve Equation: y^2 = x^3 + ax^2 + x (a = 486662)
// Prime Field: 2^255 - 19
// Base Point: x = 9
// Rangeproofs: Bulletproofs
//
// Notes: Since golang does not support operator overloading
// we rely on type methods for all ECC mathematics.

type FieldElement struct {
	num   big.Int
	prime big.Int
}

type Point struct {
	x big.Int
	y big.Int
}

type PrivateKey struct {
	sec   big.Int
	point Point
}

func NewPrivateKey(s *big.Int) (PrivateKey, error) {
	panic("Not Implemented")
}

func NewFieldElement(num big.Int, prime big.Int) (FieldElement, error) {
	f := FieldElement{num, prime}

	if num.Cmp(&prime) >= 0 {
		return f, errors.New("ecc: supplied number is not in field range")
	}

	return f, nil
}

func (f *FieldElement) EQ(e *FieldElement) bool {
	return f.num.Cmp(&e.num) == 0 && f.prime.Cmp(&e.prime) == 0
}

func (f *FieldElement) NE(e *FieldElement) bool {
	return f.num.Cmp(&e.num) != 0 && f.prime.Cmp(&e.prime) != 0
}

func (f *FieldElement) Add(e *FieldElement) error {
	if f.prime.Cmp(&e.prime) != 0 {
		return errors.New("ecc: cannot add two numbers from different fields")
	}

	f.num = *big.NewInt(0).Add(&f.num, &e.num)

	return nil
}

func (f *FieldElement) Sub(e *FieldElement) error {
	if f.prime.Cmp(&e.prime) != 0 {
		return errors.New("ecc: cannot subtract two numbers from different fields")
	}

	f.num = *big.NewInt(0).Sub(&f.num, &e.num)

	return nil
}

func (f *FieldElement) Mul(e *FieldElement) error {
	if f.prime.Cmp(&e.prime) != 0 {
		return errors.New("ecc: cannot multiply two numbers from different fields")
	}

	f.num = *big.NewInt(0).Mod(big.NewInt(0).Mul(&f.num, &e.num), &f.prime)

	return nil
}

func (f *FieldElement) Pow(e *big.Int) error {
	n := big.NewInt(0).Mod(e, big.NewInt(0).Sub(e, big.NewInt(1)))
	f.num = *big.NewInt(0).Exp(&f.num, n, &f.prime)

	return nil
}

func NewPoint(x *big.Int, y *big.Int) (Point, error) {
	p := Point{*x, *y}

	// Curve satisfaction
	y2 := big.NewInt(0).Exp(y, big.NewInt(2))
	x2 := big.NewInt(0).Exp(x, big.NewInt(2))
	x3 := big.NewInt(0).Exp(x, big.NewInt(3))
	a := big.NewInt(486662)
	ax2 := big.NewInt(0).Mul(a, x2)

	if y2.Cmp(big.NewInt(0).Add(big.NewInt(0).Add(x3, ax2), x)) != 0 {
		return p, errors.New("ecc: point is not on the curve")
	}

	return p, nil
}

func (p *Point) EQ(b *Point) bool {
	panic("Not Implemented")
}

func (p *Point) NE(b *Point) bool {
	panic("Not Implemented")
}

func (p *Point) Add(b *Point) error {
	panic("Not Implemented")
}

func (p *Point) Mul(b *Point) error {
	panic("Not Implemented")
}
