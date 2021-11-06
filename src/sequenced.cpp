
#include "sequenced.h"
#include "lighting.h"

SequencedGraphics::SequencedGraphics(LedInputInitializer& config) {
	InitializeInput(config);
	LedInput* base_ptr = this;
	shader_ = new Shader(base_ptr, "dope_shader");
}

void SequencedGraphics::UpdateInput() {
	shader_->UpdateShader();
}



