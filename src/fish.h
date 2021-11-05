#pragma once

#include "types.h"
#include "wave.h"

// foward decleration
class FishRunner;

enum class PhoenixFishState : int {
	Reviving = 0,
	Waking = 1,
	Glowing = 2,
	Dying = 3,
	Dead = 4,
};

class PhoenixFish {
public:
	PhoenixFish(
		FishRunner* daddy, const uint64_t& starting_timestamp, const uint32_t& temporal_offset, 
		const uint32_t& glorious_lifespan, const uint8_t& fish_length, const uint8_t& stream_offset,
		const uint32_t& demise_fade
	);
private:
	bool CheckFishYetBreathesTheGoodSeaWater(const uint64_t& update_timestamp);
	void ExpendFishyLifeSource();
	FishRunner* daddy_;
	std::vector<std::vector<ofPoint>> trail_;
	uint64_t last_fishy_timestamp_;
	uint64_t elapsed_life_;
	uint32_t fish_incubation_period_;
	wave* wave_;
	uint32_t firey_glow_time_;
	uint32_t glorious_lifespan_;
	uint32_t demise_fade_;
	PhoenixFishState fishy_state_;
	friend FishRunner;
};

enum class FishRunnerState : int {
	Sleeping = 0,
	Running = 1,
};

class FishRunner {
public:
	FishRunner(LedInput* input);
	FishRunner() = delete;
	~FishRunner() = delete;
public:
	void UpdateFishRunner();
private:
	void ReviveThePhoenixFish(const uint64_t& timestamp);
	void DeterminePhoenixFishResurrectionSchedule();
	bool CheckAllFishDiedInAgony(const uint64_t& timestamp);
	void RunTheFishToTheirImminentDemise();
	LedInput* led_input_;
	FishRunnerState fishy_state_;
	uint64_t fishy_timestamp_;
	uint32_t phoenix_fish_revival_period_;
	std::vector<PhoenixFish> phoenix_fish_;
	friend PhoenixFish;
};