package bus

import grpcControl "github.com/polis-interactive/slate-1/api/v1/go"

func (b *bus) HandleControlResponse(response *grpcControl.ControlResponse) {
	isOff := !response.GetOn()
	b.graphicsService.HandleSetState(isOff)
}
