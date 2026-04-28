#pragma once

#include <atomic>
#include <cstddef>
#include <optional>
#include <vector>

namespace sdr {

template <typename T>
class CircularBuffer {
public:
    explicit CircularBuffer(std::size_t capacity)
        : capacity_(capacity), storage_(capacity), write_index_(0), size_(0) {}

    void push(const T& value) {
        const auto index = write_index_.fetch_add(1, std::memory_order_relaxed) % capacity_;
        storage_[index] = value;
        auto current = size_.load(std::memory_order_relaxed);
        while (current < capacity_ && !size_.compare_exchange_weak(current, current + 1, std::memory_order_relaxed)) {
        }
    }

    std::vector<T> snapshot() const {
        const auto current_size = size_.load(std::memory_order_relaxed);
        const auto current_write = write_index_.load(std::memory_order_relaxed);
        std::vector<T> values;
        values.reserve(current_size);
        for (std::size_t offset = 0; offset < current_size; ++offset) {
            const auto index = (current_write + capacity_ - current_size + offset) % capacity_;
            if (storage_[index].has_value()) {
                values.push_back(*storage_[index]);
            }
        }
        return values;
    }

private:
    std::size_t capacity_;
    std::vector<std::optional<T>> storage_;
    std::atomic<std::size_t> write_index_;
    std::atomic<std::size_t> size_;
};

}  // namespace sdr
