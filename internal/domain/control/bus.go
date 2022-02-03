package control

import grpcControl "github.com/polis-interactive/slate-1/api/v1/go"

type Bus interface {
	HandleControlResponse(response *grpcControl.ControlResponse)
}
