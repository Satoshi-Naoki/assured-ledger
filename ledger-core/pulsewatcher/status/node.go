// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package status

import (
	"github.com/insolar/assured-ledger/ledger-core/application/api/requester"
)

type Node struct {
	URL    string
	Reply  requester.StatusResponse
	ErrStr string
}
