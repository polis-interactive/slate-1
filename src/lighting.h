#pragma once

#include "ofColor.h"
#include "ofPoint.h"

#include "ofxRpiWs281x.h"

#include "globals.h"
#include "types.h"

#include <map>
#include <vector>

class Light {
public:
	static const GridDefintions SetupLights(LedDefinitions&, ofxRpiWs281x::LedOutputGenerator*);
	static void DrawLights();
	static void UpdateLights(LedInput*);
	static Light* GetLightOrNull(const ofPoint&);
private:
	Light(LedProxy*, ofxRpiWs281x::LedOutputGenerator*);
	~Light() = delete;
	Light(const Light&) = delete;
	void operator=(const Light&) = delete;
public:
	void SetColor(const ofColor& c);
	void SetColor(const int i);
	void SetColor(uint8_t*);
	static void SetColorAll(const ofColor& c);
	static void ResetColorAll();
private:
	void DrawLight();
	ofPoint position_;
	bool is_on_;
	ofColor current_color_;
	ofColor* output_color_;
	static std::vector<Light*> all_lights;
};


