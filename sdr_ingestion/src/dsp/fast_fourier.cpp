#include "fast_fourier.hpp"

#include <algorithm>
#include <cmath>
#include <numbers>

namespace sdr {

std::vector<double> compute_magnitude_spectrum(const std::vector<std::complex<float>>& samples) {
    std::vector<double> magnitudes;
    if (samples.empty()) {
        return magnitudes;
    }

    const auto sample_count = samples.size();
    magnitudes.reserve(sample_count);
    for (std::size_t frequency_index = 0; frequency_index < sample_count; ++frequency_index) {
        std::complex<double> accumulator{0.0, 0.0};
        for (std::size_t sample_index = 0; sample_index < sample_count; ++sample_index) {
            const auto angle = -2.0 * std::numbers::pi * static_cast<double>(frequency_index * sample_index) / static_cast<double>(sample_count);
            accumulator += std::complex<double>(samples[sample_index]) * std::complex<double>(std::cos(angle), std::sin(angle));
        }
        magnitudes.push_back(20.0 * std::log10(std::abs(accumulator) + 1e-6));
    }
    return magnitudes;
}

std::vector<SpectrumBin> extract_top_peaks(const std::vector<double>& magnitudes, double center_frequency_mhz, double bandwidth_mhz, std::size_t limit) {
    std::vector<std::size_t> indexes(magnitudes.size());
    for (std::size_t index = 0; index < magnitudes.size(); ++index) {
        indexes[index] = index;
    }

    std::partial_sort(indexes.begin(), indexes.begin() + std::min(limit, indexes.size()), indexes.end(), [&](std::size_t left, std::size_t right) {
        return magnitudes[left] > magnitudes[right];
    });

    std::vector<SpectrumBin> peaks;
    const auto spacing = bandwidth_mhz / static_cast<double>(std::max<std::size_t>(magnitudes.size(), 1));
    for (std::size_t rank = 0; rank < std::min(limit, indexes.size()); ++rank) {
        const auto index = indexes[rank];
        peaks.push_back({center_frequency_mhz - bandwidth_mhz / 2.0 + spacing * static_cast<double>(index), magnitudes[index]});
    }
    return peaks;
}

}  // namespace sdr
