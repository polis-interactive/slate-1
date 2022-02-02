package bus

func (b *bus) HandleButtonPush() {
	_ = b.graphicsService.HandleButtonPush()
	// might need to let alexa know its off
}
