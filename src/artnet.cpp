
#include "artnet.h"

ArtnetGraphics::ArtnetGraphics(LedInputInitializer& config) {
	InitializeInput(config);
	receiver_ = new ofxArtnetReceiver();
	receiver_->setup(ARTNET_PORT);
}

void ArtnetGraphics::UpdateInput() {
	while (receiver_->hasMessage()) {
		ofxArtnetMessage m;
		receiver_->getNextData(m);
		handleArtnetMessage(m);
	}
}

void ArtnetGraphics::handleArtnetMessage(ofxArtnetMessage& message) {
	const int universe = message.getUniverse();
	const int size = message.getSize() / 3;
	auto pixels = pixels_->getData();
	for (int i = 0; i < size; i++) {
		const int pixel_position = 170 * universe + i;
		const int color_position = i * 3;
		const int row = pixel_position / grid_width_;
		const int column = pixel_position % grid_width_;
		const ofPoint grid_position = ofPoint(row, column);
		Light* light = Light::GetLightOrNull(grid_position);
		if (light == NULL) continue;
		uint8_t* color_data = message.dataAt(color_position);
		light->SetColor(color_data);
#ifndef __arm__
		const int pixel_x = column * segment_size_;
		const int pixel_y = (grid_height_ - 1 - row) * segment_size_;
		for (int draw_x = 0; draw_x < segment_size_; draw_x++) {
			for (int draw_y = 0; draw_y < segment_size_; draw_y++) {
				const int pixel_y_out = pixel_y + draw_y;
				const int pixel_x_out = pixel_x + draw_x;
				const int pixel_position = pixels_->getPixelIndex(pixel_x_out, pixel_y_out);
				memcpy(pixels + pixel_position, color_data, sizeof(uint8_t) * 3);
			}
		}
#endif
	}
}

void ArtnetGraphics::Teardown() {
	receiver_->~ofxArtnetReceiver();
}
