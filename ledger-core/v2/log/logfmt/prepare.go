// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package logfmt

import (
	"fmt"
	"reflect"
)

type LogStringer interface {
	LogString() string
}

type valuePrepareFn func(interface{}) (string, reflect.Kind, bool)

// WARNING! Sequence of types must match in both findValuePrepareFn() and prepareValue()
// this is a bit slower vs array scan, but may help the compiler with escape analysis
func findPrepareValueFn(t reflect.Type) valuePrepareFn {
	switch {
	case t.AssignableTo(reflect.TypeOf((*LogStringer)(nil)).Elem()):
		return func(value interface{}) (string, reflect.Kind, bool) {
			switch vv := value.(LogStringer); {
			case vv != nil:
				v := vv.LogString()
				return v, reflect.Interface, false
			default:
				return "", reflect.Interface, true
			}
		}
	case t.AssignableTo(reflect.TypeOf((func() string)(nil))):
		return func(value interface{}) (string, reflect.Kind, bool) {
			switch vv := value.(func() string); {
			case vv != nil:
				v := vv()
				return v, reflect.Func, false
			default:
				return "", reflect.Func, true
			}
		}
	case t.AssignableTo(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()):
		return func(value interface{}) (string, reflect.Kind, bool) {
			switch vv := value.(fmt.Stringer); {
			case vv != nil:
				v := vv.String()
				return v, reflect.Interface, false
			default:
				return "", reflect.Interface, true
			}
		}
	default:
		return nil
	}
}

// WARNING! Sequence of types must match in both findValuePrepareFn() and prepareValue()
func prepareValue(iv interface{}) (string, reflect.Kind, bool) {
	switch vv := iv.(type) {
	case nil:
		return "", reflect.Interface, true
	case LogStringer:
		return vv.LogString(), reflect.Interface, false
	case func() string:
		if vv == nil {
			return "", reflect.Func, true
		}
		return vv(), reflect.Func, false
	case fmt.Stringer:
		return vv.String(), reflect.Interface, false
	}
	return "", reflect.Invalid, false
}
