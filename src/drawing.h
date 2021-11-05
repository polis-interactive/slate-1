#pragma once

#include "types.h"

class DrawingOutput : public LedInput {
public:
    DrawingOutput(GridDefintions, int grid_size);
    DrawingOutput() = delete;
    ~DrawingOutput() = delete;
    DrawingOutput(const DrawingOutput&) = delete;
    DrawingOutput& operator=(const DrawingOutput&) = delete;
};