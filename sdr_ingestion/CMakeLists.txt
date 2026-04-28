cmake_minimum_required(VERSION 3.20)
project(sdr_ingestion LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

add_executable(sdr_ingestion
    src/main.cpp
    src/dsp/fast_fourier.cpp
    src/dsp/decimation.cpp
)

target_include_directories(sdr_ingestion PRIVATE include src)
