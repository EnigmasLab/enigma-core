---
sidebar_position: 1
title: Install Enigmad
---


# Install `Enigmad` 

`enigmad` is the command-line interface and daemon that connects to Enigma and enables you to interact with the Enigma blockchain. Enigma core is the official Golang reference implementation of the Enigma node software.

## Requirements

- Golang v1.16.1 - go1.17.1 linux/amd64
- Ensure your GOPATH and GOBIN environment variables are set up correctly.
- Linux users: install build-essential.
- 8 vCPU (4 physical core), x86_64 architecture processor
- 64 GB RAM (or equivalent swap file set up)
- 1 TB of storage space

## Option 1: From binary

The easiest way to install `enigmad` and Enigma core is by downloading a pre-built binary for your operating system. Download the Enigma Chain Staking-40021-1652947015 binaries from the official enigma-chain-releases.

```
wget https://github.com/EnigmasLab/enigma-chain-releases/releases/download/v0.4.19-1652947015/linux-amd64.zip
```

This zip file will contain three binaries and a virtual machine:
- **`enigmad`** - the Enigma Chain daemon
- **`peggo`** - the Enigma Chain ERC-20 bridge relayer daemon
- **`enigma-exchange`** - the Enigma Exchange daemon
- **`libwasmvm.x86_64.so`** - the wasm virtual machine which is needed to execute smart contracts.

Unzip and add `enigmad`, `enigma-exchange` and `peggo` to your `/usr/bin`. Also add `libwasmvm.x86_64.so` to user library path `/usr/lib`.

```bash
unzip linux-amd64.zip
sudo mv enigmad peggo enigma-exchange /usr/bin
sudo mv libwasmvm.x86_64.so /usr/lib
```

Check your binary version by running following commands.

```bash
enigmad version
peggo version
enigma-exchange version
```

Confirm your version matches the output below

```bash
enigmad version
Version dev (f32e524)

peggo version
Version dev (b5c188c)

enigma-exchange version
Version dev (ca1da5e)
```

## Option 2: From source

Note: you will only install `enigmad` but not `enigma-exchange`, `peggo` or `libwasmvm.x86_64.so` using this option.

### Get the Enigma core source code

Use git to retrieve [Enigma core](https://github.com/OpenDeFiFoundation/enigma-core).

Clone the enigma repo:

```bash
git clone https://github.com/OpenDeFiFoundation/enigma-core
```

### Build Enigma core from source

Build Enigma core, and install the `enigmad` executable to your GOPATH environment variable.

```bash
cd enigma-core
make install
```

### Verify your Enigma core installation

Verify that Enigma core is installed correctly.

```bash
enigmad version
```

