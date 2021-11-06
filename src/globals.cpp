#include "globals.h"


namespace globals {

	/* prog config */

	const int output_width = 750;
	const int output_height = 1000;
	const int frame_rate = 20;


	
	/* strip config */

	const float gamma = 1.2;
	const uint8_t brightness = 255;



	/* board config */
	// left is with respect to the tiles looking at them from the front

	std::vector<BoardConfiguration> board_configs = {

		// left singles
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(0, 12)), // column 0, top
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(0,5)), // column 0, bottom
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(1, 4)), // column 1, bottom
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(1, 11)), // column 1, top
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(2, 11)), // column 2, top
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(2, 4)), // column 2, bottom

		// chonkers
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_90, ofPoint(3, 0)), // column 3, bottom
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_90, ofPoint(3, 7)), // column 3, middle
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(3, 14)), // column 3, top

		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_90, ofPoint(10, 14)), // column 4, top
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(10, 7)), // column 4, middle
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(10, 0)), // column 4, bottom

		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_90, ofPoint(17, 0)), // column 5, bottom
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_90, ofPoint(17, 7)), // column 5, middle
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_90, ofPoint(17, 14)), // column 5, top

		// side singles
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(23, 14)), // column 6, top
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(23, 7)), // column 6, middle
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(23, 0)), // column 6, bottom

		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(23, 0)), // column 7, bottom
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(23, 7)), // column 7, middle
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(23, 14)), // column 7, top
	};

	std::vector<ofPoint> disallowed_positions = {
		ofPoint(0, 5), ofPoint(0, 6),
			ofPoint(0, 15), ofPoint(0, 16), ofPoint(0, 17), ofPoint(0, 18),
		ofPoint(1, 4), ofPoint(1, 17),
		ofPoint(3, 0), ofPoint(3, 1), ofPoint(3, 2),
			ofPoint(3, 19), ofPoint(3, 20),
		ofPoint(4, 0), ofPoint(4, 1),
			ofPoint(4, 20),
		ofPoint(5, 0), ofPoint(5, 1),
			ofPoint(5, 20),
		ofPoint(6, 0),
		ofPoint(7, 0),
		ofPoint(8, 0),
		ofPoint(9, 0),
		ofPoint(10, 0),
		ofPoint(11, 0),
		ofPoint(12, 0),
		ofPoint(13, 0),
		ofPoint(14, 0),
		ofPoint(15, 0),
		ofPoint(16, 0),
		ofPoint(17, 0),
		ofPoint(18, 0),
		ofPoint(19, 0),
		ofPoint(20, 0),
		ofPoint(21, 0),
		ofPoint(22, 0),
		ofPoint(23, 0),
	};

	std::vector<ofPoint> permaoff_positions = {
		ofPoint(0, 0), ofPoint(0, 1), ofPoint(0, 2), ofPoint(0, 3), ofPoint(0, 4),
			ofPoint(0, 19), ofPoint(0, 20),
		ofPoint(1, 0), ofPoint(1, 1), ofPoint(1, 2), ofPoint(1, 3),
			ofPoint(1, 18), ofPoint(1, 19), ofPoint(1, 20),
		ofPoint(2, 0), ofPoint(2, 1), ofPoint(2, 2), ofPoint(2, 3),
			ofPoint(2, 18), ofPoint(2, 19), ofPoint(2, 20),

	};

	/* input config */
#ifndef __arm__
	const int segment_size = 30;
#else
	const int segment_size = 1;
#endif
	const LedInputType input_type = LedInputType::SEQUENCED;

	/* shader config */

	const float shader_gamma = 3.75;
	const float shader_speed = 0.2;

	const float shader_scale = 1.5;
	const float shader_brightness = 0.6;
	const float shader_contrast = 0.4;

}