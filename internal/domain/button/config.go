package button

import "time"

type Config interface {
	GetButtonIsGpio() bool
	GetButtonSetup() *Setup
	GetReadFrequency() time.Duration
}
