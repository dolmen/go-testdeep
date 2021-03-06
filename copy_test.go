// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package testdeep_test

import (
	"reflect"
	"testing"

	. "github.com/maxatome/go-testdeep"
)

func checkFieldValueOK(t *testing.T,
	s reflect.Value, fieldName string, value interface{}) {
	t.Helper()

	testName := "field " + fieldName

	fieldOrig := s.FieldByName(fieldName)
	isFalse(t, fieldOrig.CanInterface(), testName+" + fieldOrig.CanInterface()")

	fieldCopy, ok := CopyValue(fieldOrig)
	if isTrue(t, ok, "Can copy "+testName) {
		if isTrue(t, fieldCopy.CanInterface(), testName+" + fieldCopy.CanInterface()") {
			CmpDeeply(t, fieldCopy.Interface(), value,
				testName+" + fieldCopy contents")
		}
	}
}

func checkFieldValueNOK(t *testing.T, s reflect.Value, fieldName string) {
	t.Helper()

	testName := "field " + fieldName

	fieldOrig := s.FieldByName(fieldName)
	isFalse(t, fieldOrig.CanInterface(), testName+" + fieldOrig.CanInterface()")

	_, ok := CopyValue(fieldOrig)
	isFalse(t, ok, "Could not copy "+testName)
}

func TestCopyValue(t *testing.T) {
	// Note that even if all the fields are public, a Struct cannot be copied
	type SubPublic struct {
		Public int
	}

	type SubPrivate struct {
		private int // nolint: unused,megacheck
	}

	type Private struct {
		boolean  bool
		integer  int
		uinteger uint
		cplx     complex128
		flt      float64
		str      string
		array    [3]interface{}
		slice    []interface{}
		hash     map[interface{}]interface{}
		pint     *int
		iface    interface{}
		fn       func()
	}

	//
	// Copy OK
	num := 123
	private := Private{
		boolean: true,
		integer: 42,
		cplx:    complex(2, -2),
		flt:     1.234,
		str:     "foobar",
		array:   [3]interface{}{1, 2, SubPublic{Public: 3}},
		slice:   append(make([]interface{}, 0, 10), 4, 5, SubPublic{Public: 6}),
		hash: map[interface{}]interface{}{
			"foo": &SubPublic{Public: 34},
			SubPublic{Public: 78}: 42,
		},
		pint:  &num,
		iface: &num,
	}
	privateStruct := reflect.ValueOf(private)

	checkFieldValueOK(t, privateStruct, "boolean", private.boolean)
	checkFieldValueOK(t, privateStruct, "integer", private.integer)
	checkFieldValueOK(t, privateStruct, "uinteger", private.uinteger)
	checkFieldValueOK(t, privateStruct, "cplx", private.cplx)
	checkFieldValueOK(t, privateStruct, "flt", private.flt)
	checkFieldValueOK(t, privateStruct, "str", private.str)
	checkFieldValueOK(t, privateStruct, "array", private.array)
	checkFieldValueOK(t, privateStruct, "slice", private.slice)
	checkFieldValueOK(t, privateStruct, "hash", private.hash)
	checkFieldValueOK(t, privateStruct, "pint", private.pint)
	checkFieldValueOK(t, privateStruct, "iface", private.iface)

	//
	// Not able to copy...
	private = Private{
		array: [3]interface{}{1, 2, SubPrivate{}},
		slice: append(make([]interface{}, 0, 10), &SubPrivate{}, &SubPrivate{}),
		hash:  map[interface{}]interface{}{"foo": &SubPrivate{}},
		iface: &SubPrivate{},
		fn:    func() {},
	}
	privateStruct = reflect.ValueOf(private)

	checkFieldValueNOK(t, privateStruct, "array")
	checkFieldValueNOK(t, privateStruct, "slice")
	checkFieldValueNOK(t, privateStruct, "hash")
	checkFieldValueNOK(t, privateStruct, "iface")
	checkFieldValueNOK(t, privateStruct, "fn")

	private.hash = map[interface{}]interface{}{SubPrivate{}: 123}
	privateStruct = reflect.ValueOf(private)
	checkFieldValueNOK(t, privateStruct, "hash")

	//
	// nil cases
	private = Private{}
	privateStruct = reflect.ValueOf(private)
	checkFieldValueOK(t, privateStruct, "slice", private.slice)
	checkFieldValueOK(t, privateStruct, "hash", private.hash)
	checkFieldValueOK(t, privateStruct, "pint", private.pint)
	checkFieldValueOK(t, privateStruct, "iface", private.iface)
}
