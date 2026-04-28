#pragma once

#include <complex>
#include <vector>

namespace sdr {

std::vector<std::complex<float>> decimate(const std::vector<std::complex<float>>& input, std::size_t factor);
double estimate_rssi_dbm(const std::vector<std::complex<float>>& samples);

}  // namespace sdr
