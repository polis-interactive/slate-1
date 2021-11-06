
#include "shader.h"

#include "globals.h"

Shader::Shader(LedInput* input, std::string shaderName): led_input_(input) {
	shader_ = new ofShader();
	shader_name_ = shaderName;
	shader_->load(shader_name_);
}

void Shader::UpdateShader() {
	static float time;
	static float width = led_input_->total_width_;
	static float height = led_input_->total_height_;
	static float resolution[] = { width, height };
	static ofFbo* fbo = led_input_->fbo_;
	static ofPixels* pixels = led_input_->pixels_;
#ifndef __arm__
	shader_->load(shader_name_);
#endif
	time = ofGetElapsedTimef();
	fbo->begin();
	shader_->begin();
	shader_->setUniform1f("time", time);
	shader_->setUniform2fv("resolution", resolution);
	shader_->setUniform1f("gamma", globals::shader_gamma);
	shader_->setUniform1f("speed", globals::shader_speed);
	shader_->setUniform1f("scale", globals::shader_scale);
	shader_->setUniform1f("brightness", globals::shader_brightness);
	shader_->setUniform1f("contrast", globals::shader_contrast);
	ofDrawRectangle(-width, -height, 2 * width, 2 * height);
	shader_->end();
	fbo->end();
	fbo->readToPixels(*pixels);
}