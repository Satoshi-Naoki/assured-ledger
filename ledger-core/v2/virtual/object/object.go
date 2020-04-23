// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package object

import (
	"github.com/insolar/assured-ledger/ledger-core/v2/conveyor"
	"github.com/insolar/assured-ledger/ledger-core/v2/conveyor/smachine"
	"github.com/insolar/assured-ledger/ledger-core/v2/conveyor/smachine/smsync"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/payload"
	"github.com/insolar/assured-ledger/ledger-core/v2/network/messagesender"
	messageSenderAdapter "github.com/insolar/assured-ledger/ledger-core/v2/network/messagesender/adapter"
	"github.com/insolar/assured-ledger/ledger-core/v2/vanilla/injector"
	"github.com/insolar/assured-ledger/ledger-core/v2/virtual/descriptor"
)

type Info struct {
	Reference     insolar.Reference
	descriptor    descriptor.ObjectDescriptor
	Deactivated   bool
	IsReadyToWork bool

	ImmutableExecute smachine.SyncLink
	MutableExecute   smachine.SyncLink
	ReadyToWork      smachine.SyncLink

	ActiveImmutablePendingCount    uint8
	ActiveMutablePendingCount      uint8
	PotentialImmutablePendingCount uint8
	PotentialMutablePendingCount   uint8
}

func (i *Info) SetDescriptor(prototype *insolar.Reference, memory []byte) {
	i.descriptor = descriptor.NewObjectDescriptor(
		i.Reference, insolar.ID{}, prototype, memory, insolar.Reference{}, nil,
	)
}

func (i *Info) Deactivate() {
	i.Deactivated = true
}

func (i *Info) Descriptor() descriptor.ObjectDescriptor {
	return i.descriptor
}

func (i *Info) SetReady() {
	i.IsReadyToWork = true
}

type SharedState struct {
	Info
}

// StackRelation shows relevance of 2 stacks
type InitReason int8

const (
	InitReasonCTConstructor InitReason = iota
	InitReasonCTMethod
	InitReasonVStateReport
)

func NewStateMachineObject(objectReference insolar.Reference, reason InitReason) *SMObject {
	return &SMObject{
		SharedState: SharedState{
			Info: Info{Reference: objectReference},
		},
		initReason: reason,
	}
}

type SMObject struct {
	smachine.StateMachineDeclTemplate

	SharedState

	readyToWorkCtl   smsync.BoolConditionalLink
	initReason       InitReason
	stateRequestSent bool

	// dependencies
	messageSender *messageSenderAdapter.MessageSender
	pulseSlot     *conveyor.PulseSlot
}

/* -------- Declaration ------------- */

func (sm *SMObject) InjectDependencies(stateMachine smachine.StateMachine, _ smachine.SlotLink, injector *injector.DependencyInjector) {
	s := stateMachine.(*SMObject)
	injector.MustInject(&s.messageSender)
	injector.MustInject(&s.pulseSlot)
}

func (sm *SMObject) GetInitStateFor(smachine.StateMachine) smachine.InitFunc {
	return sm.Init
}

/* -------- Instance ------------- */

func (sm *SMObject) GetStateMachineDeclaration() smachine.StateMachineDeclaration {
	return sm
}

func (sm *SMObject) Init(ctx smachine.InitializationContext) smachine.StateUpdate {
	sm.readyToWorkCtl = smsync.NewConditionalBool(false, "readyToWork")
	sm.ReadyToWork = sm.readyToWorkCtl.SyncLink()

	sm.ImmutableExecute = smsync.NewSemaphore(30, "immutable calls").SyncLink()
	sm.MutableExecute = smsync.NewSemaphore(1, "mutable calls").SyncLink() // TODO here we need an ORDERED queue

	sdl := ctx.Share(&sm.SharedState, 0)
	if !ctx.Publish(sm.Reference.String(), sdl) {
		return ctx.Stop()
	}

	switch sm.initReason {
	case InitReasonCTConstructor:
		return ctx.Jump(sm.stepReadyToWork)
	case InitReasonCTMethod:
		return ctx.Jump(sm.stepGetObjectState)
	case InitReasonVStateReport:
		sm.stateRequestSent = true
		return ctx.Jump(sm.stepWaitState)
	default:
		panic("Not implemented")
	}
}

// we get CallMethod but we have no object data
// we need to ask previous executor
func (sm *SMObject) stepGetObjectState(ctx smachine.ExecutionContext) smachine.StateUpdate {
	if sm.stateRequestSent {
		return ctx.Jump(sm.stepWaitState)
	}

	msg := payload.VStateRequest{
		Callee:           sm.Reference,
		RequestedContent: payload.RequestLatestDirtyState,
	}

	goCtx := ctx.GetContext()
	sm.messageSender.PrepareNotify(ctx, func(svc messagesender.Service) {
		_ = svc.SendRole(goCtx, &msg, insolar.DynamicRoleVirtualExecutor, sm.Reference, sm.pulseSlot.PulseData().PrevPulseNumber())
	}).Send()

	sm.stateRequestSent = true
	return ctx.Jump(sm.stepWaitState)
}

func (sm *SMObject) stepWaitState(ctx smachine.ExecutionContext) smachine.StateUpdate {
	if sm.SharedState.descriptor != nil {
		return ctx.Jump(sm.stepReadyToWork)
	}

	return ctx.Sleep().ThenRepeat()
}

func (sm *SMObject) stepReadyToWork(ctx smachine.ExecutionContext) smachine.StateUpdate {
	sm.SharedState.SetReady()
	return ctx.Jump(sm.stepWaitIndefinitely)
}

func (sm *SMObject) stepWaitIndefinitely(ctx smachine.ExecutionContext) smachine.StateUpdate {
	return ctx.Sleep().ThenRepeat()
}
