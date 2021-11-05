#pragma once

#include "types.h"
#include "lighting.h"

#include "ofxArtnetReceiver.h"
#include "ofxArtnetMessage.h"

#define ARTNET_PORT 6454

class ArtnetGraphics :
	public LedInput
{
public:
	ArtnetGraphics(LedInputInitializer&);
	ArtnetGraphics() = delete;
	~ArtnetGraphics() = delete;
public:
	void UpdateInput();
	void Teardown();
private:
	void handleArtnetMessage(ofxArtnetMessage&);
private:
	ofxArtnetReceiver *receiver_;
};