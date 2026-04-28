#pragma once

#include <complex>
#include <vector>

#include "sdr_types.hpp"

namespace sdr {

std::vector<double> compute_magnitude_spectrum(const std::vector<std::complex<float>>& samples);
std::vector<SpectrumBin> extract_top_peaks(const std::vector<double>& magnitudes, double center_frequency_mhz, double bandwidth_mhz, std::size_t limit);

}  // namespace sdr
