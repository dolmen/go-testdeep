// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package testdeep

import (
	"fmt"
	"reflect"
)

type tdAll struct {
	tdList
}

var _ TestDeep = &tdAll{}

// All operator compares data against several expected values. During
// a match, all of them have to match to succeed.
func All(expectedValues ...interface{}) TestDeep {
	return &tdAll{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAll) Match(ctx Context, got reflect.Value) (err *Error) {
	for idx, item := range a.items {
		origErr := deepValueEqual(
			ctx.AddDepth(fmt.Sprintf("<All#%d/%d>", idx+1, len(a.items))),
			got, item)
		if origErr != nil {
			if ctx.booleanError {
				return booleanError
			}
			err = &Error{
				Context:  ctx,
				Message:  fmt.Sprintf("compared (part %d of %d)", idx+1, len(a.items)),
				Got:      got,
				Expected: item,
				Location: a.GetLocation(),
			}

			if item.IsValid() && item.Type().Implements(testDeeper) {
				err.Origin = origErr
			}
			return
		}
	}
	return
}
