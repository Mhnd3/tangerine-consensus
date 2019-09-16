# Tangerine Consensus

[![Build Status](https://travis-ci.org/tangerine-network/tangerine-consensus.svg?branch=master)](https://travis-ci.org/tangerine-network/tangerine-consensus)

## Getting Started

### Prerequisites

- [Go 1.10](https://golang.org/dl/) or a newer version
- [dep](https://github.com/golang/dep#installation) as dependency management

### Installation

1. Clone the repo

   ```
   git clone https://github.com/tangerine-network/tangerine-consensus.git
   cd tangerine-consensus
   ```

2. Setup GOPATH, the GOPATH could be anywhere in the system. Here we use `$HOME/go`:
   ```
   export GOPATH=$HOME/go
   export PATH=$GOPATH/bin:$PATH
   ```
   You should write these settings to your `.bashrc` file.

3) Install go dependency management tool

   ```
   ./bin/install_tools.sh
   ```

4) Install all dependencies
   ```
   dep ensure
   ```

### Run Unit Tests

```
make pre-submit
```

## Simulation

### Simulation with Nodes connected by HTTP

1. Setup the configuration under `./test.toml`
2. Compile and install the cmd `dexon-simulation`

```
make
```

3. Run simulation:

```
dexcon-simulation -config test.toml -init
```

### Simulation with test.Scheduler

1. Setup the configuration under `./test.toml`
2. Compile and install the cmd `dexon-simulation-with-scheduler`

```
make
```

3. Run simulation with scheduler:

```
dexcon-simulation-with-scheduler -config test.toml
```
