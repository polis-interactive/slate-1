
#include "fish.h"
#include "utility.h"
#include "globals.h"

FishRunner::FishRunner(LedInput* input) : led_input_(input) {
	DeterminePhoenixFishResurrectionSchedule();
}

void FishRunner::DeterminePhoenixFishResurrectionSchedule() {
	// might change this to be preprogrammed instead of dynamically generated
	// step through an array of predetermined periods
	static uint8_t res_type;
	fishy_state_ = FishRunnerState::Sleeping;
	phoenix_fish_.clear();
	fishy_timestamp_ = ofGetElapsedTimeMillis();
	res_type = GetUniformRandom(0, 8);
	// 1/4 are short ressurections
	if (res_type < 2) {
		phoenix_fish_revival_period_ = GetUniformRandom(300, 1000);
	}
	// 1/2 are medium ressurections
	else if (res_type < 6) {
		phoenix_fish_revival_period_ = GetUniformRandom(3000, 5000);
	}
	// 1/4 are long ressurections
	else {
		phoenix_fish_revival_period_ = GetUniformRandom(8000, 12000);
	}
}

bool FishRunner::CheckAllFishDiedInAgony(const uint64_t& timestamp) {
	bool has_any_living_fish = false;
	for (auto& fish : phoenix_fish_) {
		if (fish.fishy_state_ == PhoenixFishState::Dead) {
			continue;
		}
		bool can_go_on = fish.CheckFishYetBreathesTheGoodSeaWater(timestamp);
		if (can_go_on) {
			has_any_living_fish = true;
		}
	}
	return !has_any_living_fish;
}

void FishRunner::UpdateFishRunner() {
	const uint64_t timestamp = ofGetElapsedTimeMillis();
	const uint64_t elapsed_time = timestamp - fishy_timestamp_;
	switch (fishy_state_) {
	case FishRunnerState::Sleeping:
		if (elapsed_time > phoenix_fish_revival_period_) {
			ReviveThePhoenixFish(timestamp);
		}
		break;
	case FishRunnerState::Running:
		if (CheckAllFishDiedInAgony(timestamp)) {
			DeterminePhoenixFishResurrectionSchedule();
		}
		else {
			RunTheFishToTheirImminentDemise();
		}
		break;
	}
}

void FishRunner::RunTheFishToTheirImminentDemise() {
	for (auto& fish : phoenix_fish_) {
		fish.ExpendFishyLifeSource();
	}
}

void FishRunner::ReviveThePhoenixFish(const uint64_t& timestamp) {
	fishy_state_ = FishRunnerState::Running;
	const uint8_t fish_count = 1;
	uint8_t cum_fish_length_ = 0;
	uint32_t cum_fish_offset_ = 0;
	for (int i = 0; i < fish_count; i++) {
		uint8_t fish_length = (uint8_t)globals::stream_positions.size() / fish_count;
		if (i == 0) {
			const uint8_t fish_remains = globals::stream_positions.size() % fish_count;
			fish_length += fish_remains;
		}
		auto fish = PhoenixFish(this, timestamp, cum_fish_offset_, 3000, fish_length, cum_fish_length_, 1000);
		phoenix_fish_.push_back(fish);
		cum_fish_length_ += fish_length;
		cum_fish_offset_ += 200;
	}
}

PhoenixFish::PhoenixFish(
	FishRunner* daddy, const uint64_t& starting_timestamp, const uint32_t& temporal_offset,
	const uint32_t& glorious_lifespan, const uint8_t& fish_length, const uint8_t& stream_offset,
	const uint32_t& demise_fade
)
	: daddy_(daddy), last_fishy_timestamp_(starting_timestamp), glorious_lifespan_(glorious_lifespan),
		fish_incubation_period_(temporal_offset), demise_fade_(demise_fade)
{
	fishy_state_ = temporal_offset != 0 ? PhoenixFishState::Reviving : PhoenixFishState::Waking;
	auto inital_segment = globals::stream_positions.begin() + stream_offset;
	auto end_segment = inital_segment + fish_length;
	std::copy(inital_segment, end_segment, std::back_inserter(trail_));
	wave_ = new wave(starting_timestamp, fish_length);
}

bool PhoenixFish::CheckFishYetBreathesTheGoodSeaWater(const uint64_t& update_timestamp) {
	elapsed_life_ = update_timestamp - last_fishy_timestamp_;
	switch (fishy_state_) {
	case PhoenixFishState::Reviving:
		if (elapsed_life_ > fish_incubation_period_) {
			wave_->setup_wave(update_timestamp);
			fishy_state_ = PhoenixFishState::Waking;
			last_fishy_timestamp_ = update_timestamp;
			elapsed_life_ = 0;
		}
		break;
	case PhoenixFishState::Waking:
		if (!wave_->should_wave_run(update_timestamp)) {
			fishy_state_ = PhoenixFishState::Glowing;
			last_fishy_timestamp_ = update_timestamp;
			elapsed_life_ = 0;
			delete wave_;
		}
		break;
	case PhoenixFishState::Glowing:
		if (elapsed_life_ > glorious_lifespan_) {
			fishy_state_ = PhoenixFishState::Dying;
			last_fishy_timestamp_ = update_timestamp;
			elapsed_life_ = 0;
		}
		break;
	case PhoenixFishState::Dying:
		if (elapsed_life_ > demise_fade_) {
			fishy_state_ = PhoenixFishState::Dead;
			elapsed_life_ = 0;
			return false;
		}
	}
	return true;
}

void PhoenixFish::ExpendFishyLifeSource() {

	switch (fishy_state_) {
	case PhoenixFishState::Waking: {
		ofColor c;
		for (unsigned int i = 0; i < wave_->node_count; i++) {
			const int wave_position = i + wave_->wave_tail_pointer;
			if (wave_position < 0 || wave_position >= wave_->wave_length) continue;
			const float wave_fraction = wave_->get_wave_fraction(wave_position, wave_position);
			if (wave_fraction >= 1.0) {
				c = ofColor::white;
			}
			else if (wave_fraction <= 0.0) {
				c = ofColor::black;
			}
			else {
				c = ofColor::black.getLerped(ofColor::white, wave_fraction);
			}
			auto& led_crew = trail_[wave_position];
			for (auto& position : led_crew) {
				daddy_->led_input_->SetProperPixelColor(position, c);
			}
		}
		int get_half_point = wave_->get_wave_half_position();
		for (int i = 0; i  < wave_->wave_length; i++) {
			if (i > get_half_point) return;
			auto& led_crew = trail_[i];
			for (auto& position : led_crew) {
				daddy_->led_input_->SetProperPixelColor(position, ofColor::white);
			}
			if (i == wave_->wave_length - 1) {
				wave_->force_finish();
			}
		}
		break;
	}
	case PhoenixFishState::Glowing:
		for (auto& white_group : trail_) {
			for (auto& position : white_group) {
				daddy_->led_input_->SetProperPixelColor(position, ofColor::white);
			}
		}
		break;
	case PhoenixFishState::Dying: {
		const float how_dead_is_he = std::min(std::max((elapsed_life_ / (float)demise_fade_), 0.0f), 1.0f);
		for (auto& white_group : trail_) {
			for (auto& position : white_group) {
				daddy_->led_input_->SetProperPixelTint(position, how_dead_is_he);
			}
		}
		break;
	}
	case PhoenixFishState::Dead:
	case PhoenixFishState::Reviving:
		break;
	}
}