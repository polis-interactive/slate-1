#pragma once

#include <cmath>

# define M_PI 3.14159265358979323846

struct wave {
    double start_time;
    float current_offset;
    unsigned int node_count;
    float node_duration;
    float max_fraction;
    int wave_tail_pointer;
    int wave_length;
    bool wave_is_backwards;
    float period;
    float time_const;
    /* moved to cpp so it can use global color */
    void setup_wave(const uint64_t& timestamp) {
        start_time = (double) timestamp;
    }
    bool is_finished() {
        return wave_tail_pointer >= wave_length;
    }
    void force_finish() {
        wave_tail_pointer = wave_length;
    }
    bool should_wave_run(const uint64_t& timestamp) {
        if (is_finished()) return false;
        return increment_pointer((double)timestamp);
    }
    bool increment_pointer(double timestamp) {
        current_offset = (timestamp - start_time);
        int tail_pointer = std::floor(current_offset / node_duration) - node_count + 1;
        if (tail_pointer > wave_tail_pointer) {
            ++wave_tail_pointer;
            if (is_finished()) return false;
        }
        const int timestep = fmodf(current_offset, node_duration);
        time_const = (float)timestep / node_duration + (float)wave_tail_pointer - 1.0;
        return true;
    }
    float get_wave_fraction(int node_number) {
        return max_fraction * sin(((float)node_number - time_const) * period);
    }
    int get_wave_half_position() {
        return wave_length / 2 + wave_tail_pointer;
    }

    float get_wave_fraction(int node_number, int position) {
        if (wave_length / 2 - position > 0) {
            return 1.0f;
        }
        else return get_wave_fraction(node_number);
    }
    wave(const uint64_t timestamp, int wave_length_) {
        current_offset = 0;
        time_const = 0;
        start_time = (double)timestamp;
        node_count = wave_length_ * 1.5;
        node_duration = 50;
        max_fraction = 1.3;
        wave_length = wave_length_;
        wave_is_backwards = false;
        wave_tail_pointer = -node_count + 1;
        period = M_PI / (float)node_count;
    }
};
