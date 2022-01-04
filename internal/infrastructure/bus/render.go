package bus

import (
	"github.com/polis-interactive/slate-1/internal/types"
	"sync"
)

func (b *bus) GetLightCount() int {
	return b.lightingService.GetLightCount()
}

func (b *bus) CopyLightsToBuffer(rawPbOut []types.Color) error {
	lights, preLockedLightsMutex := b.lightingService.GetLights()
	pbIn, preLockedGraphicsMutex := b.graphicsService.GetPb()
	defer func (lightsMu *sync.RWMutex, graphicsMu *sync.RWMutex) {
		lightsMu.RUnlock()
		graphicsMu.RUnlock()
	}(preLockedLightsMutex, preLockedGraphicsMutex)
	for _, l := range lights {
		if !l.Show {
			continue
		}
		rawPbOut[l.Pixel] = pbIn.GetPixel(&l.Position)
	}
	return nil
}