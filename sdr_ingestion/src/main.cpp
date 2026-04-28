#include <chrono>
#include <cmath>
#include <complex>
#include <iostream>
#include <numbers>
#include <thread>
#include <vector>

#include "buffer/circular_buffer.hpp"
#include "dsp/decimation.hpp"
#include "dsp/fast_fourier.hpp"
#include "sdr_types.hpp"

namespace {

std::vector<std::complex<float>> generate_iq_samples(std::size_t sample_count, double base_frequency_hz, double hop_hz) {
    std::vector<std::complex<float>> samples;
    samples.reserve(sample_count);
    for (std::size_t index = 0; index < sample_count; ++index) {
        const auto t = static_cast<double>(index) / 2'000'000.0;
        const auto theta_primary = 2.0 * std::numbers::pi * (base_frequency_hz * t);
        const auto theta_secondary = 2.0 * std::numbers::pi * ((base_frequency_hz + hop_hz) * t);
        const auto real = static_cast<float>(0.65 * std::cos(theta_primary) + 0.35 * std::cos(theta_secondary));
        const auto imag = static_cast<float>(0.65 * std::sin(theta_primary) + 0.35 * std::sin(theta_secondary));
        samples.emplace_back(real, imag);
    }
    return samples;
}

}  // namespace

int main() {
    sdr::CircularBuffer<sdr::ScanFrame> history(8);
    constexpr double center_frequency_mhz = 2440.0;
    constexpr double bandwidth_mhz = 80.0;

    for (int frame_index = 0; frame_index < 5; ++frame_index) {
        auto samples = generate_iq_samples(96, 150'000.0 + frame_index * 5'000.0, 420'000.0);
        auto decimated = sdr::decimate(samples, 2);
        auto magnitudes = sdr::compute_magnitude_spectrum(decimated);

        sdr::ScanFrame frame;
        frame.captured_at = std::chrono::system_clock::now();
        frame.center_frequency_mhz = center_frequency_mhz;
        frame.bandwidth_mhz = bandwidth_mhz;
        frame.noise_floor_db = -96.5;
        frame.rssi_dbm = sdr::estimate_rssi_dbm(decimated);
        frame.iq_samples = std::move(decimated);
        frame.magnitudes = magnitudes;
        frame.peaks = sdr::extract_top_peaks(magnitudes, center_frequency_mhz, bandwidth_mhz, 4);

        history.push(frame);
        std::cout << sdr::to_json(frame) << std::endl;
        std::this_thread::sleep_for(std::chrono::milliseconds(250));
    }

    return history.snapshot().empty() ? 1 : 0;
}
