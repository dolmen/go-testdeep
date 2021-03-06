// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// DO NOT EDIT!!! AUTOMATICALLY GENERATED!!!

package testdeep_test

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	. "github.com/maxatome/go-testdeep"
)

func ExampleCmpAll() {
	t := &testing.T{}

	got := "foo/bar"

	// Checks got string against:
	//   "o/b" regexp *AND* "bar" suffix *AND* exact "foo/bar" string
	ok := CmpAll(t, got, []interface{}{Re("o/b"), HasSuffix("bar"), "foo/bar"},
		"checks value %s", got)
	fmt.Println(ok)

	// Checks got string against:
	//   "o/b" regexp *AND* "bar" suffix *AND* exact "fooX/Ybar" string
	ok = CmpAll(t, got, []interface{}{Re("o/b"), HasSuffix("bar"), "fooX/Ybar"},
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpAny() {
	t := &testing.T{}

	got := "foo/bar"

	// Checks got string against:
	//   "zip" regexp *OR* "bar" suffix
	ok := CmpAny(t, got, []interface{}{Re("zip"), HasSuffix("bar")},
		"checks value %s", got)
	fmt.Println(ok)

	// Checks got string against:
	//   "zip" regexp *OR* "foo" suffix
	ok = CmpAny(t, got, []interface{}{Re("zip"), HasSuffix("foo")},
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpArray_array() {
	t := &testing.T{}

	got := [3]int{42, 58, 26}

	ok := CmpArray(t, got, [3]int{42}, ArrayEntries{1: 58, 2: Ignore()},
		"checks array %v", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpArray_typedArray() {
	t := &testing.T{}

	type MyArray [3]int

	got := MyArray{42, 58, 26}

	ok := CmpArray(t, got, MyArray{42}, ArrayEntries{1: 58, 2: Ignore()},
		"checks typed array %v", got)
	fmt.Println(ok)

	ok = CmpArray(t, &got, &MyArray{42}, ArrayEntries{1: 58, 2: Ignore()},
		"checks pointer on typed array %v", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpArrayEach_array() {
	t := &testing.T{}

	got := [3]int{42, 58, 26}

	ok := CmpArrayEach(t, got, Between(25, 60),
		"checks each item of array %v is in [25 .. 60]", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpArrayEach_typedArray() {
	t := &testing.T{}

	type MyArray [3]int

	got := MyArray{42, 58, 26}

	ok := CmpArrayEach(t, got, Between(25, 60),
		"checks each item of typed array %v is in [25 .. 60]", got)
	fmt.Println(ok)

	ok = CmpArrayEach(t, &got, Between(25, 60),
		"checks each item of typed array pointer %v is in [25 .. 60]", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpArrayEach_slice() {
	t := &testing.T{}

	got := []int{42, 58, 26}

	ok := CmpArrayEach(t, got, Between(25, 60),
		"checks each item of slice %v is in [25 .. 60]", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpArrayEach_typedSlice() {
	t := &testing.T{}

	type MySlice []int

	got := MySlice{42, 58, 26}

	ok := CmpArrayEach(t, got, Between(25, 60),
		"checks each item of typed slice %v is in [25 .. 60]", got)
	fmt.Println(ok)

	ok = CmpArrayEach(t, &got, Between(25, 60),
		"checks each item of typed slice pointer %v is in [25 .. 60]", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpBag() {
	t := &testing.T{}

	got := []int{1, 3, 5, 8, 8, 1, 2}

	// Matches as all items are present
	ok := CmpBag(t, got, []interface{}{1, 1, 2, 3, 5, 8, 8},
		"checks all items are present, in any order")
	fmt.Println(ok)

	// Does not match as got contains 2 times 1 and 8, and these
	// duplicates are not expected
	ok = CmpBag(t, got, []interface{}{1, 2, 3, 5, 8},
		"checks all items are present, in any order")
	fmt.Println(ok)

	got = []int{1, 3, 5, 8, 2}

	// Duplicates of 1 and 8 are expected but not present in got
	ok = CmpBag(t, got, []interface{}{1, 1, 2, 3, 5, 8, 8},
		"checks all items are present, in any order")
	fmt.Println(ok)

	// Matches as all items are present
	ok = CmpBag(t, got, []interface{}{1, 2, 3, 5, Gt(7)},
		"checks all items are present, in any order")
	fmt.Println(ok)

	// Output:
	// true
	// false
	// false
	// true
}

func ExampleCmpBetween() {
	t := &testing.T{}

	got := 156

	ok := CmpBetween(t, got, 154, 156, BoundsInIn,
		"checks %v is in [154 .. 156]", got)
	fmt.Println(ok)

	// BoundsInIn is implicit
	ok = CmpBetween(t, got, 154, 156, BoundsInIn,
		"checks %v is in [154 .. 156]", got)
	fmt.Println(ok)

	ok = CmpBetween(t, got, 154, 156, BoundsInOut,
		"checks %v is in [154 .. 156[", got)
	fmt.Println(ok)

	ok = CmpBetween(t, got, 154, 156, BoundsOutIn,
		"checks %v is in ]154 .. 156]", got)
	fmt.Println(ok)

	ok = CmpBetween(t, got, 154, 156, BoundsOutOut,
		"checks %v is in ]154 .. 156[", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// false
	// true
	// false
}

func ExampleCmpCap() {
	t := &testing.T{}

	got := make([]int, 0, 12)

	ok := CmpCap(t, got, 12, "checks %v capacity is 12", got)
	fmt.Println(ok)

	ok = CmpCap(t, got, 0, "checks %v capacity is 0", got)
	fmt.Println(ok)

	got = nil

	ok = CmpCap(t, got, 0, "checks %v capacity is 0", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
	// true
}

func ExampleCmpCap_operator() {
	t := &testing.T{}

	got := make([]int, 0, 12)

	ok := CmpCap(t, got, Between(10, 12),
		"checks %v capacity is in [10 .. 12]", got)
	fmt.Println(ok)

	ok = CmpCap(t, got, Gt(10),
		"checks %v capacity is in [10 .. 12]", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpCode() {
	t := &testing.T{}

	got := "12"

	ok := CmpCode(t, got, func(num string) bool {
		n, err := strconv.Atoi(num)
		return err == nil && n > 10 && n < 100
	},
		"checks string `%s` contains a number and this number is in ]10 .. 100[",
		got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpContains() {
	t := &testing.T{}

	got := "foobar"

	ok := CmpContains(t, got, "oob", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpContains_stringer() {
	t := &testing.T{}

	// bytes.Buffer implements fmt.Stringer
	got := bytes.NewBufferString("foobar")

	ok := CmpContains(t, got, "oob", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpContains_error() {
	t := &testing.T{}

	got := errors.New("foobar")

	ok := CmpContains(t, got, "oob", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpGt() {
	t := &testing.T{}

	got := 156

	ok := CmpGt(t, got, 155, "checks %v is > 155", got)
	fmt.Println(ok)

	ok = CmpGt(t, got, 156, "checks %v is > 156", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpGte() {
	t := &testing.T{}

	got := 156

	ok := CmpGte(t, got, 156, "checks %v is ≥ 156", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpHasPrefix() {
	t := &testing.T{}

	got := "foobar"

	ok := CmpHasPrefix(t, got, "foo", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpHasPrefix_stringer() {
	t := &testing.T{}

	// bytes.Buffer implements fmt.Stringer
	got := bytes.NewBufferString("foobar")

	ok := CmpHasPrefix(t, got, "foo", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpHasPrefix_error() {
	t := &testing.T{}

	got := errors.New("foobar")

	ok := CmpHasPrefix(t, got, "foo", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpHasSuffix() {
	t := &testing.T{}

	got := "foobar"

	ok := CmpHasSuffix(t, got, "bar", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpHasSuffix_stringer() {
	t := &testing.T{}

	// bytes.Buffer implements fmt.Stringer
	got := bytes.NewBufferString("foobar")

	ok := CmpHasSuffix(t, got, "bar", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpHasSuffix_error() {
	t := &testing.T{}

	got := errors.New("foobar")

	ok := CmpHasSuffix(t, got, "bar", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpIsa() {
	t := &testing.T{}

	type TstStruct struct {
		Field int
	}

	got := TstStruct{Field: 1}

	ok := CmpIsa(t, got, TstStruct{}, "checks got is a TstStruct")
	fmt.Println(ok)

	ok = CmpIsa(t, got, &TstStruct{},
		"checks got is a pointer on a TstStruct")
	fmt.Println(ok)

	ok = CmpIsa(t, &got, &TstStruct{},
		"checks &got is a pointer on a TstStruct")
	fmt.Println(ok)

	// Output:
	// true
	// false
	// true
}

func ExampleCmpIsa_interface() {
	t := &testing.T{}

	got := bytes.NewBufferString("foobar")

	ok := CmpIsa(t, got, (*fmt.Stringer)(nil),
		"checks got implements fmt.Stringer interface")
	fmt.Println(ok)

	errGot := fmt.Errorf("An error #%d occurred", 123)

	ok = CmpIsa(t, errGot, (*error)(nil),
		"checks errGot is a *error or implements error interface")
	fmt.Println(ok)

	// As nil, is passed below, it is not an interface but nil... So it
	// does not match
	errGot = nil

	ok = CmpIsa(t, errGot, (*error)(nil),
		"checks errGot is a *error or implements error interface")
	fmt.Println(ok)

	// BUT if its address is passed, now it is OK as the types match
	ok = CmpIsa(t, &errGot, (*error)(nil),
		"checks &errGot is a *error or implements error interface")
	fmt.Println(ok)

	// Output:
	// true
	// true
	// false
	// true
}

func ExampleCmpLen_slice() {
	t := &testing.T{}

	got := []int{11, 22, 33}

	ok := CmpLen(t, got, 3, "checks %v len is 3", got)
	fmt.Println(ok)

	ok = CmpLen(t, got, 0, "checks %v len is 0", got)
	fmt.Println(ok)

	got = nil

	ok = CmpLen(t, got, 0, "checks %v len is 0", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
	// true
}

func ExampleCmpLen_map() {
	t := &testing.T{}

	got := map[int]bool{11: true, 22: false, 33: false}

	ok := CmpLen(t, got, 3, "checks %v len is 3", got)
	fmt.Println(ok)

	ok = CmpLen(t, got, 0, "checks %v len is 0", got)
	fmt.Println(ok)

	got = nil

	ok = CmpLen(t, got, 0, "checks %v len is 0", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
	// true
}

func ExampleCmpLen_operatorSlice() {
	t := &testing.T{}

	got := []int{11, 22, 33}

	ok := CmpLen(t, got, Between(3, 8),
		"checks %v len is in [3 .. 8]", got)
	fmt.Println(ok)

	ok = CmpLen(t, got, Lt(5), "checks %v len is < 5", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpLen_operatorMap() {
	t := &testing.T{}

	got := map[int]bool{11: true, 22: false, 33: false}

	ok := CmpLen(t, got, Between(3, 8),
		"checks %v len is in [3 .. 8]", got)
	fmt.Println(ok)

	ok = CmpLen(t, got, Gte(3), "checks %v len is ≥ 3", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpLt() {
	t := &testing.T{}

	got := 156

	ok := CmpLt(t, got, 157, "checks %v is < 157", got)
	fmt.Println(ok)

	ok = CmpLt(t, got, 156, "checks %v is < 156", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpLte() {
	t := &testing.T{}

	got := 156

	ok := CmpLte(t, got, 156, "checks %v is ≤ 156", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpMap_map() {
	t := &testing.T{}

	got := map[string]int{"foo": 12, "bar": 42, "zip": 89}

	ok := CmpMap(t, got, map[string]int{"bar": 42}, MapEntries{"foo": Lt(15), "zip": Ignore()},
		"checks map %v", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpMap_typedMap() {
	t := &testing.T{}

	type MyMap map[string]int

	got := MyMap{"foo": 12, "bar": 42, "zip": 89}

	ok := CmpMap(t, got, MyMap{"bar": 42}, MapEntries{"foo": Lt(15), "zip": Ignore()},
		"checks typed map %v", got)
	fmt.Println(ok)

	ok = CmpMap(t, &got, &MyMap{"bar": 42}, MapEntries{"foo": Lt(15), "zip": Ignore()},
		"checks pointer on typed map %v", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpMapEach_map() {
	t := &testing.T{}

	got := map[string]int{"foo": 12, "bar": 42, "zip": 89}

	ok := CmpMapEach(t, got, Between(10, 90),
		"checks each value of map %v is in [10 .. 90]", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpMapEach_typedMap() {
	t := &testing.T{}

	type MyMap map[string]int

	got := MyMap{"foo": 12, "bar": 42, "zip": 89}

	ok := CmpMapEach(t, got, Between(10, 90),
		"checks each value of typed map %v is in [10 .. 90]", got)
	fmt.Println(ok)

	ok = CmpMapEach(t, &got, Between(10, 90),
		"checks each value of typed map pointer %v is in [10 .. 90]", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpN() {
	t := &testing.T{}

	got := 1.12345

	ok := CmpN(t, got, 1.1234, 0.00006,
		"checks %v = 1.1234 ± 0.00006", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpNil() {
	t := &testing.T{}

	var got fmt.Stringer // interface

	// nil value can be compared directly with nil, no need of Nil() here
	ok := CmpDeeply(t, got, nil)
	fmt.Println(ok)

	// But it works with Nil() anyway
	ok = CmpNil(t, got)
	fmt.Println(ok)

	got = (*bytes.Buffer)(nil)

	// In the case of an interface containing a nil pointer, comparing
	// with nil fails, as the interface is not nil
	ok = CmpDeeply(t, got, nil)
	fmt.Println(ok)

	// In this case Nil() succeed
	ok = CmpNil(t, got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// false
	// true
}

func ExampleCmpNone() {
	t := &testing.T{}

	got := 18

	ok := CmpNone(t, got, []interface{}{0, 10, 20, 30, Between(100, 199)},
		"checks %v is non-null, and ≠ 10, 20 & 30, and not in [100-199]", got)
	fmt.Println(ok)

	got = 20

	ok = CmpNone(t, got, []interface{}{0, 10, 20, 30, Between(100, 199)},
		"checks %v is non-null, and ≠ 10, 20 & 30, and not in [100-199]", got)
	fmt.Println(ok)

	got = 142

	ok = CmpNone(t, got, []interface{}{0, 10, 20, 30, Between(100, 199)},
		"checks %v is non-null, and ≠ 10, 20 & 30, and not in [100-199]", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
	// false
}

func ExampleCmpNoneOf() {
	t := &testing.T{}

	got := []int{4, 5, 9, 42}

	ok := CmpNoneOf(t, got, []interface{}{3, 6, 8, 41, 43},
		"checks %v contains no item listed in NoneOf()", got)
	fmt.Println(ok)

	ok = CmpNoneOf(t, got, []interface{}{3, 6, 8, 42, 43},
		"checks %v contains no item listed in NoneOf()", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpNot() {
	t := &testing.T{}

	got := 42

	ok := CmpNot(t, got, 0, "checks %v is non-null", got)
	fmt.Println(ok)

	ok = CmpNot(t, got, Between(10, 30),
		"checks %v is not in [10 .. 30]", got)
	fmt.Println(ok)

	got = 0

	ok = CmpNot(t, got, 0, "checks %v is non-null", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// false
}

func ExampleCmpNotNil() {
	t := &testing.T{}

	var got fmt.Stringer = &bytes.Buffer{}

	// nil value can be compared directly with Not(nil), no need of NotNil() here
	ok := CmpDeeply(t, got, Not(nil))
	fmt.Println(ok)

	// But it works with NotNil() anyway
	ok = CmpNotNil(t, got)
	fmt.Println(ok)

	got = (*bytes.Buffer)(nil)

	// In the case of an interface containing a nil pointer, comparing
	// with Not(nil) succeeds, as the interface is not nil
	ok = CmpDeeply(t, got, Not(nil))
	fmt.Println(ok)

	// In this case NotNil() fails
	ok = CmpNotNil(t, got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// true
	// false
}

func ExampleCmpPPtr() {
	t := &testing.T{}

	num := 12
	got := &num

	ok := CmpPPtr(t, &got, 12)
	fmt.Println(ok)

	ok = CmpPPtr(t, &got, Between(4, 15))
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpPtr() {
	t := &testing.T{}

	got := 12

	ok := CmpPtr(t, &got, 12)
	fmt.Println(ok)

	ok = CmpPtr(t, &got, Between(4, 15))
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpRe() {
	t := &testing.T{}

	got := "foo bar"
	ok := CmpRe(t, got, "(zip|bar)$", nil, "checks value %s", got)
	fmt.Println(ok)

	got = "bar foo"
	ok = CmpRe(t, got, "(zip|bar)$", nil, "checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpRe_stringer() {
	t := &testing.T{}

	// bytes.Buffer implements fmt.Stringer
	got := bytes.NewBufferString("foo bar")
	ok := CmpRe(t, got, "(zip|bar)$", nil, "checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpRe_error() {
	t := &testing.T{}

	got := errors.New("foo bar")
	ok := CmpRe(t, got, "(zip|bar)$", nil, "checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpRe_capture() {
	t := &testing.T{}

	got := "foo bar biz"
	ok := CmpRe(t, got, `^(\w+) (\w+) (\w+)$`, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	got = "foo bar! biz"
	ok = CmpRe(t, got, `^(\w+) (\w+) (\w+)$`, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpRe_compiled() {
	t := &testing.T{}

	expected := regexp.MustCompile("(zip|bar)$")

	got := "foo bar"
	ok := CmpRe(t, got, expected, nil, "checks value %s", got)
	fmt.Println(ok)

	got = "bar foo"
	ok = CmpRe(t, got, expected, nil, "checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpRe_compiledStringer() {
	t := &testing.T{}

	expected := regexp.MustCompile("(zip|bar)$")

	// bytes.Buffer implements fmt.Stringer
	got := bytes.NewBufferString("foo bar")
	ok := CmpRe(t, got, expected, nil, "checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpRe_compiledError() {
	t := &testing.T{}

	expected := regexp.MustCompile("(zip|bar)$")

	got := errors.New("foo bar")
	ok := CmpRe(t, got, expected, nil, "checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpRe_compiledCapture() {
	t := &testing.T{}

	expected := regexp.MustCompile(`^(\w+) (\w+) (\w+)$`)

	got := "foo bar biz"
	ok := CmpRe(t, got, expected, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	got = "foo bar! biz"
	ok = CmpRe(t, got, expected, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpReAll_capture() {
	t := &testing.T{}

	got := "foo bar biz"
	ok := CmpReAll(t, got, `(\w+)`, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	// Matches, but all catured groups do not match Set
	got = "foo BAR biz"
	ok = CmpReAll(t, got, `(\w+)`, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpReAll_captureComplex() {
	t := &testing.T{}

	got := "11 45 23 56 85 96"
	ok := CmpReAll(t, got, `(\d+)`, ArrayEach(Code(func(num string) bool {
		n, err := strconv.Atoi(num)
		return err == nil && n > 10 && n < 100
	})),
		"checks value %s", got)
	fmt.Println(ok)

	// Matches, but 11 is not greater than 20
	ok = CmpReAll(t, got, `(\d+)`, ArrayEach(Code(func(num string) bool {
		n, err := strconv.Atoi(num)
		return err == nil && n > 20 && n < 100
	})),
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpReAll_compiledCapture() {
	t := &testing.T{}

	expected := regexp.MustCompile(`(\w+)`)

	got := "foo bar biz"
	ok := CmpReAll(t, got, expected, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	// Matches, but all catured groups do not match Set
	got = "foo BAR biz"
	ok = CmpReAll(t, got, expected, Set("biz", "foo", "bar"),
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpReAll_compiledCaptureComplex() {
	t := &testing.T{}

	expected := regexp.MustCompile(`(\d+)`)

	got := "11 45 23 56 85 96"
	ok := CmpReAll(t, got, expected, ArrayEach(Code(func(num string) bool {
		n, err := strconv.Atoi(num)
		return err == nil && n > 10 && n < 100
	})),
		"checks value %s", got)
	fmt.Println(ok)

	// Matches, but 11 is not greater than 20
	ok = CmpReAll(t, got, expected, ArrayEach(Code(func(num string) bool {
		n, err := strconv.Atoi(num)
		return err == nil && n > 20 && n < 100
	})),
		"checks value %s", got)
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpSet() {
	t := &testing.T{}

	got := []int{1, 3, 5, 8, 8, 1, 2}

	// Matches as all items are present, ignoring duplicates
	ok := CmpSet(t, got, []interface{}{1, 2, 3, 5, 8},
		"checks all items are present, in any order")
	fmt.Println(ok)

	// Duplicates are ignored in a Set
	ok = CmpSet(t, got, []interface{}{1, 2, 2, 2, 2, 2, 3, 5, 8},
		"checks all items are present, in any order")
	fmt.Println(ok)

	// Tries its best to not raise an error when a value can be matched
	// by several Set entries
	ok = CmpSet(t, got, []interface{}{Between(1, 4), 3, Between(2, 10)},
		"checks all items are present, in any order")
	fmt.Println(ok)

	// Output:
	// true
	// true
	// true
}

func ExampleCmpShallow() {
	t := &testing.T{}

	type MyStruct struct {
		Value int
	}
	data := MyStruct{Value: 12}
	got := &data

	ok := CmpShallow(t, got, &data,
		"checks pointers only, not contents")
	fmt.Println(ok)

	// Same contents, but not same pointer
	ok = CmpShallow(t, got, &MyStruct{Value: 12},
		"checks pointers only, not contents")
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleCmpSlice_slice() {
	t := &testing.T{}

	got := []int{42, 58, 26}

	ok := CmpSlice(t, got, []int{42}, ArrayEntries{1: 58, 2: Ignore()},
		"checks slice %v", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpSlice_typedSlice() {
	t := &testing.T{}

	type MySlice []int

	got := MySlice{42, 58, 26}

	ok := CmpSlice(t, got, MySlice{42}, ArrayEntries{1: 58, 2: Ignore()},
		"checks typed slice %v", got)
	fmt.Println(ok)

	ok = CmpSlice(t, &got, &MySlice{42}, ArrayEntries{1: 58, 2: Ignore()},
		"checks pointer on typed slice %v", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpString() {
	t := &testing.T{}

	got := "foobar"

	ok := CmpString(t, got, "foobar", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpString_stringer() {
	t := &testing.T{}

	// bytes.Buffer implements fmt.Stringer
	got := bytes.NewBufferString("foobar")

	ok := CmpString(t, got, "foobar", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpString_error() {
	t := &testing.T{}

	got := errors.New("foobar")

	ok := CmpString(t, got, "foobar", "checks %s", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpStruct() {
	t := &testing.T{}

	type Person struct {
		Name        string
		Age         int
		NumChildren int
	}

	got := Person{
		Name:        "Foobar",
		Age:         42,
		NumChildren: 3,
	}

	// As NumChildren is zero in Struct() call, it is not checked
	ok := CmpStruct(t, got, Person{Name: "Foobar"}, StructFields{
		"Age": Between(40, 50),
	},
		"checks %v is the right Person")
	fmt.Println(ok)

	// Model can be empty
	ok = CmpStruct(t, got, Person{}, StructFields{
		"Name":        "Foobar",
		"Age":         Between(40, 50),
		"NumChildren": Not(0),
	},
		"checks %v is the right Person")
	fmt.Println(ok)

	// Works with pointers too
	ok = CmpStruct(t, &got, &Person{}, StructFields{
		"Name":        "Foobar",
		"Age":         Between(40, 50),
		"NumChildren": Not(0),
	},
		"checks %v is the right Person")
	fmt.Println(ok)

	// Output:
	// true
	// true
	// true
}

func ExampleCmpSubBagOf() {
	t := &testing.T{}

	got := []int{1, 3, 5, 8, 8, 1, 2}

	ok := CmpSubBagOf(t, got, []interface{}{0, 0, 1, 1, 2, 2, 3, 3, 5, 5, 8, 8, 9, 9},
		"checks at least all items are present, in any order")
	fmt.Println(ok)

	// got contains one 8 too many
	ok = CmpSubBagOf(t, got, []interface{}{0, 0, 1, 1, 2, 2, 3, 3, 5, 5, 8, 9, 9},
		"checks at least all items are present, in any order")
	fmt.Println(ok)

	got = []int{1, 3, 5, 2}

	ok = CmpSubBagOf(t, got, []interface{}{Between(0, 3), Between(0, 3), Between(0, 3), Between(0, 3), Gt(4), Gt(4)},
		"checks at least all items match, in any order with TestDeep operators")
	fmt.Println(ok)

	// Output:
	// true
	// false
	// true
}

func ExampleCmpSubMapOf_map() {
	t := &testing.T{}

	got := map[string]int{"foo": 12, "bar": 42}

	ok := CmpSubMapOf(t, got, map[string]int{"bar": 42}, MapEntries{"foo": Lt(15), "zip": 666},
		"checks map %v is included in expected keys/values", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpSubMapOf_typedMap() {
	t := &testing.T{}

	type MyMap map[string]int

	got := MyMap{"foo": 12, "bar": 42}

	ok := CmpSubMapOf(t, got, MyMap{"bar": 42}, MapEntries{"foo": Lt(15), "zip": 666},
		"checks typed map %v is included in expected keys/values", got)
	fmt.Println(ok)

	ok = CmpSubMapOf(t, &got, &MyMap{"bar": 42}, MapEntries{"foo": Lt(15), "zip": 666},
		"checks pointed typed map %v is included in expected keys/values", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpSubSetOf() {
	t := &testing.T{}

	got := []int{1, 3, 5, 8, 8, 1, 2}

	// Matches as all items are expected, ignoring duplicates
	ok := CmpSubSetOf(t, got, []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		"checks at least all items are present, in any order, ignoring duplicates")
	fmt.Println(ok)

	// Tries its best to not raise an error when a value can be matched
	// by several SubSetOf entries
	ok = CmpSubSetOf(t, got, []interface{}{Between(1, 4), 3, Between(2, 10), Gt(100)},
		"checks at least all items are present, in any order, ignoring duplicates")
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpSuperBagOf() {
	t := &testing.T{}

	got := []int{1, 3, 5, 8, 8, 1, 2}

	ok := CmpSuperBagOf(t, got, []interface{}{8, 5, 8},
		"checks the items are present, in any order")
	fmt.Println(ok)

	ok = CmpSuperBagOf(t, got, []interface{}{Gt(5), Lte(2)},
		"checks at least 2 items of %v match", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpSuperMapOf_map() {
	t := &testing.T{}

	got := map[string]int{"foo": 12, "bar": 42, "zip": 89}

	ok := CmpSuperMapOf(t, got, map[string]int{"bar": 42}, MapEntries{"foo": Lt(15)},
		"checks map %v contains at leat all expected keys/values", got)
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleCmpSuperMapOf_typedMap() {
	t := &testing.T{}

	type MyMap map[string]int

	got := MyMap{"foo": 12, "bar": 42, "zip": 89}

	ok := CmpSuperMapOf(t, got, MyMap{"bar": 42}, MapEntries{"foo": Lt(15)},
		"checks typed map %v contains at leat all expected keys/values", got)
	fmt.Println(ok)

	ok = CmpSuperMapOf(t, &got, &MyMap{"bar": 42}, MapEntries{"foo": Lt(15)},
		"checks pointed typed map %v contains at leat all expected keys/values",
		got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpSuperSetOf() {
	t := &testing.T{}

	got := []int{1, 3, 5, 8, 8, 1, 2}

	ok := CmpSuperSetOf(t, got, []interface{}{1, 2, 3},
		"checks the items are present, in any order and ignoring duplicates")
	fmt.Println(ok)

	ok = CmpSuperSetOf(t, got, []interface{}{Gt(5), Lte(2)},
		"checks at least 2 items of %v match ignoring duplicates", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
}

func ExampleCmpTruncTime() {
	t := &testing.T{}

	dateToTime := func(str string) time.Time {
		t, err := time.Parse(time.RFC3339Nano, str)
		if err != nil {
			panic(err)
		}
		return t
	}

	got := dateToTime("2018-05-01T12:45:53.123456789Z")

	// Compare dates ignoring nanoseconds and monotonic parts
	expected := dateToTime("2018-05-01T12:45:53Z")
	ok := CmpTruncTime(t, got, expected, time.Second,
		"checks date %v, truncated to the second", got)
	fmt.Println(ok)

	// Compare dates ignoring time and so monotonic parts
	expected = dateToTime("2018-05-01T11:22:33.444444444Z")
	ok = CmpTruncTime(t, got, expected, 24*time.Hour,
		"checks date %v, truncated to the day", got)
	fmt.Println(ok)

	// Compare dates exactly but ignoring monotonic part
	expected = dateToTime("2018-05-01T12:45:53.123456789Z")
	ok = CmpTruncTime(t, got, expected, 0,
		"checks date %v ignoring monotonic part", got)
	fmt.Println(ok)

	// Output:
	// true
	// true
	// true
}

func ExampleCmpZero() {
	t := &testing.T{}

	ok := CmpZero(t, 0)
	fmt.Println(ok)

	ok = CmpZero(t, float64(0))
	fmt.Println(ok)

	ok = CmpZero(t, 12) // fails, as 12 is not 0 :)
	fmt.Println(ok)

	ok = CmpZero(t, (map[string]int)(nil))
	fmt.Println(ok)

	ok = CmpZero(t, map[string]int{}) // fails, as not nil
	fmt.Println(ok)

	ok = CmpZero(t, ([]int)(nil))
	fmt.Println(ok)

	ok = CmpZero(t, []int{}) // fails, as not nil
	fmt.Println(ok)

	ok = CmpZero(t, [3]int{})
	fmt.Println(ok)

	ok = CmpDeeply(t, [3]int{0, 1}, Zero()) // fails, DATA[1] is not 0
	fmt.Println(ok)

	ok = CmpZero(t, bytes.Buffer{})
	fmt.Println(ok)

	ok = CmpZero(t, &bytes.Buffer{}) // fails, as pointer not nil
	fmt.Println(ok)

	// Output:
	// true
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
}
