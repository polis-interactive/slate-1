
#include "boards.h"
#include "globals.h"

std::pair<LedDefinitions, int> SetupLedBoards(std::vector<BoardConfiguration> &configs) {
	auto led_output = LedDefinitions();
	int last_led_asigned = 0;
	for (BoardConfiguration &config : configs) {

		const bool board_is_large = config._board_type == BoardType::BOARD_7X7;
		const int led_count = board_is_large ? 49 : 7;
		const int max_nominal_y = board_is_large ? 6 : 0;
		const ofPoint& starting_position = config._starting_point;
		const auto& orient = config._board_orientation;

		for (int i = 0; i < led_count; i++) {

			const int y_pos = (int) i / 7;
			const bool is_odd_y_pos = y_pos % 2;
			const int x_pos = is_odd_y_pos ? 6 - i % 7 : i % 7;
			ofPoint position = ofPoint(starting_position.x, starting_position.y);

			switch (orient) {
			case BoardOrientation::ORIENT_0:
				position = ofPoint(position.x + y_pos, position.y + (6 - x_pos));
				break;
			case BoardOrientation::ORIENT_90:
				position = ofPoint(position.x + x_pos, position.y + y_pos);
				
				break;
			case BoardOrientation::ORIENT_180:
				position = ofPoint(position.x + (max_nominal_y - y_pos), position.y + x_pos);
				break;
			case BoardOrientation::ORIENT_270:
				position = ofPoint(position.x + (6 - x_pos), position.y + (max_nominal_y - y_pos));
				break;
			}

			const bool is_allowed = std::find(
				globals::disallowed_positions.begin(), globals::disallowed_positions.end(), position
			) == globals::disallowed_positions.end();


			auto proxy = new LedProxy(
				last_led_asigned, position, is_allowed
			);
			led_output.push_back(proxy);
			last_led_asigned++;
		}
	}
	auto output = std::make_pair(led_output, last_led_asigned);
	return output;
}