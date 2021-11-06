#pragma once

#include "boards.h"
#include "ofxRpiWs281x.h"
#include "types.h"
#include "sequenced.h"
#include <vector>


namespace globals {

	/* prog config */
	const extern int output_width;
	const extern int output_height;
	const extern int frame_rate;


	/* strip config */

	const extern float gamma;
	const extern uint8_t brightness;



	/* board config */

	extern std::vector<BoardConfiguration> board_configs;
	extern std::vector<ofPoint> disallowed_positions;



	/* input config */

	const extern int segment_size;
	const extern LedInputType input_type;

	/* shader config */

	const extern float shader_gamma;
	const extern float shader_speed;
	const extern float shader_scale;
	const extern float shader_brightness;
	const extern float shader_contrast;

	/* output config */
	extern std::vector<ofPoint> permaoff_positions;
}
