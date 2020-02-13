// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package executor_test

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/insolar/assured-ledger/ledger-core/v2/insolar"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/bus"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/flow"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/gen"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/payload"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/record"
	"github.com/insolar/assured-ledger/ledger-core/v2/instrumentation/inslogger"
	"github.com/insolar/assured-ledger/ledger-core/v2/ledger/light/executor"
)

func TestDetachedNotifier_Notify(t *testing.T) {
	t.Parallel()

	t.Run("Simple success", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		ctx := flow.TestContextWithPulse(
			inslogger.TestContext(t),
			gen.PulseNumber(),
		)

		objectID := gen.ID()
		closedReqID := gen.ID()
		detachedReqID := gen.ID()
		detachedReq := record.Wrap(&record.OutgoingRequest{
			Reason:     *insolar.NewReference(closedReqID),
			ReturnMode: record.ReturnSaga,
		})
		detachedReqBuf, _ := detachedReq.Marshal()

		opened := []record.CompositeFilamentRecord{
			// wrong
			{
				RecordID: closedReqID,
				Record:   record.Material{Virtual: record.Wrap(&record.IncomingRequest{})},
			},
			// wrong
			{
				RecordID: gen.ID(),
				Record:   record.Material{Virtual: record.Wrap(&record.OutgoingRequest{})},
			},
			// right
			{
				RecordID: detachedReqID,
				Record:   record.Material{Virtual: detachedReq},
			},
			// wrong
			{
				RecordID: gen.ID(),
				Record: record.Material{Virtual: record.Wrap(&record.OutgoingRequest{
					Reason:     *insolar.NewReference(gen.ID()),
					ReturnMode: record.ReturnSaga,
				})},
			},
		}

		sender := bus.NewSenderMock(mc)
		expectedToVirtualMsg, _ := payload.NewMessage(&payload.SagaCallAcceptNotification{
			ObjectID:          objectID,
			DetachedRequestID: detachedReqID,
			Request:           detachedReqBuf,
		})

		sender.SendRoleMock.Inspect(func(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference) {
			require.Equal(t, expectedToVirtualMsg.Payload, msg.Payload)
			require.Equal(t, insolar.DynamicRoleVirtualExecutor, role)
			require.Equal(t, *insolar.NewReference(objectID), object)
		}).Return(make(chan *message.Message), func() {})

		dn := executor.NewDetachedNotifierDefault(sender)
		dn.Notify(ctx, opened, objectID, closedReqID)
	})
}
