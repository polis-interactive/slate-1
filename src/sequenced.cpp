
#include "sequenced.h"
#include "lighting.h"

SequencedGraphics::SequencedGraphics(LedInputInitializer& config) {
	InitializeInput(config);
	LedInput* base_ptr = this;
	// shader_ = new Shader(base_ptr, "watershader");
	// fish_runner_ = new FishRunner(base_ptr);
}

void SequencedGraphics::UpdateInput() {
	// shader_->UpdateShader();
	// fish_runner_->UpdateFishRunner();
}



