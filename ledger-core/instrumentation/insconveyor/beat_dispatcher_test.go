// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package insconveyor

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/require"

	"github.com/insolar/assured-ledger/ledger-core/appctl/beat"
	"github.com/insolar/assured-ledger/ledger-core/conveyor"
	"github.com/insolar/assured-ledger/ledger-core/insolar/payload"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
	"github.com/insolar/assured-ledger/ledger-core/vanilla/throw"
)

func newDispatcherWithConveyor(factoryFn conveyor.PulseEventFactoryFunc) beat.Dispatcher {
	ctx := context.Background()

	pulseConveyor := conveyor.NewPulseConveyor(ctx,	DefaultConfigNoEventless(),
		factoryFn, nil)

	return NewConveyorDispatcher(ctx, pulseConveyor)
}

func TestConveyorDispatcher_ErrorUnmarshalHandling(t *testing.T) {
	msgDispatcher := newDispatcherWithConveyor(nil)
	msg := message.NewMessage("", nil)
	require.False(t, isMessageAcked(msg))
	require.NotPanics(t, func() {
		require.Error(t, msgDispatcher.Process(msg))
	})
	require.True(t, isMessageAcked(msg))
}

func TestConveyorDispatcher_WrongMetaTypeHandling(t *testing.T) {
	msgDispatcher := newDispatcherWithConveyor(nil)
	req := payload.VCallRequest{}
	pl, _ := req.Marshal()
	msg := message.NewMessage("", pl)
	require.False(t, isMessageAcked(msg))
	require.NotPanics(t, func() {
		require.Error(t, msgDispatcher.Process(msg))
	})
	require.True(t, isMessageAcked(msg))
}

func TestConveyorDispatcher_PanicInAddInputHandling(t *testing.T) {
	msgDispatcher := newDispatcherWithConveyor(
		func(context.Context, conveyor.InputEvent, conveyor.InputContext) (conveyor.InputSetup, error) {
			panic(throw.E("handler panic"))
		})
	meta := payload.Meta{Pulse: pulse.Number(pulse.MinTimePulse + 1)}
	metaPl, _ := meta.Marshal()
	msg := message.NewMessage("", metaPl)
	require.False(t, isMessageAcked(msg))
	require.Panics(t, func() {
		require.NoError(t, msgDispatcher.Process(msg))
	})
	require.True(t, isMessageAcked(msg))
}

func TestConveyorDispatcher_ErrorInAddInputHandling(t *testing.T) {
	msgDispatcher := newDispatcherWithConveyor(
		func(context.Context, conveyor.InputEvent, conveyor.InputContext) (conveyor.InputSetup, error) {
			return conveyor.InputSetup{}, throw.E("handler error")
		})
	meta := payload.Meta{Pulse: pulse.Number(pulse.MinTimePulse + 1)}
	metaPl, _ := meta.Marshal()
	msg := message.NewMessage("", metaPl)
	require.False(t, isMessageAcked(msg))
	require.NotPanics(t, func() {
		require.Error(t, msgDispatcher.Process(msg))
	})
	require.True(t, isMessageAcked(msg))
}

func isMessageAcked(msg *message.Message) bool {
	select {
	case _, ok := <-msg.Acked():
		return !ok
	default:
		return false
	}
}