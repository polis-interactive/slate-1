#include "globals.h"


namespace globals {

	/* prog config */

	const int output_width = 750;
	const int output_height = 1000;


	
	/* strip config */

	const float gamma = 1.2;
	const uint8_t brightness = 255;



	/* board config */
	// left is with respect to the tiles looking at them from the front

	std::vector<BoardConfiguration> board_configs = {

		// chonky middle
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(16, 9)), // BR
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(16, 2)), // BL
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_180, ofPoint(9, 2)), // ML
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_90, ofPoint(2, 2)), // TL
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_0, ofPoint(2, 9)), // TR
		BoardConfiguration(BoardType::BOARD_7X7, BoardOrientation::ORIENT_270, ofPoint(9, 9)), // MR


		// right face
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(8, 17)), // BR, T
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(8, 16)), // TR, T
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(15, 17)), // BR, B
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_0, ofPoint(15, 16)), // TR, B

		// top face
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_90, ofPoint(0, 2)), // BR, L
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_90, ofPoint(1, 2)), // TR, L
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_90, ofPoint(0, 9)), // BR, R
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_90, ofPoint(1, 9)), // TR, R

		// bottom face
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_270, ofPoint(23, 9)), // BR, R
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_270, ofPoint(24, 9)), // TR, R
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_270, ofPoint(23, 2)), // BR, L
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_270, ofPoint(24, 2)), // TR, L

		// left face
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(3, 1)), // TR, T
		BoardConfiguration(BoardType::BOARD_1X7, BoardOrientation::ORIENT_180, ofPoint(3, 0)), // BR, T

	};

	std::vector<ofPoint> disallowed_positions = {
		// chonky middle
		ofPoint( 2,  3), ofPoint( 2,  4), ofPoint( 2, 12), ofPoint( 2, 13), ofPoint( 2, 14), ofPoint(2, 15),
		ofPoint( 3, 12), ofPoint( 3, 13), ofPoint( 3, 14), ofPoint( 3, 15),
		ofPoint( 4, 11), ofPoint( 4, 12), ofPoint( 4, 13), ofPoint( 4, 14), ofPoint( 4, 15),
		ofPoint( 5, 10), ofPoint( 5, 13), ofPoint( 5, 14), ofPoint( 5, 15),
		ofPoint( 6, 10), ofPoint( 6, 11), ofPoint( 6, 12), ofPoint( 6, 13), ofPoint( 6, 14),
		ofPoint( 7, 12),
		ofPoint( 8,  4), ofPoint( 8,  5), ofPoint( 8, 12),

		ofPoint( 9,  2), ofPoint( 9,  4), ofPoint( 9,  8), ofPoint( 9, 11), ofPoint( 9, 12),
		ofPoint(10,  3), ofPoint(10,  4), ofPoint(10,  9), ofPoint(10, 10), ofPoint(10, 11), ofPoint(10, 12),
		ofPoint(11,  2), ofPoint(11,  3), ofPoint(11, 10), ofPoint(11, 11), ofPoint(11, 12), ofPoint(11, 13),
		ofPoint(12,  2), ofPoint(12,  3), ofPoint(12,  8), ofPoint(12,  9), ofPoint(12, 10), ofPoint(12, 11), ofPoint(12, 12), ofPoint(12, 13),
		ofPoint(13,  2), ofPoint(13,  3), ofPoint(13,  4), ofPoint(13,  8), ofPoint(13,  9), ofPoint(13, 10), ofPoint(13, 11),
		ofPoint(14,  2), ofPoint(14,  3), ofPoint(14,  4), ofPoint(14,  5), ofPoint(14,  7), ofPoint(14,  8), ofPoint(14,  9), ofPoint(14, 10), ofPoint(14, 11), ofPoint(14, 12),
		ofPoint(15,  2), ofPoint(15,  3), ofPoint(15,  4), ofPoint(15,  5), ofPoint(15,  6), ofPoint(15,  7), ofPoint(15,  8), ofPoint(15,  9), ofPoint(15, 10), ofPoint(15, 11), ofPoint(15, 12),
		ofPoint(16,  2), ofPoint(16,  3), ofPoint(16,  4), ofPoint(16,  5), ofPoint(16,  6), ofPoint(16,  9), ofPoint(16, 10), ofPoint(16, 11), ofPoint(16, 12), ofPoint(16, 13),
		ofPoint(17,  2), ofPoint(17,  3), ofPoint(17,  4), ofPoint(17,  5), ofPoint(17,  6), ofPoint(17, 13), ofPoint(17, 14),
		ofPoint(18,  2), ofPoint(18,  3), ofPoint(18,  4), ofPoint(18,  5), ofPoint(18, 13), ofPoint(18, 14),
		ofPoint(19,  2), ofPoint(19,  3), ofPoint(19,  4), ofPoint(19,  8), ofPoint(19, 12), ofPoint(19, 13),
		ofPoint(20,  2), ofPoint(20,  3), ofPoint(20,  4), ofPoint(20,  5), ofPoint(20,  8), ofPoint(20,  9), ofPoint(20, 10), ofPoint(20, 12), ofPoint(20, 13),
		ofPoint(21,  2), ofPoint(21,  3), ofPoint(21,  4), ofPoint(21,  5), ofPoint(21,  7), ofPoint(21,  8), ofPoint(21,  9), ofPoint(21, 12), ofPoint(21, 13),
		ofPoint(22,  2), ofPoint(22,  3), ofPoint(22,  4), ofPoint(22,  8),
	
		// top face
		ofPoint( 0,  2), ofPoint( 0, 11), ofPoint( 0, 12), ofPoint( 0, 15),
		ofPoint( 1,  2), ofPoint( 1,  3), ofPoint( 1, 14), ofPoint( 1, 15),

		// right face
		ofPoint( 9, 16),
		ofPoint(21, 17),

		// bottom face
		ofPoint(23, 2), ofPoint(23, 10),
		ofPoint(24, 2), ofPoint(24, 3),

		// left face
		ofPoint(7, 0), ofPoint(8, 0), ofPoint(9, 0),
		ofPoint(8, 1), ofPoint(9, 1),
	};

	std::vector<ofPoint> corner_positions = {
		// Corners, TL
		ofPoint(0, 0), ofPoint(0, 1), ofPoint(1, 0), ofPoint(1, 1),
		// BL
		ofPoint(23, 0), ofPoint(24, 0), ofPoint(23, 1), ofPoint(24, 1),
		// BR
		ofPoint(23, 16), ofPoint(23, 17), ofPoint(24, 16), ofPoint(24, 17),
		// TR
		ofPoint(0, 16), ofPoint(0, 17), ofPoint(1, 16), ofPoint(1, 17),
	};

	std::vector<ofPoint> permaoff_positions = {
		// left side
		ofPoint(2, 0), ofPoint(10, 0), ofPoint(11, 0), ofPoint(12, 0),  ofPoint(13, 0), ofPoint(14, 0), ofPoint(15, 0), ofPoint(16, 0),
			ofPoint(17, 0),  ofPoint(18, 0), ofPoint(19, 0), ofPoint(20, 0), ofPoint(21, 0), ofPoint(22, 0),
		ofPoint(2, 1), ofPoint(10, 1), ofPoint(11, 1), ofPoint(12, 1),  ofPoint(13, 1), ofPoint(14, 1), ofPoint(15, 1), ofPoint(16, 1),
			ofPoint(17, 1),  ofPoint(18, 1), ofPoint(19, 1), ofPoint(20, 1), ofPoint(21, 1), ofPoint(22, 1),
		// right side
		ofPoint(22, 16), ofPoint(2, 16),  ofPoint(3, 16),  ofPoint(4, 16),  ofPoint(5, 16),  ofPoint(6, 16),  ofPoint(7, 16),
		ofPoint(22, 17), ofPoint(2, 17),  ofPoint(3, 17),  ofPoint(4, 17),  ofPoint(5, 17),  ofPoint(6, 17),  ofPoint(7, 17),
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

	/* stream config */
	std::vector<std::vector<ofPoint>> stream_positions {
		{ ofPoint( 0,  5) },
		{ ofPoint( 1,  6) },
		{ ofPoint( 1,  7), ofPoint( 2,  6) },
		{ ofPoint( 2,  7) },
		{ ofPoint( 3,  7) },
		{ ofPoint( 3,  8) },
		{ ofPoint( 3,  9), ofPoint( 4,  8) },
		{ ofPoint( 5,  9) },
		{ ofPoint( 5, 11) },
		{ ofPoint( 8, 13) },
		{ ofPoint( 9, 13) },
		{ ofPoint(10, 14), ofPoint(10, 15), ofPoint(10, 16) },
		{ ofPoint(11, 14), ofPoint(11, 15), ofPoint(11, 16) },
		{ ofPoint(12, 14) },
		{ ofPoint(13, 12), ofPoint(13, 13), ofPoint(13, 14) },
		{ ofPoint(14, 13) },
		{ ofPoint(15, 15) },
		{ ofPoint(17, 12) },
		{ ofPoint(17,  9) },
		{ ofPoint(17,  8) },
		{ ofPoint(18,  8), ofPoint(17, 7) },
		{ ofPoint(19,  5), ofPoint(19, 6) },
		{ ofPoint(22,  5) },
		{ ofPoint(23,  3), ofPoint(23, 4) },
		{ ofPoint(24,  4) },
	};
}