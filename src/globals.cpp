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

		/*

		// chonkers
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(16, 9)), // column 3, top
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(16, 2)), // column 3, middle
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(9, 2)), // column 3, bottom

		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(16, 9)), // column 4, bottom
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(16, 2)), // column 4, middle
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(9, 2)), // column 4, top

		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(16, 9)), // column 5, top
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(16, 2)), // column 5, middle
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(9, 2)), // column 5, bottom

		// side singles
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(3, 1)), // column 6, bottom
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(3, 1)), // column 6, middle
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(3, 1)), // column 6, top

		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(3, 1)), // column 7, top
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(3, 1)), // column 7, middle
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(3, 1)), // column 7, bottom

		*/
	};

	std::vector<ofPoint> disallowed_positions = {
		ofPoint(0, 5), ofPoint(0, 6),
			ofPoint(0, 15), ofPoint(0, 16), ofPoint(0, 17), ofPoint(0, 18),
		ofPoint(1, 4), ofPoint(1, 17),
	};

	std::vector<ofPoint> permaoff_positions = {
		 ofPoint(0, 4), ofPoint(1, 18), ofPoint(2, 18),
		/*
		ofPoint(0, 0), ofPoint(0, 1), ofPoint(0, 2), ofPoint(0, 3), ofPoint(0, 4),
			ofPoint(0, 19), ofPoint(0, 20),
		ofPoint(1, 0), ofPoint(1, 1), ofPoint(1, 2), ofPoint(1, 3),
			ofPoint(1, 18), // ofPoint(1, 19), ofPoint(1, 20),
		ofPoint(2, 0), ofPoint(2, 1), ofPoint(2, 2), ofPoint(2, 3),
			ofPoint(2, 18), ofPoint(2, 19), ofPoint(2, 20),
			*/
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