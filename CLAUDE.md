# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A minimal Go application that derives an Ethereum wallet address from a BIP39 mnemonic phrase. Designed as a containerized app for deployment on the EigenX Trusted Execution Environment platform.

## Build & Run Commands

```bash
# Install dependencies
go mod download

# Run locally (requires .env with MNEMONIC set; copy .env.example as starting point)
go run src/main.go

# Build binary
go build -o app ./src

# Docker build and run
docker build -t my-app .
docker run --rm --env-file .env my-app
```

There are currently no tests in the project. Standard Go test commands would apply:
```bash
go test ./...                  # run all tests
go test ./src/...              # run tests in src
go test -run TestName ./src    # run a single test
```

## Architecture

Single-file application (`src/main.go`) with no internal package structure. The data flow is:

1. Load environment variables from `.env` via `godotenv`
2. Read `MNEMONIC` environment variable (required)
3. Create HD wallet from mnemonic using `go-ethereum-hdwallet`
4. Derive first Ethereum account at path `m/44'/60'/0'/0/0`
5. Print the derived address to stdout

## Key Dependencies

- `github.com/joho/godotenv` — loads `.env` files
- `github.com/miguelmota/go-ethereum-hdwallet` — Ethereum HD wallet derivation from BIP39 mnemonics

## Deployment

Deployed via EigenX CLI (`eigenx app deploy`). See README.md for full EigenX lifecycle commands (deploy, start, stop, logs, terminate, upgrade). Optional Caddy reverse proxy can be configured for TLS via environment variables (`DOMAIN`, `APP_PORT`, `ACME_*`).
