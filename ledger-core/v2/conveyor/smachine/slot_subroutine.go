/*
 * Copyright 2020 Insolar Network Ltd.
 * All rights reserved.
 * This material is licensed under the Insolar License version 1.0,
 * available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.
 */

package smachine

import (
	"errors"

	"github.com/insolar/assured-ledger/ledger-core/v2/vanilla/throw"
)

type subroutineMarker struct {
	slotId SlotID
	_      [0]func() // force incomparable by value
}

func (s *Slot) newSubroutineMarker() *subroutineMarker {
	return &subroutineMarker{slotId: s.GetSlotID()}
}

func (s *Slot) getSubroutineMarker() *subroutineMarker {
	if s.stateStack != nil {
		return s.stateStack.childMarker
	}
	return nil
}

func (s *Slot) ensureNoSubroutine() {
	if s.hasSubroutine() {
		panic(throw.IllegalState())
	}
}

func (s *Slot) forceSubroutineUpdate(su StateUpdate, producedBy *subroutineMarker) StateUpdate {
	switch isCurrent, isValid := s.checkSubroutineMarker(producedBy); {
	case isCurrent:
		//
	case !isValid:
		s.machine.logInternal(s.NewStepLink(), "aborting routine has expired", nil)
	case !typeOfStateUpdate(su).IsSubroutineSafe():
		s._popTillSubroutine(producedBy)
	}
	return su
}

var errTerminatedByCaller = errors.New("subroutine SM is terminated by caller")

func (s *Slot) _popTillSubroutine(producedBy *subroutineMarker) {
	for ms := s.stateStack; ms != nil; ms = ms.stateStack {
		if ms.childMarker == producedBy {
			return
		}
		s.prepareSubroutineExit(errTerminatedByCaller)
	}
	if producedBy != nil {
		panic(throw.IllegalState())
	}
}

func (s *Slot) hasSubroutine() bool {
	return s.stateStack != nil
}

func wrapUnsafeSubroutineUpdate(su StateUpdate, producedBy *subroutineMarker) StateUpdate {
	return newSubroutineAbortStateUpdate(SlotStep{Transition: func(ctx ExecutionContext) StateUpdate {
		ec := ctx.(*executionContext)
		slot := ec.s
		switch isCurrent, isValid := slot.checkSubroutineMarker(producedBy); {
		case isCurrent:
			return su
		case !isValid:
			ctx.Log().Warn("aborting routine has expired")
			return su
		}
		slot._popTillSubroutine(producedBy)
		return su
	}})
}

func (s *Slot) checkSubroutineMarker(marker *subroutineMarker) (isCurrent, isValid bool) {
	switch {
	case marker == nil:
		return s.stateStack == nil, true
	case s.stateStack == nil:
		//
	case s.GetSlotID() != marker.slotId:
		//
	case s.stateStack.childMarker == marker:
		return true, true
	default:
		for sbs := s.stateStack.stateStack; sbs != nil; sbs = sbs.stateStack {
			if sbs.childMarker == marker {
				return false, true
			}
		}
	}
	return false, false
}

func (s *Slot) prepareSubroutineStart(ssm SubroutineStateMachine, exitFn SubroutineExitFunc, migrateFn MigrateFunc) SlotStep {
	if exitFn == nil {
		panic(throw.IllegalValue())
	}
	if ssm == nil {
		panic(throw.IllegalValue())
	}
	decl := ssm.GetStateMachineDeclaration()
	if decl == nil {
		panic(throw.IllegalValue())
	}

	return SlotStep{Migration: migrateFn, Transition: func(ctx ExecutionContext) StateUpdate {
		ec := ctx.(*executionContext)
		slot := ec.s

		prev := slot.stateMachineData
		if migrateFn == nil {
			migrateFn = prev.defMigrate
		}
		stackedSdd := &stackedStateMachineData{prev, migrateFn,
			exitFn, slot.newSubroutineMarker(),
			prev.stateStack != nil && prev.stateStack.hasMigrates,
		}
		if stackedSdd.stackMigrate != nil {
			stackedSdd.hasMigrates = true
		}

		slot.stateMachineData.stateStack = stackedSdd
		slot.stateMachineData.declaration = decl
		slot.stateMachineData.shadowMigrate = nil
		slot.stateMachineData.defTerminate = nil

		initFn := slot.prepareSubroutineInit(ssm, ctx.Log().GetTracerId())

		return initFn.defaultInit(ctx)
	}}
}

var defaultSubroutineStartDecl = StepDeclaration{stepDeclExt: stepDeclExt{Name: "<subroutineStart>"}}
var defaultSubroutineAbortDecl = StepDeclaration{stepDeclExt: stepDeclExt{Name: "<subroutineAbort>"}}
var defaultSubroutineExitDecl = StepDeclaration{stepDeclExt: stepDeclExt{Name: "<subroutineExit>"}}

func (s *Slot) prepareSubroutineExit(lastError error) {
	// TODO logging?
	prev := s.stateStack
	if prev == nil {
		panic(throw.IllegalState())
	}

	lastResult := s.defResult
	returnFn := prev.returnFn

	s.stateMachineData = prev.stateMachineData
	s.step = SlotStep{Transition: func(ctx ExecutionContext) StateUpdate {
		ec := ctx.(*executionContext)
		bc := subroutineExitContext{bargingInContext{ec.clone(updCtxInactive),
			lastResult, false}, lastError}

		su := bc.executeSubroutineExit(returnFn)
		su.marker = ec.getMarker()
		return su
	}}
	s.stepDecl = &defaultSubroutineExitDecl
}
