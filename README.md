# Word of Wisdom TCP Server

This project is a TCP server with DDoS protection using Proof of Work (PoW) and provides random quotes.

## Project Structure

- **cmd/server**: server entry point and its Dockerfile.
- **cmd/client**: client entry point and its Dockerfile.
- **pkg/pow**: Proof of Work implementation.
- **pkg/wisdom**: random quotes.
- **internal/server**: server logic.
- **internal/client**: client logic.

## Running the Project

```bash
make build-server
make build-client
make run-server
make run-client
```