
#include "ofApp.h"

#include "boards.h"
#include "globals.h"
#include "lighting.h"
#include "sequenced.h"


ofxRpiWs281x::LedStrip* ofApp::led_strip = NULL;
LedInput* ofApp::led_input = NULL;

//--------------------------------------------------------------
void ofApp::setup() {

	// generate data from board configuration
	std::cout << "ofApp, setup: Starting up " << std::endl;


	std::cout << "ofApp, setup: setting up boards " << std::endl;

	auto board_output = SetupLedBoards(globals::board_configs);
	LedDefinitions &led_proxys = board_output.first;
	const auto led_count = board_output.second;



	// initialize led strips
	std::cout << "ofApp, setup: setting up strips " << std::endl;

	auto strip_config = ofxRpiWs281x::LedStripConfiguration();
	strip_config.gamma = globals::gamma;
	strip_config.brightness = globals::brightness;
	strip_config.gpio_pin = ofxRpiWs281x::GpioPins::GPIO_19;
	strip_config.strip_type = ofxRpiWs281x::StripType::WS2811_GRB;
	strip_config.led_count = led_count;

	auto strip_def = ofxRpiWs281x::LedStrip::CreateLedStrip(strip_config);
	auto strip_create_ret = strip_def.second;
	if (strip_create_ret.isFailure()) {
		std::cout << "ofxRpiWs281x, CreateLedStrip: failed with error code " << strip_create_ret.Ret() << std::endl;
		exit();
	}


	std::cout << "ofApp, setup: initializing strips " << std::endl;

	led_strip = strip_def.first;
	auto strip_init_ret = led_strip->Initialize();
	if (strip_init_ret.isFailure()) {
		std::cout << "ofxRpiWs281x, Initialize: failed with error code " << strip_init_ret.Ret() << std::endl;
		exit();
	}



	// initialize led proxys
	std::cout << "ofApp, setup: initializing leds " << std::endl;

	auto grid_defs = Light::SetupLights(led_proxys, led_strip);

	std::cout << "x: " << grid_defs._min_grid_x << ", " << grid_defs._max_grid_x << std::endl;
	std::cout << "y: " << grid_defs._min_grid_y << ", " << grid_defs._max_grid_y << std::endl;



	// initalize input
	std::cout << "ofApp, setup: initializing graphics " << std::endl;

	auto input_init = LedInputInitializer(grid_defs, globals::segment_size);

	led_input = new SequencedGraphics(input_init);

	std::cout << "ofApp, setup: done " << std::endl;

}

//--------------------------------------------------------------
void ofApp::update() {
	Light::ResetColorAll();
	led_input->UpdateInput();
	Light::UpdateLights(led_input);
}

//--------------------------------------------------------------
void ofApp::draw() {

	Light::DrawLights();

	auto strip_render_ret = led_strip->Render();
	if (strip_render_ret.isFailure()) {
		std::cout << "ofxRpiWs281x, Render: failed with error code " << strip_render_ret.Ret() << std::endl;
	}

#ifndef __arm__
	led_input->DrawInput();
#endif

}

//--------------------------------------------------------------
void ofApp::exit() {
	if (led_strip != NULL) {
		led_strip->Teardown();
	}
	if (led_input != NULL) {
		led_input->Teardown();
	}
}

//--------------------------------------------------------------
void ofApp::keyPressed(int key) {
	if (key == 'q' || key == OF_KEY_ESC) {
		OF_EXIT_APP(1);
	}
}

//--------------------------------------------------------------
void ofApp::gotMessage(ofMessage msg) {

}

