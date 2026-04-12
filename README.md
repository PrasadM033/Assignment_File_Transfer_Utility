#  Assignment: File Transfer Utility (TCP-based)

📌 Overview

This project implements a cross-platform file transfer utility using TCP sockets in Go.
It supports reliable and efficient transfer of large files (up to 16GB) between two networked systems.

The system is designed with a modular architecture, ensuring scalability, maintainability, and performance.

🚀 Features
✅ Transfer large files (up to 16 GB)
✅ Chunk-based streaming (memory efficient)
✅ Optional compression (gzip, based on file type)
✅ Integrity verification using SHA-256
✅ Error handling and reporting
✅ Cross-platform compatibility (Windows/Linux/macOS)
✅ Modular architecture (protocol, transport, fileio, compression, integrity)

🏗️ Project Structure
file-transfer/
├── cmd/
│   ├── sender/        # Sender application
│   └── receiver/      # Receiver application
├── internal/
│   ├── protocol/      # Packet encoding/decoding
│   ├── transport/     # TCP connection handling
│   ├── fileio/        # File read/write in chunks
│   ├── compression/   # Compression & decompression logic
│   └── integrity/     # SHA-256 hashing
├── go.mod
└── README.txt


▶️ How to Build
go mod tidy
go build ./cmd/sender
go build ./cmd/receiver

🔄 Transfer Flow
## Sender:
Compute SHA-256 hash
Decide compression
Send metadata (filename + compression + hash)
Send file in chunks
Send END signal
Wait for ACK
## Receiver:
Receive metadata
Parse compression + hash
Receive chunks
Decompress (if required)
Write file
Compute SHA-256
Compare with sender hash
Send ACK

![alt text](image.png)