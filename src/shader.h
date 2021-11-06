#pragma once

#include "types.h"


class Shader {
public:
	Shader(LedInput* input, std::string shaderName);
	Shader() = delete;
	~Shader() = delete;
public:
	void UpdateShader();
private:
	LedInput* led_input_;
	ofShader *shader_;
	std::string shader_name_;
};

