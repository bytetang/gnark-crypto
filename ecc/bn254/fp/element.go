// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package fp

// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular,
// there is no security guarantees such as constant time implementation
// or side-channel attack resistance
// /!\ WARNING /!\

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
	"math/bits"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// Element represents a field element stored on 4 words (uint64)
// Element are assumed to be in Montgomery form in all methods
// field modulus q =
//
// 21888242871839275222246405745257275088696311157297823662689037894645226208583
type Element [4]uint64

// Limbs number of 64 bits words needed to represent Element
const Limbs = 4

// Bits number bits needed to represent Element
const Bits = 254

// Bytes number bytes needed to represent Element
const Bytes = Limbs * 8

// field modulus stored as big.Int
var _modulus big.Int

// Modulus returns q as a big.Int
// q =
//
// 21888242871839275222246405745257275088696311157297823662689037894645226208583
func Modulus() *big.Int {
	return new(big.Int).Set(&_modulus)
}

// q (modulus)
const qElementWord0 uint64 = 4332616871279656263
const qElementWord1 uint64 = 10917124144477883021
const qElementWord2 uint64 = 13281191951274694749
const qElementWord3 uint64 = 3486998266802970665

var qElement = Element{
	qElementWord0,
	qElementWord1,
	qElementWord2,
	qElementWord3,
}

// Used for Montgomery reduction. (qInvNeg) q + r'.r = 1, i.e., qInvNeg = - q⁻¹ mod r
const qInvNegLsw uint64 = 9786893198990664585

// rSquare
var rSquare = Element{
	17522657719365597833,
	13107472804851548667,
	5164255478447964150,
	493319470278259999,
}

var bigIntPool = sync.Pool{
	New: func() interface{} {
		return new(big.Int)
	},
}

func init() {
	_modulus.SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)
}

// NewElement returns a new Element from a uint64 value
//
// it is equivalent to
// 		var v NewElement
// 		v.SetUint64(...)
func NewElement(v uint64) Element {
	z := Element{v}
	z.Mul(&z, &rSquare)
	return z
}

// SetUint64 z = v, sets z LSB to v (non-Montgomery form) and convert z to Montgomery form
func (z *Element) SetUint64(v uint64) *Element {
	*z = Element{v}
	return z.Mul(z, &rSquare) // z.ToMont()
}

// Set z = x
func (z *Element) Set(x *Element) *Element {
	z[0] = x[0]
	z[1] = x[1]
	z[2] = x[2]
	z[3] = x[3]
	return z
}

// SetInterface converts provided interface into Element
// returns an error if provided type is not supported
// supported types: Element, *Element, uint64, int, string (interpreted as base10 integer),
// *big.Int, big.Int, []byte
func (z *Element) SetInterface(i1 interface{}) (*Element, error) {
	switch c1 := i1.(type) {
	case Element:
		return z.Set(&c1), nil
	case *Element:
		return z.Set(c1), nil
	case uint64:
		return z.SetUint64(c1), nil
	case int:
		return z.SetString(strconv.Itoa(c1)), nil
	case string:
		return z.SetString(c1), nil
	case *big.Int:
		return z.SetBigInt(c1), nil
	case big.Int:
		return z.SetBigInt(&c1), nil
	case []byte:
		return z.SetBytes(c1), nil
	default:
		return nil, errors.New("can't set fp.Element from type " + reflect.TypeOf(i1).String())
	}
}

// SetZero z = 0
func (z *Element) SetZero() *Element {
	z[0] = 0
	z[1] = 0
	z[2] = 0
	z[3] = 0
	return z
}

// SetOne z = 1 (in Montgomery form)
func (z *Element) SetOne() *Element {
	z[0] = 15230403791020821917
	z[1] = 754611498739239741
	z[2] = 7381016538464732716
	z[3] = 1011752739694698287
	return z
}

// Div z = x*y^-1 mod q
func (z *Element) Div(x, y *Element) *Element {
	var yInv Element
	yInv.Inverse(y)
	z.Mul(x, &yInv)
	return z
}

// Bit returns the i'th bit, with lsb == bit 0.
// It is the responsability of the caller to convert from Montgomery to Regular form if needed
func (z *Element) Bit(i uint64) uint64 {
	j := i / 64
	if j >= 4 {
		return 0
	}
	return uint64(z[j] >> (i % 64) & 1)
}

// Equal returns z == x
func (z *Element) Equal(x *Element) bool {
	return (z[3] == x[3]) && (z[2] == x[2]) && (z[1] == x[1]) && (z[0] == x[0])
}

// IsZero returns z == 0
func (z *Element) IsZero() bool {
	return (z[3] | z[2] | z[1] | z[0]) == 0
}

// IsUint64 returns true if z[0] ⩾ 0 and all other words are 0
func (z *Element) IsUint64() bool {
	return (z[3] | z[2] | z[1]) == 0
}

// Cmp compares (lexicographic order) z and x and returns:
//
//   -1 if z <  x
//    0 if z == x
//   +1 if z >  x
//
func (z *Element) Cmp(x *Element) int {
	_z := *z
	_x := *x
	_z.FromMont()
	_x.FromMont()
	if _z[3] > _x[3] {
		return 1
	} else if _z[3] < _x[3] {
		return -1
	}
	if _z[2] > _x[2] {
		return 1
	} else if _z[2] < _x[2] {
		return -1
	}
	if _z[1] > _x[1] {
		return 1
	} else if _z[1] < _x[1] {
		return -1
	}
	if _z[0] > _x[0] {
		return 1
	} else if _z[0] < _x[0] {
		return -1
	}
	return 0
}

// LexicographicallyLargest returns true if this element is strictly lexicographically
// larger than its negation, false otherwise
func (z *Element) LexicographicallyLargest() bool {
	// adapted from github.com/zkcrypto/bls12_381
	// we check if the element is larger than (q-1) / 2
	// if z - (((q -1) / 2) + 1) have no underflow, then z > (q-1) / 2

	_z := *z
	_z.FromMont()

	var b uint64
	_, b = bits.Sub64(_z[0], 11389680472494603940, 0)
	_, b = bits.Sub64(_z[1], 14681934109093717318, b)
	_, b = bits.Sub64(_z[2], 15863968012492123182, b)
	_, b = bits.Sub64(_z[3], 1743499133401485332, b)

	return b == 0
}

// SetRandom sets z to a random element < q
func (z *Element) SetRandom() (*Element, error) {
	var bytes [32]byte
	if _, err := io.ReadFull(rand.Reader, bytes[:]); err != nil {
		return nil, err
	}
	z[0] = binary.BigEndian.Uint64(bytes[0:8])
	z[1] = binary.BigEndian.Uint64(bytes[8:16])
	z[2] = binary.BigEndian.Uint64(bytes[16:24])
	z[3] = binary.BigEndian.Uint64(bytes[24:32])
	z[3] %= 3486998266802970665

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}

	return z, nil
}

// One returns 1 (in montgommery form)
func One() Element {
	var one Element
	one.SetOne()
	return one
}

// Halve sets z to z / 2 (mod p)
func (z *Element) Halve() {
	if z[0]&1 == 1 {
		var carry uint64

		// z = z + q
		z[0], carry = bits.Add64(z[0], 4332616871279656263, 0)
		z[1], carry = bits.Add64(z[1], 10917124144477883021, carry)
		z[2], carry = bits.Add64(z[2], 13281191951274694749, carry)
		z[3], _ = bits.Add64(z[3], 3486998266802970665, carry)

	}

	// z = z >> 1

	z[0] = z[0]>>1 | z[1]<<63
	z[1] = z[1]>>1 | z[2]<<63
	z[2] = z[2]>>1 | z[3]<<63
	z[3] >>= 1

}

// API with assembly impl

// Mul z = x * y mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Mul(x, y *Element) *Element {
	mul(z, x, y)
	return z
}

// Square z = x * x mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Square(x *Element) *Element {
	mul(z, x, x)
	return z
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func (z *Element) FromMont() *Element {
	fromMont(z)
	return z
}

// Add z = x + y mod q
func (z *Element) Add(x, y *Element) *Element {
	add(z, x, y)
	return z
}

// Double z = x + x mod q, aka Lsh 1
func (z *Element) Double(x *Element) *Element {
	double(z, x)
	return z
}

// Sub  z = x - y mod q
func (z *Element) Sub(x, y *Element) *Element {
	sub(z, x, y)
	return z
}

// Neg z = q - x
func (z *Element) Neg(x *Element) *Element {
	neg(z, x)
	return z
}

// Generic (no ADX instructions, no AMD64) versions of multiplication and squaring algorithms

func _mulGeneric(z, x, y *Element) {

	var t [4]uint64
	var c [3]uint64
	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, y[0])
		m := c[0] * 9786893198990664585
		c[2] = madd0(m, 4332616871279656263, c[0])
		c[1], c[0] = madd1(v, y[1], c[1])
		c[2], t[0] = madd2(m, 10917124144477883021, c[2], c[0])
		c[1], c[0] = madd1(v, y[2], c[1])
		c[2], t[1] = madd2(m, 13281191951274694749, c[2], c[0])
		c[1], c[0] = madd1(v, y[3], c[1])
		t[3], t[2] = madd3(m, 3486998266802970665, c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9786893198990664585
		c[2] = madd0(m, 4332616871279656263, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 10917124144477883021, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 13281191951274694749, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, 3486998266802970665, c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9786893198990664585
		c[2] = madd0(m, 4332616871279656263, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 10917124144477883021, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 13281191951274694749, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, 3486998266802970665, c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9786893198990664585
		c[2] = madd0(m, 4332616871279656263, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], z[0] = madd2(m, 10917124144477883021, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], z[1] = madd2(m, 13281191951274694749, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		z[3], z[2] = madd3(m, 3486998266802970665, c[0], c[2], c[1])
	}

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}
}

func _mulWGeneric(z, x *Element, y uint64) {

	var t [4]uint64
	{
		// round 0
		c1, c0 := bits.Mul64(y, x[0])
		m := c0 * 9786893198990664585
		c2 := madd0(m, 4332616871279656263, c0)
		c1, c0 = madd1(y, x[1], c1)
		c2, t[0] = madd2(m, 10917124144477883021, c2, c0)
		c1, c0 = madd1(y, x[2], c1)
		c2, t[1] = madd2(m, 13281191951274694749, c2, c0)
		c1, c0 = madd1(y, x[3], c1)
		t[3], t[2] = madd3(m, 3486998266802970665, c0, c2, c1)
	}
	{
		// round 1
		m := t[0] * 9786893198990664585
		c2 := madd0(m, 4332616871279656263, t[0])
		c2, t[0] = madd2(m, 10917124144477883021, c2, t[1])
		c2, t[1] = madd2(m, 13281191951274694749, c2, t[2])
		t[3], t[2] = madd2(m, 3486998266802970665, t[3], c2)
	}
	{
		// round 2
		m := t[0] * 9786893198990664585
		c2 := madd0(m, 4332616871279656263, t[0])
		c2, t[0] = madd2(m, 10917124144477883021, c2, t[1])
		c2, t[1] = madd2(m, 13281191951274694749, c2, t[2])
		t[3], t[2] = madd2(m, 3486998266802970665, t[3], c2)
	}
	{
		// round 3
		m := t[0] * 9786893198990664585
		c2 := madd0(m, 4332616871279656263, t[0])
		c2, z[0] = madd2(m, 10917124144477883021, c2, t[1])
		c2, z[1] = madd2(m, 13281191951274694749, c2, t[2])
		z[3], z[2] = madd2(m, 3486998266802970665, t[3], c2)
	}

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}
}

func _fromMontGeneric(z *Element) {
	// the following lines implement z = z * 1
	// with a modified CIOS montgomery multiplication
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9786893198990664585
		C := madd0(m, 4332616871279656263, z[0])
		C, z[0] = madd2(m, 10917124144477883021, z[1], C)
		C, z[1] = madd2(m, 13281191951274694749, z[2], C)
		C, z[2] = madd2(m, 3486998266802970665, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9786893198990664585
		C := madd0(m, 4332616871279656263, z[0])
		C, z[0] = madd2(m, 10917124144477883021, z[1], C)
		C, z[1] = madd2(m, 13281191951274694749, z[2], C)
		C, z[2] = madd2(m, 3486998266802970665, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9786893198990664585
		C := madd0(m, 4332616871279656263, z[0])
		C, z[0] = madd2(m, 10917124144477883021, z[1], C)
		C, z[1] = madd2(m, 13281191951274694749, z[2], C)
		C, z[2] = madd2(m, 3486998266802970665, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9786893198990664585
		C := madd0(m, 4332616871279656263, z[0])
		C, z[0] = madd2(m, 10917124144477883021, z[1], C)
		C, z[1] = madd2(m, 13281191951274694749, z[2], C)
		C, z[2] = madd2(m, 3486998266802970665, z[3], C)
		z[3] = C
	}

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}
}

func _addGeneric(z, x, y *Element) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], _ = bits.Add64(x[3], y[3], carry)

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}
}

func _doubleGeneric(z, x *Element) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], x[0], 0)
	z[1], carry = bits.Add64(x[1], x[1], carry)
	z[2], carry = bits.Add64(x[2], x[2], carry)
	z[3], _ = bits.Add64(x[3], x[3], carry)

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}
}

func _subGeneric(z, x, y *Element) {
	var b uint64
	z[0], b = bits.Sub64(x[0], y[0], 0)
	z[1], b = bits.Sub64(x[1], y[1], b)
	z[2], b = bits.Sub64(x[2], y[2], b)
	z[3], b = bits.Sub64(x[3], y[3], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 4332616871279656263, 0)
		z[1], c = bits.Add64(z[1], 10917124144477883021, c)
		z[2], c = bits.Add64(z[2], 13281191951274694749, c)
		z[3], _ = bits.Add64(z[3], 3486998266802970665, c)
	}
}

func _negGeneric(z, x *Element) {
	if x.IsZero() {
		z.SetZero()
		return
	}
	var borrow uint64
	z[0], borrow = bits.Sub64(4332616871279656263, x[0], 0)
	z[1], borrow = bits.Sub64(10917124144477883021, x[1], borrow)
	z[2], borrow = bits.Sub64(13281191951274694749, x[2], borrow)
	z[3], _ = bits.Sub64(3486998266802970665, x[3], borrow)
}

func _reduceGeneric(z *Element) {

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}
}

func mulByConstant(z *Element, c uint8) {
	switch c {
	case 0:
		z.SetZero()
		return
	case 1:
		return
	case 2:
		z.Double(z)
		return
	case 3:
		_z := *z
		z.Double(z).Add(z, &_z)
	case 5:
		_z := *z
		z.Double(z).Double(z).Add(z, &_z)
	default:
		var y Element
		y.SetUint64(uint64(c))
		z.Mul(z, &y)
	}
}

// BatchInvert returns a new slice with every element inverted.
// Uses Montgomery batch inversion trick
func BatchInvert(a []Element) []Element {
	res := make([]Element, len(a))
	if len(a) == 0 {
		return res
	}

	zeroes := make([]bool, len(a))
	accumulator := One()

	for i := 0; i < len(a); i++ {
		if a[i].IsZero() {
			zeroes[i] = true
			continue
		}
		res[i] = accumulator
		accumulator.Mul(&accumulator, &a[i])
	}

	accumulator.Inverse(&accumulator)

	for i := len(a) - 1; i >= 0; i-- {
		if zeroes[i] {
			continue
		}
		res[i].Mul(&res[i], &accumulator)
		accumulator.Mul(&accumulator, &a[i])
	}

	return res
}

func _butterflyGeneric(a, b *Element) {
	t := *a
	a.Add(a, b)
	b.Sub(&t, b)
}

// BitLen returns the minimum number of bits needed to represent z
// returns 0 if z == 0
func (z *Element) BitLen() int {
	if z[3] != 0 {
		return 192 + bits.Len64(z[3])
	}
	if z[2] != 0 {
		return 128 + bits.Len64(z[2])
	}
	if z[1] != 0 {
		return 64 + bits.Len64(z[1])
	}
	return bits.Len64(z[0])
}

// Exp z = x^exponent mod q
func (z *Element) Exp(x Element, exponent *big.Int) *Element {
	var bZero big.Int
	if exponent.Cmp(&bZero) == 0 {
		return z.SetOne()
	}

	z.Set(&x)

	for i := exponent.BitLen() - 2; i >= 0; i-- {
		z.Square(z)
		if exponent.Bit(i) == 1 {
			z.Mul(z, &x)
		}
	}

	return z
}

// ToMont converts z to Montgomery form
// sets and returns z = z * r²
func (z *Element) ToMont() *Element {
	return z.Mul(z, &rSquare)
}

// ToRegular returns z in regular form (doesn't mutate z)
func (z Element) ToRegular() Element {
	return *z.FromMont()
}

// String returns the decimal representation of z as generated by
// z.Text(10).
func (z *Element) String() string {
	return z.Text(10)
}

// Text returns the string representation of z in the given base.
// Base must be between 2 and 36, inclusive. The result uses the
// lower-case letters 'a' to 'z' for digit values 10 to 35.
// No prefix (such as "0x") is added to the string. If z is a nil
// pointer it returns "<nil>".
// If base == 10 and -z fits in a uint64 prefix "-" is added to the string.
func (z *Element) Text(base int) string {
	if base < 2 || base > 36 {
		panic("invalid base")
	}
	if z == nil {
		return "<nil>"
	}
	zz := *z
	zz.FromMont()
	if zz.IsUint64() {
		return strconv.FormatUint(zz[0], base)
	} else if base == 10 {
		var zzNeg Element
		zzNeg.Neg(z)
		zzNeg.FromMont()
		if zzNeg.IsUint64() {
			return "-" + strconv.FormatUint(zzNeg[0], base)
		}
	}
	vv := bigIntPool.Get().(*big.Int)
	r := zz.ToBigInt(vv).Text(base)
	bigIntPool.Put(vv)
	return r
}

// ToBigInt returns z as a big.Int in Montgomery form
func (z *Element) ToBigInt(res *big.Int) *big.Int {
	var b [Limbs * 8]byte
	binary.BigEndian.PutUint64(b[24:32], z[0])
	binary.BigEndian.PutUint64(b[16:24], z[1])
	binary.BigEndian.PutUint64(b[8:16], z[2])
	binary.BigEndian.PutUint64(b[0:8], z[3])

	return res.SetBytes(b[:])
}

// ToBigIntRegular returns z as a big.Int in regular form
func (z Element) ToBigIntRegular(res *big.Int) *big.Int {
	z.FromMont()
	return z.ToBigInt(res)
}

// Bytes returns the regular (non montgomery) value
// of z as a big-endian byte array.
func (z *Element) Bytes() (res [Limbs * 8]byte) {
	_z := z.ToRegular()
	binary.BigEndian.PutUint64(res[24:32], _z[0])
	binary.BigEndian.PutUint64(res[16:24], _z[1])
	binary.BigEndian.PutUint64(res[8:16], _z[2])
	binary.BigEndian.PutUint64(res[0:8], _z[3])

	return
}

// Marshal returns the regular (non montgomery) value
// of z as a big-endian byte slice.
func (z *Element) Marshal() []byte {
	b := z.Bytes()
	return b[:]
}

// SetBytes interprets e as the bytes of a big-endian unsigned integer,
// sets z to that value (in Montgomery form), and returns z.
func (z *Element) SetBytes(e []byte) *Element {
	// get a big int from our pool
	vv := bigIntPool.Get().(*big.Int)
	vv.SetBytes(e)

	// set big int
	z.SetBigInt(vv)

	// put temporary object back in pool
	bigIntPool.Put(vv)

	return z
}

// SetBigInt sets z to v (regular form) and returns z in Montgomery form
func (z *Element) SetBigInt(v *big.Int) *Element {
	z.SetZero()

	var zero big.Int

	// fast path
	c := v.Cmp(&_modulus)
	if c == 0 {
		// v == 0
		return z
	} else if c != 1 && v.Cmp(&zero) != -1 {
		// 0 < v < q
		return z.setBigInt(v)
	}

	// get temporary big int from the pool
	vv := bigIntPool.Get().(*big.Int)

	// copy input + modular reduction
	vv.Set(v)
	vv.Mod(v, &_modulus)

	// set big int byte value
	z.setBigInt(vv)

	// release object into pool
	bigIntPool.Put(vv)
	return z
}

// setBigInt assumes 0 ⩽ v < q
func (z *Element) setBigInt(v *big.Int) *Element {
	vBits := v.Bits()

	if bits.UintSize == 64 {
		for i := 0; i < len(vBits); i++ {
			z[i] = uint64(vBits[i])
		}
	} else {
		for i := 0; i < len(vBits); i++ {
			if i%2 == 0 {
				z[i/2] = uint64(vBits[i])
			} else {
				z[i/2] |= uint64(vBits[i]) << 32
			}
		}
	}

	return z.ToMont()
}

// SetString creates a big.Int with number and calls SetBigInt on z
//
// The number prefix determines the actual base: A prefix of
// ''0b'' or ''0B'' selects base 2, ''0'', ''0o'' or ''0O'' selects base 8,
// and ''0x'' or ''0X'' selects base 16. Otherwise, the selected base is 10
// and no prefix is accepted.
//
// For base 16, lower and upper case letters are considered the same:
// The letters 'a' to 'f' and 'A' to 'F' represent digit values 10 to 15.
//
// An underscore character ''_'' may appear between a base
// prefix and an adjacent digit, and between successive digits; such
// underscores do not change the value of the number.
// Incorrect placement of underscores is reported as a panic if there
// are no other errors.
//
func (z *Element) SetString(number string) *Element {
	// get temporary big int from the pool
	vv := bigIntPool.Get().(*big.Int)

	if _, ok := vv.SetString(number, 0); !ok {
		panic("Element.SetString failed -> can't parse number into a big.Int " + number)
	}

	z.SetBigInt(vv)

	// release object into pool
	bigIntPool.Put(vv)

	return z
}

// MarshalJSON returns json encoding of z (z.Text(10))
// If z == nil, returns null
func (z *Element) MarshalJSON() ([]byte, error) {
	if z == nil {
		return []byte("null"), nil
	}
	const maxSafeBound = 15 // we encode it as number if it's small
	s := z.Text(10)
	if len(s) <= maxSafeBound {
		return []byte(s), nil
	}
	var sbb strings.Builder
	sbb.WriteByte('"')
	sbb.WriteString(s)
	sbb.WriteByte('"')
	return []byte(sbb.String()), nil
}

// UnmarshalJSON accepts numbers and strings as input
// See Element.SetString for valid prefixes (0x, 0b, ...)
func (z *Element) UnmarshalJSON(data []byte) error {
	s := string(data)
	if len(s) > Bits*3 {
		return errors.New("value too large (max = Element.Bits * 3)")
	}

	// we accept numbers and strings, remove leading and trailing quotes if any
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}

	// get temporary big int from the pool
	vv := bigIntPool.Get().(*big.Int)

	if _, ok := vv.SetString(s, 0); !ok {
		return errors.New("can't parse into a big.Int: " + s)
	}

	z.SetBigInt(vv)

	// release object into pool
	bigIntPool.Put(vv)
	return nil
}

// Legendre returns the Legendre symbol of z (either +1, -1, or 0.)
func (z *Element) Legendre() int {
	var l Element
	// z^((q-1)/2)
	l.expByLegendreExp(*z)

	if l.IsZero() {
		return 0
	}

	// if l == 1
	if (l[3] == 1011752739694698287) && (l[2] == 7381016538464732716) && (l[1] == 754611498739239741) && (l[0] == 15230403791020821917) {
		return 1
	}
	return -1
}

// Sqrt z = √x mod q
// if the square root doesn't exist (x is not a square mod q)
// Sqrt leaves z unchanged and returns nil
func (z *Element) Sqrt(x *Element) *Element {
	// q ≡ 3 (mod 4)
	// using  z ≡ ± x^((p+1)/4) (mod q)
	var y, square Element
	y.expBySqrtExp(*x)
	// as we didn't compute the legendre symbol, ensure we found y such that y * y = x
	square.Square(&y)
	if square.Equal(x) {
		return z.Set(&y)
	}
	return nil
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

const updateFactorsConversionBias int64 = 0x7fffffff7fffffff // (2³¹ - 1)(2³² + 1)
const updateFactorIdentityMatrixRow0 = 1
const updateFactorIdentityMatrixRow1 = 1 << 32

func updateFactorsDecompose(c int64) (int64, int64) {
	c += updateFactorsConversionBias
	const low32BitsFilter int64 = 0xFFFFFFFF
	f := c&low32BitsFilter - 0x7FFFFFFF
	g := c>>32&low32BitsFilter - 0x7FFFFFFF
	return f, g
}

const k = 32 // word size / 2
const signBitSelector = uint64(1) << 63
const approxLowBitsN = k - 1
const approxHighBitsN = k + 1
const inversionCorrectionFactorWord0 = 11111708840330028223
const inversionCorrectionFactorWord1 = 3098618286181893933
const inversionCorrectionFactorWord2 = 756602578711705709
const inversionCorrectionFactorWord3 = 1041752015607019851

const invIterationsN = 18

// Inverse z = x⁻¹ mod q
// Implements "Optimized Binary GCD for Modular Inversion"
// https://github.com/pornin/bingcd/blob/main/doc/bingcd.pdf
func (z *Element) Inverse(x *Element) *Element {
	if x.IsZero() {
		z.SetZero()
		return z
	}

	a := *x
	b := Element{
		qElementWord0,
		qElementWord1,
		qElementWord2,
		qElementWord3,
	} // b := q

	u := Element{1}

	// Update factors: we get [u; v]:= [f0 g0; f1 g1] [u; v]
	// c_i = f_i + 2³¹ - 1 + 2³² * (g_i + 2³¹ - 1)
	var c0, c1 int64

	// Saved update factors to reduce the number of field multiplications
	var pf0, pf1, pg0, pg1 int64

	var i uint

	var v, s Element

	// Since u,v are updated every other iteration, we must make sure we terminate after evenly many iterations
	// This also lets us get away with half as many updates to u,v
	// To make this constant-time-ish, replace the condition with i < invIterationsN
	for i = 0; i&1 == 1 || !a.IsZero(); i++ {
		n := max(a.BitLen(), b.BitLen())
		aApprox, bApprox := approximate(&a, n), approximate(&b, n)

		// After 0 iterations, we have f₀ ≤ 2⁰ and f₁ < 2⁰
		// f0, g0, f1, g1 = 1, 0, 0, 1
		c0, c1 = updateFactorIdentityMatrixRow0, updateFactorIdentityMatrixRow1

		for j := 0; j < approxLowBitsN; j++ {

			if aApprox&1 == 0 {
				aApprox /= 2
			} else {
				s, borrow := bits.Sub64(aApprox, bApprox, 0)
				if borrow == 1 {
					s = bApprox - aApprox
					bApprox = aApprox
					c0, c1 = c1, c0
				}

				aApprox = s / 2
				c0 = c0 - c1

				// Now |f₀| < 2ʲ + 2ʲ = 2ʲ⁺¹
				// |f₁| ≤ 2ʲ still
			}

			c1 *= 2
			// |f₁| ≤ 2ʲ⁺¹
		}

		s = a

		var g0 int64
		// from this point on c0 aliases for f0
		c0, g0 = updateFactorsDecompose(c0)
		aHi := a.linearCombNonModular(&s, c0, &b, g0)
		if aHi&signBitSelector != 0 {
			// if aHi < 0
			c0, g0 = -c0, -g0
			aHi = a.neg(&a, aHi)
		}
		// right-shift a by k-1 bits
		a[0] = (a[0] >> approxLowBitsN) | ((a[1]) << approxHighBitsN)
		a[1] = (a[1] >> approxLowBitsN) | ((a[2]) << approxHighBitsN)
		a[2] = (a[2] >> approxLowBitsN) | ((a[3]) << approxHighBitsN)
		a[3] = (a[3] >> approxLowBitsN) | (aHi << approxHighBitsN)

		var f1 int64
		// from this point on c1 aliases for g0
		f1, c1 = updateFactorsDecompose(c1)
		bHi := b.linearCombNonModular(&s, f1, &b, c1)
		if bHi&signBitSelector != 0 {
			// if bHi < 0
			f1, c1 = -f1, -c1
			bHi = b.neg(&b, bHi)
		}
		// right-shift b by k-1 bits
		b[0] = (b[0] >> approxLowBitsN) | ((b[1]) << approxHighBitsN)
		b[1] = (b[1] >> approxLowBitsN) | ((b[2]) << approxHighBitsN)
		b[2] = (b[2] >> approxLowBitsN) | ((b[3]) << approxHighBitsN)
		b[3] = (b[3] >> approxLowBitsN) | (bHi << approxHighBitsN)

		if i&1 == 1 {
			// Combine current update factors with previously stored ones
			// [f₀, g₀; f₁, g₁] ← [f₀, g₀; f₁, g₀] [pf₀, pg₀; pf₀, pg₀]
			// We have |f₀|, |g₀|, |pf₀|, |pf₁| ≤ 2ᵏ⁻¹, and that |pf_i| < 2ᵏ⁻¹ for i ∈ {0, 1}
			// Then for the new value we get |f₀| < 2ᵏ⁻¹ × 2ᵏ⁻¹ + 2ᵏ⁻¹ × 2ᵏ⁻¹ = 2²ᵏ⁻¹
			// Which leaves us with an extra bit for the sign

			// c0 aliases f0, c1 aliases g1
			c0, g0, f1, c1 = c0*pf0+g0*pf1,
				c0*pg0+g0*pg1,
				f1*pf0+c1*pf1,
				f1*pg0+c1*pg1

			s = u
			u.linearCombSosSigned(&u, c0, &v, g0)
			v.linearCombSosSigned(&s, f1, &v, c1)

		} else {
			// Save update factors
			pf0, pg0, pf1, pg1 = c0, g0, f1, c1
		}
	}

	// For every iteration that we miss, v is not being multiplied by 2²ᵏ⁻²
	const pSq int64 = 1 << (2 * (k - 1))
	// If the function is constant-time ish, this loop will not run (probably no need to take it out explicitly)
	for ; i < invIterationsN; i += 2 {
		v.mulWSigned(&v, pSq)
	}

	z.Mul(&v, &Element{
		inversionCorrectionFactorWord0,
		inversionCorrectionFactorWord1,
		inversionCorrectionFactorWord2,
		inversionCorrectionFactorWord3,
	})
	return z
}

// approximate a big number x into a single 64 bit word using its uppermost and lowermost bits
// if x fits in a word as is, no approximation necessary
func approximate(x *Element, nBits int) uint64 {

	if nBits <= 64 {
		return x[0]
	}

	const mask = (uint64(1) << (k - 1)) - 1 // k-1 ones
	lo := mask & x[0]

	hiWordIndex := (nBits - 1) / 64

	hiWordBitsAvailable := nBits - hiWordIndex*64
	hiWordBitsUsed := min(hiWordBitsAvailable, approxHighBitsN)

	mask_ := uint64(^((1 << (hiWordBitsAvailable - hiWordBitsUsed)) - 1))
	hi := (x[hiWordIndex] & mask_) << (64 - hiWordBitsAvailable)

	mask_ = ^(1<<(approxLowBitsN+hiWordBitsUsed) - 1)
	mid := (mask_ & x[hiWordIndex-1]) >> hiWordBitsUsed

	return lo | mid | hi
}

func (z *Element) linearCombSosSigned(x *Element, xC int64, y *Element, yC int64) {
	hi := z.linearCombNonModular(x, xC, y, yC)
	z.montReduceSigned(z, hi)
}

// montReduceSigned SOS algorithm; xHi must be at most 63 bits long. Last bit of xHi may be used as a sign bit
func (z *Element) montReduceSigned(x *Element, xHi uint64) {

	const signBitRemover = ^signBitSelector
	neg := xHi&signBitSelector != 0
	// the SOS implementation requires that most significant bit is 0
	// Let X be xHi*r + x
	// note that if X is negative we would have initially stored it as 2⁶⁴ r + X
	xHi &= signBitRemover
	// with this a negative X is now represented as 2⁶³ r + X

	var t [2*Limbs - 1]uint64
	var C uint64

	m := x[0] * qInvNegLsw

	C = madd0(m, qElementWord0, x[0])
	C, t[1] = madd2(m, qElementWord1, x[1], C)
	C, t[2] = madd2(m, qElementWord2, x[2], C)
	C, t[3] = madd2(m, qElementWord3, x[3], C)

	// the high word of m * qElement[3] is at most 62 bits
	// x[3] + C is at most 65 bits (high word at most 1 bit)
	// Thus the resulting C will be at most 63 bits
	t[4] = xHi + C
	// xHi and C are 63 bits, therefore no overflow

	{
		const i = 1
		m = t[i] * qInvNegLsw

		C = madd0(m, qElementWord0, t[i+0])
		C, t[i+1] = madd2(m, qElementWord1, t[i+1], C)
		C, t[i+2] = madd2(m, qElementWord2, t[i+2], C)
		C, t[i+3] = madd2(m, qElementWord3, t[i+3], C)

		t[i+Limbs] += C
	}
	{
		const i = 2
		m = t[i] * qInvNegLsw

		C = madd0(m, qElementWord0, t[i+0])
		C, t[i+1] = madd2(m, qElementWord1, t[i+1], C)
		C, t[i+2] = madd2(m, qElementWord2, t[i+2], C)
		C, t[i+3] = madd2(m, qElementWord3, t[i+3], C)

		t[i+Limbs] += C
	}
	{
		const i = 3
		m := t[i] * qInvNegLsw

		C = madd0(m, qElementWord0, t[i+0])
		C, z[0] = madd2(m, qElementWord1, t[i+1], C)
		C, z[1] = madd2(m, qElementWord2, t[i+2], C)
		z[3], z[2] = madd2(m, qElementWord3, t[i+3], C)
	}

	// if z > q → z -= q
	// note: this is NOT constant time
	if !(z[3] < 3486998266802970665 || (z[3] == 3486998266802970665 && (z[2] < 13281191951274694749 || (z[2] == 13281191951274694749 && (z[1] < 10917124144477883021 || (z[1] == 10917124144477883021 && (z[0] < 4332616871279656263))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 4332616871279656263, 0)
		z[1], b = bits.Sub64(z[1], 10917124144477883021, b)
		z[2], b = bits.Sub64(z[2], 13281191951274694749, b)
		z[3], _ = bits.Sub64(z[3], 3486998266802970665, b)
	}
	if neg {
		// We have computed ( 2⁶³ r + X ) r⁻¹ = 2⁶³ + X r⁻¹ instead
		var b uint64
		z[0], b = bits.Sub64(z[0], signBitSelector, 0)
		z[1], b = bits.Sub64(z[1], 0, b)
		z[2], b = bits.Sub64(z[2], 0, b)
		z[3], b = bits.Sub64(z[3], 0, b)

		// Occurs iff x == 0 && xHi < 0, i.e. X = rX' for -2⁶³ ≤ X' < 0
		if b != 0 {
			// z[3] = -1
			// negative: add q
			const neg1 = 0xFFFFFFFFFFFFFFFF

			b = 0
			z[0], b = bits.Add64(z[0], qElementWord0, b)
			z[1], b = bits.Add64(z[1], qElementWord1, b)
			z[2], b = bits.Add64(z[2], qElementWord2, b)
			z[3], _ = bits.Add64(neg1, qElementWord3, b)
		}
	}
}

// mulWSigned mul word signed (w/ montgomery reduction)
func (z *Element) mulWSigned(x *Element, y int64) {
	m := y >> 63
	_mulWGeneric(z, x, uint64((y^m)-m))
	// multiply by abs(y)
	if y < 0 {
		z.Neg(z)
	}
}

func (z *Element) neg(x *Element, xHi uint64) uint64 {
	var b uint64

	z[0], b = bits.Sub64(0, x[0], 0)
	z[1], b = bits.Sub64(0, x[1], b)
	z[2], b = bits.Sub64(0, x[2], b)
	z[3], b = bits.Sub64(0, x[3], b)
	xHi, _ = bits.Sub64(0, xHi, b)

	return xHi
}

// regular multiplication by one word regular (non montgomery)
// Fewer additions than the branch-free for positive y. Could be faster on some architectures
func (z *Element) mulWRegular(x *Element, y int64) uint64 {

	// w := abs(y)
	m := y >> 63
	w := uint64((y ^ m) - m)

	var c uint64
	c, z[0] = bits.Mul64(x[0], w)
	c, z[1] = madd1(x[1], w, c)
	c, z[2] = madd1(x[2], w, c)
	c, z[3] = madd1(x[3], w, c)

	if y < 0 {
		c = z.neg(z, c)
	}

	return c
}

/*
Removed: seems slower
// mulWRegular branch-free regular multiplication by one word (non montgomery)
func (z *Element) mulWRegularBf(x *Element, y int64) uint64 {

	w := uint64(y)
	allNeg := uint64(y >> 63)	// -1 if y < 0, 0 o.w

	// s[0], s[1] so results are not stored immediately in z.
	// x[i] will be needed in the i+1 th iteration. We don't want to overwrite it in case x = z
	var s [2]uint64
	var h [2]uint64

	h[0], s[0] = bits.Mul64(x[0], w)

	c := uint64(0)
	b := uint64(0)

		{
			const curI = 1 % 2
			const prevI = 1 - curI
			const iMinusOne = 1 - 1

			h[curI], s[curI] = bits.Mul64(x[1], w)
			s[curI], c = bits.Add64(s[curI], h[prevI], c)
			s[curI], b = bits.Sub64(s[curI], allNeg & x[iMinusOne], b)
			z[iMinusOne] = s[prevI]
		}

		{
			const curI = 2 % 2
			const prevI = 1 - curI
			const iMinusOne = 2 - 1

			h[curI], s[curI] = bits.Mul64(x[2], w)
			s[curI], c = bits.Add64(s[curI], h[prevI], c)
			s[curI], b = bits.Sub64(s[curI], allNeg & x[iMinusOne], b)
			z[iMinusOne] = s[prevI]
		}

		{
			const curI = 3 % 2
			const prevI = 1 - curI
			const iMinusOne = 3 - 1

			h[curI], s[curI] = bits.Mul64(x[3], w)
			s[curI], c = bits.Add64(s[curI], h[prevI], c)
			s[curI], b = bits.Sub64(s[curI], allNeg & x[iMinusOne], b)
			z[iMinusOne] = s[prevI]
		}
	{
		const curI = 4 % 2
		const prevI = 1 - curI
		const iMinusOne = 3

		s[curI], _ = bits.Sub64(h[prevI], allNeg & x[iMinusOne], b)
		z[iMinusOne] = s[prevI]

		return s[curI] + c
	}
}*/

// Requires NoCarry
func (z *Element) linearCombNonModular(x *Element, xC int64, y *Element, yC int64) uint64 {
	var yTimes Element

	yHi := yTimes.mulWRegular(y, yC)
	xHi := z.mulWRegular(x, xC)

	carry := uint64(0)
	z[0], carry = bits.Add64(z[0], yTimes[0], carry)
	z[1], carry = bits.Add64(z[1], yTimes[1], carry)
	z[2], carry = bits.Add64(z[2], yTimes[2], carry)
	z[3], carry = bits.Add64(z[3], yTimes[3], carry)

	yHi, _ = bits.Add64(xHi, yHi, carry)

	return yHi
}
