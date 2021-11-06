
#include "types.h"
#include "globals.h"

void LedInput::InitializeInput(LedInputInitializer& config) {
	const auto grid_definitions = config._grid_definitions;
	min_y_ = grid_definitions._min_grid_y;
	min_x_ = grid_definitions._min_grid_x;
	grid_width_ = grid_definitions._max_grid_x - grid_definitions._min_grid_x + 1;
	grid_height_ = grid_definitions._max_grid_y - grid_definitions._min_grid_y + 1;
	grid_pixels_ = grid_width_ * grid_height_;
	segment_size_ = config._segment_size;
	total_width_ = grid_width_ * segment_size_;
	total_height_ = grid_height_ * segment_size_;
	pixels_ = new ofPixels();
	pixels_->allocate(total_width_, total_height_, OF_IMAGE_COLOR_ALPHA);
	pixels_->setColor(ofColor::black);
	fbo_ = new ofFbo();
	fbo_->allocate(total_width_, total_height_);
}

void LedInput::DrawInput() {
	for (const ofPoint& point : globals::permaoff_positions) {
		TurnOffRgbPixel(point, 2);
	}
	fbo_->getTexture().getTextureData().bFlipTexture = true;
	fbo_->getTexture().bind();
	fbo_->getTexture().loadData(pixels_->getData(), total_width_, total_height_, GL_RGBA);
	fbo_->getTexture().unbind();
	fbo_->draw(0, 0);
}


const ofColor LedInput::GetRGBPixel(const ofPoint& position) {
	const int mapped_x = (position.x - min_x_) * segment_size_;
	const int mapped_y = (position.y - min_y_) * segment_size_;
	ofColor c = pixels_->getColor(mapped_x, mapped_y);
	return c;
}

void LedInput::TurnOffRgbPixel(const ofPoint& position) {
	TurnOffRgbPixel(position, 0);
}
	
void LedInput::TurnOffRgbPixel(const ofPoint& position, int is_red_or_something) {
#ifndef __arm__
	auto pixels = pixels_->begin();
	for (int draw_x = 0; draw_x < segment_size_; draw_x++) {
		for (int draw_y = 0; draw_y < segment_size_; draw_y++) {
			const int pixel_x = (position.x - min_x_) * segment_size_ + draw_x;
			const int pixel_y = (position.y - min_y_) * segment_size_ + draw_y;
			const int pixel_position = pixels_->getPixelIndex(pixel_x, pixel_y);
			if (is_red_or_something == 0) {
				memset(pixels + pixel_position, 255, sizeof(uint8_t));
				memset(pixels + pixel_position + 1, 0, sizeof(uint8_t) * 2);
			}
			else if (is_red_or_something == 1) {
				memset(pixels + pixel_position, 0, sizeof(uint8_t));
				memset(pixels + pixel_position + 1, 255, sizeof(uint8_t));
				memset(pixels + pixel_position + 2, 0, sizeof(uint8_t));
			}
			else if (is_red_or_something == 2) {
				memset(pixels + pixel_position, 0, sizeof(uint8_t) * 2);
				memset(pixels + pixel_position + 2, 255, sizeof(uint8_t));
			}
			else {
				memset(pixels + pixel_position, 255, sizeof(uint8_t) * 3);
			}
		}
	}
#endif
}

void LedInput::SetProperPixelColor(const ofPoint& position, const ofColor& c) {

	auto pixel_data = pixels_->getData();

	auto pixels = pixels_->begin();
	for (int draw_x = 0; draw_x < segment_size_; draw_x++) {
		for (int draw_y = 0; draw_y < segment_size_; draw_y++) {
			const int pixel_x = (position.x - min_x_) * segment_size_ + draw_x;
			const int pixel_y = (position.y - min_y_) * segment_size_ + draw_y;
			const int pixel_position = pixels_->getPixelIndex(pixel_x, pixel_y);
			memcpy(pixel_data + pixel_position, c.v, sizeof(uint8_t) * 3);
		}
	}

}

void LedInput::SetProperPixelTint(const ofPoint& position, const float& tinter) {
	for (int draw_x = 0; draw_x < segment_size_; draw_x++) {
		for (int draw_y = 0; draw_y < segment_size_; draw_y++) {
			const int pixel_x = (position.x - min_x_) * segment_size_ + draw_x;
			const int pixel_y = (position.y - min_y_) * segment_size_ + draw_y;
			const int pixel_position = pixels_->getPixelIndex(pixel_x, pixel_y);
			const auto oc = pixels_->getColor(pixel_x, pixel_y);
			pixels_->setColor(pixel_x, pixel_y, ofColor::white.getLerped(oc, tinter));
		}
	}
}

void LedInput::SetPixelColor(const ofPoint& p, const ofColor& c) {

	auto pixel_data = pixels_->getData();

	for (int draw_x = 0; draw_x < segment_size_; draw_x++) {
		for (int draw_y = 0; draw_y < segment_size_; draw_y++) {
			const int pixel_x_out = p.x * segment_size_ + draw_x;
			const int pixel_y_out = p.y * segment_size_ + draw_y;
			const int pixel_position = pixels_->getPixelIndex(pixel_x_out, pixel_y_out);

			memcpy(pixel_data + pixel_position, c.v, sizeof(uint8_t) * 3);
		}
	}

}
