#include "ofMain.h"
#include "ofApp.h"

#include "globals.h"

#include <signal.h>

void KillHandler(int i) {
	OF_EXIT_APP(1);
}

//========================================================================
int main( ){

#ifdef TARGET_OPENGLES
	ofGLESWindowSettings settings;
	settings.glesVersion = 3;
	settings.setSize(globals::output_width, globals::output_height);
	ofCreateWindow(settings);
	signal(SIGQUIT, KillHandler);
	signal(SIGINT, KillHandler);
#else
	ofSetupOpenGL(globals::output_width, globals::output_height, OF_WINDOW);
	signal(SIGINT, KillHandler);
#endif
	ofSetFrameRate(globals::frame_rate);
	ofRunApp(new ofApp());

}
