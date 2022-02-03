package infrastructure

import "fmt"

type ErrorChannelClosed struct {
	ChannelId uint32
}

func (e *ErrorChannelClosed) Error() string {
	return fmt.Sprintf("CLOSED_CHANNEL_%d", e.ChannelId)
}
