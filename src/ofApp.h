#pragma once

#include "ofMain.h"

#include "ofxRpiWs281x.h"

#include "types.h"

class ofApp : public ofBaseApp{
public:
	void setup();
	void update();
	void draw();
	void exit();

	void keyPressed(int key);
	void gotMessage(ofMessage msg);

private:
	static ofxRpiWs281x::LedStrip *led_strip;
	static LedInput *led_input;
};
