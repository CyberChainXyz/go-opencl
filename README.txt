# go-opencl

go-opencl provides a high-level interface to the OpenCL API, allowing you to run OpenCL programs from Go.

## Development Status

**WARNING**: This project is currently under development and has not been fully tested. Use it at your own risk. We welcome any feedback and contributions.

## Requirements


```bash
sudo apt install ocl-icd-opencl-dev opencl-headers
```

## cl-info command

The cl-info command provides information about the OpenCL platforms and devices on your system.

To install cl-info, run the following command:

```bash
go get github.com/nexis-dev/go-opencl/cmd/cl-info
```

## OpenCL runner

```go
import cl "github.com/nexis-dev/go-opencl"
```

Refer to the [runner_test.go](./runner_test.go) file for usage examples of the OpenCL runner.

## Other resources

[go-cuda](https://github.com/nexis-dev/go-cuda): high-level Go interface to the CUDA API.

[ccxminer](https://github.com/cyberchain/ccxccxminer): CrytoNight-gpu miner using go-opencl and go-cuda.
