// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package execute

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/assured-ledger/ledger-core/application/builtin/proxy/testwallet"
	"github.com/insolar/assured-ledger/ledger-core/conveyor"
	"github.com/insolar/assured-ledger/ledger-core/conveyor/smachine"
	"github.com/insolar/assured-ledger/ledger-core/insolar"
	"github.com/insolar/assured-ledger/ledger-core/insolar/contract"
	"github.com/insolar/assured-ledger/ledger-core/insolar/payload"
	"github.com/insolar/assured-ledger/ledger-core/instrumentation/inslogger/instestlogger"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
	"github.com/insolar/assured-ledger/ledger-core/reference"
	"github.com/insolar/assured-ledger/ledger-core/testutils"
	"github.com/insolar/assured-ledger/ledger-core/testutils/gen"
	"github.com/insolar/assured-ledger/ledger-core/vanilla/longbits"
	"github.com/insolar/assured-ledger/ledger-core/virtual/callregistry"
	"github.com/insolar/assured-ledger/ledger-core/virtual/callsummary"
	"github.com/insolar/assured-ledger/ledger-core/virtual/testutils/shareddata"
	"github.com/stretchr/testify/require"
)

func TestSMExecute_PublishVCallResultToCallSummarySM(t *testing.T) {
	var (
		ctx = instestlogger.TestContext(t)
		mc  = minimock.NewController(t)

		pd          = pulse.NewFirstPulsarData(10, longbits.Bits256{})
		pulseSlot   = conveyor.NewPresentPulseSlot(nil, pd.AsRange())
		outgoingRef = reference.NewRecordOf(gen.UniqueGlobalRef(), gen.UniqueLocalRefWithPulse(pd.PulseNumber))

		callFlags = payload.BuildCallFlags(contract.CallTolerable, contract.CallDirty)
	)

	class := gen.UniqueGlobalRef()

	request := &payload.VCallRequest{
		CallType:            payload.CTConstructor,
		CallFlags:           callFlags,
		CallSiteDeclaration: testwallet.GetClass(),
		CallSiteMethod:      "New",
		CallOutgoing:        outgoingRef,
		Callee:              class,
		Arguments:           insolar.MustSerialize([]interface{}{}),
	}

	smExecute := SMExecute{
		Payload:   request,
		pulseSlot: &pulseSlot,
	}

	smExecute = expectedInitState(ctx, smExecute)

	res := payload.VCallResult{
		Callee:       class,
		CallOutgoing: outgoingRef,
	}

	smExecute.execution.Result = &res
	smExecute.migrationHappened = true

	{
		execCtx := smachine.NewExecutionContextMock(mc).
			JumpMock.Set(testutils.AssertJumpStep(t, smExecute.stepAwaitSMCallSummary))

		smExecute.stepFinishRequest(execCtx)
	}

	sharedCallSummary := callsummary.SharedCallSummary{
		Requests: callregistry.NewObjectRequestTable(),
	}

	{
		outgoingRef = reference.NewSelf(outgoingRef.GetLocal())

		workingTable := callregistry.NewWorkingTable()
		workingTable.Add(contract.CallTolerable, smExecute.execution.Outgoing)
		workingTable.SetActive(contract.CallTolerable, smExecute.execution.Outgoing)

		sharedCallSummary.Requests.AddObjectCallResults(outgoingRef, callregistry.ObjectCallResults{
			CallResults: workingTable.GetResults(),
		})
	}

	{
		sdlSummarySync := smachine.NewUnboundSharedData(&smachine.SyncLink{})

		getPublishedLink := func(key interface{}) (s1 smachine.SharedDataLink) {
			switch key.(type) {
			case callsummary.SummarySyncKey:
				return sdlSummarySync
			default:
				t.Fatal("Unexpected key type")
			}

			return smachine.SharedDataLink{}
		}

		execCtx := smachine.NewExecutionContextMock(mc).
			GetPublishedLinkMock.Set(getPublishedLink).
			UseSharedMock.Set(shareddata.CallSharedDataAccessor).
			AcquireForThisStepMock.Expect(smachine.SyncLink{}).Return(true).
			JumpMock.Set(testutils.AssertJumpStep(t, smExecute.stepPublishDataCallSummary))

		smExecute.stepAwaitSMCallSummary(execCtx)
	}

	{
		sdlSummaryShare := smachine.NewUnboundSharedData(&sharedCallSummary)

		getPublishedLink := func(key interface{}) (s1 smachine.SharedDataLink) {
			switch key.(type) {
			case callsummary.SummarySharedKey:
				return sdlSummaryShare
			default:
				t.Fatal("Unexpected key type")
			}

			return smachine.SharedDataLink{}
		}

		execCtx := smachine.NewExecutionContextMock(mc).
			GetPublishedLinkMock.Set(getPublishedLink).
			UseSharedMock.Set(shareddata.CallSharedDataAccessor).
			JumpMock.Set(testutils.AssertJumpStep(t, smExecute.stepSendDelegatedRequestFinished))

		smExecute.stepPublishDataCallSummary(execCtx)
	}

	workingTable, ok := sharedCallSummary.Requests.GetObjectCallResults(outgoingRef)
	require.Equal(t, 1, len(workingTable.CallResults))

	result, ok := workingTable.CallResults[smExecute.execution.Outgoing]

	require.True(t, ok)
	require.NotNil(t, result.Result)
	require.Equal(t, &res, result.Result)

	mc.Finish()
}
