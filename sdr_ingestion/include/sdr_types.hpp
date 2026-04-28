#pragma once

#include <chrono>
#include <complex>
#include <iomanip>
#include <sstream>
#include <string>
#include <vector>

namespace sdr {

struct SpectrumBin {
    double frequency_mhz{};
    double power_db{};
};

struct ScanFrame {
    std::chrono::system_clock::time_point captured_at{};
    double center_frequency_mhz{};
    double bandwidth_mhz{};
    double noise_floor_db{};
    double rssi_dbm{};
    std::vector<std::complex<float>> iq_samples;
    std::vector<double> magnitudes;
    std::vector<SpectrumBin> peaks;
};

inline std::string to_json(const ScanFrame& frame) {
    std::ostringstream out;
    const auto captured = std::chrono::duration_cast<std::chrono::milliseconds>(frame.captured_at.time_since_epoch()).count();
    out << std::fixed << std::setprecision(3);
    out << "{\"captured_at_ms\":" << captured;
    out << ",\"center_frequency_mhz\":" << frame.center_frequency_mhz;
    out << ",\"bandwidth_mhz\":" << frame.bandwidth_mhz;
    out << ",\"noise_floor_db\":" << frame.noise_floor_db;
    out << ",\"rssi_dbm\":" << frame.rssi_dbm;
    out << ",\"peaks\":[";
    for (std::size_t index = 0; index < frame.peaks.size(); ++index) {
        const auto& peak = frame.peaks[index];
        if (index > 0) {
            out << ',';
        }
        out << "{\"frequency_mhz\":" << peak.frequency_mhz << ",\"power_db\":" << peak.power_db << '}';
    }
    out << "]}";
    return out.str();
}

}  // namespace sdr
