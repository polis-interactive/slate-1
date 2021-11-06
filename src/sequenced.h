#pragma once

#include "types.h"
#include "ofGraphics.h"
#include "shader.h"


class SequencedGraphics : public LedInput {
public:
	SequencedGraphics(LedInputInitializer&);
	SequencedGraphics() = delete;
	~SequencedGraphics() = delete;
public:
	void UpdateInput();
	void Teardown() {};
private:
	Shader* shader_;
};

