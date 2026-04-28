FROM gcc:14 AS build
WORKDIR /app

COPY include ./include
COPY src ./src

RUN g++ -std=c++20 -O2 -Iinclude -Isrc src/main.cpp src/dsp/fast_fourier.cpp src/dsp/decimation.cpp -o /sdr-ingestion

FROM debian:bookworm-slim
COPY --from=build /sdr-ingestion /usr/local/bin/sdr-ingestion

ENTRYPOINT ["/usr/local/bin/sdr-ingestion"]
