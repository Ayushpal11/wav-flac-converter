# WAV to FLAC Audio Conversion Service

This is a Go-based backend service that takes WAV audio files, converts them to FLAC format in real-time, and returns the output back to the client using WebSockets. This project is designed to support high-fidelity, low-latency audio conversion, suitable for applications requiring audio conversion.

## Table of Contents
- [Features](#features)
- [Technologies](#technologies)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Docker Setup](#docker-setup)
- [Troubleshooting](#troubleshooting)

## Features
- Real-time audio conversion from WAV to FLAC
- High-fidelity audio conversion
- Low-latency audio conversion

## Technologies
- Go
- WebSockets
- FFmpeg

## Prerequisites
- Go 1.16 or higher
- FFmpeg

## Installation
1. Clone the repository
2. Run `go mod download` to install dependencies
3. Run `go run main.go` to start the server

## Usage
1. Start the server using `go run main.go`
2. Connect to the WebSocket server using the appropriate client
3. Send a WAV audio file to the server
4. Receive the converted FLAC audio file from the server

## API Endpoints
- `/ws`: WebSocket endpoint for audio conversion

## Docker Setup
1. Build the Docker image using `docker build -t wav-flac-converter .`
2. Run the Docker container using `docker run -p 8080:8080 wav-flac-converter`

## Troubleshooting
- If you encounter any issues with the audio conversion, make sure FFmpeg is installed and accessible in your PATH.
- If you encounter any issues with the server, make sure the correct port is being used and no other services are running on that port.


## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
