#pragma once

#include "ofPoint.h"

#include "types.h"

#include <utility>
#include <vector>

enum class BoardType : int {
    BOARD_7X7 = 0,
    BOARD_1X7 = 1
};

// cc rotation
enum class BoardOrientation : int {
    ORIENT_0 = 0,
    ORIENT_90 = 1,
    ORIENT_180 = 2,
    ORIENT_270 = 3
};

struct BoardConfiguration {
    // some sane defaults
    BoardConfiguration(BoardType board_type, BoardOrientation board_orientation, ofPoint starting_point) :
        _board_type(board_type), _board_orientation(board_orientation), _starting_point(starting_point)
    {};
    BoardType _board_type;
    BoardOrientation _board_orientation;
    ofPoint _starting_point;
};

std::pair<LedDefinitions, int> SetupLedBoards(std::vector<BoardConfiguration> &configs);