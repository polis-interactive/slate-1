package bus

func (b *bus) GetLightCount() int {
	return b.lightingService.GetLightCount()
}

