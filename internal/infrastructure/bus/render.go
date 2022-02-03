package bus

import (
	"github.com/polis-interactive/slate-1/internal/types"
	"sync"
)

func (b *bus) GetLightCount() int {
	return b.lightingService.GetLightCount()
}

func (b *bus) CopyLightsToColorBuffer(rawPbOut []types.Color) error {
	lights, preLockedLightsMutex := b.lightingService.GetLights()
	pbIn, preLockedGraphicsMutex := b.graphicsService.GetPb()
	defer func(lightsMu *sync.RWMutex, graphicsMu *sync.RWMutex) {
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

func (b *bus) CopyLightsToUint32Buffer(rawUint32BuffOut []uint32) (isOff bool, err error) {
	lights, preLockedLightsMutex := b.lightingService.GetLights()
	pbIn, preLockedGraphicsMutex := b.graphicsService.GetPb()
	defer func(lightsMu *sync.RWMutex, graphicsMu *sync.RWMutex) {
		lightsMu.RUnlock()
		graphicsMu.RUnlock()
	}(preLockedLightsMutex, preLockedGraphicsMutex)
	isOff = true
	for _, l := range lights {
		if !l.Show {
			continue
		}
		bits := pbIn.GetPixelPointer(&l.Position).ToBits()
		if isOff && bits != 0 {
			isOff = false
		}
		rawUint32BuffOut[l.Pixel] = bits
	}
	return isOff, nil
}
