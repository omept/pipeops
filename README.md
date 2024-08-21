# Go Pipeline Operations

## Overview

This project demonstrates a data processing pipeline in Go using goroutines and channels. The pipeline consists of three stages: a generator, a squaring function, and a summation function.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.22+)

### Running the Program

1. Clone the repository:

    ```bash
    git clone https://github.com/omept/pipeops.git
    cd pipeops
    ```

2. Build and run the program:

    ```bash
    go run main.go
    ```

   You can also specify the number of integers to generate using the `-n` flag:

    ```bash
    go run main.go -n 5000
    ```

### Running Tests

To run the unit tests, execute the following command:

```bash
go test -v
