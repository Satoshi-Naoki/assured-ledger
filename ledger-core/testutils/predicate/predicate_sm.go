// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package predicate

import (
	"reflect"

	"github.com/insolar/assured-ledger/ledger-core/conveyor/smachine"
	"github.com/insolar/assured-ledger/ledger-core/testutils/debuglogger"
)

func AfterAnyStopOrError(event debuglogger.UpdateEvent) bool {
	updateType := event.Update.UpdateType
	return updateType == "stop" || updateType == "panic" || updateType == "error"
}

func NewSMTypeFilter(sample smachine.StateMachine, andPredicate Func) Func {
	var smType = reflect.TypeOf(sample)

	return func(event debuglogger.UpdateEvent) bool {
		switch {
		case event.SM == nil:
		case reflect.TypeOf(event.SM) != smType:
		case andPredicate != nil && !andPredicate(event):
		default:
			return true
		}
		return false
	}
}