package message

import (
	"time"

	"github.com/ZalgoNoise/sysprobe/types"
)

// MakeMsg function - it builds a new Message struct,
// containing the data in Internet, Battery, as well
// as the current Unix timestamp
func MakeMsg(i *types.Internet, b *types.Battery) *types.Message {
	msg := &types.Message{}

	msg.Internet = *i
	msg.Battery = *b
	msg.Timestamp = int32(time.Now().Unix())

	return msg
}
