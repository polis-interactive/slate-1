#pragma once

#include "ofPoint.h"
#include "ofColor.h"
#include "ofFbo.h"
#include "ofPixels.h"
#include "ofGraphics.h"
#include "ofUtils.h"
#include "ofShader.h"
#include "ofAppRunner.h"
#include "ofImage.h"


#include <tuple>

struct LedProxy {
	LedProxy(int pixel_number, ofPoint pixel_position, bool is_on) :
		_pixel_number(pixel_number), _pixel_position(pixel_position), _is_on(is_on)
	{}
	const int _pixel_number;
	const ofPoint _pixel_position;
	const bool _is_on;
};

typedef std::vector<LedProxy*> LedDefinitions;

struct GridDefintions {
	int _min_grid_x = 0;
	int _max_grid_x = 0;
	int _min_grid_y = 0;
	int _max_grid_y = 0;

	void TrySetMinMax(const int x_cord, const int y_cord) {
		if (!has_been_set_) {
			_min_grid_x = x_cord;
			_max_grid_x = x_cord;
			_min_grid_y = y_cord;
			_max_grid_y = y_cord;
			has_been_set_ = true;
		}
		else {
			_min_grid_x = std::min(_min_grid_x, x_cord);
			_max_grid_x = std::max(_max_grid_x, x_cord);
			_min_grid_y = std::min(_min_grid_y, y_cord);
			_max_grid_y = std::max(_max_grid_y, y_cord);
		}
	}
private:
	bool has_been_set_ = false;
};


enum class LedInputType : int {
	ARTNET = 0,
	SEQUENCED = 1,
};

struct LedInputInitializer {
	LedInputInitializer(GridDefintions grid_definitions, int segment_size)
		: _grid_definitions(grid_definitions), _segment_size(segment_size)
	{}
	const GridDefintions _grid_definitions;
	const int _segment_size;
};


/* 
	might think about ripping implementaiton into types.cpp
*/

class LedInput {
public:
	void InitializeInput(LedInputInitializer& config);
	virtual void UpdateInput() = 0;
	virtual void Teardown() = 0;
public:
	void DrawInput();
	const ofColor GetRGBPixel(const ofPoint& position);
	void TurnOffRgbPixel(const ofPoint& position);
	void TurnOffRgbPixel(const ofPoint& position, int is_red_or_something);
	void SetPixelColor(const ofPoint& p, const ofColor& c);
	void SetProperPixelColor(const ofPoint& position, const ofColor& c);
	void SetProperPixelTint(const ofPoint& position, const float &tinter);
public:
	ofFbo* fbo_;
	ofPixels* pixels_;
	ofImage* img_;
	int segment_size_;
	int min_x_;
	int min_y_;
	int grid_width_;
	int grid_height_;
	int grid_pixels_;
	int total_width_;
	int total_height_;
};

const float DEG_TO_UINT_MULTIPLIER = 255.0 / 360.0;
