# slurpy
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Version](https://img.shields.io/badge/version-0.1.5.2-brightgreen.svg)

Early build 0.1.5 is available, it has limited functionality however it covers the core use case available in the description bellow, do not use for the mainnet and only store disposable wallets that don't have any real currency in them!


slurpy is an open-source tool written in Go for quickly deploying smart contracts based on bytecode and ABI. It streamlines the deployment process, when copying existing network infrastructure from one chain to another.

## Description

I created this tool out of necessity while working on my latest project, a shared financial responsibility wallet. Due to QA requirements, I had to run a local chain because of constraints on testnet tokens.

Manually compiling and downloading all smart contract infrastructure on other blockchains can be a tedious process. Managing swaps, smart contracts, currencies, and mimicking the behavior of interactions on a local network can be time-consuming and impractical, especially when trying to maintain a copy of the infrastructure.

To address these challenges, I opted for this tool, which allows me to retrieve the bytecode of existing infrastructure from a preferred chain, set up a migration file, and deploy a copy of the chain with nothing more than the bytecode of all the smart contracts.

While tools like Genache have a forking option in a lot of cases especially when you have a bit bigger team you either need to pay for your own RPC when forking so you don't hit a rate limit or simply wait tedious amount of time before a block is mined.

With this tool i solved my issue and hopefully did that for least one more person trying to set up an easily relocatable development environment that can be run both locally or on a server in minutes just using simple migration files and slurpy for deployments.

## Features

- Deploy smart contracts using only the bytecode and ABI.
- Lightweight and efficient, leveraging the power of Go.
- Supports Ethereum-based networks.

## Installation

To install slurpy, clone the repository and build the project:

```bash
git clone https://github.com/KristiforMilchev/slurpy.git
cd slurpy
go build
