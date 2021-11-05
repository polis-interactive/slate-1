
#include "lighting.h"
#include <tuple>

std::vector<Light*> Light::all_lights;

const GridDefintions Light::SetupLights(LedDefinitions &led_defs, ofxRpiWs281x::LedOutputGenerator *output_generator) {
	auto grid_def = GridDefintions();
	for (auto led_def : led_defs) {
		new Light(led_def, output_generator);
		const auto& point = led_def->_pixel_position;
		grid_def.TrySetMinMax(point.x, point.y);
	}
	return grid_def;
}

Light::Light(LedProxy *proxy, ofxRpiWs281x::LedOutputGenerator *output_generator) {
	SetColor(0);
	output_color_ = output_generator->GetPixel(proxy->_pixel_number);
	position_ = proxy->_pixel_position;
	is_on_ = proxy->_is_on;
	all_lights.push_back(this);
}

void Light::UpdateLights(LedInput* input) {
	for (auto light : all_lights) {
		if (!light->is_on_) {
			input->TurnOffRgbPixel(light->position_);
		}
		else {
			const ofColor c = input->GetRGBPixel(light->position_);
			light->SetColor(c);
		}
	}
}

Light* Light::GetLightOrNull(const ofPoint &p) {
	auto got_light = std::find_if(all_lights.begin(), all_lights.end(), [&](Light* light) {
		return light->position_ == p;
	});
	if (got_light == all_lights.end()) {
		return NULL;
	}
	auto light = *got_light;
	if (!light->is_on_) {
		return NULL;
	}
	else {
		
		return light;
	}
}

void Light::SetColor(const ofColor& c) {
	current_color_ = c;
}


void Light::SetColor(const int i) {
	current_color_.set(i);
}

void Light::SetColor(uint8_t* color_data) {
	memcpy(current_color_.v, color_data, sizeof(uint8_t) * 3);
}

void Light::SetColorAll(const ofColor& c) {
	for (auto light : all_lights) {
		light->SetColor(c);
	}
}

void Light::ResetColorAll() {
	for (auto light : all_lights) {
		light->SetColor(0);
	}
}

void Light::DrawLights() {
	for (auto light : all_lights) {
		light->DrawLight();
	}
}

void Light::DrawLight() {
	memcpy(output_color_->v, current_color_.v, sizeof(uint8_t) * 3);
}
