# 📡 C-UAS Radio Interceptor

## Description
The C-UAS (Counter Unmanned Aerial Systems) Radio Interceptor is a high-performance Signals Intelligence (SIGINT) platform designed to detect and classify drone telemetry and control links. The architecture utilizes a C++ engine for high-speed raw IQ data capture via SDR hardware, passes the stream to a Go-based pattern recognition brain for signature matching, and visualizes the RF environment in a tactical React/WebGL dashboard.

## 📑 Table of Contents
- [Features](#-features)
- [Technologies Used](#-technologies-used)
- [Installation](#-installation)
- [Usage](#-usage)
- [Project Structure](#-project-structure)
- [Contributing](#-contributing)
- [License](#-license)

## 🚀 Features
* **High-Speed SDR Ingestion:** Built in C++ to interface with `librtlsdr` or `libhackrf`. Utilizes lock-free circular buffers for raw IQ storage and performs fast fourier transforms via FFTW or Liquid-DSP.
* **Pattern Recognition Engine:** A Go-based classifier that identifies specific hopping sequences and drone protocols, including ExpressLRS and DJI OcuSync.
* **Distance Estimation:** Employs the Friis Path Loss model alongside burst detection to estimate the range of incoming drone signals.
* **Tactical WebGL Spectrogram:** A React and TypeScript frontend that renders a rolling waterfall display using custom GLSL shaders to map signal power to color scales. 
* **Real-time Threat HUD:** Instantly pops up threat alerts when a recognized drone signature is matched, displaying RSSI-based distance metrics.

## 🛠️ Technologies Used
* **SDR Ingestion:** C++, FFTW / Liquid-DSP, librtlsdr / libhackrf
* **Signal Classifier:** Go, ZeroMQ (for stream ingestion)
* **Frontend Dashboard:** React, TypeScript, WebGL
* **Infrastructure:** Docker, Kubernetes, TimescaleDB

## ⚙️ Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mtepenner/cuas-radio-interceptor.git
   cd cuas-radio-interceptor
   ```

2. Ensure your SDR hardware (e.g., RTL-SDR or HackRF) is plugged into the host machine. 

3. Launch the stack using Docker Compose. *Note: The SDR gateway container requires privileged access to communicate with the USB hardware.*
   ```bash
   docker-compose up -d
   ```

## 💻 Usage

* **Access the Dashboard:** Navigate to `http://localhost:3000` to access the tactical SIGINT spectrogram and observe the real-time RF waterfall.
* **Monitor Alerts:** Watch the Threat Alerts HUD for positive identifications of OcuSync or ExpressLRS signals within the monitored airspace.
* **Historical Review:** Utilize the connected TimescaleDB instance to review and analyze historical signal logs for post-action reports.

## 📂 Project Structure
* `/sdr_ingestion`: C++ layer handling hardware interfacing, decimation, and FFT processing.
* `/signal_classifier`: Go application orchestrating the signal analysis pipeline and protocol signature matching.
* `/tactical_dashboard`: React frontend featuring the WebGL waterfall display and signal strength meters.
* `/infrastructure`: YAML manifests for deploying the SDR gateway, compute nodes, and database.
* `/.github/workflows`: CI/CD pipelines to validate DSP algorithms and FFT math.

## 🤝 Contributing
Contributions are welcome! If you are modifying the Digital Signal Processing (DSP) logic, please ensure that all FFT and frequency hopping detection math passes the included GitHub Actions CI/CD workflows.

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
