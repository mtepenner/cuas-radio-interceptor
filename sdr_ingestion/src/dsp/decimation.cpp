#include "decimation.hpp"

#include <cmath>

namespace sdr {

std::vector<std::complex<float>> decimate(const std::vector<std::complex<float>>& input, std::size_t factor) {
    if (factor <= 1 || input.empty()) {
        return input;
    }

    std::vector<std::complex<float>> output;
    output.reserve(input.size() / factor + 1);
    for (std::size_t index = 0; index < input.size(); index += factor) {
        output.push_back(input[index]);
    }
    return output;
}

double estimate_rssi_dbm(const std::vector<std::complex<float>>& samples) {
    if (samples.empty()) {
        return -120.0;
    }

    double power_sum = 0.0;
    for (const auto& sample : samples) {
        power_sum += std::norm(sample);
    }

    const auto mean_power = power_sum / static_cast<double>(samples.size());
    return 10.0 * std::log10(mean_power + 1e-9) + 30.0;
}

}  // namespace sdr
